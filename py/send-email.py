#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2012.03.19

# Copied and modified version of code at
# http://www.doughellmann.com/PyMOTW/smtpd/

import commands
import email.utils
import smtplib
import sys

from email.mime.text import MIMEText

# Only run on port 25 if root

PORT = 1025
# PORT = 25

if PORT == 25:
    _, whoami = commands.getstatusoutput('whoami')

    if whoami != 'root':
        print "Must run as root, or else you can't listen on port 25!"
        sys.exit(1)


# Create the message
msg = MIMEText('This is the body of the message.')
msg['To'] = email.utils.formataddr(('Recipient', 'recipient@example.com'))
msg['From'] = email.utils.formataddr(('Author', 'author@example.com'))
msg['Subject'] = 'Simple test message'

server = smtplib.SMTP('127.0.0.1', PORT)
server.set_debuglevel(True) # show communication with the server
try:
    server.sendmail('author@example.com', ['recipient@example.com'], msg.as_string())
finally:
    server.quit()
