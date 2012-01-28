#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2012.01.28

import commands

cmd = 'ls'
status, output = commands.getstatusoutput(cmd)

if status:
    print "Error:", output
else: # status == 0 when no error returned
    print output
