package sprite

// Line is line
type Line struct {
	Start, End Point
}

// Length return length
func (l Line) Length() Length {
	return CalcLength(l.Start, l.End)
}
