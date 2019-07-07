# Tesla Cam automation with a Raspberry Pi Zero W

## Description

Ideas:

- Use a Raspberry Pi Zero W (USB powered device) to act as USB storage for a Tesla.
- Use the Pi to serve a web interface accessible from the in car browser.
- Let the user setup sync of video clips stored by the car to external services (OneDrive, home local NAS/PC device)
- Reverse syncing of music to the Pi, for in car usage.
- Let the setup process be a painless as possible - There are users without Linux experience out there.
- Provide streamable access to the video clips via WiFi on the move; Create WiFi in Ad-Hoc AP mode, web interface to display/delete/archive videos.

## Inspiration

There are solutions out there that does this, but they have little traction, because they either are too complicated for non Linux users or, in the case where there has been an attempt to create some UI, aren't as functional. My wish is to combine those.

This could potentially also be used as a music USB storage for other brands of cars, accessing the UI from a smartphone to manage it.

## How it will work/How it will be build

Stuff needed:

- Raspberry Pi Zero W. This specific model of the Pi has support for USB OTG (On-The-Go, https://en.wikipedia.org/wiki/USB_On-The-Go ) and WiFi, the other models have a USB HUB, which disabled this functionality.
- A USB-A to USB-Micro cable
- A Micro-SD card (32 GB or larger, good quality, fast to be able to handle 3 simultaneous streams) Ie SanDisk Extreme Micro/SDXC A2/U3/V30 256GB
- A way to attach a Micro-SD card to your Mac/PC
- Optionally: A Tesla

### Setup

- Download and install [Raspbian Stretch Lite](https://downloads.raspberrypi.org/raspbian_lite_latest) on the SD card using Rasbians instructions (That is using [Etcher](https://www.balena.io/etcher/)). `./fetch_os_image.sh` will help you get the latest and unpacking it into the `cache` directory.
- Configure WiFi (Instructions later)
- Install the software from here (Instructions later)
- Insert SD card into the Raspberry Pi
- Attach Pi to your car.


## Developing/building the zerostick

### Setup, once:

```
make setup
```

### Build for your own maschine

```
make
```

### To run it

```
make run
```

### Create image for the Pi

```
make image
```
