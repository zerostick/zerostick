#!/bin/bash -e

on_chroot << EOF
/sbin/dphys-swapfile swapoff
/sbin/dphys-swapfile uninstall
apt purge -y triggerhappy dphys-swapfile logrotate
apt autoremove -y
EOF
