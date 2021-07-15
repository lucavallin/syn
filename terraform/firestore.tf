#
# Enable Firestore :|
#
resource "google_app_engine_application" "this" {
  provider      = google-beta
  project       = data.google_project.this.project_id
  location_id   = "europe-west"
  database_type = "CLOUD_FIRESTORE"
}

#
# Let functions use Firestore
#
resource "google_project_iam_member" "firestore_user" {
  project = data.google_project.this.project_id
  role    = "roles/datastore.user"
  member  = "serviceAccount:${google_service_account.functions.email}"
}