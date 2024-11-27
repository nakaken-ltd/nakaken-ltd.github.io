package main

import (
	"bytes"
	"cmp"
	_ "embed"
	"reflect"
	"slices"
	"testing"

	"golang.org/x/net/html"
)

func TestPolygon(t *testing.T) {
	// p represents the following polygon:
	// (2,  3)-----(9,   0)
	//   |            |
	//   |            |
	//   |            |
	// (1, 11)-----(10, 10)
	p := polygon{
		S: []stroke{
			{2, 3, 1, 11},
			{1, 11, 10, 10},
			{10, 10, 9, 0},
			{9, 0, 2, 3},
		},
	}

	// Boundary test

	if p.Left() != 1 {
		t.Errorf("left: expected 1, got %d", p.Left())
	}

	if p.Right() != 10 {
		t.Errorf("right: expected 10, got %d", p.Right())
	}

	if p.Top() != 0 {
		t.Errorf("top: expected 0, got %d", p.Top())
	}

	if p.Bottom() != 11 {
		t.Errorf("bottom: expected 11, got %d", p.Bottom())
	}

	if p.Width() != 9 {
		t.Errorf("width: expected 9, got %d", p.Width())
	}

	if p.Height() != 11 {
		t.Errorf("height: expected 11, got %d", p.Height())
	}

	// Inside/outside calculation test

	ss := []struct {
		X, Y   int
		Inside bool
	}{
		{2, 2, false},
		{2, 4, true},
		{2, 10, true},
		{2, 12, false},
		{3, 2, false},
		{3, 3, true},
		{3, 10, true},
		{3, 11, false},
		{4, 2, false},
		{4, 3, true},
		{4, 10, true},
		{4, 11, false},
		{5, 1, false},
		{5, 2, true},
		{5, 10, true},
		{5, 11, false},
		{6, 1, false},
		{6, 2, true},
		{6, 10, true},
		{6, 11, false},
		{7, 0, false},
		{7, 1, true},
		{7, 10, true},
		{7, 11, false},
		{8, 0, false},
		{8, 1, true},
		{8, 10, true},
		{8, 11, false},
		{9, -1, false},
		{9, 1, true},
		{9, 10, true},
		{9, 11, false},
	}

	for _, s := range ss {
		actual := p.Inside(s.X, s.Y)
		if actual != s.Inside {
			t.Errorf("inside (%d, %d): expected %v, got %v", s.X, s.Y, s.Inside, actual)
		}
	}

	// Convert to SVG test

	ss2 := []struct {
		P      polygon
		Expect string
	}{
		{
			P:      polygon{},
			Expect: "",
		},
		{
			P:      polygon{S: []stroke{{3, 9, 4, 10}}},
			Expect: "M3,9",
		},
		{
			P:      p,
			Expect: "M1,3L0,11L9,10L8,0Z",
		},
	}

	for _, s := range ss2 {
		actual := s.P.SVGClipPath()
		if actual != s.Expect {
			t.Errorf("SVG: expected %s, got %s", s.Expect, actual)
		}
	}

	// Convert to polygon text

	expect := "11.11%_27.27%,0.00%_100.00%,100.00%_90.91%,88.89%_0.00%"
	actual := p.PolygonClipPath()
	if actual != expect {
		t.Errorf("Polygon: expected %s, got %s", expect, actual)
	}
}

//go:embed test.svg
var svgStr []byte

