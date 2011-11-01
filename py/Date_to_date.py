## Written by avb_wkyhu
## Know right away that this is a buggy and inexact program made for practice purposes
print "\n"
print "Date to Date v1.0.\n"
print "*****Know right away that this is a buggy and inexact program made for practice purposes*****\n"
print """Input the first date then the second date.
This program (JKJKJKJKJKJK) knows the names of the months.
For leap years do not trust this program,
it's calculation does not include leap years.
As of now this program views each month as ~30.
So it can have a ~5 day error to it.
"""
print "\n"

print """Jan : 1
Feb: 2
March: 3
Apr: 4
May: 5
June: 6
July: 7
Aug: 8
Sept: 9
Oct: 10
Nov: 11
Dec: 12
\n
"""

#Gets all the numbers for the Start Date
print "Starting Date: "
day1 = int(raw_input("Day: "))
month1 = int(raw_input("Month: "))
year1 = int(raw_input("Year: "))

#Gets all the numbers for the End Date
print "\nEnding Date: "
day2 = int(raw_input("Day: "))
month2 = int(raw_input("Month: "))
year2 = int(raw_input("Year: "))

#Gets all of the months in days and adds it with the days. IE for 23/9/1994 it will do ((9-1)*30) + 23
if month1 == 1:
    startd = 0
elif month1 > 1:
    startd = (month1 - 1) * 30 

startd = startd + day1

#Same as above but for the end date.
if month2 == 1:
    enddate = 0
elif month2 > 1:
    enddate = (month2 - 1) * 30

enddate = enddate + day2

#Gets the difference between the start and end date. IE from 1994 to 2011 it will do (2011-1994) * 365
endyear = year2 - year1

endyeardays = endyear * 365


#
if startd < enddate:
    daysdiff = enddate - startd
## If for example the first date is 10/30/11 (300 days intotheyear) and second is 2/25/12 (55 days intotheyear) then you would want to
## go 365 - (300-55) = 120 which is pretty accurate. On paper I got 115. This is because there is a ~5 day error atm since each month is 30days
elif day1 > day2:
    daystosub = day1 - day2
    daysdiff = 365 - daystosub
#

exactdays = (endyear * 365) + daysdiff

if daysdiff > 30:
    monthsdiff = daysdiff / 30
    daysdiff = daysdiff - monthsdiff
elif daysdiff < 30:
    monthsdiff = 0


print "\nFrom the starting date to the ending date: \n"
print "It has been %r days, %r months, and %r years" % (daysdiff, monthsdiff, endyear)
print "In only days it has been: %r days" % exactdays
