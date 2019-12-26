package svg

import (
	// "encoding/base64"
	"text/template"

	"github.com/toukii/bezier"
	// "github.com/toukii/bytes"
	log "github.com/sirupsen/logrus"
	"github.com/toukii/goutils"
	"github.com/toukii/icat"
)

var (
	svgtpl = `<svg width="{{.width}}" height="{{.height}}" version="1.1" xmlns="http://www.w3.org/2000/svg">
<path d="{{.PathContent}}" stroke="green" stroke-width="2" fill="none"></path>
</svg>`
)

func Display(width, height int, points ...*bezier.Point) error {
	pathTpl, err := template.New("Path").Parse(svgtpl)
	if err != nil {
		log.Errorf("err:%+v", err)
		return err
	}

	data := map[string]interface{}{
		"PathContent": goutils.ToString(bezier.Trhs(2, points...)),
		"width":       width,
		"height":      height,
	}
	bs := Excute(pathTpl, data)
	err = icat.DisplaySVG(bs)
	if err != nil {
		log.Errorf("err:%+v", err)
		return err
	}
	return nil
}
