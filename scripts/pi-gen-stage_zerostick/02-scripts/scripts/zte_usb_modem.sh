#!/bin/bash
# Support for (at least) ZTE MF823 USB LTE Modem
# Reference material: https://wiki.archlinux.org/index.php/ZTE_MF_823_(Megafon_M100-3)_4G_Modem

# Connect to network:
# curl --header "Referer: http://192.168.0.1/index.html" http://192.168.0.1/goform/goform_set_cmd_process?goformId=CONNECT_NETWORK

# Detect ZTE WCDMA thing
lsusb | grep -q "ZTE WCDMA"
if [ $? -eq 0 ]; then
    echo ZTE WCDMA Modem on board. Starting device.
    # Connect to network at boot
    curl --header "Referer: http://192.168.0.1/index.html" "http://192.168.0.1/goform/goform_set_cmd_process?goformId=SET_CONNECTION_MODE&ConnectionMode=auto_dial"

    # Set the interface metric on usb0 to be higher that wlan0 (Making wlan0 preferred)
    ifmetric usb0 1000
fi
