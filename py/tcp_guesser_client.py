#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2016.10.06

# Code derived from https://pymotw.com/2/socket/tcp.html

import socket
import sys

if len(sys.argv) != 3:
    print "Usage: {} <host> <port>".format(sys.argv[0])
    sys.exit(1)

host = sys.argv[1]
port = int(sys.argv[2])


# Create a TCP/IP socket
sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

# Connect the socket to the port where the server is listening
server_address = (host, port)
sock.connect(server_address)

try:
    # Get guess from stdin, then send to server
    guess = raw_input("")
    sock.sendall(guess)

    # Receive and print response from server
    resp_data = sock.recv(100)
    print resp_data,

finally:
    sock.close()
