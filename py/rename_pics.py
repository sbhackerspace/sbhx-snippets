#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.11.19

import os

old_names = ['apple.jpg', 'banana3.jpg', 'carrot.jpg',
             'carrot_stem.jpg', 'jacket_frnt.jpg', 'jacket_back.jpg',
             'jacket.jpg']
new_names = ['apple4.jpg', 'banana35.jpg', 'carrot03.jpg',
             'carrot_stem7.jpg', 'jacket_frnt8.jpg', 'jacket_back0.jpg',
             'jacket3.jpg']

# OLD_FILE_DIR = ''  # '/Users/username/Documents/old_pics'
# NEW_FILE_DIR = ''  # '/Users/username/Documents/NEW_pics'

# old_names = os.listdir(OLD_FILE_DIR)
# new_names = os.listdir(NEW_FILE_DIR)

olds = [word.split('.')[0] for word in old_names]
#news = [word.split('.')[0] for word in new_names]

for new in new_names:
    matches = [word for word in olds if new.startswith(word)]
    old = sorted(matches, key=len)[-1]
    #print "Old name was", old, "and new name is", new
    os.rename('/tmp/rename_pics/'+old+'.jpg',
              '/tmp/rename_pics/'+new)
