#!/bin/bash

# Update hostapd.conf with last 3 octets from MAC
LAST_OF_MAC=`ifconfig wlan0 | grep -ioh '[0-9A-F]\{2\}\(:[0-9A-F]\{2\}\)\{5\}' | cut -f 4,5,6 -d : | tr -d :`
cat /opt/zerostick/templates/hostapd.conf | sed -e "s/MAC/${LAST_OF_MAC}/" > /etc/hostapd/hostapd.conf


