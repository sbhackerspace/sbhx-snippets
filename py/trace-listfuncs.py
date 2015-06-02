#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2015.06.02

# Run this with
#   python -m trace --listfuncs trace-listfuncs.py

def f():
    return True

def g():
    # Not to be confused with f()
    return f

def h():
    # Will the trace show i()?
    def i():
        return g()
    return i()


print h()()
