#
# Setup project
#
module "project" {
  source  = "terraform-google-modules/project-factory/google"
  version = "~> 10.3.1"

  name                        = "${var.company_code}-whopooped"
  org_id                      = data.google_organization.this.org_id
  billing_account             = data.google_billing_account.this.id
  disable_services_on_destroy = false

  activate_apis = [
    "admin.googleapis.com",
    "cloudidentity.googleapis.com",
    "cloudresourcemanager.googleapis.com",
    "cloudbilling.googleapis.com",
    "iam.googleapis.com",
    "iamcredentials.googleapis.com",
    "logging.googleapis.com",
    "monitoring.googleapis.com",
    "storage-api.googleapis.com",
    "billingbudgets.googleapis.com",
    "sourcerepo.googleapis.com",
    "cloudfunctions.googleapis.com"
  ]
}