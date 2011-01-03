#!/usr/bin/env python
# Steve Phillips / fraktil
# 2010.12.13

import sys
import os

if len(sys.argv) > 1:
    number = sys.argv[1]
else:
    number = raw_input("Give me an integer: ")

try:
    number = int(number)
except:
    print "Not an int! Exiting"
    sys.exit(1)

print number,

while number != 1:
    print ' -> ',
    if number % 2 == 0:
        print number/2,
        number /= 2
        #number >>= 1
    elif number != 1 and number % 2 == 1:
        print number*3 + 1,
        number = number*3 + 1
