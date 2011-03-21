#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.03.20

import commands, feedparser, re, time

REPO_HOME = '/home/steve/sbhx/'
SLEEP_SECONDS = 10

repo_names = [
    #'sbhx-androidapp'
    'sbhx-ircbot',
    #'sbhx-projecteuler'
    'sbhx-rov',
    'sbhx-sicp',
    'sbhx-snippets'
    ]

def main():
    last_updated = {}
    for repo in repo_names:
         # Download and parse feed
        d = feedparser.parse('https://github.com/sbhackerspace/' + repo + '/commits/master.atom')
        #print repo, '\t Last updated', d.feed.updated #d['feed']['updated']

         # Grab commit message
        cmd = "cd " + REPO_HOME + repo + " && git log | head -n 6"
        (status, output) = commands.getstatusoutput(cmd)

        #print "OUTPUT:", output

        if not 'merge' in output.lower():
            commit, author, date, _, msg, _ = output.split('\n')
        else:
            commit, merge, author, date, _, msg = output.split('\n')

        print "Repo:  ", repo
        print author
        print msg
        print

    while True:
        for repo in repo_names:
            if not repo in last_updated:  # first run
                last_updated[repo] = d.feed.updated
            elif last_updated[repo] != d.feed.updated:
                print repo, "has been updated"
                cmd = "cd " + REPO_HOME + repo + " && git log | head -n 6"
                (status, output) = commands.getstatusoutput(cmd)

                if not 'merge' in output.lower():
                    commit, author, date, _, msg = output.split()
                else:
                    commit, merge, author, date, _, msg = output.split()

                print author
                print "Message", msg
                print
            else:
                #pass
                print repo + ":\t Nothing new"
                print

        time.sleep(SLEEP_SECONDS)

if __name__ == '__main__':
    main()
