#!/bin/bash -e

install -v -o 0 -g 0 -m 755 -d "${ROOTFS_DIR}/opt/zerostick/boot_scripts"
install -v -o 0 -g 0 -m 755 scripts/create_zero_data_partition.sh  "${ROOTFS_DIR}/opt/zerostick/boot_scripts/"
install -v -o 0 -g 0 -m 755 scripts/mount_partitions.sh "${ROOTFS_DIR}/opt/zerostick/boot_scripts/"
install -v -o 0 -g 0 -m 755 scripts/zte_usb_modem.sh "${ROOTFS_DIR}/opt/zerostick/boot_scripts/"
install -v -o 0 -g 0 -m 755 scripts/rc.local  "${ROOTFS_DIR}/etc/"

install -v -o 0 -g 0 -m 755 -d "${ROOTFS_DIR}/opt/zerostick/scripts"
install -v -o 0 -g 0 -m 755 scripts/keep_net_alive.sh "${ROOTFS_DIR}/opt/zerostick/scripts/"
install -v -o 0 -g 0 -m 755 services/keep-net-alive.service "${ROOTFS_DIR}/etc/systemd/system/"

on_chroot << EOF
/usr/bin/systemctl daemon-reload
/usr/bin/systemctl enable keep-net-alive
EOF