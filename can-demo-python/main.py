#!/usr/bin/python3
#  Copyright (c) 2021 Bosch Rexroth AG
#  All rights reserved. See LICENSE file for details.
import logging
import subprocess
from subprocess import PIPE


def run():
    """Read two CAN messages from can1 and write back their merged payload in new message.

    :return: None
    """
    first_frame, second_frame = _read_two_can_messages()
    merged_message = _merge_can_frames(first_frame, second_frame)
    _send_can_message(merged_message)


def _read_two_can_messages():
    """Read two messages from can1 channel

    :return: tuple of two CAN messages.
    """
    cmd = [
        "candump",
        "can1",
        "-L",
        "-n 2",
    ]
    run_process = subprocess.run(cmd, stdout=PIPE, stderr=PIPE, check=False)
    output = run_process.stdout.decode()
    # output = "(1611830256.616413) can1 18FF2703#0000000011111111\n" \
    #         "(1611830256.617581) can1 18FF2803#1100111100000000"
    first_frame, second_frame = output.splitlines()
    logging.debug(f"Received 2 CAN messages.")
    logging.debug(f"First: {first_frame}")
    logging.debug(f"Second: {second_frame}")
    return first_frame, second_frame


def _merge_can_frames(first_frame: str, second_frame: str):
    """Merge two frames with bitwise OR operator into a new CAN message.

    :param first_frame: first message to merge.
    :param second_frame: second message to merge.
    :return: new CAN message with merged payload of two input messages.
    """
    first_payload = _get_payload(first_frame)
    second_payload = _get_payload(second_frame)

    merged_frames = []
    for first, second in zip(first_payload, second_payload):
        merged_frames.append(first | second)

    payload = []
    for b in merged_frames:
        if len(hex(b)[2:]) == 1:
            payload.append("0" + hex(b)[2:])
            continue
        payload.append(hex(b)[2:])

    merged_message = f"003#{''.join(payload)}"
    return merged_message


def _get_payload(frame: str):
    """Parse CAN message and return its payload.

    :param frame: message to get payload from.
    :return: CAN message payload.
    """
    payload = frame.split(" ")[-1]
    payload = payload.split("#")[1]

    return bytes.fromhex(payload)


def _send_can_message(message):
    """Send a CAN message on can1 channel.

    :param message: message string in format <message-id>#<message-payload>.
    :return: None
    """
    cmd = [
        "cansend",
        "can1",
        message
    ]
    logging.debug(f"Running: {' '.join(cmd)}")
    _ = subprocess.run(cmd, stdout=PIPE, stderr=PIPE, check=False)


if __name__ == '__main__':
    """Entry point for the snap."""
    logging.basicConfig(level=logging.DEBUG)
    logging.info("Starting CAN processing")
    while True:
        run()
