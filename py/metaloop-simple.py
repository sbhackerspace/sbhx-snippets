#!/usr/bin/env python

direction = 1
printme = 3
times = 20

def mymeta(direction, printme, times):
    printme += direction

    if printme == 1 or printme == 9:
        direction *= -1

    print printme
    if times > 0:
        mymeta(direction, printme, times-1)

mymeta(direction, printme, times)
    

# if '__init__' == '__main__':
#     while True:
#         mymeta(printme)
