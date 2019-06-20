#!/bin/bash

echo Building latest version of rclone for the Pi Zero
if [ ! -f ../build/rclone ]
then
    git clone https://github.com/ncw/rclone.git
    cd rclone
    GOOS=linux GOARM=5 GOARCH=arm go build
    if [ $? == 0 ]
    then
        echo RClone build OK
        mv rclone ../../build/rclone
    fi
    cd ..
    rm -rf rclone
else
    echo rclone already build, skipping
fi
