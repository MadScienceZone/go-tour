// Demonstration of thread-safe ID generation using a channel and goroutine
package main

import (
    "fmt"
    "sync"
)

// an internal "service" that provides an infinite stream
// of incrementing ID numbers
func serveMessageIDs(c chan<- int) {
    var id int
    for {
        c <- id
        id++
    }
}

func main() {
    var wg sync.WaitGroup

    IDSource := make(chan int)
    go serveMessageIDs(IDSource)

    for i := 0; i < 100; i++ {
        id := i  // needed prior to Go 1.22
                 // from 1.22 onward we can just use i directly
        wg.Add(1)
        go func() {
            defer wg.Done()
            fmt.Printf("Goroutine #%d, ID=%d\n", id, <-IDSource)
        }()
    }

    // Wait for all goroutines to finish
    wg.Wait()
}

