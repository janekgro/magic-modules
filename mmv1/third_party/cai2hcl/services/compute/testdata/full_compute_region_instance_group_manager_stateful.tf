resource "google_compute_region_instance_group_manager" "test-stateful-rmig" {
  base_instance_name               = "test-instance"
  description                      = "Full stateful regional MIG"
  distribution_policy_target_shape = "BALANCED"
  distribution_policy_zones        = ["us-central1-a", "us-central1-b", "us-central1-c"]
  instance_lifecycle_policy {
    default_action_on_failure = "REPAIR"
    force_update_on_repair    = "NO"
  }
  name   = "test-stateful-rmig"
  region = "us-central1"
  stateful_disk {
    delete_rule = "ON_PERMANENT_INSTANCE_DELETION"
    device_name = "data-disk-1"
  }
  stateful_internal_ip {
    delete_rule    = "NEVER"
    interface_name = "nic0"
  }
  target_size = 3
  update_policy {
    instance_redistribution_type = "NONE"
    max_surge_fixed              = 3
    minimal_action               = "REPLACE"
    type                         = "PROACTIVE"
  }
  version {
    instance_template = "projects/test-project/global/instanceTemplates/test-template"
  }
}
