#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.01.06

import re
import urllib2

html = urllib2.urlopen('https://clojure.github.io/clojure/clojure.core-api.html').read()
html = html.replace('&gt;', '>').replace('&lt;', '<')

for func in re.findall(r'<a href="#clojure.core/(.*?)" class="toc-entry-anchor">', html, re.DOTALL):
    print func

