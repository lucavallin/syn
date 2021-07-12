resource "null_resource" "raspberrypi_motion" {
  provisioner "remote-exec" {
    inline = [
      "sudo apt install -y motion",
      "sudo chown pi /etc/motion/motion.conf"
    ]

    connection {
      type = "ssh"
      user = var.raspberry_pi_user
      password = var.raspberry_pi_password
      host = var.raspberry_pi_host
    }
  }
}

resource "null_resource" "raspberrypi_motion_config" {
  provisioner "file" {
    source = "raspberrypi/motion.conf"
    destination = "/etc/motion/motion.conf"

    connection {
      type = "ssh"
      user = var.raspberry_pi_user
      password = var.raspberry_pi_password
      host = var.raspberry_pi_host
    }
  }
}