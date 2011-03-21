#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.03.20

import feedparser, time

SLEEP_SECONDS = 2

repo_names = ['sbhx-ircbot', 'sbhx-rov', 'sbhx-sicp', 'sbhx-snippets',
              #'sbhx-androidapp',
              #'sbhx-projecteuler'
              ## These two don't get parsed correctly for some reason...
              ]

def main():
    last_updated = {}
    while True:
        for repo in repo_names:
            old = feedparser.parse('https://github.com/sbhackerspace/' + \
                                       repo + '/commits/master.atom')
            time.sleep(SLEEP_SECONDS)
            d = feedparser.parse('https://github.com/sbhackerspace/' + \
                                         repo + '/commits/master.atom')

            if d.entries[0] != old.entries[0]:
                print d.entries[0].author.split()[0], "committed to", repo
                print "   ", d.entries[0].title
                print

if __name__ == '__main__':
    main()
