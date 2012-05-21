#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2012.05.20

import os

print '/'.join(os.path.realpath(__file__).split('/')[:-1])
print os.path.abspath(os.path.join(os.path.dirname(__file__)))
