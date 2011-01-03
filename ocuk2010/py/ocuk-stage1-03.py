#!/usr/bin/env python
# Steve Phillips / fraktil
# 2010.12.13

import ctypes

MAX_NUM = 1500000
one = ctypes.c_uint(1).value
longest_chain = longest_n = ctypes.c_uint(0).value

for num in range(2, MAX_NUM + 1):
    number = ctypes.c_uint(num).value
    chain_length = ctypes.c_uint(0).value

    while number != 1:
        if not number & one: # Even
            number >>= 1
        else:              # Odd
            number += (number << 1) + 1
        chain_length += 1

        if chain_length > longest_chain:
            longest_chain = chain_length
            longest_n = num

print 'n =', longest_n, 'generates chain length', longest_chain
