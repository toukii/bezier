package bezier_in_svg

import (
	"fmt"
	"math/rand"
	"testing"
	"text/template"
	"time"

	"github.com/toukii/bezier"
	"github.com/toukii/bytes"
	"github.com/toukii/goutils"
	"github.com/toukii/icat"
)

var (
	PolylineFmt = `<polyline points="{{range . }}{{.X}},{{.Y}} {{end}}" stroke="#33b5e5" stroke-width="1" fill="none"></polyline>`
	PathFmt     = `<path d="{{.Path}}" stroke="{{.Color}}" stroke-width="{{.StrokeWidth}}" fill="none"></path>`
	SvgFmt      = `<svg width="550" height="550" version="1.1" xmlns="http://www.w3.org/2000/svg">
{{.Polyline}}
{{.Path2}}
{{.Path1}}
{{.Path11}}
<path d="{{.Ctl1}}" stroke="purple" stroke-width="2" fill="none"></path>
<path d="{{.Ctl2}}" stroke="green" stroke-width="2" fill="none"></path>
</svg>`

	PolylineTpl *template.Template
	PathTpl     *template.Template
	SvgTpl      *template.Template
	rd          *rand.Rand
)

func init() {
	var err error

	PolylineTpl, err = template.New("Polyline").Parse(PolylineFmt)
	if err != nil {
		panic(err)
	}

	PathTpl, err = template.New("Path").Parse(PathFmt)
	if err != nil {
		panic(err)
	}

	SvgTpl, err = template.New("Svg").Parse(SvgFmt)
	if err != nil {
		panic(err)
	}

	rd = rand.New(rand.NewSource(time.Now().Unix()))
}

func Excute(tpl *template.Template, item interface{}) []byte {
	wr := bytes.NewWriter(make([]byte, 1024))
	err := tpl.Execute(wr, item)
	if goutils.CheckErr(err) {
		return nil
	}
	return wr.Bytes()
}

func MultiExcute(tpl *template.Template, item ...*bezier.Point) []byte {
	wr := bytes.NewWriter(make([]byte, 1024))
	err := tpl.Execute(wr, item)
	if goutils.CheckErr(err) {
		return nil
	}
	return wr.Bytes()
}

func randomPoints(n int) []*bezier.Point {
	points := make([]*bezier.Point, n)
	for i := 0; i < n; i++ {
		points[i] = bezier.NewPoint(rd.Intn(550), rd.Intn(550))
	}
	return points
}

func noSmoothPoints() []*bezier.Point {
	return []*bezier.Point{
		bezier.NewPoint(243, 216),
		bezier.NewPoint(217, 241),
		bezier.NewPoint(124, 512),
		bezier.NewPoint(502, 432),
		bezier.NewPoint(547, 476),
		bezier.NewPoint(309, 123),
		bezier.NewPoint(418, 161),
		bezier.NewPoint(542, 377),
	}
}

func samplePoints() []*bezier.Point {
	return []*bezier.Point{
		bezier.NewPoint(110, 105),
		bezier.NewPoint(220, 240),
		bezier.NewPoint(130, 250),
		bezier.NewPoint(180, 350),
		bezier.NewPoint(280, 450),
		bezier.NewPoint(480, 150),
		bezier.NewPoint(111, 211),
		bezier.NewPoint(222, 122),
		bezier.NewPoint(333, 433),
		bezier.NewPoint(444, 344),
		bezier.NewPoint(555, 655),
		bezier.NewPoint(666, 566),
		bezier.NewPoint(777, 877),
		bezier.NewPoint(888, 788),
		bezier.NewPoint(999, 999),
	}
}

func TestBezierSvg(t *testing.T) {
	points := randomPoints(8)
	// points := noSmoothPoints()
	// points := samplePoints()

	data1 := map[string]string{
		"Path":        goutils.ToString(bezier.Trhs(2, points...)),
		"Color":       "red",
		"StrokeWidth": "3",
	}
	fmt.Println()

	bezier.ShortenTh = 1.2
	data11 := map[string]string{
		"Path":        goutils.ToString(bezier.Trhs(2, points...)),
		"Color":       "purple",
		"StrokeWidth": "2",
	}
	fmt.Println()

	data2 := map[string]string{
		"Path":        goutils.ToString(bezier.Trhs(1, points...)),
		"Color":       "green",
		"StrokeWidth": "4",
	}
	// fmt.Println(data1)

	ctls := bezier.TrhCtls(points...)
	ctlSize := len(ctls) / 2
	ctlWr := bytes.NewWriter(make([]byte, 0, 1024))
	for i := 0; i < ctlSize; i++ {
		ctlWr.Write(bezier.PathTuple(ctls[i*2], ctls[i*2+1]))
	}

	ctls2 := bezier.TrhCtls(points...)
	ctlSize2 := len(ctls2) / 2
	ctlWr2 := bytes.NewWriter(make([]byte, 0, 1024))
	for i := 0; i < ctlSize2; i++ {
		ctlWr2.Write(bezier.PathTuple(ctls2[i*2], ctls2[i*2+1]))
	}

	path1 := Excute(PathTpl, data1)
	path11 := Excute(PathTpl, data11)
	path2 := Excute(PathTpl, data2)
	polyline := MultiExcute(PolylineTpl, points...)

	// path2 = nil

	svgData := map[string]string{
		"Path1":    goutils.ToString(path1),
		"Path11":   goutils.ToString(path11),
		"Path2":    goutils.ToString(path2),
		"Polyline": goutils.ToString(polyline),
		"Ctl1":     goutils.ToString(ctlWr.Bytes()),
		"Ctl2":     goutils.ToString(ctlWr2.Bytes()),
	}

	// fmt.Printf("svgData: %+v", svgData)

	svgOutput := Excute(SvgTpl, svgData)
	fmt.Printf("%s", svgOutput)
	icat.DisplaySVG(svgOutput)
}
