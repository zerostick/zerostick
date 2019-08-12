#!/bin/bash -e

install -v -o 0 -g 0 -d "${ROOTFS_DIR}/opt/zerostick"
install -v -o 0 -g 0 -m 755 bin/zerostick "${ROOTFS_DIR}/opt/zerostick/"
install -v -o 0 -g 0 -m 755 bin/rclone "${ROOTFS_DIR}/usr/local/bin/"
install -v -o 0 -g 0 -m 755 bin/dms "${ROOTFS_DIR}/usr/local/bin/"
install -v -o 0 -g 0 -m 644 files/zerostick.service "${ROOTFS_DIR}/etc/systemd/system/"
install -v -o 0 -g 0 -d "${ROOTFS_DIR}/usr/share/plymouth/themes/pix/"
install -v -o 0 -g 0 -m 644 files/ZeroStick.png "${ROOTFS_DIR}/usr/share/plymouth/themes/pix/splash.png"

# Skip this when the vfs stuff works in zerostick UI
cp -r zerostick_web "${ROOTFS_DIR}/opt/zerostick/"

on_chroot << EOF
/usr/bin/systemctl daemon-reload
/usr/bin/systemctl enable zerostick
EOF
