#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.01.10

import hashlib, sys

if sys.argv[1] == '-f':
    for f in sys.argv[2:]:
        for line in open(f, 'r').readlines():
            line = line.strip('\n')
            print hashlib.md5(line).hexdigest(), line
else:
    for arg in sys.argv[1:]:
        print hashlib.md5(arg).hexdigest(), arg
