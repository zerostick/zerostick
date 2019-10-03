#!/bin/bash

BASEPATH=$(realpath $(dirname $0)/..)
mkdir -p "${BASEPATH}/build/bin"

echo Building latest version of rclone for the Pi Zero
if [ ! -f "${BASEPATH}/build/bin/rclone" ]
then
    cd "${BASEPATH}/build"
    git clone https://github.com/ncw/rclone.git
    cd rclone
    GOOS=linux GOARM=6 GOARCH=arm go build
    if [ $? == 0 ]
    then
        echo RClone build OK
        mv rclone "${BASEPATH}/build/bin/rclone"
    fi
else
    echo rclone already build, skipping
fi
