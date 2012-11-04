#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2012.05.20

import os

print 1, '/'.join(os.path.realpath(__file__).split('/')[:-1])
print 2, os.path.join(os.path.realpath(__file__).split('/')[:-1])
print 3, os.path.abspath(os.path.dirname(__file__))
print 4, os.path.realpath(os.path.dirname(__file__)) + '/'
print 5, os.path.realpath(__file__)
print 6, os.path.dirname(os.path.realpath(__file__))

#print open(__file__).read()
