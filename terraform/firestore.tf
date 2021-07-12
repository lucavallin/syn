#
# Enable Firestore :|
#
resource "google_app_engine_application" "this" {
  provider      = google-beta
  location_id   = "europe-west4"
  database_type = "CLOUD_FIRESTORE"
}

#
# Let functions use Firestore
#
resource "google_project_iam_member" "firestore_user" {
  role   = "roles/datastore.user"
  member = "serviceAccount:${google_service_account.functions.email}"
}