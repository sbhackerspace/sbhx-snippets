import os
import sys
import urllib

if len(sys.argv) == 0:
    print 'Usage:', sys.argv[0], '<url>'
    os.exit(1)

url = sys.argv[1]

filename = url.split('/')[-1]
if not filename.endswith('.pdf'):
    filename += '.pdf'

pdf = 'application/pdf'

resp = urllib.urlopen(url)

if resp.headers.dict['content-type'] == pdf:
    local_file = open(filename, 'w')
    local_file.write(resp.read())
    local_file.close()
    resp.close()
