// Demonstration of Go structures and methods
package main

import "fmt"

type Triangle struct {
	Base   int
	Height int

	// Reference point
	X      int
	Y      int
}

func main() {

	var t1 Triangle
	t1.Base = 37
	t1.Height = 15
	t1.X = 11
	t1.Y = 22
	fmt.Println("Triangle t1 =", t1)

	var t2 Triangle = Triangle{Base: 3, Height: 1}
	fmt.Println("Triangle t2 =", t2)
	fmt.Println("t2's base is", t2.Base)

	t3 := Triangle{
		Base:   100,
		Height: 42,
		X:      -3,
		Y:      14,
	}
	fmt.Println("Triangle t3 =", t3)

	// Using functions
	fmt.Printf("T1 starting location (%d,%d), area %v\n", t1.X, t1.Y, Area(t1))
	t1 = Translate(t1, -2,+3)
	fmt.Printf("T1 ending location (%d,%d), area %v\n", t1.X, t1.Y, Area(t1))

	// Using functions with pointer to triangle
	TranslatePtr(&t1, -1, -1)
	fmt.Printf("T1 ending location (ptrs) (%d,%d), area %v\n", t1.X, t1.Y, Area(t1))

	// Using methods
	fmt.Printf("T2 starting location (%d,%d), area %v\n", t2.X, t2.Y, t2.Area())
	t2.Translate(-5,+13)
	fmt.Printf("T2 ending location (%d,%d), area %v\n", t2.X, t2.Y, t2.Area())
}

func Area(t Triangle) float64 {
    return (float64(t.Base) * float64(t.Height)) / 2.0
}

func Translate(t Triangle, dx, dy int) Triangle {
    t.X += dx
    t.Y += dy
    return t
}

func TranslatePtr(t *Triangle, dx, dy int) {
    t.X += dx
    t.Y += dy
}

func (t Triangle) Area() float64 {
    return (float64(t.Base) * float64(t.Height)) / 2.0
}

func (t *Triangle) Translate(dx, dy int) {
    t.X += dx
    t.Y += dy
}
