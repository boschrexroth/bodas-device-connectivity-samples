#! /usr/bin/env node

/*
###The MIT License

###Copyright 2020 Bosch Rexroth AG

###Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

###The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

###THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

const mqtt  = require('mqtt');
const fs = require('fs');


//Configure variables for running the code as an application the RCU. These comments shall be removed before packaging the application in snap
const DATA_FOLDER = process.env.SNAP_DATA
const RCU_HOST = "127.0.0.1"

//configure variables for running the code as an application on your computer. The next 2 lines of code shall be commented before packaging the application in snap
//const DATA_FOLDER = "."
//const RCU_HOST = "192.168.188.20"

//topic for publishing a message that shall be sent through can2
const PUB_TOPIC = "hardware/can/channel/can2/tx"

//message be sent through can2
const PUB_PAYLOAD = "(1586260384.286194) can2 18FE563D#8877665544332211"

//topic for receiving messages that are coming through can2
const SUB_TOPIC_CAN_DATA = "hardware/can/channel/can1/rx"

//topic for receiving can status messages
const SUB_TOPIC_CAN_STATUS = "hardware/can/status"

//MQTT Connection options
var connectOptions = {
  host: RCU_HOST,
  port: 8883,
  protocol: "mqtts",
  key: fs.readFileSync(DATA_FOLDER + "/certs/can.key"),
  cert: fs.readFileSync(DATA_FOLDER + "/certs/can.crt"),
  rejectUnauthorized: false,
};


//connect the client to the mosquitto broker
console.log("Sending connection request");
var client = mqtt.connect(connectOptions);

//handle incoming messages
client.on('message', function (topic, message, packet) {
  console.log("message is " + message);
  console.log("topic is " + topic);
});


//handle successfull connection
client.on("connect", function () {
  console.log("connected  " + client.connected);

  //subscribe to topics
  console.log("subscribing to topics");
  client.subscribe(SUB_TOPIC_CAN_DATA, { qos: 1 }); //single topic
  client.subscribe(SUB_TOPIC_CAN_STATUS, { qos: 1 }); //single topic
})


//handle errors
client.on("error", function (error) {
  console.log("Can't connect" + error);
  process.exit(1)
});


//Define an a method that publish a message on the broker every 3 seconds
setInterval(function() {
  console.log("Publishing a message");
  client.publish(PUB_TOPIC, PUB_PAYLOAD)
}, 3000);