# Started by: AJ B / avb_wkyhu
# Forked by:  Steve Phillips / elimisteve
print "\n"
print "Date to Date v2.0\n"
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

months2days = { 1 : 31 , 2 : 28, 3 : 31, 4 : 30, 5 : 31, 6 : 30, 7 : 31, 8 : 31, 9 : 30, 10 : 31, 11 : 30, 12 : 31}

##

##
##               
def date_from_user(date_name):
    """Gets a date from the user via stdin and returns a tuple in the
    form of (month, day, year)"""

    user_input = raw_input("%s Date (month/day/year): " % date_name)
    month, day, year = [int(x) for x in user_input.split('/')]

    # If "11" is entered for the year, interpret that as 2011 by adding 2000
    if year < 100:
        year += 2000

    return (month, day, year)
##


print "Input two dates and view the duration between them\n"


##
## Improvement #2: Keep it DRY ("Don't Repeat Yourself"). Remove
## duplicate code by creating a function and calling it twice.
##

# Gets Start Date and Finishing Date values
month1, day1, year1 = date_from_user("Starting")
month2, day2, year2 = date_from_user("Finishing")



#

#This is how you can get the days from starting to ending date correctly
if month1 < month2:
    daysdiff = months2days[month1] - day1 + sum([months2days[x] for x in range(month1+1,month2)]) + day2
if month1 > month2:
    daysdiff = sum([months2days[x] for x in range(month1,13)]) - day1 + sum([months2days[x] for x in range(1,month2)]) + day2


# Gets the difference between the start and end date. IE from 1994 to
# 2011 it will do (2011-1994) * 365

endyear = year2 - year1
endyeardays = endyear * 365

exactdays = (endyear * 365) + daysdiff

if daysdiff > 30:
    monthsdiff = daysdiff / 30
    daysdiff = daysdiff - (monthsdiff * 30)
elif daysdiff < 30:
    monthsdiff = 0

print "\nFrom the starting date to the ending date: \n"
print "It has been %r days, %r months, and %r years" % (daysdiff, monthsdiff, endyear)
print "In only days it has been: %r days" % exactdays
