#!/bin/bash

while true; do
    cansend can2 001#0102030405060708
	sleep 1
    cansend can2 002#1122334455667788
    sleep 1
    cansend can2 003#9988776655443322
    sleep 1
    cansend can2 004#0000111122223333
    sleep 1
done