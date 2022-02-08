// Copyright (c) 2022 Bosch Rexroth AG
// All rights reserved. See LICENSE file for details.
package main

import (
	"fmt"
	"time"

	"boschrexroth.com/can-demo-golang/pkg/can"
)

func main() {
	fmt.Println("Go Snap Sample started...")
	canDevice := can.NewCanDevice("can1")

	for {
		fmt.Println("---")
		frame := canDevice.CanRecv()
		frame = can.InvertEndianness(frame)
		canDevice.CanSend(frame)

		time.Sleep(2 * time.Second)
	}
}
