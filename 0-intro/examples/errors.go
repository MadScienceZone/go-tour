// Example showing the checking of error return values
package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    var intval int
    var err    error

    for i, arg := range os.Args[1:] {
        intval, err = strconv.Atoi(arg)
        if err != nil {
            fmt.Printf("Arg #%d (\"%s\"): %v.\n",
                       i, arg, err)
        } else {
            fmt.Printf("Arg #%d == %d\n", i, intval)
        }
    }
}
