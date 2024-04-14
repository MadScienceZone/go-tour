// A first look at goroutines
package main

import (
    "fmt"
    "time"
)

func countdown() {
    for i := 10; i >= 0; i-- {
        fmt.Printf(">>> %d <<<\n", i)
        time.Sleep(1 * time.Second)
    }
}

func main() {
    go countdown()
    fmt.Println("Starting a long-running task...")
    time.Sleep(15 * time.Second)
    fmt.Println("Done. Exiting.")
}
