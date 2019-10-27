#!/bin/bash

while [ ! -f /tmp/shutdown_keep_net_alive ]; do
    # Ping google DNS to see if we can get online via wifi
    ping -q -I wlan0 -c 2 8.8.8.8 2>&1 > /dev/null
    if [ $? -eq 0 ]; then
        # We are connected via wlan0, downgrade usb0
	    ifmetric usb0 1000
    else
        # We are NOT connected via wlan0, upgrade usb0
	    ifmetric usb0 100
    fi
    sleep 120 # Sleep 2 minutes before checking again
done
