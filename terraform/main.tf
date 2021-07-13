data "google_organization" "this" {
  domain = var.company_domain
}

data "google_billing_account" "this" {
  display_name = var.company_domain
}

data "google_project" "this" {
  project_id = module.project.project_id
}

#
# Setup project
#
module "project" {
  source  = "terraform-google-modules/project-factory/google"
  version = "~> 10.3.1"

  name                        = "${var.company_code}-syn"
  org_id                      = data.google_organization.this.org_id
  billing_account             = data.google_billing_account.this.id
  disable_services_on_destroy = false

  activate_apis = [
    "logging.googleapis.com",
    "monitoring.googleapis.com",
    "billingbudgets.googleapis.com",
    "sourcerepo.googleapis.com",
    "cloudfunctions.googleapis.com",
    "cloudbuild.googleapis.com",
    "vision.googleapis.com",
    "firebase.googleapis.com",
    "serviceusage.googleapis.com"
  ]
}

#
# Give owner permissions to your user and terraform
#
resource "google_organization_iam_binding" "owner" {
  org_id = data.google_organization.this.org_id
  role   = "roles/owner"
  members = [
    "user:${var.google_cloud_owner_email}",
  ]
}