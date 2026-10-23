resource "google_compute_instance_group" "test-unmanaged-group" {
  description = "Full unmanaged instance group"
  instances   = ["projects/test-project/zones/us-central1-a/instances/test-inst-1"]
  name        = "test-unmanaged-group"
  named_port {
    name = "http"
    port = 80
  }
  named_port {
    name = "https"
    port = 443
  }
  network = "projects/test-project/global/networks/default"
  zone    = "us-central1-a"
}