func TestSVGParser(t *testing.T) {
	doc, err := html.Parse(bytes.NewReader(svgStr))
	if err != nil {
		t.Fatal(err)
	}

	cut := findCutLineLayer(doc)
	if cut == nil {
		t.Fatalf("cut line is nil")
	}

	actualPoly, err := findPolygons(cut)
	if err != nil {
		t.Fatalf("failed finding/converting polygons: %s", err)
	} else if len(actualPoly) != 7 {
		t.Errorf("incorrect polygons length: expected 7, got %d", len(actualPoly))
	}

	actualPoly = slices.SortedFunc(
		slices.Values(actualPoly),
		func(p1 polygon, p2 polygon) int {
			return cmp.Compare(p1.Name, p2.Name)
		},
	)

	expectedPoly := []polygon{
		{
			Name:       "c1-lt",
			OffsetTop:  0,
			OffsetLeft: 0,
			S: []stroke{
				stroke{X0: 0, Y0: 0, X1: 0, Y1: 543},
				stroke{X0: 0, Y0: 543, X1: 550, Y1: 543},
				stroke{X0: 550, Y0: 543, X1: 784, Y1: 0},
				stroke{X0: 784, Y0: 0, X1: 0, Y1: 0},
			},
		},
		{
			Name:       "c2-rt",
			OffsetTop:  0,
			OffsetLeft: 0,
			S: []stroke{
				{X0: 2560, Y0: 0, X1: 2560, Y1: 543},
				{X0: 2560, Y0: 543, X1: 2010, Y1: 543},
				{X0: 2010, Y0: 543, X1: 1776, Y1: 0},
				{X0: 1776, Y0: 0, X1: 2560, Y1: 0},
			},
		},
		{
			Name:       "c3-lb",
			OffsetTop:  0,
			OffsetLeft: 0,
			S: []stroke{
				{X0: 0, Y0: 1440, X1: 0, Y1: 897},
				{X0: 0, Y0: 897, X1: 550, Y1: 897},
				{X0: 550, Y0: 897, X1: 784, Y1: 1440},
				{X0: 784, Y0: 1440, X1: 0, Y1: 1440},
			},
		},
		{
			Name:       "c4-rb",
			OffsetTop:  0,
			OffsetLeft: 0,
			S: []stroke{
				{X0: 2560, Y0: 1440, X1: 2560, Y1: 897},
				{X0: 2560, Y0: 897, X1: 2010, Y1: 897},
				{X0: 2010, Y0: 897, X1: 1776, Y1: 1440},
				{X0: 1776, Y0: 1440, X1: 2560, Y1: 1440},
			},
		},
		{
			Name:       "c5-s",
			OffsetTop:  0,
			OffsetLeft: 0,
			S: []stroke{
				{X0: 1054, Y0: 186, X1: 1157, Y1: 329},
				{X0: 1157, Y0: 329, X1: 1325, Y1: 382},
				{X0: 1325, Y0: 382, X1: 1221, Y1: 524},
				{X0: 1221, Y0: 524, X1: 1222, Y1: 701},
				{X0: 1222, Y0: 701, X1: 1054, Y1: 645},
				{X0: 1054, Y0: 645, X1: 887, Y1: 701},
				{X0: 887, Y0: 701, X1: 888, Y1: 524},
				{X0: 888, Y0: 524, X1: 784, Y1: 382},
				{X0: 784, Y0: 382, X1: 951, Y1: 329},
				{X0: 951, Y0: 329, X1: 1054, Y1: 186},
			},
		},
		{
			Name:       "c6-p",
			OffsetTop:  0,
			OffsetLeft: 0,
			S: []stroke{
				{X0: 1592, Y0: 632, X1: 1300, Y1: 1110},
				{X0: 1300, Y0: 1110, X1: 1889, Y1: 942},
				{X0: 1889, Y0: 942, X1: 1984, Y1: 576},
				{X0: 1984, Y0: 576, X1: 1592, Y1: 632},
			},
		},
		{
			Name:       "c7-r",
			OffsetTop:  0,
			OffsetLeft: 0,
			S: []stroke{
				{X0: 708, Y0: 841, X1: 708, Y1: 1200},
				{X0: 708, Y0: 1200, X1: 1244, Y1: 1200},
				{X0: 1244, Y0: 1200, X1: 1244, Y1: 841},
				{X0: 1244, Y0: 841, X1: 708, Y1: 841},
			},
		},
	}

	if !reflect.DeepEqual(expectedPoly, actualPoly) {
		t.Errorf("unexpected polygons")
		t.Logf("expected: %#v", expectedPoly)
		t.Logf("got: %#v", actualPoly)
	}
}
