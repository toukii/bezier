package svg

import (
	// "encoding/base64"
	"text/template"

	"github.com/toukii/bezier"
	"github.com/toukii/bytes"
	"github.com/toukii/goutils"
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
}

func ExcutePath(color string, stroke string, points ...*bezier.Point) []byte {
	data := map[string]string{
		"Path":        goutils.ToString(bezier.Trhs(2, points...)),
		"Color":       color,
		"StrokeWidth": stroke,
	}
	return Excute(PathTpl, data)
}

func Excute(tpl *template.Template, item interface{}) []byte {
	wr := bytes.NewWriter(make([]byte, 1024))
	err := tpl.Execute(wr, item)
	if goutils.CheckErr(err) {
		return nil
	}
	return wr.Bytes()
}
