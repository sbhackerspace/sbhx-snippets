#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2012.05.04

for num in xrange(1, 101):
    line = ""
    if num % 3 == 0:
        line += "Fizz"
    if num % 5 == 0:
        line += "Buzz"
    print line if line else num
