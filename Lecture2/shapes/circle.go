package shapes

import "math"

type Circle struct {
	Figure
	Radius float64
}

func NewCircle(radius float64) Circle {
	return Circle{
		Figure: Figure{Color: "blue"},
		Radius: radius,
	}
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}