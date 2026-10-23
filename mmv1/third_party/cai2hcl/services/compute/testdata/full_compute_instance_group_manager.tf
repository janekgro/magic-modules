resource "google_compute_instance_group_manager" "test-mig" {
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
  base_instance_name = "test-instance"
  description        = "test description"
  instance_lifecycle_policy {
    default_action_on_failure = "REPAIR"
    force_update_on_repair    = "YES"
  }
  list_managed_instances_results = "PAGELESS"
  name                           = "test-mig"
  named_port {
    name = "http"
    port = 80
  }
  standby_policy {
    initial_delay_sec = 0
    mode              = "MANUAL"
  }
  target_pools          = ["projects/test-project/regions/us-central1/targetPools/test-tp"]
  target_size           = 3
  target_stopped_size   = 1
  target_suspended_size = 1
  update_policy {
    max_surge_fixed       = 3
    max_unavailable_fixed = 0
    minimal_action        = "REPLACE"
    type                  = "PROACTIVE"
  }
  version {
    instance_template = "projects/test-project/global/instanceTemplates/test-template"
    name              = "primary"
  }
  zone = "us-central1-a"
}
