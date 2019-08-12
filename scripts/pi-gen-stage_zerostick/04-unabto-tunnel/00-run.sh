#!/bin/bash -e

install -v -o 0 -g 0 -d "${ROOTFS_DIR}/opt/unabto-tunnel"
install -v -o 0 -g 0 -m 644 files/unabto-tunnel.service "${ROOTFS_DIR}/opt/unabto-tunnel/"

on_chroot << EOF

apt install -y cmake ninja-build git

cd /tmp/
git clone https://github.com/nabto/unabto.git
mkdir unabto/apps/tunnel/build
cd unabto/apps/tunnel/build
perl -p -i -e 's:#define AMP_DEVICE_NAME_DEFAULT "Tunnel":#define AMP_DEVICE_NAME_DEFAULT "ZeroStick":g' ../src/main.c
perl -p -i -e 's:static const char* device_product = "uNabto Video";:static const char* device_product = "ZeroStick device";:g' ../src/main.c
perl -p -i -e 's:uNabto Video:ZeroStick device:g' ../src/main.c
cmake -GNinja -DCMAKE_BUILD_TYPE=Release -DUNABTO_CRYPTO_MODULE=openssl_armv4 -DUNABTO_RANDOM_MODULE=openssl_armv4 ..
ninja
mv unabto_tunnel /opt/unabto-tunnel/
apt purge -y cmake ninja-build git
apt autoremove -y
EOF
