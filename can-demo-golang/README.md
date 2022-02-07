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

    

## COPYRIGHT

The file [COPYRIGHT](COPYRIGHT) contains a comma-separated list of all dependencies and their licenses.
This list is auto-generated using [go-licenses](https://github.com/google/go-licenses) (see command below).

```
go-licenses csv ./...
```

Further, all dependency versions can be retrieved from [go.mod](go.mod).
