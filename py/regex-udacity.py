#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2012.08.09

import re
import urllib2

# Download the Udacity homepage HTML
udacity_homepage = urllib2.urlopen("http://www.udacity.com").read()

# Use this regex string to parse out the names of all of Udacity's classes
regex = '<h4 class="course-shortcut-grid-course-name"><a href="(?:.*?)">(.*?)</br>'

for class_name in re.findall(regex, udacity_homepage, re.DOTALL):
    print class_name
