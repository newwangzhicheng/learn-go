package perimeter

import "math"

//Shape interface
type Shape interface {
	Area() float64
}

//Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

//Area of Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

//Circle
type Circle struct {
	Radius float64
}

//Area of Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//Triangle
type Triangle struct {
	Base   float64
	Height float64
}

//Area of Triangle
func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

//Perimeter calculates perimeter
func Perimeter(rectangle Rectangle) float64 {
	return rectangle.Width*2 + rectangle.Height*2
}

//Area calculates area
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
