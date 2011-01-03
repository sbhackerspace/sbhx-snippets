#!/usr/bin/python

import time
import sys

greeting = "Welcome to SpeedReader!"
print greeting

wpl = float(raw_input("Words per line? "))
wpm = float(raw_input("Words per minute? "))
fromfile = str(raw_input("Read from file (y/n)? "))
wps = wpm / 60.0
pause = wpl / wps # words/line / words/second = seconds/line

if fromfile[0].lower() == "y":
    readfromfile = True
else:
    readfromfile = False

filename = str(raw_input("File name (leave blank if none)? "))


def speedreader(wpl, wpm, filename):
    """Prints words at desired rate"""
    if readfromfile:
        try:
            file = open(filename, 'r')
        except IOError:
            print "Cannot open file '%s'" % filename
            sys.exit(0)

        file = file.read()
    else:
        print "Input text to speedread (newline, ctrl-d when finished): "
        file = sys.stdin.read()

    print "\n" * 10 # clear screen
    bigstring = ""
    whitespace = '\t\r\n '
    for char in file:
        if char not in whitespace:
            bigstring += char # appends non-newlines to bigstring
        else:
            if bigstring and bigstring[-1] not in whitespace:
                bigstring += " " # turns consecutive whitespace chars
                                 # into one space

    # Print every wpl words
    words = 0
    printstring = ""
    for char in bigstring:
        printstring += char
        if char == " ":
            words += 1
            if words % wpl == 0 and not set(printstring) in set(whitespace):
                print printstring
                time.sleep(pause)
                printstring = ""


speedreader(wpl, wpm, filename)
