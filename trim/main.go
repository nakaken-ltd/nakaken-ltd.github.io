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

	"github.com/BurntSushi/toml"
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
	DPR                   int
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

// SVGClipPath generates a clip-path shape in `path()` notation.
// This notation is not recommended because the shape won't be scaled along with the <img>'s scale.
func (p polygon) SVGClipPath() string {
	switch len(p.S) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("M%d,%d", p.S[0].X0, p.S[0].Y0)
	default:
		ox, oy := p.Left(), p.Top()
		commands := []string{fmt.Sprintf("M%d,%d", p.S[0].X0-ox, p.S[0].Y0-oy)}
		for _, s := range p.S[1:] {
			commands = append(commands, fmt.Sprintf("L%d,%d", s.X0-ox, s.Y0-oy))
		}
		return strings.Join(append(commands, "Z"), "")
	}
}

// PolygonClipPath generates a clip-path shape in `polygon()` notation.
// This notation is recommended because the shape will be scaled along with the <img>'s scale.
func (p polygon) PolygonClipPath() string {
	if len(p.S) < 3 {
		return ""
	}

	ox, oy := float64(p.Left()), float64(p.Top())
	w, h := float64(p.Width()), float64(p.Height())

	var commands []string
	for _, s := range p.S {
		x, y := float64(s.X0)-ox, float64(s.Y0)-oy
		commands = append(commands, fmt.Sprintf("%.2f%%_%.2f%%", x/w*100, y/h*100))
	}
	return strings.Join(commands, ",")
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
				DPR:  1,
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
				v := strings.ReplaceAll(attr.Val, "\n", "")
				v = strings.ReplaceAll(v, "\t", "")
				v = strings.ReplaceAll(v, ",", " ")
				coords := strings.Split(strings.TrimSpace(v), " ")
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

				// Append the first point if it is not present
				if ps[0].X != ps[len(ps)-1].X || ps[0].Y != ps[len(ps)-1].Y {
					ps = append(ps, ps[0])
				}

				var poly polygon
				poly.Name = eid
				poly.DPR = 1
				for i := 0; i < len(ps)-1; i++ {
					poly.S = append(poly.S, stroke{ps[i].X, ps[i].Y, ps[i+1].X, ps[i+1].Y})
				}
				polygons = append(polygons, poly)
			}
		}
	}
	return
}

func scalePolygons(polygons []polygon, dpr int) (newPolygons []polygon) {
	for i := range polygons {
		poly := polygon{
			Name:       polygons[i].Name,
			DPR:        dpr,
			OffsetTop:  polygons[i].OffsetTop * dpr,
			OffsetLeft: polygons[i].OffsetLeft * dpr,
		}

		for j := range polygons[i].S {
			poly.S = append(poly.S, stroke{
				X0: polygons[i].S[j].X0 * dpr,
				Y0: polygons[i].S[j].Y0 * dpr,
				X1: polygons[i].S[j].X1 * dpr,
				Y1: polygons[i].S[j].Y1 * dpr,
			})
		}

		newPolygons = append(newPolygons, poly)
	}

	return
}

// Template parameters

type templateParamsPolygon struct {
	Link          string
	Src           string
	Alt           string
	Width, Height int
	Top, Left     int
	SVGCP         string
	PolyCP        string
}

type templateParams struct {
	CompName      string
	Width, Height int
	MaxWidth      int
	Polygons      []templateParamsPolygon
}

type trimConfig struct {
	LinksHref map[string]string `toml:"links_href"`
	Alts      map[string]string `toml:"alts"`
}

func main() {
	os.Exit(main_())
}

