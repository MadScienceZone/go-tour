// Demonstration of contexts in Go
package main

import (
    "context"
    "log"
    "time"
)

func doSomething(_ string) error {
    time.Sleep(500 * time.Millisecond)
    return nil
}

func collectData(ctx context.Context, 
                 stream <-chan string) error {
    for {
        select {
        case <-ctx.Done():
            log.Print("done")
            return nil
        case <-time.After(2 * time.Second):
            // warning: you aren't GUARANTEED this will happen if there's
            // always other cases that may proceed without blocking as well.
            log.Print("collectData taking too long")
        case data, ok := <-stream:   
            log.Print("process")
            if !ok { log.Print("closed"); return nil }
            if err := doSomething(data); err != nil {
                return err
            }
        }
    }
}

func main() {
    stream := make(chan string)

    // Feed an infinite data stream so our context will terminate
    // the reader.
    // 
    // Experiment: try changing this so the channel has a finite amount
    // of data written to it, then have the writer close the channel.
    // The reader should stop when the channel is closed.
    go func() {
        for {
            stream <- `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut 
            labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi 
            ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse 
            cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa 
            qui officia deserunt mollit anim id est laborum.`
            time.Sleep(2500 * time.Millisecond)
        }
    }()

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)

    // Experiment: try actually calling cancel() to stop the data collection early
    defer cancel()

    if err := collectData(ctx, stream); err != nil {
        log.Fatalf("collectData reported error %v\n", err)
    }
}

