package main

import (
	"fmt"

	"boschrexroth.com/can-demo-golang/pkg/can"
)

func main() {
	fmt.Println("Hello world")
	canDevice := can.NewCanDevice("can1")

	frame := canDevice.CanRecv()
	frame = can.InveretEndianness(frame)
	canDevice.CanSend(frame)
}
