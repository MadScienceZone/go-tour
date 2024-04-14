// Demonstration of composition and polymorphism for object-oriented
// Go programming
package main

import (
	"fmt"
	"math"
)

// Common attributes and methods for all shapes are
// represented in the BaseShape type
type BaseShape struct {
    X int
    Y int
}

func (s BaseShape) ReferencePoint() (int, int) {
	return s.X, s.Y
}

// A Triangle is a BaseShape with a base and height
type Triangle struct {
    BaseShape
    Base   int
    Height int
}

func (t Triangle) Area() float64 {
    return (float64(t.Base) * float64(t.Height)) / 2.0
}

// A Rectangle is a BaseShape with a width and height
type Rectangle struct {
    BaseShape
    Width  int
    Height int
}

func (r Rectangle) Area() float64 {
    return float64(r.Width * r.Height)
}

// A (regular) Polygon is defined by a number of sides, the side length and inscribed radius
type Polygon struct {
    BaseShape
    Sides  int
    Length float64  // Length of each side
    Radius float64  // Radius of inscribed circle
}

func (p Polygon) Area() float64 {
    return ((float64(p.Sides) / 2.0) * p.Length * p.Radius)
}

// A Circle just needs a radius to define it
type Circle struct {
    BaseShape
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * math.Pow(c.Radius, 2)
}

type Shape interface {
    Area() float64
	ReferencePoint() (int,int)
}

func main() {
	shapes := []Shape{
		Triangle{
			BaseShape: BaseShape{
				X: 3, 
				Y: 12,
			}, 
			Base: 3, 
			Height: 2,
		},
		Circle{
			BaseShape: BaseShape{
				Y: 22,
			}, 
			Radius: 1.5,
		},
		Rectangle{
			Height: 100, 
			Width: 50,
		},
	}

	for i, shape := range shapes {
		x, y := shape.ReferencePoint()
		fmt.Printf("#%d at (%d,%d), area=%f\n", i, x, y, shape.Area())
	}

	c1 := Circle{Radius: 10}
	reportArea(c1)
}

func reportArea(s Shape) {
    fmt.Printf("The area is %f\n", s.Area())
}
