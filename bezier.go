package bezier

import (
	"bytes"
	"fmt"
	"math"

	"github.com/toukii/goutils"
)

type IPoint interface {
	GetX() int
	GetY() int
}

type Point struct {
	X, Y int
}

func NewPoint(x, y int) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

func ParsePoint(p IPoint) *Point {
	return NewPoint(p.GetX(), p.GetY())
}

// manhattan metric
func (p1 *Point) MahatMetric(p2 *Point) float64 {
	var x, y int
	if p2 != nil {
		x, y = p2.X, p2.Y
	}
	return math.Abs(float64(p1.X-x)) + math.Abs(float64(p1.Y-y))
}

// euclidean metric
func (p1 *Point) EucMetric(p2 *Point) float64 {
	var x, y int
	if p2 != nil {
		x, y = p2.X, p2.Y
	}
	return math.Abs(float64(p1.X*p1.X-x*x)) + math.Abs(float64(p1.Y*p1.Y-y*y))
}

func (p *Point) Key() string {
	return fmt.Sprintf("%d-%d", p.X, p.Y)
}

func (p *Point) PathFmt() string {
	return fmt.Sprintf("%d %d", p.X, p.Y)
}

func (p1 *Point) Center(p2 *Point) *Point {
	return NewPoint((p1.X+p2.X)>>1, (p1.Y+p2.Y)>>1)
}

func (p1 *Point) Dlt(p2 *Point) *Point {
	return NewPoint((p1.X-p2.X)>>1, (p1.Y-p2.Y)>>1)
}

func (p *Point) Shorten(th float64) {
	p.X, p.Y = shorten(p.X, th), shorten(p.Y, th)
}

func shorten(i int, th float64) int {
	return int(float64(i) * th)
}

// 2 control points
func (xy *Point) CtlPoints(dlt *Point) [2]*Point {
	return [2]*Point{
		NewPoint(xy.X+dlt.X, xy.Y+dlt.Y),
		NewPoint(xy.X-dlt.X, xy.Y-dlt.Y),
	}
}

func (p *Point) Spilt() bool {
	return p.X == -1 && p.Y == -1
}

func Trhs(ps ...*Point) []byte {
	size := len(ps)
	buf := bytes.NewBuffer(make([]byte, 0, 2048))
	for i := 3; i <= size; i++ {
		trh := Trh(ps[i-3:i], i == 3, i == size)
		buf.Write(trh)
	}
	return buf.Bytes()
}

func Trh(ps []*Point, start, end bool) []byte {
	size := len(ps)
	if size > 3 {
		return Trhs(ps...)
	} else if size == 2 {
		return goutils.ToByte(fmt.Sprintf("M%sL", ps[0].PathFmt(), ps[1].PathFmt()))
	} else if size <= 1 {
		return nil
	}

	p1 := ps[0].Center(ps[1])                            // p1
	p2 := ps[1].Center(ps[2])                            // p2
	dlt := p1.Dlt(p2)                                    // dlt
	c12 := p1.Center(p2)                                 // center point of p1 and p2
	th_ := c12.MahatMetric(ps[1]) / dlt.MahatMetric(nil) // metric threshold
	th := th_
	if th > 0.8 {
		th = 1.0 / math.Pow(math.E, th_+0.2) // shorten
	}
	// fmt.Printf("%+v %+v\n", th_, th)
	dlt.Shorten(th) // shorten the dlt

	ctl := ps[1].CtlPoints(dlt) // reflect the 2 control points

	// start or end point
	var startP, endP *Point
	if start {
		startP = ps[0]
	} else {
		startP = p1
	}
	if end {
		endP = ps[2]
	} else {
		endP = p2
	}

	return goutils.ToByte(fmt.Sprintf("M%s C%s, %s, %s", startP.PathFmt(), ctl[0].PathFmt(), ctl[1].PathFmt(), endP.PathFmt()))
}
