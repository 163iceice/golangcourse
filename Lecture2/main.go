// main.go
package main

import (
	"fmt"
	"Lecture2/shapes"
)

func main() {
	rect := shapes.NewRectangle(5, 4)
	circle := shapes.NewCircle(3)

	shapesList := []shapes.Shape{rect, circle}

	for _, shape := range shapesList {
		fmt.Printf("Type: %T\n", shape)
		fmt.Printf("Area: %.2f\n", shape.Area())
		fmt.Printf("Perimeter: %.2f\n\n", shape.Perimeter())
	}
}