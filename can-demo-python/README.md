# Python CAN Example

This is a sample CAN snap written with Python. It demonstrates how a snap can interact with the CAN reading two messages
from can1. Subsequently, the snap writes back a new message to can1. The new message contains the merged payload (
bitwise OR is applied) of both received messages.

## Setup and Build

For building and installing the snap a certain build environment needs to be set up. As described in the
main [README.md](../README.md), it requires a Raspberry Pi 3 as a build computer.

### Raspberry Pi

The Raspberry Pi (RasPi) should run a 32-bit Ubuntu 18. Specifically, we
recommend [ubuntu-18.04.4-preinstalled-server-armhf+raspi3.img.xz](http://old-releases.ubuntu.com/releases/18.04.4/ubuntu-18.04.4-preinstalled-server-armhf+raspi3.img.xz)
from the [Ubuntu release archive](http://old-releases.ubuntu.com/releases/18.04.4/). Flash it onto the RasPi's SD card
using a tool of your choice (e.g. [Rufus](https://rufus.ie/) (recommended)
or [Balena Etcher](https://www.balena.io/etcher/)).

Subsequently, plug the card in the RasPi and boot it (default credentials should be `user: ubuntu`
and `password: ubuntu`).
[Set up wifi](https://netplan.io/examples/) or plug the RasPi into your router via ethernet. Before building the snap,
you need to install snapcraft as follows.

```
sudo snap install snapcraft --classic
```

### Build

The build computer (the RasPi) is now ready to compile your snap. Just copy the whole snap folder [can-demo-python](.)
onto the RasPi. Using a USB drive this can be done by following instructions.

```
# Mount the USB drive at the moint point /mnt (note: the device may vary e.g. sda, sdb, ...).
sudo mount /dev/sda1 /mnt
# Copy the folder recursively into the current directory.
sudo cp -R /mnt/can-demo-python ./can-demo-python
```

Make sure the folder and its contents are owned by your user `ubuntu`.

```
sudo chown -R ubuntu:ubuntu ./can-demo-python
```

It's also important that the snapcraft environment variable `SNAPCRAFT_BUILD_ENVIRONMENT` is set to "host" before building.

```
export SNAPCRAFT_BUILD_ENVIRONMENT=host
```

Then just run the [build script](build.sh).

```
sh build.sh
```

After successful execution, there will be the resulting
snap [python-can-example_1.0_armhf.snap](python-can-example_1.0_armhf.snap). You can now copy it onto the RCU and
install it.

```
snap install python-can-example_1.0_armhf.snap --dangerous --devmode
```

## Content

The following section briefly explains the content of this sample.

The sample is a minimal working example without external dependencies. In order to read and write from `can1`, the
utilities `candump` and `cansend` are directly embedded in [usr/bin/](usr/bin/).

The file [main.py](main.py) contains all the business logic required. See the function comments for a more detailed
explanation of its workflow. It represents the main entry point of the snap and is called on snap start.

The snap is defined in [snap/snapcraft.yaml](snap/snapcraft.yaml). See comments for a few detail information about the
key instructions.

The file [build.sh](build.sh) is used to simplify building the snap. It ensures that `candump` and `cansend` binaries
are actually executable. Eventually it cleans possibly remaining build artifacts and starts a new build.

In [setup.py](setup.py), some meta information required by the `python` snap plugin are given.

## Usage

Once the snap is installed on the RCU, it automatically starts. Check its status with the following command.

```
snap info python-can-example
```

Review its logs (the optional `-f` let's you keep listening to the logs, exit with `Ctrl` + `C`).

```
snap logs python-can-example [-f]
```

Test the functionality by sending two CAN messages on can1, e.g.:

```
can1 001#00000000AAAAAAAA
can1 002#BBBBBBBB00000000
```

The snap will merge the payloads and return a new message on can1.

```
can1 003#BBBBBBBBAAAAAAAA
```
