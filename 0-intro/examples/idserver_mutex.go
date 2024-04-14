// Demonstration of thread-safe ID generation using a mutex
package main

import (
    "fmt"
    "sync"
)

type GameState struct {
    nextMessageID int
    lock          sync.Mutex
}

func (state *GameState) GetNextID() int {
    state.lock.Lock()
    defer state.lock.Unlock()

    state.nextMessageID++
    return state.nextMessageID
}

func main() {
    var gameServer GameState
    var wg sync.WaitGroup

    for i := 0; i < 100; i++ {
        id := i  // needed prior to Go 1.22
                 // from 1.22 onward we can just use i directly
        wg.Add(1)
        go func() {
            defer wg.Done()
            fmt.Printf("Goroutine #%d, ID=%d\n", id, gameServer.GetNextID())
        }()
    }

    // Wait for all goroutines to finish
    wg.Wait()
}

