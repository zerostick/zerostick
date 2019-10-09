#!/bin/bash
# Mount the partitions read-only, after a fsck check and repair.

function prep_probe_and_mount() {
    PART=$1
    mkdir -p /${PART}

    # Probe the partitions, check and mount
    if [ -L /dev/zeroVG/${PART} ]
    then
        logger -t ZeroStick -s Probing partitions on /dev/zeroVG/${PART}
        partprobe /dev/zeroVG/${PART}
        logger -t ZeroStick -s Running fsck on /dev/mapper/zeroVG-${PART}1
        fsck.msdos -a /dev/mapper/zeroVG-${PART}1
        logger -t ZeroStick -s Mounting /dev/mapper/zeroVG-${PART}1 to /${PART}
        mount -o ro -t vfat /dev/mapper/zeroVG-${PART}1 /${PART}

    fi

}

prep_probe_and_mount cam
prep_probe_and_mount music
