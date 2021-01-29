# Python CAN Example

This is a sample CAN snap written with Python.
It demonstrates how a snap can interact with the CAN reading two messages from can1.
Subsequently, the snap writes back a new message to can1.
The new message contains the merged payload (bitwise OR is applied) of both received messages.

## Setup

## Content

The following section briefly explains the content of this sample.

The sample is a minimal working example without external dependencies. 
In order to read and write from `can1`, the utilities `candump` and `cansend` are directly embedded in [usr/bin/](usr/bin/).

The file [main.py](main.py) contains all the business logic required.
See the function comments for a more detailed explanation of its workflow.
It represents the main entry point of the snap and is called on snap start.

The snap is defined in [snap/snapcraft.yaml](snap/snapcraft.yaml).
See comments for a few detail information about the key instructions.

The file [build.sh](build.sh) is used to simplify building the snap.
It ensures that `candump` and `cansend` binaries are actually executable.
Eventually it cleans possibly remaining build artifacts and starts a new build.

In [setup.py](setup.py), some meta information required by the `python` snap plugin are given.
