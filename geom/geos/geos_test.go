package geos

import "testing"

func TestFoo(t *testing.T) {
	_ = NewGeos()
}

func BenchmarkWKB(b *testing.B) {
	g := NewGeos()
	defer g.Finish()

	for i := 0; i < b.N; i++ {
		cs, _ := g.CreateCoordSeq(10, 2)
		for i := 0; i < 10; i++ {
			cs.SetXY(g, uint32(i), float64(i*10), 0)
		}
		geom, _ := cs.AsLineString(g)
		if g.IsValid(geom) == false {
			b.Fail()
		}
		g.AsWkb(geom)
		g.Destroy(geom)
	}
}
