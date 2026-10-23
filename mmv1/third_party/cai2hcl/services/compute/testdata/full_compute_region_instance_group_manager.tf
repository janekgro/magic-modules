resource "google_compute_region_instance_group_manager" "test-rmig" {
  all_instances_config {
    labels = {
      env = "prod"
    }
    metadata = {
      foo = "bar"
    }
  }
  auto_healing_policies {
    health_check      = "projects/test-project/global/healthChecks/test-hc"
    initial_delay_sec = 300
  }
  base_instance_name               = "test-rinstance"
  description                      = "test regional description"
  distribution_policy_target_shape = "EVEN"
  distribution_policy_zones        = ["us-central1-a", "us-central1-b"]
  instance_flexibility_policy {
    instance_selections {
      disks {
        auto_delete  = true
        boot         = true
        device_name  = "boot-disk"
        disk_size_gb = 50
        disk_type    = "pd-ssd"
        source_image = "projects/debian-cloud/global/images/family/debian-11"
      }
      disks {
        auto_delete  = false
        boot         = false
        device_name  = "data-disk"
        disk_size_gb = 100
        disk_type    = "pd-balanced"
      }
      machine_types    = ["c3-standard-4"]
      min_cpu_platform = "Intel Sapphire Rapids"
      name             = "c3-standard-4-selection"
      rank             = 0
    }
  }
  instance_lifecycle_policy {
    default_action_on_failure = "REPAIR"
    force_update_on_repair    = "YES"
  }
  list_managed_instances_results = "PAGELESS"
  name                           = "test-rmig"
  named_port {
    name = "http"
    port = 80
  }
  region = "us-central1"
  standby_policy {
    initial_delay_sec = 0
    mode              = "MANUAL"
  }
  target_pools          = ["projects/test-project/regions/us-central1/targetPools/test-tp"]
  target_size           = 5
  target_stopped_size   = 1
  target_suspended_size = 1
  update_policy {
    instance_redistribution_type = "PROACTIVE"
    max_surge_fixed              = 0
    max_unavailable_fixed        = 4
    minimal_action               = "RESTART"
    type                         = "PROACTIVE"
  }
  version {
    instance_template = "projects/test-project/global/instanceTemplates/test-template"
    name              = "primary"
  }
}
