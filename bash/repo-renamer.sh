#!/bin/bash
# Steve Phillips / elimisteve
# 2012.03.24

for repo in {django-crud-gen,django-projectbuilder,go-helpers}; do
    mv proto-$repo/ $repo/ && cd $repo && git remote rm origin && git remote add origin git@github.com:prototypemagic/$repo.git && git pull origin master && cd ..
done
