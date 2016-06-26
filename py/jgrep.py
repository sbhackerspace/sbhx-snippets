#!/usr/bin/env python
# Steven Phillips / elimisteve
# 2016.06.25

import json
import sys

def main():
    if len(sys.argv) == 1:
        print "Usage: jgrep <pattern> < input.json"
        return

    json_array = json.loads(sys.stdin.read())
    if type(json_array) is dict:
        json_array = [json_array]

    pattern = sys.argv[1].lower()

    for obj in json_array:
        for k in obj:
            if pattern in k.lower() or in_or_equal(pattern, obj[k]):
                print "{}: {}".format(k, json.dumps(obj[k]))


collection_types = (list, tuple, set)
def is_collection(obj):
    return any([ isinstance(obj, typ) for typ in collection_types ])

def in_or_equal(obj, container):
    if obj == container:
        return True

    if isinstance(obj, basestring) and isinstance(container, basestring):
        if obj in container.lower():
            return True

    if is_collection(container):
        for collection_obj in container:
            if in_or_equal(obj, collection_obj):
                return True

    return False


if __name__ == '__main__':
    main()
