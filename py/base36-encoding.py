#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2012.02.02

# From http://en.wikipedia.org/wiki/Base_36#Python_implementation

def base36encode(number, alphabet='0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ'):
    """Convert positive integer to a base36 string."""
    if not isinstance(number, (int, long)):
        raise TypeError('number must be an integer')
 
    # Special case for zero
    if number == 0:
        return alphabet[0]
 
    base36 = ''
 
    sign = ''
    if number < 0:
        sign = '-'
        number = - number
 
    while number != 0:
        number, i = divmod(number, len(alphabet))
        base36 = alphabet[i] + base36
 
    return sign + base36
 
def base36decode(number):
    return int(number, 36)
 

print base36encode(3111111111)
print base36decode('1FG9WP3')
