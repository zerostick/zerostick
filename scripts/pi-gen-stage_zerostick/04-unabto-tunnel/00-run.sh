#!/bin/bash -e

install -v -o 0 -g 0 -d "${ROOTFS_DIR}/opt/unabto-tunnel"
install -v -o 0 -g 0 -m 644 files/unabto-tunnel.service "${ROOTFS_DIR}/etc/systemd/system/unabto-tunnel.service"
install -v -o 0 -g 0 -m 644 files/zerostick.env "${ROOTFS_DIR}/etc/zerostick.env"


on_chroot << EOF

apt install -y cmake ninja-build git

cd /tmp/
git clone https://github.com/nabto/unabto.git
mkdir unabto/apps/tunnel/build
cd unabto/apps/tunnel/build
cmake -GNinja -DCMAKE_BUILD_TYPE=Release -DUNABTO_CRYPTO_MODULE=openssl_armv4 -DUNABTO_RANDOM_MODULE=openssl_armv4 ..
ninja
mv unabto_tunnel /opt/unabto-tunnel/
apt purge -y cmake ninja-build git
apt autoremove -y
# Enable service
systemctl enable unabto-tunnel
EOF
