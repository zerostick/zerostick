#!/bin/bash -e

install -v  -o 0 -g 0 -m 755 scripts/create_zero_data_partition.sh  "${ROOTFS_DIR}/usr/local/bin/"
install -v  -o 0 -g 0 -m 755 scripts/rc.local  "${ROOTFS_DIR}/etc/"
