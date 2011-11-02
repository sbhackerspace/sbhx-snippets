#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.07.16

text = open('gettysburg.txt', 'r').read()
answers = []

for ndx in range(len(text)):
    for rest in range( len(text)-ndx ):
        if len(text[ndx:rest+1+ndx]) > 0:
            answers.append( text[ndx:rest+1+ndx] )

print sorted([x for x in answers if x == x[::-1] ], None, lambda x: len(x))[-1]
