#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2012.05.04

# Harder to read than fizzbuzz.py

for num in xrange(1, 101):
    if num % 3 == 0:
        if num % 5 == 0:
            print "FizzBuzz"
            continue
        else:
            print "Fizz"
            continue
    if num % 5 == 0:
        print "Buzz"
        continue
    print num
