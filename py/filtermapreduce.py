#!/usr/bin/env python
# 2009.04.24

def catstrings(x, y):
    print str(x)+str(y)

catstrings(1, 5)

print "reduce(lambda x, y: str(x)+str(y), range(9)) =>", reduce(lambda x, y: str(x)+str(y), range(9))
