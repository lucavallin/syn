#
# Create bucket for storing Terraform state
#
resource "google_storage_bucket" "terraform_state" {
  project                     = data.google_project.this.project_id
  name                        = "${var.company_code}-syn-tfstate"
  location                    = "europe-west4"
  force_destroy               = false
  uniform_bucket_level_access = true

  versioning {
    enabled = true
  }
}

#
# Create bucket for storing Raspberry Pi images
#
resource "google_storage_bucket" "uploads" {
  project                     = data.google_project.this.project_id
  name                        = "${var.company_code}-syn-uploads"
  location                    = "europe-west4"
  uniform_bucket_level_access = true
}