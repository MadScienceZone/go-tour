// Demonstration of type assertions and type switches
package main

import "fmt"

func f(mystery any) {    // any == interface{}
    var v int

    // unsafe: if mystery isn't an int, the program will panic!
    // x := mystery.(int)
    //
    // instead, we should use the two-value version of the assertion.
    // this guarantees we'll get a valid value, but it will be the
    // "zero" value if the assertion failed, but we can distinguish
    // that case by looking at the second value returned.
    //
    x, ok := mystery.(int)
    v = x + 15
    if ok {
        fmt.Println("int mystery is", x)
        fmt.Println("result is", v)
    } else {
        fmt.Printf("The value passed was of type %T, so we just got the default %v.\n", mystery, x)
    }
}

func main() {
    fmt.Println("=== Type Assertions ===")
    f(42)
    f("hello")

    fmt.Println("=== Type Switch ===")
    g(42)
    g("hello")
    g(12.34)
    g([]string{"hello", "世界"})
}

func g(mystery any) {
    var v int
    
    switch x := mystery.(type) {
    case int:
        v = x + 15
        fmt.Println("int value", x, "+ 15 =", v)
    case string:
        fmt.Println("string", x)
    default:
        fmt.Printf("We don't know how to handle data \"%v\" of type %T.\n", mystery, mystery)
    }
}
