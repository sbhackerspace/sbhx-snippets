#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.02.07

from math import sqrt

def primes_between(low, high):
    "Returns list of primes between low and high, inclusive"
    sqrt_low  = int(sqrt(low))
    sqrt_high = int(sqrt(high))
    sieve     = range(sqrt_low, high+1)

    for i in xrange(sqrt_low, sqrt_high):
        if sieve[i] != 0:
            sieve[i*i: high+1 : i] = [0] * len(sieve[i*i: high+1 : i])

    return [x for x in sieve if x]


print primes_between(10, 20)

num_ranges = 0#int(raw_input())

for _ in range(num_ranges):
    low, high = [int(x) for x in raw_input().split()]

