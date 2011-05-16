#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.05.06

import sys

lines = open(sys.argv[1], 'r').readlines()
#print "\n".join([line.strip() for line in lines])
output_file = open(sys.argv[2], 'w')
answer = []

gen = (line.strip() for line in lines)
for it in range(int(gen.next())):
    items = [x for x in gen.next().split(' ')]

    num_pairs = int(items[0])
    letters = items[1::2]
    numbers = [int(x) for x in items[2::2]]
    pairs = zip(letters, numbers)

    orange, blue = {"position": 1}, {"position": 1}

    for ndx in range(num_pairs):
        let, num = pairs[ndx]
        blue["next"] = [n for n,l in pairs if l == 'B'][ndx]
        orange["next"] = [n for n,l in pairs if l == 'O'][ndx]

        if  blue["next"] > blue["position"]:
            blue["position"] += 1
        elif blue["next"] < blue["position"]:
            blue["position"] -= 1



    pre = "Case #%d: " % (it+1,)
    answer.append(pre + ' ')

output_file.write("\n".join(answer))
