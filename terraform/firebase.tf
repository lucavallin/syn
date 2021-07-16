data "google_firebase_web_app_config" "this" {
  web_app_id = data.google_project.this.project_id
}

output "firebase_config" {
  value = data.google_firebase_web_app_config.this
}
