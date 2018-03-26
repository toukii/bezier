package bezier

import (
	"fmt"
	"testing"
)

type Str struct {
	V string
}

func (s *Str) Key() string {
	return s.V
}

func TestCacheStr(t *testing.T) {
	cah := NewCache(3)
	fmt.Println(cah)
	cah.Put(&Str{"100"})
	fmt.Println(cah)
	cah.Put(&Str{"200"})
	fmt.Println(cah)
	cah.Put(&Str{"300"})
	fmt.Println(cah)
	cah.Put(&Str{"400"})
	fmt.Println(cah)
	cah.Put(&Str{"500"})
	fmt.Println(cah)
	cah.Put(&Str{"600"})
	fmt.Println(cah)
	cah.Put(&Str{"700"})
	fmt.Println(cah)
}

type Point struct {
	X, Y int
}

func (p *Point) Key() string {
	return fmt.Sprintf("%d-%d", p.X, p.Y)
}

func TestCachePoint(t *testing.T) {
	cah := NewCache(3)
	fmt.Println(cah)
	cah.Put(&Point{100, 100})
	fmt.Println(cah)
	cah.Put(&Point{200, 200})
	fmt.Println(cah)
	cah.Put(&Point{300, 300})
	fmt.Println(cah)
	cah.Put(&Point{400, 400})
	fmt.Println(cah)
	cah.Put(&Point{500, 500})
	fmt.Println(cah)
	cah.Put(&Point{600, 600})
	fmt.Println(cah)
	cah.Put(&Point{700, 700})
	fmt.Println(cah)
	fmt.Printf("%s: %+v\n", "500-500", cah.Get("500-500"))
}

func TestCacheGet(t *testing.T) {
	cah := NewCache(3)
	fmt.Printf("%s: %+v\n", "100-100", cah.Get("100-100"))
	fmt.Println(cah.Get("100-100"))
	cah.Put(&Point{100, 100})
	fmt.Printf("%s: %+v\n", "100-100", cah.Get("100-100"))
}
func BenchmarkCacheGet(b *testing.B) {
	cah := NewCache(3)
	cah.Put(&Str{"key"})
	for i := 0; i < b.N; i++ {
		cah.Get("key")
	}
}
