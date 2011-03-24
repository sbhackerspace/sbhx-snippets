#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.02.07

user_input = raw_input()

while user_input != '42':
    print user_input
    user_input = raw_input()


# From http://www.spoj.pl/forum/viewtopic.php?t=49#p20419 (bottom) --

# for line in iter(raw_input, '42'):
#     print line
