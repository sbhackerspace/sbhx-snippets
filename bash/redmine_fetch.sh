#!/bin/bash

# Remember to use
#   git clone --mirror $url
# when cloning repo

repo_dir=/home/ubuntu/git_repos

for repo in $(ls $repo_dir); do
    cd $repo_dir/$repo && git fetch -q --all
done
