package shapes


type Rectangle struct {
	Figure
	Width  float64
	Height float64
}

func NewRectangle(width, height float64) Rectangle {
	return Rectangle{
		Figure: Figure{Color: "red"},
		Width:  width,
		Height: height,
	}
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}