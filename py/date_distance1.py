#!/usr/bin/env python
# Started by: AJ B / avb_wkyhu
# Forked by:  Steve Phillips / elimisteve
# 2011.11.01

print "Input two dates and view the duration between them\n"

##
## Improvement #1: Ask user fewer questions. List comprehensions can
## make code more concise, but still highly readable.
##

# Gets all the numbers for the Start Date
user_input = raw_input("Starting Date (month/day/year): ")
month1, day1, year1 = [int(x) for x in user_input.split('/')]

# If "11" is entered for the year, interpret that as 2011 by adding 2000
if year1 < 100:
    year1 += 2000

# Gets all the numbers for the End Date
user_input = raw_input("Ending Date (month/day/year): ")
month2, day2, year2 = [int(x) for x in user_input.split('/')]

if year2 < 100:
    year2 += 2000

# Gets all of the months in days and adds it with the days. IE for
# 23/9/1994 it will do ((9-1)*30) + 23
if month1 == 1:
    startd = 0
elif month1 > 1:
    startd = (month1 - 1) * 30

startd = startd + day1

# Same as above but for the end date.
if month2 == 1:
    enddate = 0
elif month2 > 1:
    enddate = (month2 - 1) * 30

enddate = enddate + day2

# Gets the difference between the start and end date. IE from 1994 to
# 2011 it will do (2011-1994) * 365

endyear = year2 - year1

endyeardays = endyear * 365


# If for example the first date is 10/30/11 (300 days intotheyear)
# and second is 2/25/12 (55 days intotheyear) then you would want to
# go 365 - (300-55) = 120 which is pretty accurate. On paper I got
# 115. This is because there is a ~5 day error atm since each month
# is 30days
if startd < enddate:
    daysdiff = enddate - startd
elif day1 > day2:
    daystosub = day1 - day2
    daysdiff = 365 - daystosub
# else:
#     daysdiff = enddate - startd  # Just copied from if-statement

exactdays = (endyear * 365) + daysdiff

if daysdiff > 30:
    monthsdiff = daysdiff / 30
    daysdiff = daysdiff - monthsdiff
elif daysdiff < 30:
    monthsdiff = 0

print "\nFrom the starting date to the ending date: \n"
print "It has been %r days, %r months, and %r years" % (daysdiff, monthsdiff, endyear)
print "In only days it has been: %r days" % exactdays
