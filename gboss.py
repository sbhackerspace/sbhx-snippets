#!/usr/bin/env python
# Steve Phillips / elimisteve
# 2011.10.22

import paramiko

#cmd = "sudo /etc/init.d/mediatomb restart"
cmd = "ls /etc"

ssh = paramiko.SSHClient()
ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
ssh.connect('microwave', username='matt') 
stdin, stdout, stderr = ssh.exec_command(cmd)
print stdout.readlines()
ssh.close()
