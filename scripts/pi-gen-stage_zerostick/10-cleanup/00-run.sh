#!/bin/sh -e

on_chroot << EOF
/sbin/dphys-swapfile swapoff
/sbin/dphys-swapfile disable
apt purge -y triggerhappy dphys-swapfile logrotate
apt autoremove -y
EOF
