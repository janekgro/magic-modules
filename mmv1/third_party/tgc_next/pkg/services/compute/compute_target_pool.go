package compute

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/registry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registry.Schema{
		Name:        ComputeTargetPoolSchemaName,
		ProductName: "compute",
		Type:        registry.SchemaTypeResource,
		Schema:      ResourceComputeTargetPool(),
	}.Register()
}

// ComputeTargetPoolAssetType is the CAI asset type name for compute target pool.
const ComputeTargetPoolAssetType string = "compute.googleapis.com/TargetPool"

// ComputeTargetPoolSchemaName is the TF resource schema name for compute target pool.
const ComputeTargetPoolSchemaName string = "google_compute_target_pool"

func ResourceComputeTargetPool() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"instances": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"health_checks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"session_affinity": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NONE",
			},
			"failover_ratio": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"backup_pool": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_policy": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
