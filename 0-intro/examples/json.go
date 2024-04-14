// Demonstration of build-in JSON support in Go.
package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    Name   string   `json:"name"`
    ID     int      `json:",omitempty"`
    Attrs  []string `json:"attributes,omitempty"`
    Secret []byte   `json:"-"`
}

func main() {
    sdata := []byte("Super-secret data that should no go in the JSON string")
    
    data := User{
        Name: "steve", 
        ID: 42, 
        Attrs: []string{"foo","bar"}, 
        Secret: sdata,
    }

    encoded, err := json.Marshal(data)
    if err != nil {
        panic(err)
    }
    fmt.Println("JSON-marshalled data:", string(encoded))

    var rdata User
    err = json.Unmarshal(encoded, &rdata)
    if err != nil {
        panic(err)
    }
    fmt.Println("Structure after unmarshalling:", rdata)
}
