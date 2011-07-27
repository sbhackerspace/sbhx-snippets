#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.01.11

filename = '/tmp/sometextfile.txt'

try:
    line_list = open(filename, 'r').readlines()
except IOError:
    _ = open(filename, 'w').close()  # Creates blank file
    line_list = []

line_list.append('Appended to ' + filename + '\n')
line_list.append('Last line of ' + filename + '\n')

# Write file
f = open(filename, 'w')
f.write(''.join(line_list))
f.close()

print line_list
