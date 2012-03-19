#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2012.03.19

# Copied and modified version of code at
# http://www.doughellmann.com/PyMOTW/smtpd/

import asyncore
import commands
import smtpd
import sys

# Only run on port 25 if root

PORT = 1025
# PORT = 25

if PORT == 25:
    _, whoami = commands.getstatusoutput('whoami')

    if whoami != 'root':
        print "Must run as root, or else you can't listen on port 25!"
        sys.exit(1)


class CustomSMTPServer(smtpd.SMTPServer):

    def process_message(self, peer, mailfrom, rcpttos, data):
        print 'Receiving message from:', peer
        print 'Message addressed from:', mailfrom
        print 'Message addressed to  :', rcpttos
        print 'Message length        :', len(data)
        return


server = CustomSMTPServer(('0.0.0.0', PORT), None)

asyncore.loop()
