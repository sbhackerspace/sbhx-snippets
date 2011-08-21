#!/usr/bin/env python

direction = -1
printme = 8

def mymeta(direction, printme):
    print printme

    # Read file
    import sys
    f = open(sys.argv[0], "r")
    lines = f.readlines()
    f.close()

    # Modify printme
    lines[3] = lines[3].replace(lines[3][-2],str(int(lines[3][-2])+direction))

    # Extra space problem: solution #1
    # Append a space before '1' but not before '-1' when changing direction
    # The output of ' '*-1 is ''; The output of ' '*1 is (of course) ' '
    space = ' ' * direction

    # Modify direction
    if int(lines[3][-2]) == 1 or int(lines[3][-2]) == 9:
        lines[2] = lines[2].replace(lines[2][-3:-1],space+str(int(lines[2][-3:-1])*-1))

    # Re-open and save changes
    f = open(sys.argv[0], "w")
    f.writelines(lines)
    f.close()

mymeta(direction, printme)
