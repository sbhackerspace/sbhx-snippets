#!/usr/bin/python -tt
#
# Copyright 2014 Garrett D Holmstrom
#
# This program is free software; you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation; either version 3 of the License or (at
# your option) any later version accepted by the membership of the Santa
# Barbara Hackerspace (or its successor approved by the membership of
# the Santa Barbara Hackerspace), which shall act as a proxy as defined
# in Section 14 of version 3 of the license.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
# General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program. If not, see http://www.gnu.org/licenses/.

import base64
import binascii
import getpass
import hashlib
import random


__all__ = ('generate_ntlm_hash', 'generate_ssha_hash')


def generate_ntlm_hash(password):
    """
    Generate the NTLM hash of a password.

    Note that NTLM is generally regarded as insecure.  Nonetheless,
    large amounts of equipment still requires it, so it is included
    to help facilitate interoperability.
    """
    md4 = hashlib.new('md4')
    md4.update(password.encode('utf-16le'))
    return binascii.hexlify(md4.digest()).upper()


def generate_ssha_hash(password):
    """
    Generate a salted SHA1 hash of a password.
    """
    salt = base64.b64encode(bytes(random.getrandbits(8)))
    sha = hashlib.sha1()
    sha.update(password)
    sha.update(salt)
    return base64.b64encode(sha.digest() + salt)


def _main():
    """
    Prompt for a password and print its NTLM and SSHA hashes.
    """
    password = getpass.getpass()
    print 'NTLM:', generate_ntlm_hash(password)
    print 'SSHA:', generate_ssha_hash(password)


if __name__ == '__main__':
    _main()
