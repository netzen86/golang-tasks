package areacalc

import "fmt"

const pi = 3.14159

type Shape interface {
	Area() float64
	Type() string
}

// type ShapeType struct {
// 	Shape string
// }

// func (st ShapeType) Type() string {
// 	return st.Shape
// }

type Rectangle struct {
	Hypotenuse float64
	Height     float64
	Shape      string
}

func NewRectangle(hypotenuse float64, height float64, shape string) *Rectangle {
	return &Rectangle{Hypotenuse: hypotenuse, Height: height, Shape: shape}
}

func (ra Rectangle) Area() float64 {
	return ra.Hypotenuse * ra.Height
}

func (rt Rectangle) Type() string {
	return rt.Shape
}

type Circle struct {
	Radius float64
	Shape  string
}

func NewCircle(radius float64, shape string) *Circle {
	return &Circle{Radius: radius, Shape: shape}
}

func (r Circle) Area() float64 {
	return pi * (r.Radius * r.Radius)
}

func (ct Circle) Type() string {
	return ct.Shape
}

func AreaCalculator(figures []Shape) (string, float64) {
	var commonArea float64 = 0.0
	listOfShapes := ""
	for i, figure := range figures {
		commonArea += figure.Area()
		if i == 0 {
			listOfShapes = figure.Type()
			continue
		}
		listOfShapes = fmt.Sprintf("%s-%s", listOfShapes, figure.Type())
	}
	return listOfShapes, commonArea
}
