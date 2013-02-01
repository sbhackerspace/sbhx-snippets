#!/usr/bin/env python

import re, urllib2

udacity_homepage = urllib2.urlopen("http://www.udacity.com").read()
regex = '<div class="course-title">(.*?)</div>'
for class_name in re.findall(regex, udacity_homepage, re.DOTALL):
    print class_name
