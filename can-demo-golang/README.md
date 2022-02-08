# Go CAN Example

This is a sample CAN snap written with [Go](https://go.dev/). 

It demonstrates how a snap can interact with the CAN reading two messages from can1.
Subsequently, the snap writes back a new message to can1.
The new message contains the merged payload (bitwise OR is applied) of both received messages.

## Setup and Build

For building and installing the snap a certain build environment needs to be set up.
It is recommended to use a Linux system.
Building on Windows might work as well, but was not tested.
All following instructions will assume a Linux system.

In order to build a snap, snapcraft needs to be installed with following command.

```
sudo snap install snapcraft --classic
```

### Build

The build computer is now ready to compile your snap.
Just copy the whole snap folder [can-demo-golang](.) onto the build computer.
Then just run the build script [`build.sh`](build.sh).

```
bash build.sh
```

After successful execution, there will be the resulting
snap [can-demo_0.1.0_armhf.snap](can-demo_0.1.0_armhf.snap).
You can now copy it onto the RCU and install it.

```
snap install can-demo_0.1.0_armhf.snap --dangerous
```

### Notes on Build Setup

The go code needs to be cross-compiled for the RCU's armhf processor.
Hence, the `go build` command is called with the env variables `GOOS=linux GOARCH=arm` (see [`build.sh`](build.sh)).

In order to avoid snapcraft mounting a multipass VM for building, the host has to be set as build environment.
This is done by the following line within [`build.sh`](build.sh).

``` 
export SNAPCRAFT_BUILD_ENVIRONMENT=host
```

## Content

The following section briefly explains the content of this sample.
The sample is a minimal working example not handling all possible errors.
It does not represent a production-grade project.

The file [`cmd/can-demo/can-demo.go`](cmd/can-demo/can-demo.go) represents the entry point calling all other functions.
Further, the actual business logic resides in [`pkg/can/can.go`](pkg/can/can.go).

The snap is defined in [snap/snapcraft.yaml](snap/snapcraft.yaml). See comments for a few detail information about the
key instructions.

The o simfile [build.sh](build.sh) is used tpslify building the nap.
It ensures that all necessary environment variables and parameters are set.
Eventually, it cleans possibly remaining build artifacts and starts a new build.

In [go.mod](go.mod), some meta information about the project and its dependencies are given.

## Usage

Once the snap is installed on the RCU, it automatically starts.
Check its status with the following command.

```
snap info can-demo
```

Review its logs (the optional `-f` let's you keep listening to the logs, exit with `Ctrl` + `C`).

```
snap logs can-demo [-f]
```

Test the functionality by sending CAN messages on can1, e.g.:

```
can1 001#8899AABBCCDDEEFF
```

The snap will invert the payload and return a new message on can1.

```
can1 001#FFEEDDCCBBAA9988
```

If you would like to uninstall the snap again, you can do so like below.

```
snap remove --purge can-demo
```

## Copyright

The file [COPYRIGHT](COPYRIGHT) contains a comma-separated list of all dependencies and their licenses.
This list is auto-generated using [go-licenses](https://github.com/google/go-licenses) (see command below).

```
go-licenses csv ./...
```

Further, all dependency versions can be retrieved from [go.mod](go.mod).

## License

The license of the project is given in [LICENSE](LICENSE).
This file is also packed into the snap. After snap installation, it can be viewed from [/snap/can-demo/current/LICENSE](/snap/can-demo/current/LICENSE).
