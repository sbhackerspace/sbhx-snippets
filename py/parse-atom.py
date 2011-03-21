#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.03.20

import commands, feedparser, re, time

SLEEP_SECONDS = 10

repo_names = ['sbhx-ircbot', 'sbhx-rov', 'sbhx-sicp', 'sbhx-snippets',
              #'sbhx-androidapp',
              #'sbhx-projecteuler'
              ]

def main():
    last_updated = {}
    for repo in repo_names:
         # Download and parse feed
        d = feedparser.parse('https://github.com/sbhackerspace/' + repo + '/commits/master.atom')
        print repo, '\t Last updated', d.feed.updated #d['feed']['updated']

    while True:
        for repo in repo_names:
            if not repo in last_updated:  # first run
                last_updated[repo] = d.feed.updated
            elif last_updated[repo] != d.feed.updated:
                print d.entries[0].author.split[0], "committed to", repo
                print "   ", d.entries[0].title
            else:
                pass
                #print repo + ":\t Nothing new"

        time.sleep(SLEEP_SECONDS)

if __name__ == '__main__':
    main()
