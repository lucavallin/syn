#
# Create the Raspberry Pi service account
#
resource "google_service_account" "functions" {
  project      = data.google_project.this.project_id
  account_id   = "functions"
  display_name = "Cloud Functions service account"
}

resource "google_storage_bucket_iam_binding" "functions" {
  bucket = google_storage_bucket.uploads.name
  role   = "roles/storage.objectAdmin"
  members = [
    "serviceAccount:${google_service_account.functions.email}",
  ]
}

resource "google_project_iam_member" "functions" {
  project = data.google_project.this.project_id
  role    = "roles/serviceusage.serviceUsageViewer"
  member  = "serviceAccount:${google_service_account.functions.email}"
}

resource "google_sourcerepo_repository" "syn" {
  project = data.google_project.this.project_id
  name    = "syn"
}

resource "google_cloudfunctions_function" "process_upload" {
  project               = data.google_project.this.project_id
  region                = "europe-west1"
  name                  = "ProcessUpload"
  description           = "Processes new uploads to ${google_storage_bucket.uploads.name}"
  service_account_email = google_service_account.functions.email
  runtime               = "go113"
  ingress_settings      = "ALLOW_INTERNAL_ONLY"
  available_memory_mb   = 128

  entry_point = "ProcessUpload"

  source_repository {
    url = "https://source.developers.google.com/projects/${data.google_project.this.project_id}/repos/syn/moveable-aliases/master/paths/functions"
  }

  event_trigger {
    event_type = "google.storage.object.finalize"
    resource   = google_storage_bucket.uploads.name
  }

  environment_variables = {
    "GOOGLE_CLOUD_PROJECT_ID" : data.google_project.this.project_id
    "ACCEPTED_LABELS" : var.accepted_labels
  }
}

resource "google_cloudfunctions_function" "notify" {
  project               = data.google_project.this.project_id
  region                = "europe-west1"
  name                  = "Notify"
  description           = "Notifies of newly labeled uploads"
  service_account_email = google_service_account.functions.email
  runtime               = "go113"
  ingress_settings      = "ALLOW_INTERNAL_ONLY"
  available_memory_mb   = 128

  entry_point = "Notify"

  source_repository {
    url = "https://source.developers.google.com/projects/${data.google_project.this.project_id}/repos/syn/moveable-aliases/master/paths/functions"
  }

  event_trigger {
    event_type = "providers/cloud.firestore/eventTypes/document.create"
    resource   = "Events/{ids}"
  }

  environment_variables = {
    "IFTTT_WEBHOOK_URL" : var.ifttt_webhook_url
  }
}