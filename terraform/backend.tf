terraform {
  backend "gcs" {
    # Replace with your unique bucket name
    bucket = "cvln-syn-tfstate"
    prefix = "terraform/syn/state"
  }
}