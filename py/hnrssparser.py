#!/usr/bin/env python

import collections
import feedparser
import string

# This is using feedparser to parse the xml from HN's raw RSS feed
hn_entries = feedparser.parse("http://news.ycombinator.com/rss").entries

# A list of commonwords to rule out. Will use this later in the code.
commonwords = '''all go its had day to
                 has do them his get they
                 not now him like did these
                 she each people some see are
                 out what said for find may
                 be we were water come by
                 on about her of could or
                 first into number one down long
                 your use from would there two
                 been their call way was more
                 that but part with than he
                 made word look this up will
                 can many my and then is
                 it an as at have in
                 if no make when write how
                 other which you oil I who
                 a so time the you'll
                 our can't & am give back
                 why only too'''.split()

# This is putting all of the 'entries' part of the xml code into a
# list.
#entries = ycomb.entries
hn_titles = []
hn_title_words = []

for entry in hn_entries:
        hn_titles.append(entry.title)

#This is splitting each title seperating all the words into one long
#list.
for title in hn_titles:
    for word in title.split():
	    word = word.strip(string.punctuation).lower()
            if word and word.lower() not in commonwords:
                    hn_title_words.append(word)

print '\n\n'
print ', '.join(hn_title_words)
print '\n\n'

word_counter = collections.Counter(hn_title_words)

word_counts = [(count, word) for (word, count) in word_counter.iteritems()]
word_counts.sort()

for (count, word) in word_counts:
	print count, word
