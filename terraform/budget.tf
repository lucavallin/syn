
resource "google_billing_budget" "budget" {
  billing_account = data.google_billing_account.this.id
  display_name    = "Billing Budget"

  budget_filter {
    projects = ["projects/${data.google_project.this.number}"]
  }

  amount {
    specified_amount {
      currency_code = "EUR"
      units         = "10"
    }
  }

  threshold_rules {
    threshold_percent = 1.0
  }
  threshold_rules {
    threshold_percent = 1.0
    spend_basis       = "FORECASTED_SPEND"
  }

  all_updates_rule {
    monitoring_notification_channels = [
      google_monitoring_notification_channel.notification_channel.id,
    ]
    disable_default_iam_recipients = true
  }
}

resource "google_monitoring_notification_channel" "notification_channel" {
  project      = data.google_project.this.id
  display_name = "Email Notification Channel"
  type         = "email"

  labels = {
    email_address = var.google_cloud_owner
  }
}