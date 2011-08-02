#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.07.29

import random, sys

if len(sys.argv) < 4:
    name = sys.argv[0].split('/')[-1]
    print "Usage:", name, "[how many nums] [smallest] [biggest]"
    sys.exit(1)

how_many, small, big = [ int(x) for x in sys.argv[1:] ]

for _ in range(how_many):
    print random.randint(small, big),
