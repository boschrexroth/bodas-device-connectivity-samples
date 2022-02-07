package can

import (
	"context"
	"fmt"

	"go.einride.tech/can"
	"go.einride.tech/can/pkg/socketcan"
)

type CanDevice struct {
	channel string
}

func NewCanDevice(channel string) CanDevice {
	return CanDevice{channel: channel}
}

func (c CanDevice) CanSend(frame can.Frame) {
	conn, _ := socketcan.DialContext(context.Background(), "can", c.channel)

	tx := socketcan.NewTransmitter(conn)
	_ = tx.TransmitFrame(context.Background(), frame)
	fmt.Printf("Sent CAN Frame: \t\t%s\n", frame.String())
}

func (c CanDevice) CanRecv() can.Frame {
	conn, _ := socketcan.DialContext(context.Background(), "can", c.channel)
	recv := socketcan.NewReceiver(conn)

	recv.Receive()
	frame := recv.Frame()
	fmt.Printf("Received CAN Frame: \t\t%s\n", frame.String())

	return frame
}

func InveretEndianness(frame can.Frame) can.Frame {
	fmt.Printf("Original CAN Frame: \t\t%s\n", frame.String())
	be := frame.Data.PackBigEndian()
	frame.Data.UnpackLittleEndian(be)
	fmt.Printf("Inverted Endianness: \t\t%s\n", frame.String())
	return frame
}
