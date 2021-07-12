data "google_organization" "this" {
  domain = var.company_domain
}

data "google_billing_account" "this" {
  display_name = var.company_domain
}

data "google_project" "this" {
  project_id = module.project.project_id
}