#!/bin/bash
# Builds the image to put on the Pi
# Requires docker and docker-compose

BASEPATH=$(realpath $(dirname $0)/..)
echo Building the image for the ZeroStick

# Cache for pi-gen
cd ${BASEPATH}/cache
if [ -d pi-gen.git ]
then
  echo Updating pi-gen.git
  cd pi-gen.git
  git fetch
else
  git clone https://github.com/RPi-Distro/pi-gen.git --bare
fi

mkdir -p ${BASEPATH}/build/
cd ${BASEPATH}/build/
rm -rf pi-gen
git clone ${BASEPATH}/cache/pi-gen.git
cd pi-gen

# Start apt-cacher in case we need to rerun this build
docker-compose up -d 

# Remove X11, desktop etc from image
rm -rf stage[345] stage2/EXPORT_NOOBS # stage2/EXPORT_IMAGE

cp -r ${BASEPATH}/scripts/pi-gen-stage7 stage7
cp -r ${BASEPATH}/scripts/pi-gen-config config
./build-docker.sh

exit 0

# Stop and clean docker
docker-compose stop
