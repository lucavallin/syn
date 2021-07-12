resource "google_sourcerepo_repository" "syn" {
  project = data.google_project.this.project_id
  name    = "syn"
}

resource "google_cloudfunctions_function" "process_upload" {
  project             = data.google_project.this.project_id
  region              = "europe-west1"
  name                = "ProcessUpload"
  description         = "Processes new uploads to ${google_storage_bucket.uploads.name}"
  runtime             = "go113"
  ingress_settings    = "ALLOW_INTERNAL_ONLY"
  available_memory_mb = 128

  entry_point = "ProcessUpload"

  source_repository {
    url = "https://source.developers.google.com/projects/${data.google_project.this.project_id}/repos/syn/moveable-aliases/master/paths/functions"
  }

  event_trigger {
    event_type = "google.storage.object.finalize"
    resource   = google_storage_bucket.uploads.name
  }
}