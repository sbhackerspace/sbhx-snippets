#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.04.28

import sys

lines = open(sys.argv[1], 'r').readlines()
#print "\n".join([line.strip() for line in lines])
output_file = open(sys.argv[2], 'w')
answer = []

# it == iteration
gen = (x.strip() for x in lines)
for it in range(int(gen.next())):
    total, length = int(gen.next()), int(gen.next())
    items = [int(x) for x in gen.next().split(' ')]

    match = False
    for ndx1 in range(len(items)):
        for ndx2 in range(len(items)):
            if items[ndx1] + items[ndx2] == total and ndx1 != ndx2:
                match = True
            if match:
                answer.append("Case #%d: %d %d" % (it+1, ndx1+1, ndx2+1))
                break
        if match:
            break

output_file.write("\n".join(answer))
