#!/bin/bash
# Builds the image to put on the Pi
# Requires docker and docker-compose

BASEPATH=$(realpath $(dirname $0)/..)
echo Building the image for the ZeroStick

# Cache for pi-gen
mkdir -p "${BASEPATH}/cache"
cd "${BASEPATH}/cache"
if [ -d pi-gen.git ]
then
  echo Updating pi-gen.git
  cd pi-gen.git
  git fetch
else
  git clone https://github.com/RPi-Distro/pi-gen.git --bare --mirror
fi

mkdir -p "${BASEPATH}/build/"
cd "${BASEPATH}/build/"
rm -rf pi-gen
git clone "${BASEPATH}/cache/pi-gen.git"
cd pi-gen

# Start apt-cacher in case we need to rerun this build
docker-compose up -d

cp -r "${BASEPATH}/scripts/pi-gen-stage_zerostick" stage_zerostick
cp -r "${BASEPATH}/build/bin" stage_zerostick/03-binaries/
cp stage2/prerun.sh stage_zerostick/prerun.sh
rm -f stage2/EXPORT_IMAGE stage2/EXPORT_NOOBS
cp stage4/EXPORT_IMAGE stage_zerostick/EXPORT_IMAGE
cp "${BASEPATH}/scripts/pi-gen-config" config

# Skip this when the vfs stuff works in zerostick UI
cp -r "${BASEPATH}/zerostick_web" stage_zerostick/03-binaries/

time ./build-docker.sh

exit 0

# Stop and clean docker
docker-compose stop
