#!/bin/bash
# Steve Phillips / fraktil
# 2010.12.18

lynx -dump checkip.dyndns.org 2>&1 | awk '{print $4}' | grep ^[0-9]
