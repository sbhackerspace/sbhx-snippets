# From https://github.com/Yelp/mrjob

# After you've run
#    sudo pip install mrjob
# run this:
#     python ~/sbhx/sbhx-snippets/py/mrjob-example.py ~/.bashrc -r emr > counts.txt
# to run WordCount on your .bashrc file and redirect the results into 'counts.txt'

"""The classic MapReduce job: count the frequency of words.
"""
from mrjob.job import MRJob
import re

WORD_RE = re.compile(r"[\w']+")


class MRWordFreqCount(MRJob):

    def mapper(self, _, line):
        for word in WORD_RE.findall(line):
            yield (word.lower(), 1)

    def combiner(self, word, counts):
        yield (word, sum(counts))

    def reducer(self, word, counts):
        yield (word, sum(counts))


if __name__ == '__main__':
    MRWordFreqCount.run()
