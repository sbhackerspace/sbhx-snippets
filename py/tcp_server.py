#!/usr/bin/env python

# Code sample from https://wiki.python.org/moin/TcpCommunication

import socket

TCP_IP = '127.0.0.1'
TCP_PORT = 5005

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.bind((TCP_IP, TCP_PORT))
s.listen(1)

conn, addr = s.accept()
print 'Connection address:', addr
while True:
    data = conn.recv(1024)
    if not data:
        break
    print data,
    conn.send(data)  # echo
conn.close()
