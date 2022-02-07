package can

import (
	"context"
	"fmt"

	"go.einride.tech/can/pkg/socketcan"
)

func CanSend() {
	fmt.Println("here we want to send CAN messages")
}
func CanRecv() {
	fmt.Println("here we want to recieve CAN messages")
	channel := "can1"

	fmt.Println("HALLPOA")

	// Error handling omitted to keep example simple
	conn, _ := socketcan.DialContext(context.Background(), "can", channel)

	recv := socketcan.NewReceiver(conn)
	fmt.Printf("Start listening on %s", channel)
	for recv.Receive() {
		frame := recv.Frame()
		fmt.Println(frame.String())
		return
	}
}
