locals {
  motion_config_source_path      = "raspberrypi/motion.conf"
  motion_config_destination_path = "/etc/motion/motion.conf"
  sa_keys_source_path            = "raspberrypi/keys.json"
  sa_keys_destination_path       = "/home/pi/keys.json"
}

#
# Create the Raspberry Pi service account
#
resource "google_service_account" "raspberrypi" {
  project      = data.google_project.this.project_id
  account_id   = "raspberry-pi"
  display_name = "Raspberry Pi"
}

#
# Install the gcloud SDK on the Raspberry Pi
#
resource "null_resource" "raspberrypi_gcloud" {
  provisioner "remote-exec" {
    inline = [
      "sudo echo \"deb [signed-by=/usr/share/keyrings/cloud.google.gpg] http://packages.cloud.google.com/apt cloud-sdk main\" | sudo tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && sudo curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key --keyring /usr/share/keyrings/cloud.google.gpg  add - && sudo apt-get update -y && sudo apt-get install google-cloud-sdk -y"
    ]

    connection {
      type     = "ssh"
      user     = var.raspberry_pi_user
      password = var.raspberry_pi_password
      host     = var.raspberry_pi_host
    }
  }
}

#
# Generate service account keys for the Raspberry Pi
#
resource "null_resource" "raspberrypi_sa_keys" {
  provisioner "local-exec" {
    command = "gcloud iam service-accounts keys create ${local.sa_keys_source_path} --iam-account=${google_service_account.raspberrypi.email}"
  }
}

#
# Copy service account keys to the Raspberry Pi
#
resource "null_resource" "raspberrypi_copy_sa_keys" {
  provisioner "file" {
    source      = local.sa_keys_source_path
    destination = local.sa_keys_destination_path

    connection {
      type     = "ssh"
      user     = var.raspberry_pi_user
      password = var.raspberry_pi_password
      host     = var.raspberry_pi_host
    }
  }
}

#
# Generate service account keys for the Raspberry Pi
#
resource "null_resource" "delete_raspberrypi_sa_keys" {
  depends_on = [null_resource.raspberrypi_copy_sa_keys]
  provisioner "local-exec" {
    command = "rm ${local.sa_keys_source_path}"
  }
}

#
# Authenticate the Raspberry Pi
#
resource "null_resource" "raspberrypi_sa_auth" {
  provisioner "remote-exec" {
    inline = [
      "gcloud auth activate-service-account ${google_service_account.raspberrypi.email} --key-file=${local.sa_keys_destination_path}",
    ]

    connection {
      type     = "ssh"
      user     = var.raspberry_pi_user
      password = var.raspberry_pi_password
      host     = var.raspberry_pi_host
    }
  }
}

#
# Install motion on the Raspberry Pi and set permissions
#
resource "null_resource" "raspberrypi_motion" {
  provisioner "remote-exec" {
    inline = [
      "sudo apt install -y motion",
      "sudo chown pi ${local.motion_config_destination_path}",
      "sudo chown pi /var/log/motion",
      "sudo chown pi /var/lib/motion",
    ]

    connection {
      type     = "ssh"
      user     = var.raspberry_pi_user
      password = var.raspberry_pi_password
      host     = var.raspberry_pi_host
    }
  }
}

#
# Copy motion configuration to the Raspberry Pi
#
resource "null_resource" "raspberrypi_motion_config" {
  provisioner "file" {
    source      = local.motion_config_source_path
    destination = local.motion_config_destination_path

    connection {
      type     = "ssh"
      user     = var.raspberry_pi_user
      password = var.raspberry_pi_password
      host     = var.raspberry_pi_host
    }
  }
}