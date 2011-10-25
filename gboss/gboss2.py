#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.10.22

import paramiko
import os
privatekeyfile = os.path.expanduser('~/.ssh/id_rsa.pub')
mykey = paramiko.RSAKey.from_private_key_file(privatekeyfile)
ssh.connect('sbhx.org', username = 'steve', pkey = mykey)

stdin, stdout, stderr = ssh.exec_command('df -h')
print stdout.readlines()
ssh.close()
