<!--
*** Thanks for checking out this README Template. If you have a suggestion that would
*** make this better, please fork the repo and create a pull request or simply open
*** an issue with the tag "enhancement".
*** Thanks again! Now go create something AMAZING! :D
-->





<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->

<!-- PROJECT LOGO -->
<br />
<p align="center">
  
  <h3 align="center">BODAS Connect Samples Code</h3>

</p>



<!-- TABLE OF CONTENTS -->
## Table of Contents

* [Introduction](#introduction)
* [Getting Started](#getting-started)
  * [Prerequisites](#prerequisites)
  * [Development Infrastructure](#development-infrastructure)
  * [Connecting your computer with the RCU](#connecting-your-computer-with-the-rcu)
    * [Serial connection](#serial-connection)
    * [Windows Hyperterminal](#windows-hyperterminal)
    * [Linux Minicom](#linux-minicom)
    * [Ethernet connections](#ethernet-connections)
* [Developing Snaps for RCUs](#developing-snaps-for-rcus)
  * [Development and Testing](#development-and-testing)
  * [Building and Packaging](#building-and-packaging)
  * [Installation and Testing](#installattion-and-testing)
* [License](#license)
* [Contact](#contact)
* [Acknowledgements](#acknowledgements)



<!-- ABOUT THE PROJECT -->
## Introduction

This guide describes how to create your own application for the Rexroth Connectivity Unit (RCU). It is going to start out by explaining how to set up your development infrastructure. This guide will not show you how to implement your functionality, but will focus on how to build, package and deploy your applications for/on the Rexroth Connectivity Unit. The examples are written in C++, Golang, Javascript (NodeJS) and Python. Example written in other languages will be added soon.


<!-- GETTING STARTED -->
## Getting Started

Some topics are not covered by this programming guides and are prerequisites for being able to start using it. Moreover you have to setup your development infrastructure.

### Prerequisites

* [Snapcraft overview](https://snapcraft.io/)
* [Tutorial on creating a SNAP](https://snapcraft.io/docs/creating-a-snap)
* [MQTT](https://mqtt.org/)
* [Eclipse Mosquitto](https://mosquitto.org/)
* [Eclipse Mosquitto TLS](https://mosquitto.org/man/mosquitto-tls-7.html)
* [Secure Shell](https://en.wikipedia.org/wiki/Secure_Shell)
* [Can-utils](https://github.com/linux-can/can-utils)
* C++ Programming know-how if you want to use the C++ sample code
* Golang Programming know-how if you want to use the Golang sample code
* Javascript (NodeJS) Programming know-how if you want to use the Javascript (NodeJS) sample code
* Python3 Programming know-how if you want to use the Python3 sample code

### Development Infrastructure

In order to efficiently develop for the RCU you need (additionally to your developer computer) a build computer and a RCU for final testing. They have to be connected in the same TCP/IP network. The following picture shows an example infrastructure where the Developer Computer (IPv4 Address 1), the Build Computer (IPv4 Address 2) and the RCU (IPv4 Address 3) are connected to the same network.
* `Developer Computer`: The developer computer is any computer that you would use for developing your application.
* `Build Computer`: Applications will run on the RCU armhf processor architecture. It is therefore necessary to build and package the application on a computer with the same architecture. The build computer that you use shall therefore have an armhf processor architecture. Additionally SNAP application are always based on an Ubuntu core snap. On RCUs we mainly use core and core18. We also recommend to have a version of Ubuntu 18 installed on the build computer. The build computer we use is a Raspberry Pi 3, on which the version 18.05 of Ubuntu Server 32 bits is installed.
* `RCU`: We recommend to use an RCU4-3W for your development activities.

![Development Infrastructure][img-infrastructure]

### Connecting your computer with the RCU
#### Serial connection

Either Linux OS or Windows OS can be used to connect to the RCU from a Personal Computer (PC) through serial port. The required configuration parameters are the following:
* Bit Rate: 115200 bps
* Data Bits: 8
* Parity: none
* Bit Stop: 1
* Flow Control: None

#### Windows Hyperterminal

Use Windows HyperTerminal to connect the RCU to the PC configuring the serial port parameters to the values indicated in previous section. Switch on the RCU. Once the Kernel is loaded in RAM memory and the system is up, the device waits for the user to enter a valid user name to log in.  Once logged in, the user is in the RCU file system which has the directory structure. To transfer a file from the PC to the RCU, change to /home directory or to the directory where the file is to be stored (cd /home or cd /directory_name), type rz command and choose Transfer-> Send File… option of the HyperTerminal. To transfer a file from the device to the PC, change to the directory where the file is, then type sz command indicating the name of the file (sz file_name) and choose the Transfer -> Receive File… option of the HyperTerminal. File transfer protocol is zmodem in both cases.

#### Linux Minicom
Run the minicom program and configure the serial port parameters to the values indicated in previous section. Minicom help is shown by typing Ctrl-A Z. Serial port device files (/dev/ttyS0, /dev/ttyS1…) must have reading and writing permissions for all users. Log in as root and type chmod a+rw /dev/ttySx to change permissions. Switch on the RCU and wait until login prompt appears. Log in and enter into the device Operating system. To transfer a file from the local PC to the device, change to /home directory (or to the directory where the file is to be stored), type rz command in RCU OS, type Control-A S so that the minicom knows the file that is to be transferred. The file transfer protocol is zmodem. To transfer a file from the device to the local PC, change to the directory where the file to be transferred is stored, type sz command indicating the name of the file (sz file_name) and type Control-A R so that the minicom starts to receive the file. The file transfer protocol is also zmodem.

#### Ethernet connections
To communicate with the RCU a SSH connection can be established too using the Ethernet interface, if the unit features this option. The system get the SSH daemon up by default. Default ip to reach the RCU via SSH: 192.168.10.1 on Port 22. Only login as jarvis is permitted. To get elevated rights the user has to become superuser from the jarvis console.

## Developing Snaps for RCUs

### Development and Testing

The following picture shows the high level architecture of our sample code during development on the developer computer. CAN1 is connected to CAN2 and we have an internal shell script which is generating CAN messages and sending those messages to CAN 2. Two 120 Ohm are are used for ensuring the correct termination (60 Ohm between High and Low).

![Development Architecture][dev-architecture]

#### MQTT Broker on RCU

The mosquitto broker is the central communication hub on the RCU. It handles asynchronous communication between software components on the RCU. Each component has to be authenticated by the Broker in order to be able to publish or receive message to/from the broker. Please get in touch with us at Connect.BODAS@boschrexroth.de to get the right configuration (mosquitto.conf), certificates and keys for communicating with the broker.

#### CAN Snap on RCU

This software module provides communication to the CAN bus based on Linux socket CAN.
It provides the content of the CAN bus to other software modules via MQTT.

##### Parameterization
To change a parameter run

```shell script
snap set can <parameter_key>=<parameter_value>
```
This Snap will filter the can messages on a given CAN IDs. 
To change the filter, create (or modify) a filter.txt and set the path to the file via 
```shell script
snap set can can.<interface>.fiter=<path/to/file/filter.txt>
```

##### Usage example

The snap can be managed using `snapd` like follows.

```shell script
snap start can
snap stop can
snap remove can
snap info can
```

To review the log output, use following command.
The optional `-f` flag allows to review the live logs.

```shell script
journalctl -u snap.snap.can.cand [-f]
```
Alternatively you can use the following command, too. 

```shell script
snap logs can [-f]
```

The CAN messages received by RCU will be published to the broker on the defined topic. 
Per default the topic is `hardware/can/channel/<interface>/rx`
The format of the messages is `(<timestamp>) <interface> <CAN id>#<data>` while the timestamp is in `seconds.microseconds`. 

Example:
```shell script
(1586260384.286194) can1 21A#FE3612FE690507AD
```

The CAN message sending from RCU will be sent to CAN bus on the defined topic. 
Per default the topic is `hardware/can/channel/<interface>/tx` 
The format of message is same as above, while the timestamp field will be ignored, because the CAN message will be send once the mqtt topic is received by this module. So zero '0' could be used as placeholder. Example:
```shell script
(0) can1 21A#FE3612FE690507AD
```
##### Configuration Parameters

The following sample json snippet contains all parameters of the software module snap. 
Add this section to the `/etc/rexroth/rcu.json` configuration file on the RCU.

```
{
    "module": "can",
    "private": {
        "status.interval": 60,
        "logging.level": "info",
        "can1.baudrate": 250000,
        "can1.filter": "/etc/rexroth/can1.filter",
        "can2.baudrate": 250000,
        "can2.filter": "/etc/rexroth/can2.filter",
        "can3.baudrate": 250000,
        "can3.filter": "/etc/rexroth/can3.filter"
    },
    "public": {}
}
```

#### Sample Code on the Developer Computer

The sample code is structured as follow:
* Definition of variables for running the code as an application on the RCU. These definitions are commented out when the code is not run on the RCU.
* Definition of variables for running the code as an application on your computer or on the build. These definitions are commented out when the code is run on the RCU.
* Definition of variables containing the path to the certificates and keys needed for the TLS communication with the MQTT Broker
* Definition of variables containing:
  * the topic for publishing a message that shall be sent through can2
  * the message be sent through can2
  * the topic for receiving messages that are coming through can1
  * the topic for receiving can status messages
* Inititialization of the communication with the mqtt broker
* Main functionality:
  * All messages received by the broker are printed on the console
  * A message is sent to the broker every 5 seconds


Use your preferred IDE for the chosen programming language. Develop your application on your computer using the remote access to the RCU, over <<IPv4 Address 3>>, to communicate with the MQTT Broker.
Test your application. Sample results of testing the application from the developer computer are shown on the following picture. We see that the code is printing received messages from the can1 interface and publishing messages that are forwarded to can2.

![Testing Results from developer computer][test-dev]


### Building and Packaging

When the functionality of your application is successfully tested on your computer, follow following steps for building and packaging the application into a snap:
* configure your application to be packaged as a snap (change the host address of the MQTT Broker to 127.0.0.1, ...)
* create your snapcraft.yaml (learn more on how to create a snapcraft yaml file here: https://snapcraft.io/docs/creating-a-snap)
* copy the source code from the developer computer to the build computer
```sh
scp -r <<source_folder>> <<user_name>>@<<IPv4 Address 2>>:/home/<<user_name>>/
```
* on the build computer:
  * install all necessary libraries needed for building the sample on the build computer
    * E.g. for the C++ sample code you will beed to install:
      * Build tools and libraries
	  ```sh
	  apt-get -y install build-essential git libudev-dev zlib1g-dev libssl-dev gcc g++ cmake
	  ```
	  * [Eclipse Paho MQTT C++ Client Library](#https://github.com/eclipse/paho.mqtt.cpp) which requires the [Eclipse Paho C Client Library for the MQTT Protocol](#https://github.com/eclipse/paho.mqtt.c)
  * test your code natively on the build computer to make sure that the functionality can be executed on a processor with armhf architecture
  * install the snapcraft tool
  ```sh
  sudo snap install snapcraft --classic
  ```
  * define your build computer as the host for building
  ```sh
  export SNAPCRAFT_BUILD_ENVIRONMENT=host
  ```
  * build and package your application.
  ```sh
  snapcraft clean && snapcraft
  ```
After building and packaging, a snap with the ending "_armhf.snap" is created.

![Build Results][build-results]
  
#### Installation and Testing

After sucessfully building and packaging your snap:

* Copy the generated snap to the RCU
```sh
scp <<snapname_version>>_armhf.snap <<user_name>>@<<IPv4 Address 3>>:/home/<<user_name>>/
```
* On the RCU:
  * Install the generated snap
  ```sh
  snap install <<snapname_version>>_armhf.snap --devmode --dangerous
  ```
  * Check that the installation was successful
  ```sh
  snap list
  ```
  * Test your snap by showing the output of its console:
  ```sh
  snap logs <<snapname>> -f
  ```
  * After testing you may want to remove the installed snaps
  ```sh
  snap remove <<snapname_version>>_armhf.snap --purge
  ```

![Testing Results from the RCU][test-rcu]


<!-- LICENSE -->
## License

All sample codes are distributed under the MIT License. See `LICENSE` for more information.


<!-- CONTACT -->
## Contact

BODAS Connect Team - Connect.BODAS@boschrexroth.de

Project Link: [BODAS Connect](https://apps.boschrexroth.com/rexroth/en/transforming-mobile-machines/bodas-connect/)

<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements
* [Eclipse Paho C Client Library for the MQTT Protocol](https://github.com/eclipse/paho.mqtt.c), Eclipse Public License v2.0
and Eclipse Distribution License v2.0
* [Eclipse Paho MQTT C++ Client Library](https://github.com/eclipse/paho.mqtt.cpp), Eclipse Public License v1.0 and Eclipse Distribution License v1.0
* [Eclipse Paho MQTT Go client](https://github.com/eclipse/paho.mqtt.golang), Eclipse Public License - v 1.0 and Eclipse Distribution License - v 1.0
* [Eclipse Paho™ MQTT Python Client](https://github.com/eclipse/paho.mqtt.python), Eclipse Public License 1.0 and the Eclipse Distribution License 1.0
* [mqtt NPM Package](https://www.npmjs.com/package/mqtt), The MIT License
* [Best-README-Template](https://github.com/othneildrew/Best-README-Template), The MIT License

## Image sources
* https://pixabay.com/vectors/device-electronic-open-hardware-pi-1295187/
* https://pixabay.com/vectors/wlan-telecommunications-router-2007682/
* https://pixabay.com/illustrations/macbook-vector-macbook-pro-4515471/


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[img-infrastructure]: images/dev_infrastructure.png
[dev-architecture]: images/dev_architecture.png
[test-dev]: images/test_from_dev.png
[build-results]: images/build_results.png
[test-rcu]: images/test_from_rcu.png
