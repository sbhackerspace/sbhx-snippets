#!/bin/bash

cd ~ && mkdir sbhx && cd sbhx && for repo in {ircbot,snippets,rubyquiz,signin,projecteuler,sicp,rov,androidapp}; do git clone git@github.com:sbhackerspace/sbhx-$repo; done
