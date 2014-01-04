#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2013.12.29

for case in xrange(int(raw_input())):
    print "Case #%d: %s" % (case+1, ' '.join(raw_input().split(' ')[::-1]))
