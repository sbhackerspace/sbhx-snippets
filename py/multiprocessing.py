#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.03.21

from multiprocessing import Process

def f(name):
    print 'hello', name

if __name__ == '__main__':
    p = Process(target=f, args=('bob',))
    p.start()
    p.join()
