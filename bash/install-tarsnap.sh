#!/bin/bash
# Steve Phillips / elimisteve
# 2014.07.14

if [ $# -lt 1 ]; then
    echo "`basename $0` version (e.g., 1.0.35)"
    exit 1
fi

version=$1

wget https://www.tarsnap.com/download/tarsnap-autoconf-${version}.tgz
wget https://www.tarsnap.com/download/tarsnap-sigs-${version}.asc
wget https://www.tarsnap.com/tarsnap-signing-key.asc

sudo apt-get update
sudo apt-get install hashalot libssl-dev zlib1g-dev e2fslibs-dev

gpg --import < tarsnap-signing-key.asc
gpg --decrypt tarsnap-sigs-${version}.asc
sha256 tarsnap-autoconf-${version}.tgz
tar xvf tarsnap-autoconf-${version}.tgz
cd tarsnap-autoconf-${version}/
./configure 
make all install clean
sudo make all install clean
