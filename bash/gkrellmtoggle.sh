#!/bin/bash
# Steve Phillips / elimisteve
# 2010.12.18

# If gkrellm is running, kill it.
# If it's not running, start it.

if grep -q gkrellm <(echo `ps aux | grep -v grep`); then
	killall gkrellm &
else
	gkrellm &
fi
