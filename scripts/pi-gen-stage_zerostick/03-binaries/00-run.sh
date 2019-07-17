#!/bin/sh -e

install -v -o 0 -g 0 -m 755 bin/zerostick  "${ROOTFS_DIR}/usr/local/bin/"
install -v -o 0 -g 0 -m 755 bin/rclone  "${ROOTFS_DIR}/usr/local/bin/"
install -v -o 0 -g 0 -m 755 bin/dms "${ROOTFS_DIR}/usr/local/bin/"
install -v -o 0 -g 0 -m 655 files/zerostick.service "${ROOTFS_DIR}/etc/systemd/system/"
