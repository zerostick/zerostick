#!/bin/bash -eu

DISK_DEV=/dev/mmcblk0

# Does zeroVG exist?
if [ ! -h /dev/zeroVG/cam ] || [ ! -h /dev/zeroVG/music ]
then # Create partitioning

    # suck in disk layout
    sfdisk -d $DISK_DEV > /tmp/partitions
    #ORIGINAL_DISK_IDENTIFIER=$( cat /tmp/partitions | grep label-id | sed 's/^.*0x//g' )

    # Add a 3rd partition from the free space
    parted -m /dev/mmcblk0 unit s print free | grep "free;" | tail -n 1 | sed -e 's/s//g' | awk -F':' '{ print "/dev/mmcblk0p3 : start= " $2 ", size= " $3 ", type=8e" }' >> /tmp/partitions
    sfdisk --no-reread --force $DISK_DEV < /tmp/partitions

    #NEW_DISK_IDENTIFIER=$( sfdisk -d $DISK_DEV | grep label-id | sed 's/^.*0x//g' )

    # Fix booting
    #sed -i "s/${ORIGINAL_DISK_IDENTIFIER}/${NEW_DISK_IDENTIFIER}/g" /etc/fstab
    #sed -i "s/${ORIGINAL_DISK_IDENTIFIER}/${NEW_DISK_IDENTIFIER}/" /boot/cmdline.txt

    # Reread partition table
    partprobe $DISK_DEV

    # Setup LVM
    #apt update -y
    #apt-get install -y lvm2
    pvcreate /dev/mmcblk0p3
    vgcreate zeroVG /dev/mmcblk0p3

    vfree=$(vgs --rows --units b | grep VFree | awk '{print $2}' | sed 's/B//')
    # Use 50% for now for cam
    CAM_SIZE=$(( $vfree / 2))
    lvcreate -L ${CAM_SIZE}b -n cam zeroVG
    # The rest for music
    MUSIC_SIZE=$( vgs --rows --units b | grep VFree | awk '{print $2}' | sed 's/B//')
    lvcreate -L ${MUSIC_SIZE}b -n music zeroVG

    # Partition the LVs again
    echo ",,c;" | sfdisk --no-reread --force /dev/zeroVG/cam
    echo ",,c;" | sfdisk --no-reread --force /dev/zeroVG/music
    kpartx -av /dev/zeroVG/cam
    kpartx -av /dev/zeroVG/music

    logger "Formatting new partitions..."
    mkfs.vfat -n ZEROMUSIC /dev/mapper/zeroVG-music1
    mkfs.vfat -n ZEROCAM   /dev/mapper/zeroVG-cam1

    # Create mountpoints
    mkdir -p /cam /music


    # Create TeslaCam folder
    mount /dev/mapper/zeroVG-cam1 /cam
    mkdir -p /cam/TeslaCam
    umount /cam
    
    # Create g_mass_storage config
    echo "options g_mass_storage file=/dev/mapper/zeroVG-cam,/dev/mapper/zeroVG-music removable=1,1 ro=0,0 stall=0 iSerialNumber=123456" > /etc/modprobe.d/g_mass_storage.conf
    echo "g_mass_storage" > /etc/modules-load.d/g_mass_storage.conf

fi
