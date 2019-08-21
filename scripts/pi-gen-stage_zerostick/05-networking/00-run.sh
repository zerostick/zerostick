#!/bin/bash

install -v -o 0 -g 0 -d "${ROOTFS_DIR}/opt/zerostick/templates"
install -v -o 0 -g 0 -m 644 files/hostapd.conf "${ROOTFS_DIR}/opt/zerostick/templates/"
install -v -o 0 -g 0 -m 755 files/hostapd_render.sh "${ROOTFS_DIR}/opt/zerostick/boot_scripts/"
install -v -o 0 -g 0 -m 644 files/dnsmasq.conf "${ROOTFS_DIR}/etc/"

on_chroot << EOF
/usr/bin/systemctl unmask hostapd.service
/usr/bin/systemctl enable hostapd.service
/usr/bin/systemctl enable dnsmasq.service
/usr/bin/systemctl disable dhcpcd
EOF
