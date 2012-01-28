#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2012.01.28

import commands

cmd = 'ls'
status, output = commands.getstatusoutput(cmd)

if not status: # status == 0 when no error returned
    print output
else:
    print "Error:", status
