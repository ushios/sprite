package sprite

import "testing"

func TestCalcLength(t *testing.T) {
	test := func(p1, p2 Point, e float64) {
		d := CalcLength(p1, p2)

		if int64(e) != d.Int() {
			t.Errorf("length between (%s) and (%s) expected (%d) but (%d)", p1, p2, int(e), d.Int())
		}
	}

	test(Point{X: 10, Y: 0}, Point{X: 20, Y: 0}, 10)
	test(Point{X: -10, Y: 0}, Point{X: 20, Y: 0}, 30)
}
