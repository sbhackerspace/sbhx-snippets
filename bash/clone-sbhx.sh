#!/bin/bash

cd ~ && mkdir sbhx && cd sbhx && for repo in {ai,androidapp,ircbot,projecteuler,rov,rubyquiz,sicp,signin,snippets,website}; do git clone git@github.com:sbhackerspace/sbhx-$repo; done
