#!/usr/bin/env python

def qsort(list):
    if len(list) <= 1:
        return list
    else:
        pivot = list[0]
        less = [x for x in list if x < pivot]
        more = [x for x in list if x > pivot]
        return qsort(less) + [pivot] + qsort(more)

print qsort([1, 16, -3, 5, 2, 27, 14, 29, -17, 20, -9])
