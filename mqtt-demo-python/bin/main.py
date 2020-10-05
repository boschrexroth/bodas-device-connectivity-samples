#!/usr/bin/env python

###The MIT License

###Copyright 2020 Bosch Rexroth AG

###Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

###The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

###THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

import time
import paho.mqtt.client as paho
import ssl
import threading
import os

#Configure variables for running the code as an application the RCU. These comments shall be removed before packaging the application in snap
DATA_FOLDER = os.environ.get('SNAP_DATA')
RCU_HOST = "127.0.0.1"

#configure variables for running the code as an application on your computer. The next 2 lines of code shall be commented before packaging the application in snap
#DATA_FOLDER = ".."
#RCU_HOST = "192.168.188.20" #IP Address of the RCU

#topic for publishing a message that shall be sent through can2
PUB_TOPIC = "hardware/can/channel/can2/tx"

#message be sent through can2
PUB_PAYLOAD = "(1586260384.286194) can2 18FE563D#8877665544332211"

#topic for receiving messages that are coming through can2
SUB_TOPIC_CAN_DATA = "hardware/can/channel/can1/rx"

#topic for receiving can status messages
SUB_TOPIC_CAN_STATUS = "hardware/can/status"

#Callbacks when a message is received
def on_message(client, userdata, message):
  print("received message =",str(message.payload.decode("utf-8")))

#Callbacks for logging
def on_log(client, userdata, level, buf):
  print("log: ",buf)

#Callbacks when a the client is connected to the MQTT Broker
def on_connect(client, userdata, flags, rc):
  print("Connected ")
  client.subscribe([(SUB_TOPIC_CAN_STATUS, 1), (SUB_TOPIC_CAN_DATA, 1)])

#Client initialization
client=paho.Client() 
client.on_message=on_message
client.on_log=on_log
client.on_connect=on_connect

#Setup the certificates that shall be used for the communication
ssl_context = ssl.create_default_context()
ssl_context.load_verify_locations(cafile=DATA_FOLDER + "/certs/boschrexroth-pjiot.crt")
ssl_context.load_cert_chain(certfile=DATA_FOLDER + "/certs/can.crt", keyfile=DATA_FOLDER + "/certs/can.key")
client.tls_set_context(ssl_context)
client.tls_insecure_set(True)

#Connect to the broker
print("connecting to broker")
client.connect(RCU_HOST, 8883)

#start loop to process received messages
client.loop_start()

#Define an a method that publish a message on the broker every 3 seconds
def publish():
  threading.Timer(3.0, publish).start()
  client.publish(topic=PUB_TOPIC, payload=PUB_PAYLOAD, qos=1)
publish()
