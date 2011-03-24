#!/bin/bash
# Steve Phillips / fraktil
# 2010.02.08

ifconfig | grep 'inet ' | awk '{print $2}'| cut -d: -f2 | grep -v ^127 | grep -v ^172
#ifconfig | grep 'inet ' | awk '{print $2}'| cut -d: -f2 | grep -v ^127
#ifconfig |grep 'inet '|cut -d\  -f12|cut -d: -f2| grep -v ^127
