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

  activate_apis = []
}