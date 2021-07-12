# whopooped
Litterbox monitoring using the Raspberry Pi & Camera, Motion, React Native, Firebase.

## Introduction
This IoT project is designed to monitor litterbox usage by my cats Ake & Runa, and to give me something to work on while sharpening my IoT, Google Cloud, AI and React Native skills.

## Raspberry Pi
I have used the following components for the Raspberry Pi:
- `Raspberry Pi & Camera`
- `Zero View camera mount`: https://thepihut.com/products/zeroview
- `Raspberry Pi Zero Camera Adapter`: https://thepihut.com/products/raspberry-pi-zero-camera-adapter

### Setup
These steps are required to get started with the project:
- Install Raspberry Pi OS on the Raspberry Pi. See https://www.raspberrypi.org/software/operating-systems/
- For local development purposes, the default `ssh` username and password can be left as they are
- Enable the Raspberry Pi Camera by changing the correct settings after running `sudo raspi-config`

### Motion 
Motion is a highly configurable program that monitors video signals from many types of cameras.

- Install `motion` with `sudo apt-get install motion`. See https://motion-project.github.io/index.html for further information
- Copy the `motion.conf` file included in this repository in the `raspberrypi` directory to `/etc/motion/motion.conf`

I have made the following changes to the included `motion.conf` file (compared to the default configuration):

- Uncomment the `mmalcam_name vc.ril.camera` parameter
- Enable and set `target_dir` to `/home/pi/Documents/motion`
- Set `ffmpeg_output_movies` to `off`
- Set `stream_localhost` to `off`
- Set `webcontrol_localhost` to `off`
- Set `width` to `640` and `height` to `480`
- Set `locate_motion_mode` to `preview`
- Set `locate_motion_style` to `redbox`
- Set `event_gap` to `10`
- Set `output_pictures` to `center`
- Set `quality` to `80`
- Set `text_changes` to `on`

## Infrastructure
The `terraform` directory contains all of the infrastructure configuration required by the project.
Make sure you already have a Google Cloud organization and billing account before going forward.

### Configure the gcloud SDK
Setup the gcloud tool in order to easily deploy changes to the infrastructure.
- Create and switch to a new `gcloud` configuration: `gcloud config configurations create whopooped`
- Authenticate with `gcloud auth login`
- Get application default credentials with `gcloud auth application-default login`
- Set the project name with `gcloud config set project <your-prefix>-whopooped` (replace `<your-prefix>`)

### Terraform
Terraform is configured to use a Google Cloud Storage bucket for saving state. Furthermore, the `owner` role is given to the organization admin for simplicity (far from being a best practice, but ok for a fun project).
The following aspects are taken care of:
- Project creation
- Setup of IAM bindings
- Billing budget and notifications
- Terraform state storage bucket

## Makefile
The Makefile contains the following commands to make development easier:
- `make copy-motion-config`: Copies the local `motion.conf` file in the `raspberrypi` directory of the repository to `/etc/motion/motion.conf` on the Raspberry Pi.

# TODO
Step 1:
- Tweak `motion.conf` and send relevant images to Google Cloud Storage
- Setup infrastructure
