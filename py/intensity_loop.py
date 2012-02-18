#!/usr/bin/env python
# 2012.02.18

import datetime

previous_date = datetime.datetime.now()
previous_intensity = -1

while True:
    intensity = int(raw_input("Distraction intensity?  "))
    now = datetime.datetime.now()
    diff = (now - previous_date)
    print "Previous intensity:", previous_intensity, "; Elapsed time since:", diff
    previous_date = now
    previous_intensity = intensity

