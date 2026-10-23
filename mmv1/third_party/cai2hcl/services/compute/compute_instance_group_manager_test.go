package compute_test

import (
	"testing"

	cai2hcl_testing "github.com/GoogleCloudPlatform/terraform-google-conversion/v7/cai2hcl/testing"
)

func TestComputeInstanceGroupManager(t *testing.T) {
	cai2hcl_testing.AssertTestFiles(
		t,
		"./testdata",
		[]string{
			"full_compute_instance_group_manager",
			"full_compute_region_instance_group_manager",
			"full_compute_region_instance_group_manager_stateful",
		},
	)
}
