#!/bin/bash
# Steve Phillips / elimisteve
# 2014.07.14

wget https://www.tarsnap.com/download/tarsnap-autoconf-1.0.35.tgz
wget https://www.tarsnap.com/download/tarsnap-sigs-1.0.35.asc
wget https://www.tarsnap.com/tarsnap-signing-key.asc

sudo apt-get update
sudo apt-get install hashalot libssl-dev zlib1g-dev e2fslibs-dev

gpg --import < tarsnap-signing-key.asc
gpg --decrypt tarsnap-sigs-1.0.35.asc
sha256 tarsnap-autoconf-1.0.35.tgz
tar xvf tarsnap-autoconf-1.0.35.tgz 
cd tarsnap-autoconf-1.0.35/
./configure 
make all install clean
sudo make all install clean
