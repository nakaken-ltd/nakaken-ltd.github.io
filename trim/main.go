package main

import (
	"bytes"
	"cmp"
	"context"
	_ "embed"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"sync"
	"text/template"

	"github.com/kanrichan/resvg-go"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

//go:embed tsx.tmpl
var tsxTemplate string

type point struct {
	X, Y int
}

type stroke struct {
	X0, Y0, X1, Y1 int
}

type polygon struct {
	Name                  string
	OffsetTop, OffsetLeft int
	S                     []stroke
}

func (p polygon) Left() int {
	return slices.MinFunc(p.S, func(a, b stroke) int { return cmp.Compare(a.X0, b.X0) }).X0
}

func (p polygon) Right() int {
	return slices.MaxFunc(p.S, func(a, b stroke) int { return cmp.Compare(a.X0, b.X0) }).X0
}

func (p polygon) Top() int {
	return slices.MinFunc(p.S, func(a, b stroke) int { return cmp.Compare(a.Y0, b.Y0) }).Y0
}

func (p polygon) Bottom() int {
	return slices.MaxFunc(p.S, func(a, b stroke) int { return cmp.Compare(a.Y0, b.Y0) }).Y0
}

func (p polygon) Width() int {
	return p.Right() - p.Left()
}

func (p polygon) Height() int {
	return p.Bottom() - p.Top()
}

// Inside tests if the given point (x, y) is inside the polygon or not.
// It implements Crossing Number Algorithm.
// Ref: https://www.nttpc.co.jp/technology/number_algorithm.html
func (p polygon) Inside(x, y int) bool {
	count := 0
	for _, s := range p.S {
		if (s.Y0 <= y && s.Y1 > y) || (s.Y0 > y && s.Y1 <= y) {
			vt := float64(y-s.Y0) / float64(s.Y1-s.Y0)
			if float64(x) < (float64(s.X0) + (vt * float64(s.X1-s.X0))) {
				count++
			}
		}
	}
	return count%2 != 0
}

// Main procedures split into funcs for testability

func findCutLineLayer(doc *html.Node) *html.Node {
	for n := range doc.Descendants() {
		if n.Data != "g" {
			continue
		}
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == "cutline" {
				return n
			}
		}
	}
	return nil
}

func findPolygons(cutLineLayer *html.Node) (polygons []polygon, err error) {
	for n := range cutLineLayer.Descendants() {
		switch n.Data {
		case "rect":
			// x and y will be zero if omitted
			// width and height must be a valid float
			xf, yf, wf, hf, eid := 0.0, 0.0, math.NaN(), math.NaN(), ""

			for _, attr := range n.Attr {
				switch attr.Key {
				case "x":
					xf, err = strconv.ParseFloat(attr.Val, 64)
					if err != nil {
						err = fmt.Errorf("failed to parse rect x: %s (value = '%s')", err, attr.Val)
						return
					}
				case "y":
					yf, err = strconv.ParseFloat(attr.Val, 64)
					if err != nil {
						err = fmt.Errorf("failed to parse rect y: %s (value = '%s')", err, attr.Val)
						return
					}
				case "width":
					wf, err = strconv.ParseFloat(attr.Val, 64)
					if err != nil {
						err = fmt.Errorf("failed to parse rect width: %s (value = '%s')", err, attr.Val)
						return
					}
				case "height":
					hf, err = strconv.ParseFloat(attr.Val, 64)
					if err != nil {
						err = fmt.Errorf("failed to parse rect height: %s (value = '%s')", err, attr.Val)
						return
					}
				case "id":
					eid = attr.Val
				}
			}

			if math.IsNaN(wf) || math.IsNaN(hf) {
				err = fmt.Errorf("width or height was not found for a rect")
				return
			}

			x, y, w, h := int(math.Round(xf)), int(math.Round(yf)), int(math.Round(wf)), int(math.Round(hf))

			// Convert the rect into a polygon like the following:
			// (x,   y)---(x+w,   y)
			//   |            |
			//   |            |
			//   |            |
			// (x, y+h)---(x+w, y+h)

			polygons = append(polygons, polygon{
				Name: eid,
				S: []stroke{
					{x, y, x, y + h},
					{x, y + h, x + w, y + h},
					{x + w, y + h, x + w, y},
					{x + w, y, x, y},
				},
			})
		case "polygon":
			var ps []point
			var eid string
			for _, attr := range n.Attr {
				switch attr.Key {
				case "id":
					eid = attr.Val
					continue
				case "points":
					// Go ahead!
				default:
					continue
				}

				// Split into string ints of multiple points' coordinates
				coords := strings.Split(attr.Val, " ")
				if len(coords)%2 != 0 {
					err = fmt.Errorf("points' length is invalid: %d", len(coords))
					return
				}

				for i := 0; i < len(coords); i += 2 {
					xs, ys := coords[i], coords[i+1]
					xf, yf := 0.0, 0.0

					xf, err = strconv.ParseFloat(xs, 64)
					if err != nil {
						err = fmt.Errorf("points has invalid position (failed to parse X): %s", xs)
						return
					}

					yf, err = strconv.ParseFloat(ys, 64)
					if err != nil {
						err = fmt.Errorf("points has invalid position (failed to parse Y): %s", ys)
						return
					}

					ps = append(ps, point{int(math.Round(xf)), int(math.Round(yf))})
				}

				var poly polygon
				poly.Name = eid
				for i := 0; i < len(ps)-1; i++ {
					poly.S = append(poly.S, stroke{ps[i].X, ps[i].Y, ps[i+1].X, ps[i+1].Y})
				}
				polygons = append(polygons, poly)
			}
		}
	}
	return
}

