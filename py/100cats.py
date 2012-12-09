#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Steve Phillips / elimisteve
# 2012.12.09

"""
You have 100 cats (You are quite lucky to own so many cats!).

You have arranged all your cats in a line. Initially, none of your
cats have any hats. You take 100 rounds walking around the cats,
always starting with the first cat. Every time you stop at a cat,
you put a hat on it if it doesn't have one, and you take take its
hat off if it has one on.

The first round, you stop at every cat. The second round, you only stop at
every 2nd cat (#2, #4, #6, #8, etc.). The third round, you only stop at
every 3rd cat (#3, #6, #9, #12, etc.). You continue this process until the
100th pass (i.e. you only visit the 100th cat).

At the end, which cats have hats?
"""
NUM_CATS = 100

cat_nums = xrange(1, NUM_CATS+1)

#cats = [False for _ in cat_nums]  # Equivalent alternative
cats = [False] * NUM_CATS

for n in cat_nums:
    for cat in xrange(len(cats)):
        if cat % n == 0:
            cats[cat] = not cats[cat]

cats_with_hats = []
for cat in xrange(len(cats)):
    if cats[cat]:
        cats_with_hats.append(cat)

print cats_with_hats
