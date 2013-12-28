#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2013.12.28

num_lines = int(raw_input())
for case in xrange(num_lines):
    words = raw_input().split(' ')
    print "Case #%d: %s" % (case+1, ' '.join(words[::-1]))
