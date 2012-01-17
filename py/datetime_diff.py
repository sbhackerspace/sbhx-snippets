#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2012.01.16

import datetime

now              = datetime.datetime.now()
tomorrow_at_1215 = datetime.datetime(now.year, now.month, now.day+1, 12, 15)

diff_in_seconds  = (tomorrow_at_1215 - now).seconds
diff_in_hours    = diff_in_seconds / 60. / 60. # 60 secs/min, 60 mins/hour

print "Tomorrow at 1215 is %0.2f hours away" % (diff_in_hours,)

hours = int(diff_in_hours)
mins  = int( 60 * (diff_in_hours - hours) ) # multiply fraction hour by 60

print "Also known as %d hours and %d minutes" % (hours, mins)
