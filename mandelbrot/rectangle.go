package main

type Rectangle struct {
	Left   float64
	Right  float64
	Top    float64
	Bottom float64
}

func (r *Rectangle) FromRelative(p *Point) Point {
	var x = r.Left + (r.Right-r.Left)*p.X
	var y = r.Bottom + (r.Top-r.Bottom)*p.Y

	return Point{x, y}
}

func NewRectangle(center *Point, ratio float64, width float64) *Rectangle {
	height := width / ratio
	left := center.X - width/2
	right := center.X + width/2
	bottom := center.Y - height/2
	top := center.Y + height/2

	return &Rectangle{
		Left:   left,
		Right:  right,
		Bottom: bottom,
		Top:    top,
	}
}
