/*
###The MIT License

###Copyright 2020 Bosch Rexroth AG

###Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

###The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

###THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

//Configure variables for running the code as an application the RCU. These comments shall be removed before packaging the application in snap
var DATA_FOLDER = os.Getenv("SNAP_DATA")
var RCU_HOST = "127.0.0.1:8883"

//configure variables for running the code as an application on your computer. The next 2 lines of code shall be commented before packaging the application in snap
//var DATA_FOLDER = "."
//var RCU_HOST = "192.168.188.20:8883" //IP Address of the RCU

//topic for publishing a message that shall be sent through can2
var PUB_TOPIC = "hardware/can/channel/can2/tx"

//message be sent through can2
var PUB_PAYLOAD = "(1586260384.286194) can2 18FE563D#8877665544332211"

//topic for receiving messages that are coming through can2
var SUB_TOPIC_CAN_DATA = "hardware/can/channel/can1/rx"

//topic for receiving can status messages
var SUB_TOPIC_CAN_STATUS = "hardware/can/status"

//connect the client to the mosquitto broker
func connect(clientId string) mqtt.Client {
	opts := createClientOptions(clientId)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

//MQTT Connection options
func createClientOptions(clientId string) *mqtt.ClientOptions {
	cer, _ := tls.LoadX509KeyPair(DATA_FOLDER+"/certs/can.crt", DATA_FOLDER+"/certs/can.key")
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cer}}
	tlsConfig.InsecureSkipVerify = true

	opts := mqtt.NewClientOptions()
	opts.SetTLSConfig(tlsConfig).AddBroker(fmt.Sprintf("tls://%s", RCU_HOST))
	opts.SetClientID(clientId)
	return opts
}

//handle incoming messages
func listen() {
	client := connect("sub")
	client.Subscribe(SUB_TOPIC_CAN_DATA, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})
	client.Subscribe(SUB_TOPIC_CAN_STATUS, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})
}

func main() {

	go listen()

	//Publish a message on the broker every 3 seconds
	client := connect("pub")
	timer := time.NewTicker(3 * time.Second)
	for range timer.C {
		fmt.Println("Publishing message")
		client.Publish(PUB_TOPIC, 0, false, PUB_PAYLOAD)
	}
}
