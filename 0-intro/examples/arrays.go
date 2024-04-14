/* Demonstration of arrays and slices in Go */
package main

import (
	"fmt"
	"slices"
)

// creation of a fixed-length array using an array literal value and :=
// as well as retrieving elements and subsets of the array using slices
func arrays() {
	things := [5]string{
		"raindrops on roses",
		"whiskers on kittens",
		"copper kettles",
		"woolen mittens",
		"wild geese",
	}

	fmt.Println("I like", things[2])
	fmt.Println("I also like", things)
	fmt.Println("I know", len(things), "things.")
	fmt.Println("Some things:", things[1:3])
	fmt.Println("Some things:", things[:3])
	fmt.Println("Some things:", things[1:])
	fmt.Println("Some things:", things[:])
}

// creation of a slice by appending values to it
func slices1() {
	var things []string

	things = append(things, "doorbells")
	things = append(things, "sleighbells", "schnitzel")
	fmt.Println(len(things), things) 
}

// creation of a slice by assigning a slice literal value
// also shows deleting part of a slice
func slices2() {
	things := []string{
		"doorbells",
		"sleighbells",
		"schnitzel",
	}
	fmt.Println(things)

	primes := []int{2, 3, 5, 7, 11, 13}

	lowPrimes := slices.Delete(primes, 3, len(primes))
	fmt.Println(lowPrimes)
}

func main() {
	fmt.Println("=== Arrays ===")
	arrays()

	fmt.Println("=== Slices example 1 ===")
	slices1()

	fmt.Println("=== Slices example 2 ===")
	slices2()
}
