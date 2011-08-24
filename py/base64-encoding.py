#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.08.14

import base64, sys

src_filename = sys.argv[1]

source = open(src_filename, 'r')
dest_filename = src_filename+'.base64'
destination = open(dest_filename, 'w')

base64.encode(source, destination)

print open(dest_filename, 'r').read()
