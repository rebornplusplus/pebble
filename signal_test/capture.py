#!/usr/bin/python3

import signal, time

def handle_signal(signum, frame):
    signame = signal.Signals(signum).name
    print(f"Python Signal handler called with signal {signame} ({signum})")
    exit(signum)

signal.signal(signal.SIGINT, handle_signal)
signal.signal(signal.SIGILL, handle_signal)
signal.signal(signal.SIGTERM, handle_signal)
signal.signal(signal.SIGSEGV, handle_signal)

while True:
    time.sleep(1)
