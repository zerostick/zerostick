#!/bin/bash

install -v -o 0 -g 0 -d "${ROOTFS_DIR}/opt/zerostick/templates"
install -v -o 0 -g 0 -m 644 files/hostapd.conf "${ROOTFS_DIR}/opt/zerostick/templates/"
install -v -o 0 -g 0 -m 755 files/hostapd_render.sh 
install -v -o 0 -g 0 -m 644 files/dnsmasq.conf "${ROOTFS_DIR}/etc/"

on_chroot << EOF
# Maybe not the right way to do it, but make sure that there is no /etc/wpa_supplicant/wpa_supplicant.conf
# but instead a /etc/wpa_supplicant/wpa_supplicant-wlan0.conf file, so we only try to connect via wlan0 interface.
# There might be a wpa_supplicant.conf file there if ZeroStick is build with params to setup wifi in the first place.
# if [ -f /etc/wpa_supplicant/wpa_supplicant.conf ]
# then
#   mv /etc/wpa_supplicant/wpa_supplicant.conf /etc/wpa_supplicant/wpa_supplicant-wlan0.conf
# fi
/usr/bin/systemctl unmask hostapd.service
/usr/bin/systemctl enable hostapd.service
/usr/bin/systemctl enable dnsmasq.service
/usr/bin/systemctl disable wpa_supplicant
EOF

