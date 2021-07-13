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

#
# Setup the Terraform service account
#
resource "google_service_account" "terraform" {
  project      = data.google_project.this.project_id
  account_id   = "terraform"
  display_name = "Terraform"
}

#
# Set permissions for Terraform service account
#
resource "google_organization_iam_member" "terraform" {
  for_each = toset([
    "roles/billing.user",
    "roles/orgpolicy.policyAdmin",
    "roles/resourcemanager.organizationViewer",
    "roles/serviceusage.serviceUsageAdmin",
    "roles/firebase.admin"
  ])
  org_id = data.google_organization.this.org_id
  role   = each.value
  member = "serviceAccount:${google_service_account.terraform.email}"
}

resource "google_billing_account_iam_member" "tf_sa" {
  billing_account_id = data.google_billing_account.this.id
  role               = "roles/billing.user"
  member             = "serviceAccount:${google_service_account.terraform.email}"
}

#
# Let terraform change/read buckets
#
resource "google_storage_bucket_iam_member" "terraform_state" {
  bucket = google_storage_bucket.terraform_state.name
  role   = "roles/storage.admin"
  member = "serviceAccount:${google_service_account.terraform.email}"
}

#
# Let your user impersonate the Terraform service account
#
resource "google_service_account_iam_member" "admin" {
  service_account_id = google_service_account.terraform.name
  role               = "roles/iam.serviceAccountTokenCreator"
  member             = "user:${var.google_cloud_owner_email}"
}