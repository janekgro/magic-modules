package compute_test

import (
	"testing"

	cai2hcl_testing "github.com/GoogleCloudPlatform/terraform-google-conversion/v7/cai2hcl/testing"
)

func TestComputeTargetPool(t *testing.T) {
	cai2hcl_testing.AssertTestFiles(
		t,
		"./testdata",
		[]string{
			"full_compute_target_pool",
		},
	)
}
