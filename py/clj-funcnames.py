#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.01.06

import re, urllib2

html = urllib2.urlopen('http://richhickey.github.com/clojure/clojure.core-api.html').read()
html = html.replace('&gt;', '>').replace('&lt;', '<')

for func in re.findall(r"<a href=\"#clojure.core/(.*?)\">", html, re.DOTALL):
    print func

