#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2014.05.29

# Solution for http://opengarden.com/jobs

import itertools

target_sum = 10**8

populations = [18897109, 12828837, 9461105, 6371773, 5965343, 5946800, 5582170,
               5564635, 5268860, 4552402, 4335391, 4296250, 4224851, 4192887,
               3439809, 3279833, 3095313, 2812896, 2783243, 2710489, 2543482,
               2356285, 2226009, 2149127, 2142508, 2134411]

all_combos = [itertools.combinations(populations, num_cities)
              for num_cities in xrange(1, len(populations)+1)]

for combos_of_length_n in all_combos:
    for combo in combos_of_length_n:
        if sum(combo) == target_sum:
            print combo
