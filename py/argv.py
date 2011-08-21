#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.01.09

import sys

for arg in sys.argv:
    print arg,
print "\n"

print "Found", len(sys.argv), "args"
print "sys.argv[0] ==", sys.argv[0]
