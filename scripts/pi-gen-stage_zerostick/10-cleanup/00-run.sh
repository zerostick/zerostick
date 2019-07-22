#!/bin/bash -e

on_chroot << EOF
/sbin/dphys-swapfile swapoff || true
/sbin/dphys-swapfile uninstall || true
apt purge -y triggerhappy dphys-swapfile logrotate exim4-base exim4-config exim4-daemon-light cron alsa-utils
apt autoremove -y
EOF
