resource "google_compute_region_instance_template" "test-rtemplate" {
  can_ip_forward = true
  description    = "test regional template description"
  disk {
    auto_delete = true
    boot        = true
  }
  disk {
    auto_delete = false
    boot        = false
    mode        = "READ_WRITE"
  }
  enable_display             = true
  instance_description       = "Regional instance description"
  key_revocation_action_type = "STOP"
  machine_type               = "n2-standard-4"
  metadata = {
    enable-oslogin = "TRUE"
  }
  min_cpu_platform = "Intel Ice Lake"
  name             = "test-rtemplate"
  network_interface {
    access_config {
      network_tier = "PREMIUM"
    }
    network    = "projects/test-project/global/networks/default"
    subnetwork = "projects/test-project/regions/us-central1/subnetworks/default"
  }
  region = "us-central1"
  scheduling {
    automatic_restart   = true
    on_host_maintenance = "MIGRATE"
  }
  service_account {
    email  = "default"
    scopes = ["https://www.googleapis.com/auth/cloud-platform"]
  }
  shielded_instance_config {
    enable_integrity_monitoring = true
    enable_secure_boot          = true
    enable_vtpm                 = true
  }
  tags = ["http-server", "https-server"]
}
