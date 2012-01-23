#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2012.01.22

# Simple example
def incrementor(n=0):
    inc = n
    while True:
        yield n
        n += 1

start_at_5 = incrementor(5)
print "A bunch of integers, starting with 5"
for _ in range(10):
    print start_at_5.next()


print 

# Lucky numbers consist of only 4s and 7s
def lucky_gen():
    num = '3'
    while True:
        num = str( int(num)+1 )
        if num[-1] not in '47':
            continue
        if num.count('4') + num.count('7') != len(num):
            continue
        yield num

lucky = lucky_gen()

print "The first 10 lucky numbers"
for _ in range(10):
    print lucky.next()