// Template parameters

type templateParamsPolygon struct {
	Alt           string
	Src           string
	Width, Height int
	Top, Left     int
}

type templateParams struct {
	CompName      string
	Width, Height int
	Polygons      []templateParamsPolygon
}

func main() {
	os.Exit(main_())
}

func main_() int {
	var noCut bool
	var generateTSX string

	flag.BoolVar(&noCut, "nocut", true, "Skip cutting image")
	flag.StringVar(&generateTSX, "generate", "", "Generate TSX and save it at the specified path")
	flag.Parse()

	if flag.NArg() != 1 {
		log.Errorf("Usage: %s INPUT_SVG", os.Args[0])
		return 1
	}

	fn := flag.Arg(0)
	log.Infof("reading SVG from: %s", fn)

	svg, err := os.ReadFile(fn)
	if err != nil {
		log.Errorf("failed to read the input file: %s", err)
		return 1
	}

	fnBase := filepath.Base(fn)
	fnName := strings.TrimSuffix(fnBase, filepath.Ext(fnBase))

	worker, _ := resvg.NewDefaultWorker(context.Background())
	defer worker.Close()

	log.Infof("parsing SVG")
	doc, err := html.Parse(bytes.NewReader(svg))
	if err != nil {
		log.Errorf("failed to parse SVG: %s", err)
		return 1
	}

	log.Infof("finding the 'cutline' layer")
	cutLineLayer := findCutLineLayer(doc)
	if cutLineLayer == nil {
		log.Errorf("Layer 'cutline' was not found")
		return 1
	}

	polygons, err := findPolygons(cutLineLayer)
	if err != nil {
		log.Errorf("failed to find polygons: %s", err)
		return 1
	}
	log.Infof("found %d rects/polygons", len(polygons))

	if !noCut {
		log.Infof("rendering SVG")
		pngBuf, err := worker.Render(svg)
		if err != nil {
			log.Errorf("failed to render SVG: %s", err)
			return 1
		}

		img, _, err := image.Decode(bytes.NewReader(pngBuf))
		if err != nil {
			log.Errorf("failed to decode rendered PNG: %s", err)
			return 1
		}

		save := func(wg *sync.WaitGroup, fn string, poly polygon) {
			defer wg.Done()

			frame := image.NewRGBA(image.Rect(0, 0, poly.Width(), poly.Height()))
			left, top := poly.Left(), poly.Top()
			bg := color.RGBA{R: 255, G: 255, B: 255, A: 0}

			for y := poly.Top(); y <= poly.Bottom(); y++ {
				for x := poly.Left(); x <= poly.Right(); x++ {
					if poly.Inside(x, y) {
						frame.Set(x-left, y-top, img.At(x, y))
					} else {
						frame.Set(x, y-top, bg)
					}
				}
			}

			var f *os.File
			if poly.Name == "" {
				f, err = os.Create(fmt.Sprintf("%s_%d_%d.png", fn, poly.Left(), poly.Top()))
			} else {
				f, err = os.Create(fmt.Sprintf("%s_%s.png", fn, poly.Name))
			}
			if err != nil {
				log.Errorf("failed to create %s: %s", fn, err)
			}
			defer f.Close()

			err = png.Encode(f, frame)
			if err != nil {
				log.Errorf("failed to encode a frame into PNG: %s", err)
			}
		}

		wg := sync.WaitGroup{}
		for i, poly := range polygons {
			if poly.Name == "" {
				log.Infof("processing and saving %02d/%02d", i+1, len(polygons))
			} else {
				log.Infof("processing and saving %02d/%02d (%s)", i+1, len(polygons), poly.Name)
			}
			wg.Add(1)
			go save(&wg, fnName, poly)
		}
		wg.Wait()
	}

	if generateTSX != "" {
		tsx, err := template.New("tsx").Parse(tsxTemplate)
		if err != nil {
			log.Errorf("failed to parse TSX template: %s", err)
			return 1
		}

		leftmost := slices.MinFunc(polygons, func(a, b polygon) int { return cmp.Compare(a.Left(), b.Left()) }).Left()
		rightmost := slices.MaxFunc(polygons, func(a, b polygon) int { return cmp.Compare(a.Right(), b.Right()) }).Right()
		topmost := slices.MinFunc(polygons, func(a, b polygon) int { return cmp.Compare(a.Top(), b.Top()) }).Top()
		bottommost := slices.MaxFunc(polygons, func(a, b polygon) int { return cmp.Compare(a.Bottom(), b.Bottom()) }).Bottom()

		var pargs []templateParamsPolygon
		for _, poly := range polygons {
			pargs = append(pargs, templateParamsPolygon{
				Alt:    poly.Name,
				Src:    fmt.Sprintf("%s_%s.png", fnName, poly.Name),
				Width:  poly.Width(),
				Height: poly.Height(),
				Top:    poly.Top(),
				Left:   poly.Left(),
			})
		}

		args := templateParams{
			CompName: "MangaFrames",
			Width:    rightmost - leftmost,
			Height:   bottommost - topmost,
			Polygons: pargs,
		}

		f, err := os.Create(generateTSX)
		if err != nil {
			log.Errorf("failed to create TSX output file: %s", err)
			return 1
		}

		if err = tsx.Execute(f, args); err != nil {
			log.Errorf("failed to render TSX output file: %s", err)
			return 1
		}
	}

	return 0
}
