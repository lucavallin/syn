variable "company_code" {
    type = string
    description = "Company code to prefix resources where needed"
}

variable "company_domain" {
    type = string
    description = "Google Cloud organization domain"
}

variable "google_cloud_owner" {
    type = string
    description = "Google Cloud user email of the organization owner"
}