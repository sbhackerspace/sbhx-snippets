#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.06.13

import sys

contents = open(sys.argv[1], 'r').read()
print sum([1 if char == 'x' else 0 for char in contents])
