#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2012.01.22

import string

print string.letters[:26]

print ''.join(chr(ord('a')+inc) for inc in range(26))

