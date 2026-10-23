package compute

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/registry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registry.Schema{
		Name:        ComputeInstanceGroupManagerSchemaName,
		ProductName: "compute",
		Type:        registry.SchemaTypeResource,
		Schema:      ResourceComputeInstanceGroupManager(),
	}.Register()

	registry.Schema{
		Name:        ComputeRegionInstanceGroupManagerSchemaName,
		ProductName: "compute",
		Type:        registry.SchemaTypeResource,
		Schema:      ResourceComputeRegionInstanceGroupManager(),
	}.Register()
}

// ComputeInstanceGroupManagerAssetType is the CAI asset type name for compute instance group manager.
const ComputeInstanceGroupManagerAssetType string = "compute.googleapis.com/InstanceGroupManager"

// ComputeRegionInstanceGroupManagerAssetType is the CAI asset type name for regional compute instance group manager.
const ComputeRegionInstanceGroupManagerAssetType string = "compute.googleapis.com/RegionInstanceGroupManager"

// ComputeInstanceGroupManagerSchemaName is the TF resource schema name for compute instance group manager.
const ComputeInstanceGroupManagerSchemaName string = "google_compute_instance_group_manager"

// ComputeRegionInstanceGroupManagerSchemaName is the TF resource schema name for regional compute instance group manager.
const ComputeRegionInstanceGroupManagerSchemaName string = "google_compute_region_instance_group_manager"

func ResourceComputeInstanceGroupManager() *schema.Resource {
	return &schema.Resource{
		Schema: instanceGroupManagerSchema(true),
	}
}

func ResourceComputeRegionInstanceGroupManager() *schema.Resource {
	return &schema.Resource{
		Schema: instanceGroupManagerSchema(false),
	}
}

func instanceGroupManagerSchema(isZonal bool) map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"base_instance_name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"project": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ForceNew: true,
		},
		"target_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"target_stopped_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"target_suspended_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"target_pools": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"fingerprint": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"instance_group": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"self_link": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"operation": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"wait_for_instances": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"wait_for_instances_status": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "STABLE",
		},
		"list_managed_instances_results": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "PAGELESS",
		},
		"version": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"instance_template": {
						Type:     schema.TypeString,
						Required: true,
					},
					"target_size": {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"fixed": {
									Type:     schema.TypeInt,
									Optional: true,
								},
								"percent": {
									Type:     schema.TypeInt,
									Optional: true,
								},
							},
						},
					},
				},
			},
		},
		"named_port": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:     schema.TypeString,
						Required: true,
					},
					"port": {
						Type:     schema.TypeInt,
						Required: true,
					},
				},
			},
		},
		"auto_healing_policies": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"health_check": {
						Type:     schema.TypeString,
						Required: true,
					},
					"initial_delay_sec": {
						Type:     schema.TypeInt,
						Required: true,
					},
				},
			},
		},
		"update_policy": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"minimal_action": {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
					"most_disruptive_allowed_action": {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
					"type": {
						Type:     schema.TypeString,
						Required: true,
					},
					"instance_redistribution_type": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"max_surge_fixed": {
						Type:     schema.TypeInt,
						Optional: true,
						Computed: true,
					},
					"max_surge_percent": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"max_unavailable_fixed": {
						Type:     schema.TypeInt,
						Optional: true,
						Computed: true,
					},
					"max_unavailable_percent": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"min_ready_sec": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"replacement_method": {
						Type:     schema.TypeString,
						Optional: true,
						Default:  "SUBSTITUTE",
					},
				},
			},
		},
		"all_instances_config": {
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
					"labels": {
						Type:     schema.TypeMap,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},
				},
			},
		},
		"standby_policy": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"initial_delay_sec": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"mode": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		"resource_policies": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"workload_policy": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		"stateful_disk": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"device_name": {
						Type:     schema.TypeString,
						Required: true,
					},
					"delete_rule": {
						Type:     schema.TypeString,
						Optional: true,
						Default:  "NEVER_ON_PERMANENT_INSTANCE_DELETION",
					},
				},
			},
		},
		"stateful_internal_ip": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"interface_name": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"delete_rule": {
						Type:     schema.TypeString,
						Optional: true,
						Default:  "NEVER_ON_PERMANENT_INSTANCE_DELETION",
					},
				},
			},
		},
		"stateful_external_ip": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"interface_name": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"delete_rule": {
						Type:     schema.TypeString,
						Optional: true,
						Default:  "NEVER_ON_PERMANENT_INSTANCE_DELETION",
					},
				},
			},
		},
		"instance_lifecycle_policy": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"force_update_on_repair": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"default_action_on_failure": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"on_failed_health_check": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"on_repair": {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"allow_changing_zone": {
									Type:     schema.TypeString,
									Optional: true,
								},
							},
						},
					},
				},
			},
		},
		"params": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"resource_manager_tags": {
						Type:     schema.TypeMap,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},
				},
			},
		},
		"target_size_policy": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"mode": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		"instance_flexibility_policy": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"instance_selections": {
						Type:     schema.TypeSet,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"name": {
									Type:     schema.TypeString,
									Required: true,
								},
								"rank": {
									Type:     schema.TypeInt,
									Optional: true,
								},
								"machine_types": {
									Type:     schema.TypeSet,
									Optional: true,
									Elem:     &schema.Schema{Type: schema.TypeString},
								},
								"min_cpu_platform": {
									Type:     schema.TypeString,
									Optional: true,
								},
								"disks": {
									Type:     schema.TypeList,
									Optional: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"auto_delete": {
												Type:     schema.TypeBool,
												Optional: true,
												Default:  true,
											},
											"boot": {
												Type:     schema.TypeBool,
												Optional: true,
												Default:  false,
											},
											"device_name": {
												Type:     schema.TypeString,
												Optional: true,
											},
											"disk_size_gb": {
												Type:     schema.TypeInt,
												Optional: true,
											},
											"disk_type": {
												Type:     schema.TypeString,
												Optional: true,
											},
											"source_image": {
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
	}

	if isZonal {
		s["zone"] = &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ForceNew: true,
		}
	} else {
		s["region"] = &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ForceNew: true,
		}
		s["distribution_policy_zones"] = &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		}
		s["distribution_policy_target_shape"] = &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		}
	}

	return s
}
