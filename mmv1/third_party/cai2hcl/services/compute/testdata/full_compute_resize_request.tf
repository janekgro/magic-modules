resource "google_compute_resize_request" "test-resize-1" {
  description            = "Resize request for scale-out"
  instance_group_manager = "test-mig"
  name                   = "test-resize-1"
  requested_run_duration {
    seconds = "3600"
  }
  resize_by = 10
  zone      = "us-central1-a"
}
