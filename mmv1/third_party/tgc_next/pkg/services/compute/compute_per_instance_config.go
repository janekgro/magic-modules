package compute

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/registry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registry.Schema{
		Name:        ComputePerInstanceConfigSchemaName,
		ProductName: "compute",
		Type:        registry.SchemaTypeResource,
		Schema:      ResourceComputePerInstanceConfig(),
	}.Register()
}

// ComputePerInstanceConfigAssetType is the CAI asset type name for per instance config.
const ComputePerInstanceConfigAssetType string = "compute.googleapis.com/PerInstanceConfig"

// ComputePerInstanceConfigSchemaName is the TF resource schema name for per instance config.
const ComputePerInstanceConfigSchemaName string = "google_compute_per_instance_config"

func ResourceComputePerInstanceConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_group_manager": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"zone": {
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
			"minimal_action": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"most_disruptive_allowed_action": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"remove_instance_on_destroy": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"remove_instance_state_on_destroy": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"preserved_state": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metadata": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"disk": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"device_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"source": {
										Type:     schema.TypeString,
										Required: true,
									},
									"mode": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "READ_WRITE",
									},
									"delete_rule": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "NEVER_ON_PERMANENT_INSTANCE_DELETION",
									},
								},
							},
						},
						"internal_ip": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"auto_delete": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "NEVER",
									},
									"ip_address": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"address": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"external_ip": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"auto_delete": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "NEVER",
									},
									"ip_address": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"address": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
