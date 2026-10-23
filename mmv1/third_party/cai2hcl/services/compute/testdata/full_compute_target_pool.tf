resource "google_compute_target_pool" "test-tp" {
  backup_pool      = "projects/test-project/regions/us-central1/targetPools/backup-tp"
  description      = "test target pool description"
  failover_ratio   = 0.5
  health_checks    = ["projects/test-project/global/httpHealthChecks/test-hc"]
  instances        = ["projects/test-project/zones/us-central1-a/instances/test-inst-1"]
  name             = "test-tp"
  region           = "us-central1"
  security_policy  = "projects/test-project/regions/us-central1/securityPolicies/test-policy"
  session_affinity = "CLIENT_IP"
}
