import contextlib
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

with contextlib.closing(urllib.urlopen(url)) as resp:
    if resp.headers.dict['content-type'] != pdf:
        os.exit(0)

    with open(filename, 'w') as local_file:
        buf = resp.read(10**6)
        while buf:
            local_file.write(buf)
            buf = resp.read(10**6)
