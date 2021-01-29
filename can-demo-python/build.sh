#!/usr/bin/env bash

sudo chmod +x snap/hooks/*
snapcraft clean && snapcraft
