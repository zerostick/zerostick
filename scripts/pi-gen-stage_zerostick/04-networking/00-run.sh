#!/bin/sh -e

install -v -o 0 -g 0 -m 644 files/90-wifi-devices.rules "${ROOTFS_DIR}/etc/udev/rules.d/"
install -v -o 0 -g 0 -m 644 files/dnsmasq.conf "${ROOTFS_DIR}/etc/"
install -v -o 0 -g 0 -m 644 files/hostapd.conf "${ROOTFS_DIR}/etc/hostapd/"

on_chroot << EOF
/usr/bin/systemctl unmask hostapd.service
/usr/bin/systemctl enable hostapd.service
/usr/bin/systemctl enable dnsmasq.service
EOF