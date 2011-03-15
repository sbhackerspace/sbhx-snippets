#!/bin/bash

#This script toggles the gui on and off.
#Tested in ubuntu 10.10
#####USE AT YOUR OWN RISK!!!#####

#check if gdm.conf file exists and rename it.
if [ -f /etc/init/gdm.conf ]; then
	sudo mv /etc/init/gdm.conf /etc/init/gdm.bak; 
	echo "GUI Disabled";
	sudo service gdm stop | less > /dev/null 2>&1;
else
	sudo mv /etc/init/gdm.bak /etc/init/gdm.conf; 
	echo "GUI Enabled";
	sudo service gdm start
exit

fi
