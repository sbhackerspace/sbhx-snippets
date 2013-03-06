#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Steve Phillips / elimisteve
# 2013.03.05

# Forked from http://akshitshah.com/?p=7

import boto
import os


def s3_dir_uploader(bucket, bucket_dir, src_dir):
    conn = boto.connect_s3()
    b = conn.get_bucket(bucket)
    k = boto.s3.key.Key(b)
    for path, dir, files in os.walk(src_dir):
        for f in files:
            full_filename = os.path.join(path, f)
            print "Uploading", f,
            k.key = bucket_dir + "/" + os.path.relpath(full_filename, src_dir)
            print "to", k.key
            k.set_contents_from_filename(full_filename)
            b.set_acl('public-read', k.key)
            print "File should be visible here:",
            print "https://%s.s3.amazonaws.com/%s" % (bucket, k.key)
            print


bucket = 'mcc.prototypemagic.com'
bucket_dir = 'images'
src_dir = '/home/steve/django_projects/mycollectionconnection_site/media/images'

s3_dir_uploader(bucket, bucket_dir, src_dir)
