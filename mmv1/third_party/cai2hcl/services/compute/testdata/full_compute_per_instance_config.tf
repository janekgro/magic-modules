resource "google_compute_per_instance_config" "instance-1" {
  instance_group_manager = "test-mig"
  name                   = "instance-1"
  preserved_state {
    disk {
      delete_rule = "NEVER"
      device_name = "disk-1"
      mode        = "READ_WRITE"
      source      = "projects/test-project/zones/us-central1-a/disks/disk-1"
    }
    external_ip {
      auto_delete    = "NEVER"
      interface_name = "nic0"
      ip_address {
        address = "projects/test-project/regions/us-central1/addresses/ext-ip-1"
      }
    }
    internal_ip {
      auto_delete    = "NEVER"
      interface_name = "nic0"
      ip_address {
        address = "projects/test-project/regions/us-central1/addresses/ip-1"
      }
    }
    metadata = {
      foo = "bar"
    }
  }
  zone = "us-central1-a"
}
