package bezier_in_svg

import (
	// "fmt"
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
	PolylineFmt = `<polyline points="{{range . }}{{.X}},{{.Y}} {{end}}" stroke="yellow" stroke-width="8" fill="none"></polyline>`
	PathFmt     = `<path d="{{.Path}}" stroke="red" stroke-width="3" fill="none"></path>`
	SvgFmt      = `<svg width="350" height="350" version="1.1" xmlns="http://www.w3.org/2000/svg">
{{.Polyline}}
{{.Path}}
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
		points[i] = bezier.NewPoint(rd.Intn(350), rd.Intn(350))
	}
	return points
}

func TestBezierSvg(t *testing.T) {
	points := randomPoints(8)

	data := map[string]string{
		"Path": goutils.ToString(bezier.Trhs(points...)),
	}
	// fmt.Println(data)

	path := Excute(PathTpl, data)
	polyline := MultiExcute(PolylineTpl, points...)

	svgData := map[string]string{
		"Path":     goutils.ToString(path),
		"Polyline": goutils.ToString(polyline),
	}

	// fmt.Printf("svgData: %+v", svgData)

	svgOutput := Excute(SvgTpl, svgData)
	// fmt.Printf("svg: %s", svgOutput)
	icat.DisplaySVG(svgOutput)
}
