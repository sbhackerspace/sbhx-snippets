
import feedparser

# This is using feedparser to parse the xml from HN's raw RSS feed
urll = "http://news.ycombinator.com/rss"
ycomb = feedparser.parse(urll)

# A list of commonwords to rule out. Will use this later in the code.
commonwords = ['all', 'go', 'its', 'had', 'day', 'to', 'has', 'do', 'them', 'his', 'get', 'they', 'not', 'now', 'him', 'like', 'did', 'these', 'she', 'each', 'people', 'some', 'see', 'are', 'out', 'what', 'said', 'for', 'find', 'may', 'be', 'we', 'were', 'water', 'come', 'by', 'on', 'about', 'her', 'of', 'could', 'or', 'first', 'into', 'number', 'one', 'down', 'long', 'your', 'use', 'from', 'would', 'there', 'two', 'been', 'their', 'call', 'way', 'was', 'more', 'that', 'but', 'part', 'with', 'than', 'he', 'made', 'word', 'look', 'this', 'up', 'will', 'can', 'many', 'my', 'and', 'then', 'is', 'it', 'an', 'as', 'at', 'have', 'in', 'if', 'no', 'make', 'when', 'write', 'how', 'other', 'which', 'you', 'oil', 'I', 'who', 'a', 'so', 'time', 'the', "you'll", 'our', "can't", "&", "am", 'give', "back", "why", 'only', 'too']


ycli = []
ycli = ycomb.entries #This is putting all of the 'entries' part of the xml code into a list.
ycombtitles = []
hntitlessplit = []

n = len(ycli) #Counting the number of entries
n = n - 1
x = 0

#This is grabbing only the titles of each entry
while n > x: 
        ycti = ycomb.entries[x].title
        ycombtitles.append(ycti)
        x = x + 1

#This is splitting each title seperating all the words into one long list.
for title in ycombtitles:
    for word in title.split(' '):
        hntitlessplit.append(word)

##for x in range(len(hntitlessplit)):
##    hntitlessplit = hntitlessplit[x].strip('?')


#This deletes each word in the list that is in the 'commonwords' list at the beginning.
hntitlessplit = [x for x in hntitlessplit if x.lower() not in commonwords]

print '\n\n'    
print hntitlessplit
print '\n\n'

