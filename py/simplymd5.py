#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.01.10

import sys, hashlib

for string in sys.argv[1:]:
    print hashlib.md5(string).hexdigest()
