locals {
  motion_config_source_path      = "raspberrypi/motion.tmpl"
  motion_config_destination_path = "/etc/motion/motion.conf"
  sa_keys_destination_path       = "/home/pi/keys.json"
  init_script_source_path        = "raspberrypi/init.tmpl"
  init_script_destination_path   = "/home/pi/init.sh"
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
# Generate service account keys for the Raspberry Pi
#
resource "google_service_account_key" "raspberrypi" {
  service_account_id = google_service_account.raspberrypi.name
}

#
# Allow the Raspberry Pi service account to write to the bucket
#
resource "google_storage_bucket_iam_binding" "raspberry_pi" {
  bucket = google_storage_bucket.uploads.name
  role   = "roles/storage.objectCreator"
  members = [
    "serviceAccount:${google_service_account.raspberrypi.email}",
  ]
}

#
# Copy service account keys to the Raspberry Pi
#
resource "null_resource" "raspberrypi_keys" {
  triggers = {
    timestamp = timestamp()
  }

  provisioner "file" {
    content     = base64decode(google_service_account_key.raspberrypi.private_key)
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
# Copy motion configuration to the Raspberry Pi
#
resource "null_resource" "raspberrypi_motion_config" {
  triggers = {
    timestamp = timestamp()
  }

  provisioner "file" {
    content     = templatefile(local.motion_config_source_path, { bucket_name = google_storage_bucket.uploads.url })
    destination = local.motion_config_destination_path

    connection {
      type     = "ssh"
      user     = var.raspberry_pi_user
      password = var.raspberry_pi_password
      host     = var.raspberry_pi_host
    }
  }
}

#
# Install the gcloud SDK on the Raspberry Pi
#
resource "null_resource" "raspberrypi_init_script" {
  depends_on = [null_resource.raspberrypi_motion_config]

  triggers = {
    timestamp = timestamp()
  }

  provisioner "file" {
    content = templatefile(local.init_script_source_path, {
      email              = google_service_account.raspberrypi.email
      key_file           = local.sa_keys_destination_path
      motion_config_path = local.motion_config_destination_path
    })
    destination = local.init_script_destination_path

    connection {
      type     = "ssh"
      user     = var.raspberry_pi_user
      password = var.raspberry_pi_password
      host     = var.raspberry_pi_host
    }
  }
}

#
# Run init.sh script
#
resource "null_resource" "raspberrypi_init" {
  depends_on = [null_resource.raspberrypi_init_script]

  triggers = {
    timestamp = timestamp()
  }

  provisioner "remote-exec" {
    inline = [
      "sudo chmod +x ${local.init_script_destination_path}",
      "sudo bash ${local.init_script_destination_path}"
    ]

    connection {
      type     = "ssh"
      user     = var.raspberry_pi_user
      password = var.raspberry_pi_password
      host     = var.raspberry_pi_host
    }
  }
}