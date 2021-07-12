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