#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.03.20

import feedparser, time

account_name = 'sbhackerspace'
branch = 'master'
repo_names = ['sbhx-ircbot', 'sbhx-rov', 'sbhx-sicp', 'sbhx-snippets',
              #'sbhx-androidapp', 'sbhx-projecteuler'
              ## These two don't get parsed correctly for some reason...
              ]
SLEEP_SECONDS = float(60*2)/len(repo_names)  # Check each repo once/2 minutes

def main():
    old_version = {}
    while True:
        for repo in repo_names:
            old_version[repo] = feedparser.parse(
                'https://github.com/' + account_name +
                '/' + repo + '/commits/' + branch + '.atom'
                )

        time.sleep(SLEEP_SECONDS)  # Wait then compare

        for repo in repo_names:
            new = feedparser.parse('https://github.com/' + account_name +
                                   '/' + repo + '/commits/' + branch + '.atom')

            if new.entries[0] != old_version[repo].entries[0]:
                #author = new.entries[0].author.split()[0]  # First name
                author = new.entries[0].author_detail.href.split('/')[-1]
                commit_msg = new.entries[0].title
                print author, "committed to", repo
                print "   ", commit_msg
                print

if __name__ == '__main__':
    main()
