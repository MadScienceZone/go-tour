// Demonstration of typed constants with iota

package main

import "fmt"

type MessageType byte

const (
	ServerCommand MessageType = iota
	ServerReply
	ServerError
	UrgentMessage
)

func main() {
	fmt.Println(ServerCommand, ServerReply, ServerError, UrgentMessage)
	fmt.Printf("Type of ServerCommand is %T\n", ServerCommand)
}

