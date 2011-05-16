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
    items = [x for x in gen.next().split(' ')]
    pre = "Case #%d: " % (it+1,)
    answer.append(pre + ' '.join(items[::-1]))


output_file.write("\n".join(answer))
