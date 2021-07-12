#
# Give owner permissions to your user
#
resource "google_organization_iam_binding" "owner" {
  org_id = data.google_organization.this.org_id
  role   = "roles/owner"
  members = [
    "user:${var.google_cloud_owner}",
  ]
}