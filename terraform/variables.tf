variable "company_code" {
  type        = string
  description = "Company code to prefix resources where needed"
}

variable "company_domain" {
  type        = string
  description = "Google Cloud organization domain"
}

variable "google_cloud_owner_email" {
  type        = string
  description = "Google Cloud user email of the organization owner"
}

variable "raspberry_pi_user" {
  type        = string
  description = "Username for SSH connection to the Raspberry Pi"
  default     = "pi"
}

variable "raspberry_pi_password" {
  type        = string
  description = "Password for SSH connection to the Raspberry Pi"
  default     = "raspberry"
}

variable "raspberry_pi_host" {
  type        = string
  description = "IP for SSH connection to the Raspberry Pi"
}

variable "allowed_labels" {
  type        = list(string)
  description = "The ProcessUpload Cloud Function will only store uploads in Firestore if at least one of these labels is found by Vision API"
}