/*
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces

What we learn:
- Declaring structs to create your own data types which lets you bundle related data together and make the intent of your code clearer
- Declaring interfaces so you can define functions that can be used by different types (parametric polymorphism)
- Adding methods so you can add functionality to your data types and so you can implement interfaces
- Table based tests to make your assertions clearer and your suites easier to extend & maintain
*/

package structsmethodsinterfaces

import "math"

//Shape define the methods that all shapes should have
type Shape interface {
	Area() float64
}

//Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

//Circle struct
type Circle struct {
	Radius float64
}

//Triangle struct
type Triangle struct {
	Base   float64
	Height float64
}

//Perimeter calcs the perimeter based in width and height passed by parameter
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

//Area calcs the area of a Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

//Area calcs the area of a Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//Area calcs the area of a Triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
