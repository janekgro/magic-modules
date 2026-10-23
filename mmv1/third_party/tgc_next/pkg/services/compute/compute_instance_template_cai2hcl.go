package compute

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/cai2hcl/converters/utils"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/cai2hcl/models"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/caiasset"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/tpgresource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ComputeInstanceTemplateCai2hclConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewComputeInstanceTemplateCai2hclConverter(provider *schema.Provider) models.Cai2hclConverter {
	return newComputeInstanceTemplateCai2hclConverter(provider, ComputeInstanceTemplateSchemaName)
}

func NewComputeRegionInstanceTemplateCai2hclConverter(provider *schema.Provider) models.Cai2hclConverter {
	return newComputeInstanceTemplateCai2hclConverter(provider, ComputeRegionInstanceTemplateSchemaName)
}

func newComputeInstanceTemplateCai2hclConverter(provider *schema.Provider, schemaName string) models.Cai2hclConverter {
	schema := provider.ResourcesMap[schemaName].Schema

	return &ComputeInstanceTemplateCai2hclConverter{
		name:   schemaName,
		schema: schema,
	}
}

func (c *ComputeInstanceTemplateCai2hclConverter) Convert(assets []caiasset.Asset, options *models.ResourceConverterOptions) ([]*models.TerraformResourceBlock, error) {
	if len(assets) == 0 {
		return nil, nil
	}

	var blocks []*models.TerraformResourceBlock
	for _, asset := range assets {
		if asset.Resource == nil || asset.Resource.Data == nil {
			continue
		}
		data := asset.Resource.Data
		name, _ := data["name"].(string)
		if name == "" {
			name = tpgresource.GetResourceNameFromSelfLink(asset.Name)
		}

		resourceType := c.name
		hclData := make(map[string]any)

		outputFields := map[string]struct{}{"project": {}}
		if c.name == ComputeRegionInstanceTemplateSchemaName {
			utils.ParseUrlParamValuesFromAssetName(asset.Name, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/instanceTemplates/{{name}}", outputFields, hclData)
		} else {
			utils.ParseUrlParamValuesFromAssetName(asset.Name, "//compute.googleapis.com/projects/{{project}}/global/instanceTemplates/{{name}}", outputFields, hclData)
		}

		if name != "" {
			hclData["name"] = name
		}
		if description, ok := data["description"].(string); ok && description != "" {
			hclData["description"] = description
		}
		if region, ok := data["region"].(string); ok && region != "" {
			hclData["region"] = tpgresource.GetResourceNameFromSelfLink(region)
		}

		properties, ok := data["properties"].(map[string]any)
		if !ok {
			properties = data
		}

		if machineType, ok := properties["machineType"].(string); ok && machineType != "" {
			hclData["machine_type"] = machineType
		}
		if instanceDesc, ok := properties["description"].(string); ok && instanceDesc != "" {
			hclData["instance_description"] = instanceDesc
		}
		if displayDevice, ok := properties["displayDevice"].(map[string]any); ok {
			if enableDisplay, ok := displayDevice["enableDisplay"].(bool); ok && enableDisplay {
				hclData["enable_display"] = true
			}
		}
		if minCpuPlatform, ok := properties["minCpuPlatform"].(string); ok && minCpuPlatform != "" {
			hclData["min_cpu_platform"] = minCpuPlatform
		}
		if canIpForward, ok := properties["canIpForward"].(bool); ok {
			hclData["can_ip_forward"] = canIpForward
		}
		if keyRevocationActionType, ok := properties["keyRevocationActionType"].(string); ok && keyRevocationActionType != "" {
			hclData["key_revocation_action_type"] = keyRevocationActionType
		}

		if resourcePoliciesList, ok := properties["resourcePolicies"].([]any); ok && len(resourcePoliciesList) > 0 {
			var policies []string
			for _, p := range resourcePoliciesList {
				if s, ok := p.(string); ok && s != "" {
					policies = append(policies, s)
				}
			}
			if len(policies) > 0 {
				hclData["resource_policies"] = policies
			}
		}

		if labels, ok := properties["labels"].(map[string]any); ok && len(labels) > 0 {
			hclData["labels"] = labels
		}
		if metadataMap, ok := properties["metadata"].(map[string]any); ok {
			if items, ok := metadataMap["items"].([]any); ok && len(items) > 0 {
				m := make(map[string]string)
				for _, rawItem := range items {
					if item, ok := rawItem.(map[string]any); ok {
						if k, ok := item["key"].(string); ok && k != "" {
							v, _ := item["value"].(string)
							m[k] = v
						}
					}
				}
				if len(m) > 0 {
					hclData["metadata"] = m
				}
			} else if metaMap, ok := metadataMap["metadata"].(map[string]any); ok && len(metaMap) > 0 {
				hclData["metadata"] = metaMap
			}
		}

		if tagsMap, ok := properties["tags"].(map[string]any); ok {
			if itemsList, ok := tagsMap["items"].([]any); ok && len(itemsList) > 0 {
				var tags []string
				for _, item := range itemsList {
					if s, ok := item.(string); ok && s != "" {
						tags = append(tags, s)
					}
				}
				if len(tags) > 0 {
					hclData["tags"] = tags
				}
			}
		}

		// Flatten nested blocks
		if disks := flattenTemplateDisks(properties["disks"]); len(disks) > 0 {
			hclData["disk"] = disks
		}
		if networkInterfaces := flattenTemplateNetworkInterfaces(properties["networkInterfaces"]); len(networkInterfaces) > 0 {
			hclData["network_interface"] = networkInterfaces
		}
		if serviceAccounts := flattenTemplateServiceAccounts(properties["serviceAccounts"]); len(serviceAccounts) > 0 {
			hclData["service_account"] = serviceAccounts
		}
		if scheduling := flattenTemplateScheduling(properties["scheduling"]); len(scheduling) > 0 {
			hclData["scheduling"] = scheduling
		}
		if shielded := flattenTemplateShieldedInstanceConfig(properties["shieldedInstanceConfig"]); len(shielded) > 0 {
			hclData["shielded_instance_config"] = shielded
		}
		if confidential := flattenTemplateConfidentialInstanceConfig(properties["confidentialInstanceConfig"]); len(confidential) > 0 {
			hclData["confidential_instance_config"] = confidential
		}
		if reservationAffinity := flattenTemplateReservationAffinity(properties["reservationAffinity"]); len(reservationAffinity) > 0 {
			hclData["reservation_affinity"] = reservationAffinity
		}
		if guestAccelerator := flattenTemplateGuestAccelerator(properties["guestAccelerators"]); len(guestAccelerator) > 0 {
			hclData["guest_accelerator"] = guestAccelerator
		}
		if advancedFeatures := flattenTemplateAdvancedMachineFeatures(properties["advancedMachineFeatures"]); len(advancedFeatures) > 0 {
			hclData["advanced_machine_features"] = advancedFeatures
		}

		ctyVal, err := utils.MapToCtyValWithSchema(hclData, c.schema)
		if err != nil {
			return nil, fmt.Errorf("error converting map to cty.Value for InstanceTemplate: %w", err)
		}

		blocks = append(blocks, &models.TerraformResourceBlock{
			Labels: []string{resourceType, name},
			Value:  ctyVal,
		})
	}

	return blocks, nil
}

func flattenTemplateDisks(v any) []map[string]any {
	if v == nil {
		return nil
	}
	disksList, ok := v.([]any)
	if !ok || len(disksList) == 0 {
		return nil
	}
	var result []map[string]any
	for _, rawDisk := range disksList {
		diskMap, ok := rawDisk.(map[string]any)
		if !ok {
			continue
		}
		item := make(map[string]any)
		if autoDelete, ok := diskMap["autoDelete"].(bool); ok {
			item["auto_delete"] = autoDelete
		}
		if boot, ok := diskMap["boot"].(bool); ok {
			item["boot"] = boot
		}
		if deviceName, ok := diskMap["deviceName"].(string); ok && deviceName != "" {
			item["device_name"] = deviceName
		}
		if mode, ok := diskMap["mode"].(string); ok && mode != "" {
			item["mode"] = mode
		}
		if typeVal, ok := diskMap["type"].(string); ok && typeVal != "" {
			item["type"] = typeVal
		}
		if source, ok := diskMap["source"].(string); ok && source != "" {
			item["source"] = source
		}

		if initParams, ok := diskMap["initializeParams"].(map[string]any); ok {
			if diskSizeGb, ok := initParams["diskSizeGb"].(float64); ok {
				item["disk_size_gb"] = int64(diskSizeGb)
			} else if diskSizeGbStr, ok := initParams["diskSizeGb"].(string); ok && diskSizeGbStr != "" {
				var gb int64
				fmt.Sscanf(diskSizeGbStr, "%d", &gb)
				if gb > 0 {
					item["disk_size_gb"] = gb
				}
			}
			if diskType, ok := initParams["diskType"].(string); ok && diskType != "" {
				item["disk_type"] = diskType
			}
			if sourceImage, ok := initParams["sourceImage"].(string); ok && sourceImage != "" {
				item["source_image"] = sourceImage
			}
		}

		if len(item) > 0 {
			result = append(result, item)
		}
	}
	return result
}

func flattenTemplateNetworkInterfaces(v any) []map[string]any {
	if v == nil {
		return nil
	}
	niList, ok := v.([]any)
	if !ok || len(niList) == 0 {
		return nil
	}
	var result []map[string]any
	for _, rawNi := range niList {
		niMap, ok := rawNi.(map[string]any)
		if !ok {
			continue
		}
		item := make(map[string]any)
		if network, ok := niMap["network"].(string); ok && network != "" {
			item["network"] = network
		}
		if subnetwork, ok := niMap["subnetwork"].(string); ok && subnetwork != "" {
			item["subnetwork"] = subnetwork
		}
		if stackType, ok := niMap["stackType"].(string); ok && stackType != "" {
			item["stack_type"] = stackType
		}

		if accessConfigs, ok := niMap["accessConfigs"].([]any); ok && len(accessConfigs) > 0 {
			var acResult []map[string]any
			for _, rawAc := range accessConfigs {
				acMap, ok := rawAc.(map[string]any)
				if !ok {
					continue
				}
				acItem := make(map[string]any)
				if natIP, ok := acMap["natIP"].(string); ok && natIP != "" {
					acItem["nat_ip"] = natIP
				}
				if networkTier, ok := acMap["networkTier"].(string); ok && networkTier != "" {
					acItem["network_tier"] = networkTier
				}
				if len(acItem) > 0 {
					acResult = append(acResult, acItem)
				}
			}
			if len(acResult) > 0 {
				item["access_config"] = acResult
			}
		}

		if len(item) > 0 {
			result = append(result, item)
		}
	}
	return result
}

func flattenTemplateServiceAccounts(v any) []map[string]any {
	if v == nil {
		return nil
	}
	saList, ok := v.([]any)
	if !ok || len(saList) == 0 {
		return nil
	}
	var result []map[string]any
	for _, rawSa := range saList {
		saMap, ok := rawSa.(map[string]any)
		if !ok {
			continue
		}
		item := make(map[string]any)
		if email, ok := saMap["email"].(string); ok && email != "" {
			item["email"] = email
		}
		if scopesList, ok := saMap["scopes"].([]any); ok && len(scopesList) > 0 {
			var scopes []string
			for _, s := range scopesList {
				if str, ok := s.(string); ok && str != "" {
					scopes = append(scopes, str)
				}
			}
			if len(scopes) > 0 {
				item["scopes"] = scopes
			}
		}
		if len(item) > 0 {
			result = append(result, item)
		}
	}
	return result
}

func flattenTemplateScheduling(v any) []map[string]any {
	if v == nil {
		return nil
	}
	schedMap, ok := v.(map[string]any)
	if !ok {
		return nil
	}
	item := make(map[string]any)
	if automaticRestart, ok := schedMap["automaticRestart"].(bool); ok {
		item["automatic_restart"] = automaticRestart
	}
	if onHostMaintenance, ok := schedMap["onHostMaintenance"].(string); ok && onHostMaintenance != "" {
		item["on_host_maintenance"] = onHostMaintenance
	}
	if provisioningModel, ok := schedMap["provisioningModel"].(string); ok && provisioningModel != "" {
		item["provisioning_model"] = provisioningModel
	}
	if preemptible, ok := schedMap["preemptible"].(bool); ok {
		item["preemptible"] = preemptible
	}
	if len(item) == 0 {
		return nil
	}
	return []map[string]any{item}
}

func flattenTemplateShieldedInstanceConfig(v any) []map[string]any {
	if v == nil {
		return nil
	}
	shieldedMap, ok := v.(map[string]any)
	if !ok {
		return nil
	}
	item := make(map[string]any)
	if enableSecureBoot, ok := shieldedMap["enableSecureBoot"].(bool); ok {
		item["enable_secure_boot"] = enableSecureBoot
	}
	if enableVtpm, ok := shieldedMap["enableVtpm"].(bool); ok {
		item["enable_vtpm"] = enableVtpm
	}
	if enableIntegrityMonitoring, ok := shieldedMap["enableIntegrityMonitoring"].(bool); ok {
		item["enable_integrity_monitoring"] = enableIntegrityMonitoring
	}
	if len(item) == 0 {
		return nil
	}
	return []map[string]any{item}
}

func flattenTemplateConfidentialInstanceConfig(v any) []map[string]any {
	if v == nil {
		return nil
	}
	confMap, ok := v.(map[string]any)
	if !ok {
		return nil
	}
	item := make(map[string]any)
	if enableConfidentialCompute, ok := confMap["enableConfidentialCompute"].(bool); ok {
		item["enable_confidential_compute"] = enableConfidentialCompute
	}
	if len(item) == 0 {
		return nil
	}
	return []map[string]any{item}
}

func flattenTemplateReservationAffinity(v any) []map[string]any {
	if v == nil {
		return nil
	}
	resMap, ok := v.(map[string]any)
	if !ok {
		return nil
	}
	item := make(map[string]any)
	if consumeReservationType, ok := resMap["consumeReservationType"].(string); ok && consumeReservationType != "" {
		item["type"] = consumeReservationType
	}
	if len(item) == 0 {
		return nil
	}
	return []map[string]any{item}
}

func flattenTemplateGuestAccelerator(v any) []map[string]any {
	if v == nil {
		return nil
	}
	accelList, ok := v.([]any)
	if !ok || len(accelList) == 0 {
		return nil
	}
	var result []map[string]any
	for _, rawAccel := range accelList {
		accelMap, ok := rawAccel.(map[string]any)
		if !ok {
			continue
		}
		item := make(map[string]any)
		if acceleratorType, ok := accelMap["acceleratorType"].(string); ok && acceleratorType != "" {
			item["type"] = acceleratorType
		}
		if acceleratorCount, ok := accelMap["acceleratorCount"].(float64); ok {
			item["count"] = int64(acceleratorCount)
		}
		if len(item) > 0 {
			result = append(result, item)
		}
	}
	return result
}

func flattenTemplateAdvancedMachineFeatures(v any) []map[string]any {
	if v == nil {
		return nil
	}
	advMap, ok := v.(map[string]any)
	if !ok {
		return nil
	}
	item := make(map[string]any)
	if enableNestedVirtualization, ok := advMap["enableNestedVirtualization"].(bool); ok {
		item["enable_nested_virtualization"] = enableNestedVirtualization
	}
	if threadsPerCore, ok := advMap["threadsPerCore"].(float64); ok {
		item["threads_per_core"] = int64(threadsPerCore)
	}
	if visibleCoreCount, ok := advMap["visibleCoreCount"].(float64); ok {
		item["visible_core_count"] = int64(visibleCoreCount)
	}
	if len(item) == 0 {
		return nil
	}
	return []map[string]any{item}
}
