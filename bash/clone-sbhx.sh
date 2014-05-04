#!/bin/bash

cd ~; mkdir sbhx; cd sbhx; for repo in {ai,androidapp,boardinator,cryptag,firefoxos,gomandrill,ircbot,ntlanman,projecteuler,rov,rubyquiz,sicp,signin,sbitter,snippets,spaceapi,the-fist,website,webtech}; do git clone git@github.com:sbhackerspace/sbhx-$repo; echo; echo; done
