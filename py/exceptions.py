#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2014.05.04

# Catch all exceptions and print out the error that occurred

try:
    x = 1/0
except Exception, e:
    print e
