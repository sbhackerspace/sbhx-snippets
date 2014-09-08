#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2014.09.01

import sys

print "sys.argv ==", sys.argv

for i in xrange(len(sys.argv)):
    print i, sys.argv[i]
