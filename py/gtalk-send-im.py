#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.08.31

# From http://www.sakibiqbal.com/2011/programming/sending-google-talk-im-using-python/

import xmpp
import time

# Credentials
username = 'sbhackerspace@gmail.com'
password = 'secret'

recipient = 'elimisteve@gmail.com'
msg = 'Test message sent at ' + time.ctime()

# Create client
client = xmpp.Client('gmail.com')
client.connect( ('talk.google.com', 5223) )

# Connect/Authenticate
jid = xmpp.protocol.JID(username)
client.auth(jid.getNode(), password)

# Send IM
client.send( xmpp.protocol.Message(recipient, msg, typ = "chat") )

time.sleep(1)

print "Message '" + msg + "' sent."
