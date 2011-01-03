#!/usr/bin/env python
# Steve Phillips / fraktil
# 2010.12.13

MAX_NUM = 1500000

longest_chain = longest_n = 0

for num in range(2, MAX_NUM + 1):
    number = num
    chain_length = 0

    while number != 1:
        if not number & 1: # Even
            number >>= 1
        else:              # Odd
            number += (number << 1) + 1
        chain_length += 1

        if chain_length > longest_chain:
            longest_chain = chain_length
            longest_n = num

print 'n =', longest_n, 'generates chain length', longest_chain
