#!/usr/bin/env python
# Steve Phillips / fraktil
# 2010.12.28

import sys

try:
    print sum(int(x) for x in sys.argv[1:])
except:
    print "Give me some ints!"

    
    
