#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.03.21
# 2011.08.24

from multiprocessing import Process
import os, time

def func(num):
    print 'Loop', num
    time.sleep(num+1)
    print 'Loop %d done sleeping' % num

if __name__ == '__main__':
    procs = []
    for num in range(3):
        p = Process(target=func, args=(num,))
        p.start()
        procs.append(p)

    print "Way before", os.getpid()

    for p in procs:
        print "Before", os.getpid(), os.getppid()
        p.join()
        print "After", os.getpid(), os.getppid()
        
