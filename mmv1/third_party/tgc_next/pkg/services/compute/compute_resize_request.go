package compute

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/registry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registry.Schema{
		Name:        ComputeResizeRequestSchemaName,
		ProductName: "compute",
		Type:        registry.SchemaTypeResource,
		Schema:      ResourceComputeResizeRequest(),
	}.Register()
}

// ComputeResizeRequestAssetType is the CAI asset type name for resize request.
const ComputeResizeRequestAssetType string = "compute.googleapis.com/InstanceGroupManagerResizeRequest"

// ComputeResizeRequestSchemaName is the TF resource schema name for resize request.
const ComputeResizeRequestSchemaName string = "google_compute_resize_request"

func ResourceComputeResizeRequest() *schema.Resource {
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
			"resize_by": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
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
			"requested_run_duration": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"nanos": {
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
						},
						"seconds": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
			"creation_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"error": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"errors": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"code": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"location": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"message": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"last_attempt": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"errors": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"code": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"location": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"message": {
																Type:     schema.TypeString,
																Computed: true,
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
				},
			},
		},
	}
}
