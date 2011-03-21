#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.03.20

import feedparser, time

SLEEP_SECONDS = 1

account_name = 'sbhackerspace'
branch = 'master'
repo_names = ['sbhx-ircbot', 'sbhx-rov', 'sbhx-sicp', 'sbhx-snippets',
              #'sbhx-androidapp', 'sbhx-projecteuler'
              ## These two don't get parsed correctly for some reason...
              ]

def main():
    last_updated = {}
    while True:
        for repo in repo_names:
            old = feedparser.parse('https://github.com/' + account_name +
                                   '/' + repo + '/commits/' + branch + '.atom')
            time.sleep(SLEEP_SECONDS)  # Wait then compare

            d = feedparser.parse('https://github.com/' + account_name +
                                   '/' + repo + '/commits/' + branch + '.atom')

            if d.entries[0] != old.entries[0]:
                #author = d.entries[0].author.split()[0]  # First name
                author = d.entries[0].author_detail.href.split('/')[-1]
                commit_msg = d.entries[0].title
                print author, "committed to", repo
                print "   ", commit_msg
                print

if __name__ == '__main__':
    main()
