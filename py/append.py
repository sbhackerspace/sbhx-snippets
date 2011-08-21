#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.01.11

filename = '/tmp/mytextfile.txt'

# Open and read file
f = open(filename, 'a+')
line_list = f.readlines()

# Append to line_list
line_list.append('Appended to ' + filename + '\n')
line_list.append('Last line of ' + filename + '\n')

# Write file
f.write(''.join(line_list))
f.close()

print ''.join(line_list),
