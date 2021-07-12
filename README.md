# whopooped
Litterbox monitoring using the Raspberry Pi & Camera, Motion, React Native, Firebase.

## Introduction
This IoT project is designed to monitor litterbox usage by my cats Ake & Runa, and to give me something to work on while sharpening my IoT, Google Cloud, AI and React Native skills.

## Hardware
- `Raspberry Pi & Camera`
- `Zero View camera mount`: https://thepihut.com/products/zeroview
- `Raspberry Pi Zero Camera Adapter`: https://thepihut.com/products/raspberry-pi-zero-camera-adapter

## Raspberry Pi Setup
These steps are required to get started with the project:
- Install Raspberry Pi OS on the Raspberry Pi. See https://www.raspberrypi.org/software/operating-systems/
- For local development purposes, the default `ssh` username and password can be left as they are
- Enable the Raspberry Pi Camera by changing the correct settings after running `sudo raspi-config`
- Install `motion` with `sudo apt-get install motion`. See https://motion-project.github.io/index.html for further information
- Copy the `motion.conf` file included in this repository in the `raspberrypi` directory to `/etc/motion/motion.conf`

## Architecture
This project makes use of the following:
- `Firebase`: for providing functionality to the mobile app and to process incoming images from the Raspberry Pi
- `Terraform`: to setup the Firebase resources
- `React Native`: to develop a multi-platform mobile app for managing the system

## Terraform
The `terraform` directory contains the infrastructure configuration for the projects. To make things easier, I have created the Google Cloud project and the bucket that contains the state manually.

## Makefile
The Makefile contains the following commands to make development easier:
- `make copy-motion-config`: Copies the local `motion.conf` file in the `raspberrypi` directory of the repository to `/etc/motion/motion.conf` on the Raspberry Pi.