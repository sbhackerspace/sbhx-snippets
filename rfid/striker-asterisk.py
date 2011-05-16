#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.05.11

from pylibftdi import Device
import asterisk.manager, binascii, time
#'\x01\xe7\x6e'
USER_DATA = {
    '\x01\xe7l\x02': {'nick': 'swiss',       'song': 'custom/youngbuk2'},
    '\x01\xe7_Y':    {'nick': 'Joule_Thief', 'song': 'custom/smokeonwater'},
    '\x01\xe7ma':    {'nick': 'elimisteve',  'song': 'custom/loop'},
    '\x01\xe7f\x8d': {'nick': 'luckyaba',    'song': 'custom/goodfoot'}
}

KEY_LENGTH = 6

dev = Device(mode='b')
dev.baudrate = 9600

while True:
    string = dev.read(KEY_LENGTH).strip()
    old_string = '' # Not yet used to compare string and old_string

    if string != "":
        if string in USER_DATA:
            print "Hex:", binascii.b2a_hex(string)
            dev.write('U')

            # Play ringtone/theme song
            manager = asterisk.manager.Manager()
            manager.connect('192.168.1.200')
            manager.login('steve', 'amp111')

            # manager.originate('', '')
            data = {"Action": "Originate",
                    "Channel": "SIP/180",
                    "Application": "Playback",
                    "Data": USER_DATA[string]['song']
                    }

            manager.send_action(data)
            response = manager.status()
            #print response
            manager.logoff()
            old_string = string

            time.sleep(4)
            dev.flush()
        else:
            print binascii.hexlify(string), "DENIED"
