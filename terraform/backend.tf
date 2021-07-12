terraform {
  backend "gcs" {
    # Replace with your unique bucket name
    bucket = "cvln-whopooped-tfstate"
    prefix = "terraform/whopooped/state"
  }
}