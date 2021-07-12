#
# Create bucket for storing Terraform state
#
resource "google_storage_bucket" "org_terraform_state" {
  project                     = data.google_project.this.project_id
  name                        = "${var.company_code}-whopooped-tfstate"
  location                    = "EU"
  force_destroy               = false
  uniform_bucket_level_access = true

  versioning {
    enabled = true
  }
}

#
# Create bucket for storing Raspberry Pi images
#
resource "google_storage_bucket" "images" {
  project                     = data.google_project.this.project_id
  name                        = "${var.company_code}-whopooped-images"
  location                    = "europe-west4"
  uniform_bucket_level_access = true
}

#
# Allow the Raspberry Pi service account to write to the bucket
#
resource "google_storage_bucket_iam_binding" "binding" {
  bucket = google_storage_bucket.images.name
  role   = "roles/storage.objectCreator"
  members = [
    "serviceAccount:${google_service_account.raspberrypi.email}",
  ]
}