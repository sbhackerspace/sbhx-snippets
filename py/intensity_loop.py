#!/usr/bin/env python
# 2012.02.18

while True:
    intensity = int(raw_input("Distraction intensity?  "))
    now = datetime.datetime.now()
    diff = (now - previous_date)
    print "Time elapsed:", diff, "; Previous intensity:", previous_intensity
    previous_date = now
    previous_intensity = intensity

