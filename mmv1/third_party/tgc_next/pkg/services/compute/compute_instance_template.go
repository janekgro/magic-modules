package compute

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/registry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registry.Schema{
		Name:        ComputeInstanceTemplateSchemaName,
		ProductName: "compute",
		Type:        registry.SchemaTypeResource,
		Schema:      ResourceComputeInstanceTemplate(),
	}.Register()

	registry.Schema{
		Name:        ComputeRegionInstanceTemplateSchemaName,
		ProductName: "compute",
		Type:        registry.SchemaTypeResource,
		Schema:      ResourceComputeRegionInstanceTemplate(),
	}.Register()
}

// ComputeInstanceTemplateAssetType is the CAI asset type name for compute instance template.
const ComputeInstanceTemplateAssetType string = "compute.googleapis.com/InstanceTemplate"

// ComputeRegionInstanceTemplateAssetType is the CAI asset type name for regional compute instance template.
const ComputeRegionInstanceTemplateAssetType string = "compute.googleapis.com/RegionInstanceTemplate"

// ComputeInstanceTemplateSchemaName is the TF resource schema name for compute instance template.
const ComputeInstanceTemplateSchemaName string = "google_compute_instance_template"

// ComputeRegionInstanceTemplateSchemaName is the TF resource schema name for regional compute instance template.
const ComputeRegionInstanceTemplateSchemaName string = "google_compute_region_instance_template"

func ResourceComputeInstanceTemplate() *schema.Resource {
	return &schema.Resource{
		Schema: instanceTemplateSchema(false),
	}
}

func ResourceComputeRegionInstanceTemplate() *schema.Resource {
	return &schema.Resource{
		Schema: instanceTemplateSchema(true),
	}
}

