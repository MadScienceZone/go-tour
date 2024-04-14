// First look at channels
package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan byte)
    go func(c chan byte) {
        x := <-c
        fmt.Println("Read", x, "from channel")
    }(ch)

    fmt.Println("Writing to channel")
    ch <- 42

    // Don't do this in a real program, but for this simple example,
	// we delay a little because the main routine may end here before
	// the goroutine has a chance to print its output.
	time.Sleep(1 * time.Second)   
}
