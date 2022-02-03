package main

import (
	"fmt"

	"boschrexroth.com/can-demo-golang/pkg/can"
)

func main() {
	fmt.Println("Hello world")
	can.CanSend()
	can.CanRecv()
}
