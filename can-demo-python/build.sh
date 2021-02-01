#!/usr/bin/env bash
#  Copyright (c) 2021 Bosch Rexroth AG
#  All rights reserved. See LICENSE file for details.

sudo chmod +x usr/bin/*
snapcraft clean && snapcraft
