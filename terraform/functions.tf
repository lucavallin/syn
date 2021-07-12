resource "google_sourcerepo_repository" "functions" {
  project = data.google_project.this.project_id
  name    = "functions"
}

resource "google_cloudfunctions_function" "process_uploaded_image" {
  project             = data.google_project.this.project_id
  region              = "europe-west1"
  name                = "process-uploaded-image"
  description         = "Processes new images being uploaded to ${google_storage_bucket.images.name}"
  runtime             = "go113"
  ingress_settings    = "ALLOW_INTERNAL_ONLY"
  available_memory_mb = 128

  source_repository {
    url = "${google_sourcerepo_repository.functions.url}/functions/processNewImage.go"
  }

  event_trigger {
    event_type = "google.storage.object.finalize"
    resource   = google_storage_bucket.images.name
  }
}