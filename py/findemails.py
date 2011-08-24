#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.08.01

import re, sys

if len(sys.argv) < 2:
    name = sys.argv[0].split('/')[-1]
    print "Usage:", name, "contains-emails.csv [moreemails.xml] ..."
    sys.exit(1)

# The contents of all command line-specified files separated by a space
contents = ' '.join([ open(x, 'r').read() for x in sys.argv[1:] ])

emails = []
for email in re.findall(r'[\w\-+][\w\-+.]+@[\w\-][\w\-.]+[a-zA-Z]{1,4}',
                        contents, re.DOTALL):
    emails.append(email)
    #print email

# Converting to a set removes duplicates
print '\n'.join(sorted(set(emails)))
#print ', '.join(emails)
