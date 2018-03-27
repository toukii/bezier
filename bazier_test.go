package bezier

import (
	"testing"
)

func TestCenter(t *testing.T) {
	p1 := NewPoint(10, 15)
	p2 := NewPoint(20, 22)
	center := p1.Center(p2)
	t.Logf("Center:%+v", center)
	dlt := p1.Dlt(p2)
	t.Logf("Dlt:%+v", dlt)
}

func TestCtl(t *testing.T) {
	p1 := NewPoint(10, 14)
	p2 := NewPoint(10, 6)
	xy := NewPoint(0, 10)
	dlt := p1.Dlt(p2)

	ps := xy.CtlPoints(dlt)
	t.Logf("ps:%+v, %+v", ps[0], ps[1])
}

func TestCtl2(t *testing.T) {
	p1 := NewPoint(10, 10)
	p2 := NewPoint(6, 6)
	xy := NewPoint(6, 10)
	dlt := p1.Dlt(p2)

	ps := xy.CtlPoints(dlt)
	t.Logf("ps:%+v, %+v", ps[0], ps[1])
}

func TestTrh(t *testing.T) {
	ps := []*Point{
		NewPoint(111, 211),
		NewPoint(222, 122),
		NewPoint(333, 433),
	}

	trh := Trh(ps, false, false)
	t.Logf("trh: %+s", trh)

	ps = []*Point{
		NewPoint(222, 122),
		NewPoint(333, 433),
		NewPoint(444, 344),
	}

	trh = Trh(ps, false, false)
	t.Logf("trh: %+s", trh)

	ps = []*Point{
		NewPoint(333, 433),
		NewPoint(444, 344),
		NewPoint(555, 655),
	}

	trh = Trh(ps, false, false)
	t.Logf("trh: %+s", trh)
}

func TestTrhs(t *testing.T) {
	trhs := Trhs(
		NewPoint(111, 211),
		NewPoint(222, 122),
		NewPoint(333, 433),
		NewPoint(444, 344),
		NewPoint(555, 655),
		NewPoint(666, 566),
		NewPoint(777, 877),
		NewPoint(888, 788),
		NewPoint(999, 999),
	)
	t.Logf("trhs: %+s", trhs)
}