func main_() int {
	cutFrames := flag.String("cut", "", "Cut image into the frames and save it in the specified path")
	generateTSX := flag.String("generate", "", "Generate TSX and save it at the specified path")
	dpr := flag.Int("dpr", 1, "device pixel ratio")
	maxWidth := flag.Int("maxwidth", 1024, "Max height width")
	lrBlank := flag.Int("lrblank", 1, "left blank and right blank")
	configFn := flag.String("config", "", "Configuration file")
	flag.Parse()

	config := trimConfig{}
	if *configFn != "" {
		_, err := toml.DecodeFile(*configFn, &config)
		if err != nil {
			log.Errorf("failed to decode config file: %s", err)
			return 1
		}
	}

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

	if *cutFrames != "" {
		fnPath, err := filepath.EvalSymlinks(*cutFrames)
		if err != nil {
			log.Errorf("failed to evaluate symlinks (invalid output path?): %s", err)
			return 1
		}

		log.Infof("rendering SVG")
		tree, err := worker.NewTreeFromData(svg, &resvg.Options{})
		if err != nil {
			log.Errorf("failed to generate a new tree from data: %s", err)
			return 1
		}
		defer tree.Close()

		width, height, err := tree.GetSize()
		if err != nil {
			log.Errorf("failed to get the dimension: %s", err)
			return 1
		}
		log.Infof("width: %.1f, height: %.1f", width, height)

		if *dpr > 1 {
			log.Infof("scaling %dx", *dpr)
		}

		pm, err := worker.NewPixmap(uint32(width*float32(*dpr)), uint32(height*float32(*dpr)))
		if err != nil {
			log.Errorf("failed to generate a new pixmap: %s", err)
			return 1
		}
		defer pm.Close()

		err = tree.Render(resvg.TransformFromScale(float32(*dpr), float32(*dpr)), pm)
		if err != nil {
			log.Errorf("failed to render SVG: %s", err)
			return 1
		}

		pngBuf, err := pm.EncodePNG()
		if err != nil {
			log.Errorf("failed to encode PNG: %s", err)
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
				f, err = os.Create(fmt.Sprintf("%s_%d_%d_%dx.png", fn, poly.Left()/poly.DPR, poly.Top()/poly.DPR, poly.DPR))
			} else {
				f, err = os.Create(fmt.Sprintf("%s_%s_%dx.png", fn, poly.Name, poly.DPR))
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

		scaledPolygons := scalePolygons(polygons, *dpr)

		wg := sync.WaitGroup{}
		for i, poly := range scaledPolygons {
			if poly.Name == "" {
				log.Infof("processing and saving %02d/%02d", i+1, len(scaledPolygons))
			} else {
				log.Infof("processing and saving %02d/%02d (%s)", i+1, len(scaledPolygons), poly.Name)
			}
			wg.Add(1)
			go save(&wg, filepath.Join(fnPath, fnName), poly)
		}
		wg.Wait()
	}

	if *generateTSX != "" {
		tsx, err := template.New("tsx").Parse(tsxTemplate)
		if err != nil {
			log.Errorf("failed to parse TSX template: %s", err)
			return 1
		}

		leftmost := slices.MinFunc(polygons, func(a, b polygon) int { return cmp.Compare(a.Left(), b.Left()) }).Left()
		rightmost := slices.MaxFunc(polygons, func(a, b polygon) int { return cmp.Compare(a.Right(), b.Right()) }).Right()
		topmost := slices.MinFunc(polygons, func(a, b polygon) int { return cmp.Compare(a.Top(), b.Top()) }).Top()
		bottommost := slices.MaxFunc(polygons, func(a, b polygon) int { return cmp.Compare(a.Bottom(), b.Bottom()) }).Bottom()

		fnName = "top"

		var pargs []templateParamsPolygon
		for _, poly := range polygons {
			link := "/" + poly.Name
			if l, ok := config.LinksHref[poly.Name]; ok {
				link = l
			}

			alt := poly.Name
			if a, ok := config.Alts[poly.Name]; ok {
				alt = a
			}

			pargs = append(pargs, templateParamsPolygon{
				Link:   link,
				Src:    fmt.Sprintf("%s_%s.avif", fnName, poly.Name),
				Alt:    alt,
				Width:  poly.Width() - leftmost + *lrBlank*2, // Add leftmost (left blank width) to centerize the frames
				Height: poly.Height(),
				Top:    poly.Top() - topmost, // Subtract topmost to remove the top blank
				Left:   poly.Left() - leftmost + *lrBlank,
				SVGCP:  poly.SVGClipPath(),
				PolyCP: poly.PolygonClipPath(),
			})
		}

		args := templateParams{
			CompName: "MangaFrames",
			Width:    rightmost - leftmost + *lrBlank*2, // Add leftmost (left blank width) to centerize the frames
			Height:   bottommost - topmost,              // Subtract topmost to remove the top blank
			MaxWidth: *maxWidth,
			Polygons: pargs,
		}

		f, err := os.Create(*generateTSX)
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
