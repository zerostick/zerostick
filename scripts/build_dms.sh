#!/bin/bash

BASEPATH=$(realpath $(dirname $0)/..)
mkdir -p "${BASEPATH}/build/bin"

echo Building latest version of dms for the Pi Zero
if [ ! -f "${BASEPATH}/build/bin/dms" ]
then
    cd "${BASEPATH}/build"
    git clone https://github.com/anacrolix/dms.git
    cd dms
    GOOS=linux GOARM=6 GOARCH=arm go build
    if [ $? == 0 ]
    then
        echo dms build OK
        mv dms "${BASEPATH}/build/bin/"
    fi
else
    echo dms already build, skipping
fi