func instanceTemplateSchema(isRegional bool) map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"name": {
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ForceNew:      true,
			ConflictsWith: []string{"name_prefix"},
		},
		"name_prefix": {
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ForceNew:      true,
			ConflictsWith: []string{"name"},
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"enable_display": {
			Type:     schema.TypeBool,
			Optional: true,
			ForceNew: true,
		},
		"instance_description": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"tags": {
			Type:     schema.TypeSet,
			Optional: true,
			ForceNew: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"labels": {
			Type:     schema.TypeMap,
			Optional: true,
			ForceNew: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"terraform_labels": {
			Type:     schema.TypeMap,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"effective_labels": {
			Type:     schema.TypeMap,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"project": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ForceNew: true,
		},
		"region": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ForceNew: true,
		},
		"machine_type": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"min_cpu_platform": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"can_ip_forward": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
			ForceNew: true,
		},
		"metadata": {
			Type:     schema.TypeMap,
			Optional: true,
			ForceNew: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"metadata_startup_script": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"disk": {
			Type:     schema.TypeList,
			Required: true,
			ForceNew: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"auto_delete": {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  true,
						ForceNew: true,
					},
					"boot": {
						Type:     schema.TypeBool,
						Optional: true,
						ForceNew: true,
					},
					"device_name": {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
						ForceNew: true,
					},
					"disk_name": {
						Type:     schema.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"source_image": {
						Type:     schema.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"interface": {
						Type:     schema.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"mode": {
						Type:     schema.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"source": {
						Type:     schema.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"disk_type": {
						Type:     schema.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"disk_size_gb": {
						Type:     schema.TypeInt,
						Optional: true,
						ForceNew: true,
					},
					"type": {
						Type:     schema.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"labels": {
						Type:     schema.TypeMap,
						Optional: true,
						ForceNew: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},
					"resource_policies": {
						Type:     schema.TypeList,
						Optional: true,
						ForceNew: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},
					"provisioned_iops": {
						Type:     schema.TypeInt,
						Optional: true,
						ForceNew: true,
					},
					"provisioned_throughput": {
						Type:     schema.TypeInt,
						Optional: true,
						ForceNew: true,
					},
					"disk_encryption_key": {
						Type:     schema.TypeList,
						Optional: true,
						ForceNew: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"kms_key_self_link": {
									Type:     schema.TypeString,
									Required: true,
									ForceNew: true,
								},
							},
						},
					},
					"source_image_encryption_key": {
						Type:     schema.TypeList,
						Optional: true,
						ForceNew: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"kms_key_self_link": {
									Type:     schema.TypeString,
									Optional: true,
									ForceNew: true,
								},
								"kms_key_service_account": {
									Type:     schema.TypeString,
									Optional: true,
									ForceNew: true,
								},
							},
						},
					},
					"source_snapshot_encryption_key": {
						Type:     schema.TypeList,
						Optional: true,
						ForceNew: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"kms_key_self_link": {
									Type:     schema.TypeString,
									Optional: true,
									ForceNew: true,
								},
								"kms_key_service_account": {
									Type:     schema.TypeString,
									Optional: true,
									ForceNew: true,
								},
							},
						},
					},
				},
			},
		},
		"network_interface": {
			Type:     schema.TypeList,
			Optional: true,
			ForceNew: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"network": {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
						ForceNew: true,
					},
					"subnetwork": {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
						ForceNew: true,
					},
					"subnetwork_project": {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
						ForceNew: true,
					},
					"network_ip": {
						Type:     schema.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"stack_type": {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
						ForceNew: true,
					},
					"queue_count": {
						Type:     schema.TypeInt,
						Optional: true,
						ForceNew: true,
					},
					"nic_type": {
						Type:     schema.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"access_config": {
						Type:     schema.TypeList,
						Optional: true,
						ForceNew: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"nat_ip": {
									Type:     schema.TypeString,
									Optional: true,
									Computed: true,
									ForceNew: true,
								},
								"network_tier": {
									Type:     schema.TypeString,
									Optional: true,
									Computed: true,
									ForceNew: true,
								},
								"public_ptr_domain_name": {
									Type:     schema.TypeString,
									Optional: true,
									ForceNew: true,
								},
							},
						},
					},
					"ipv6_access_config": {
						Type:     schema.TypeList,
						Optional: true,
						ForceNew: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"network_tier": {
									Type:     schema.TypeString,
									Required: true,
									ForceNew: true,
								},
								"public_ptr_domain_name": {
									Type:     schema.TypeString,
									Optional: true,
									ForceNew: true,
								},
							},
						},
					},
					"alias_ip_range": {
						Type:     schema.TypeList,
						Optional: true,
						ForceNew: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"ip_cidr_range": {
									Type:     schema.TypeString,
									Required: true,
									ForceNew: true,
								},
								"subnetwork_range_name": {
									Type:     schema.TypeString,
									Optional: true,
									ForceNew: true,
								},
							},
						},
					},
				},
			},
		},
		"service_account": {
			Type:     schema.TypeList,
			Optional: true,
			ForceNew: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"email": {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
						ForceNew: true,
					},
					"scopes": {
						Type:     schema.TypeSet,
						Required: true,
						ForceNew: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},
				},
			},
		},
		"scheduling": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			ForceNew: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"automatic_restart": {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  true,
						ForceNew: true,
					},
					"on_host_maintenance": {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
						ForceNew: true,
					},
					"preemptible": {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  false,
						ForceNew: true,
					},
					"provisioning_model": {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
						ForceNew: true,
					},
					"instance_termination_action": {
						Type:     schema.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"maintenance_interval": {
						Type:     schema.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"max_run_duration": {
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
									Type:     schema.TypeInt,
									Required: true,
									ForceNew: true,
								},
							},
						},
					},
					"node_affinities": {
						Type:     schema.TypeSet,
						Optional: true,
						ForceNew: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"key": {
									Type:     schema.TypeString,
									Required: true,
									ForceNew: true,
								},
								"operator": {
									Type:     schema.TypeString,
									Required: true,
									ForceNew: true,
								},
								"values": {
									Type:     schema.TypeSet,
									Required: true,
									ForceNew: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
							},
						},
					},
				},
			},
		},
		"shielded_instance_config": {
			Type:     schema.TypeList,
			Optional: true,
			ForceNew: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"enable_secure_boot": {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  false,
						ForceNew: true,
					},
					"enable_vtpm": {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  true,
						ForceNew: true,
					},
					"enable_integrity_monitoring": {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  true,
						ForceNew: true,
					},
				},
			},
		},
		"confidential_instance_config": {
			Type:     schema.TypeList,
			Optional: true,
			ForceNew: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"enable_confidential_compute": {
						Type:     schema.TypeBool,
						Required: true,
						ForceNew: true,
					},
					"confidential_instance_type": {
						Type:     schema.TypeString,
						Optional: true,
						ForceNew: true,
					},
				},
			},
		},
		"guest_accelerator": {
			Type:     schema.TypeList,
			Optional: true,
			ForceNew: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"type": {
						Type:     schema.TypeString,
						Required: true,
						ForceNew: true,
					},
					"count": {
						Type:     schema.TypeInt,
						Required: true,
						ForceNew: true,
					},
				},
			},
		},
		"reservation_affinity": {
			Type:     schema.TypeList,
			Optional: true,
			ForceNew: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"type": {
						Type:     schema.TypeString,
						Required: true,
						ForceNew: true,
					},
					"specific_reservation": {
						Type:     schema.TypeList,
						Optional: true,
						ForceNew: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"key": {
									Type:     schema.TypeString,
									Required: true,
									ForceNew: true,
								},
								"values": {
									Type:     schema.TypeList,
									Required: true,
									ForceNew: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
							},
						},
					},
				},
			},
		},
		"advanced_machine_features": {
			Type:     schema.TypeList,
			Optional: true,
			ForceNew: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"enable_nested_virtualization": {
						Type:     schema.TypeBool,
						Optional: true,
						ForceNew: true,
					},
					"threads_per_core": {
						Type:     schema.TypeInt,
						Optional: true,
						ForceNew: true,
					},
					"visible_core_count": {
						Type:     schema.TypeInt,
						Optional: true,
						ForceNew: true,
					},
					"performance_monitoring_unit": {
						Type:     schema.TypeString,
						Optional: true,
						ForceNew: true,
					},
					"enable_uefi_networking": {
						Type:     schema.TypeBool,
						Optional: true,
						ForceNew: true,
					},
				},
			},
		},
		"network_performance_config": {
			Type:     schema.TypeList,
			Optional: true,
			ForceNew: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"total_egress_bandwidth_tier": {
						Type:     schema.TypeString,
						Required: true,
						ForceNew: true,
					},
				},
			},
		},
		"resource_policies": {
			Type:     schema.TypeList,
			Optional: true,
			ForceNew: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"key_revocation_action_type": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
	}

	return s
}
