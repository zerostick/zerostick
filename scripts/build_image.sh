#!/bin/bash
# Builds the image to put on the Pi
# Requires docker and docker-compose

BASEPATH=$(realpath $(dirname $0)/..)
echo Building the image for the ZeroStick

# Cache for pi-gen
mkdir -p ${BASEPATH}/cache
cd ${BASEPATH}/cache
if [ -d pi-gen.git ]
then
  echo Updating pi-gen.git
  cd pi-gen.git
  git fetch
else
  git clone https://github.com/RPi-Distro/pi-gen.git --bare --mirror
fi

mkdir -p ${BASEPATH}/build/
cd ${BASEPATH}/build/
rm -rf pi-gen
git clone ${BASEPATH}/cache/pi-gen.git
cd pi-gen

# Fix build-docker.sh until PR is merged upstream
git cherry-pick cccce273fbb8e4d3c1a0f0fadf99a99f93b5157d

# Start apt-cacher in case we need to rerun this build
docker-compose up -d

cp -r ${BASEPATH}/scripts/pi-gen-stage_zerostick stage_zerostick
cp stage2/prerun.sh stage_zerostick/prerun.sh
touch stage_zerostick/EXPORT_IMAGE
cp ${BASEPATH}/scripts/pi-gen-config config

time ./build-docker.sh

exit 0

# Stop and clean docker
docker-compose stop
