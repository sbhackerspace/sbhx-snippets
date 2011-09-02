#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.09.02

# From http://devinfee.com/blog/2010/11/08/make-json-requests-with-python/

import json, urllib2

data = {'key': 'value', 'hello': 'world'}
data_json = json.dumps(data)
host = "http://localhost:9999/"
req = urllib2.Request(host, data_json,
                      {'content-type': 'application/json'})
response_stream = urllib2.urlopen(req)
response = response_stream.read()
print response
