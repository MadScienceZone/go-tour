// Demonstration of maps in Go
package main

import "fmt"

func main() {
	// a map of names to ages (string to integer)
	// creation of map variable by assigning a map literal via :=
	Ages := map[string]int{
		"Alice": 14,
		"Bob": 22,
		"Charlie": 27,
		"Daria": 42,
	}
	// Println can print maps natively
	fmt.Println(Ages)

	// Iterate over (key, value) pairs using the range operator in a for loop
	for name, age := range Ages {
		if age >= 18 {
			fmt.Printf("%s may vote.\n", name)
		} else {
			fmt.Printf("%s is not eligible.\n", name)
		}
	}

	// Retrieve values by indexing the map with a key
	aliceAge := Ages["Alice"]
	fmt.Println("The value for key \"Alice\" is", aliceAge)

	// Non-existent keys return the "zero" value
	eveAge := Ages["Eve"]
	fmt.Println("There isn't a value for key \"Eve\" so we get", eveAge)

	// We can test for existence by using the two-return-value form of value retrieval from the map
	aliceAge, aliceExists := Ages["Alice"]
    eveAge, eveExists := Ages["Eve"]
	fmt.Println("Alice: age:", aliceAge, "exists:", aliceExists)
	fmt.Println("Eve: age:", eveAge, "exists:", eveExists)

	// We can change a value or add new (key, value) pairs by assigning into the map
	Ages["Eve"] = 20
	fmt.Println("Now Eve is", Ages["Eve"])

	// We can remove (key, value) pairs using the delete function
	delete(Ages, "Bob")

	showAge(Ages, "Alice")
	showAge(Ages, "Bob")
}

func showAge(ages map[string]int, name string) {
	// If we just want to know if a key exists we can discard the value
	if _, exists := ages[name]; exists {
		fmt.Println("We do know about", name)
	}

	if age, exists := ages[name]; exists {
		fmt.Printf("We know %s's age is %d.\n", name, age)
	} else {
		fmt.Println("We don't know", name)
	}
}
