#!/bin/bash
# Fetch and unpack Rasbian Lite
CACHE_DIR=cache
IMAGE=rasbian_lite_latest.zip

echo Fetching Rasbian Lite image
mkdir -p $CACHE_DIR
curl -L https://downloads.raspberrypi.org/raspbian_lite_latest > $CACHE_DIR/$IMAGE
cd $CACHE_DIR
unzip $IMAGE
rm $IMAGE
echo Done fetching and unpacking Rasbian Lite image