#!/usr/bin/env python
# Authors: elimisteve, katbon
# 2014.09.01

import re
import requests

def place_sounds_good(post_html):
    return 'male' in post_html and not 'female' in post_html and \
        not 'ellwood' in post_html
        # ' 420 ' not in post_html and \

BASE_URL = 'http://santabarbara.craigslist.org'

# This would normally be scraped from a CL search result, such as
# https://santabarbara.craigslist.org/search/roo?query=male&sale_date=-
listings_html = ['<a href="/roo/4647511762.html" data-id="4647511762" class="hdrlnk">IV  Spot for Chill Male Roommate </a>'] * 5

#post_urls = re.findall(r'<a href="(.*?)"', ''.join( listings_html ))

# Only include URLs that consist of 1 or more non-whitespace characters
post_urls = re.findall(r'<a href="([^\s]+?)"', ''.join( listings_html ))

urls = [BASE_URL + post_url for post_url in post_urls]

for url in urls:
    cl_post = requests.get(url).content.lower()
    if place_sounds_good(cl_post):
        print url
