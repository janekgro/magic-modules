package compute

import (
	"fmt"
	"sort"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/cai2hcl/converters/utils"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/cai2hcl/models"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/caiasset"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/tpgresource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ComputeInstanceGroupManagerCai2hclConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewComputeInstanceGroupManagerCai2hclConverter(provider *schema.Provider) models.Cai2hclConverter {
	return newComputeInstanceGroupManagerCai2hclConverter(provider, ComputeInstanceGroupManagerSchemaName)
}

func NewComputeRegionInstanceGroupManagerCai2hclConverter(provider *schema.Provider) models.Cai2hclConverter {
	return newComputeInstanceGroupManagerCai2hclConverter(provider, ComputeRegionInstanceGroupManagerSchemaName)
}

func newComputeInstanceGroupManagerCai2hclConverter(provider *schema.Provider, schemaName string) models.Cai2hclConverter {
	schema := provider.ResourcesMap[schemaName].Schema

	return &ComputeInstanceGroupManagerCai2hclConverter{
		name:   schemaName,
		schema: schema,
	}
}

func (c *ComputeInstanceGroupManagerCai2hclConverter) Convert(assets []caiasset.Asset, options *models.ResourceConverterOptions) ([]*models.TerraformResourceBlock, error) {
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
		if c.name == ComputeRegionInstanceGroupManagerSchemaName {
			utils.ParseUrlParamValuesFromAssetName(asset.Name, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/instanceGroupManagers/{{name}}", outputFields, hclData)
		} else {
			utils.ParseUrlParamValuesFromAssetName(asset.Name, "//compute.googleapis.com/projects/{{project}}/zones/{{zone}}/instanceGroupManagers/{{name}}", outputFields, hclData)
		}

		if name != "" {
			hclData["name"] = name
		}
		if desc, ok := data["description"]; ok && desc != nil {
			hclData["description"] = desc
		}
		if bin, ok := data["baseInstanceName"]; ok && bin != nil {
			hclData["base_instance_name"] = bin
		}
		if reg, ok := data["region"].(string); ok && reg != "" {
			hclData["region"] = tpgresource.GetResourceNameFromSelfLink(reg)
		}
		if z, ok := data["zone"].(string); ok && z != "" {
			hclData["zone"] = tpgresource.GetResourceNameFromSelfLink(z)
		}
		hclData["list_managed_instances_results"] = data["listManagedInstancesResults"]

		if targetSize, ok := data["targetSize"].(float64); ok {
			hclData["target_size"] = int64(targetSize)
		}
		if targetStoppedSize, ok := data["targetStoppedSize"].(float64); ok {
			hclData["target_stopped_size"] = int64(targetStoppedSize)
		}
		if targetSuspendedSize, ok := data["targetSuspendedSize"].(float64); ok {
			hclData["target_suspended_size"] = int64(targetSuspendedSize)
		}

		hclData["version"] = flattenVersions(data["versions"], data["instanceTemplate"])
		hclData["distribution_policy_target_shape"] = flattenDistributionPolicyTargetShape(data["distributionPolicy"])
		hclData["distribution_policy_zones"] = flattenDistributionPolicyZones(data["distributionPolicy"])
		hclData["all_instances_config"] = flattenAllInstancesConfig(data["allInstancesConfig"])
		hclData["auto_healing_policies"] = flattenAutoHealingPolicies(data["autoHealingPolicies"])
		hclData["update_policy"] = flattenUpdatePolicy(data["updatePolicy"], data["instanceFlexibilityPolicy"] != nil)
		hclData["named_port"] = flattenNamedPorts(data["namedPorts"])
		if tp, ok := data["targetPools"].([]any); ok && len(tp) > 0 {
			hclData["target_pools"] = tp
		}
		hclData["stateful_disk"] = flattenStatefulDisks(data["statefulPolicy"])
		hclData["stateful_internal_ip"] = flattenStatefulIPs(data["statefulPolicy"], "internalIPs")
		hclData["stateful_external_ip"] = flattenStatefulIPs(data["statefulPolicy"], "externalIPs")
		hclData["standby_policy"] = flattenStandbyPolicy(data["standbyPolicy"])
		hclData["instance_lifecycle_policy"] = flattenInstanceLifecyclePolicy(data["instanceLifecyclePolicy"])
		hclData["resource_policies"] = flattenResourcePolicies(data["resourcePolicies"])
		hclData["params"] = flattenParams(data["params"])
		hclData["target_size_policy"] = flattenTargetSizePolicy(data["targetSizePolicy"])
		hclData["instance_flexibility_policy"] = flattenInstanceFlexibilityPolicy(data["instanceFlexibilityPolicy"])

		ctyVal, err := utils.MapToCtyValWithSchema(hclData, c.schema)
		if err != nil {
			return nil, fmt.Errorf("error converting map to cty.Value: %w", err)
		}

		block := &models.TerraformResourceBlock{
			Labels: []string{resourceType, name},
			Value:  ctyVal,
		}
		blocks = append(blocks, block)
	}

	return blocks, nil
}

func flattenVersions(v any, instanceTemplate any) []map[string]any {
	if versionsList, ok := v.([]any); ok && len(versionsList) > 0 {
		var result []map[string]any
		for _, rawVersion := range versionsList {
			versionMap, ok := rawVersion.(map[string]any)
			if !ok {
				continue
			}
			item := make(map[string]any)
			if versionName, ok := versionMap["name"].(string); ok && versionName != "" {
				item["name"] = versionName
			}
			if it, ok := versionMap["instanceTemplate"].(string); ok && it != "" {
				item["instance_template"] = it
			}
			if targetSize, ok := versionMap["targetSize"].(map[string]any); ok {
				sizeMap := make(map[string]any)
				if fixed, ok := targetSize["fixed"].(float64); ok {
					sizeMap["fixed"] = int64(fixed)
				}
				if percent, ok := targetSize["percent"].(float64); ok {
					sizeMap["percent"] = int64(percent)
				}
				if len(sizeMap) > 0 {
					item["target_size"] = []map[string]any{sizeMap}
				}
			}
			if len(item) > 0 {
				result = append(result, item)
			}
		}
		if len(result) > 0 {
			return result
		}
	}
	if itStr, ok := instanceTemplate.(string); ok && itStr != "" {
		return []map[string]any{
			{"instance_template": itStr},
		}
	}
	return nil
}

func flattenAllInstancesConfig(v any) []map[string]any {
	original := toMap(v)
	if original == nil {
		return nil
	}
	props := toMap(original["properties"])
	if props == nil {
		props = original
	}
	configMap := make(map[string]any)
	if labels, ok := props["labels"].(map[string]any); ok && len(labels) > 0 {
		configMap["labels"] = labels
	}
	if metadata, ok := props["metadata"].(map[string]any); ok && len(metadata) > 0 {
		configMap["metadata"] = metadata
	}
	if len(configMap) == 0 {
		return nil
	}
	return []map[string]any{configMap}
}

func flattenAutoHealingPolicies(v any) []map[string]any {
	policiesList := toSlice(v)
	if len(policiesList) == 0 {
		return nil
	}
	var result []map[string]any
	for _, rawPolicy := range policiesList {
		policyMap := toMap(rawPolicy)
		if policyMap == nil {
			continue
		}
		item := make(map[string]any)
		if healthCheck, ok := policyMap["healthCheck"].(string); ok && healthCheck != "" {
			item["health_check"] = healthCheck
		}
		if initialDelaySec, ok := policyMap["initialDelaySec"].(float64); ok {
			item["initial_delay_sec"] = int64(initialDelaySec)
		}
		if len(item) > 0 {
			result = append(result, item)
		}
	}
	return result
}

func toMap(v any) map[string]any {
	if m, ok := v.(map[string]any); ok {
		return m
	}
	return nil
}

func toSlice(v any) []any {
	if s, ok := v.([]any); ok {
		return s
	}
	return nil
}

func flattenDistributionPolicyTargetShape(v any) any {
	m := toMap(v)
	if m == nil {
		return nil
	}
	if targetShape, ok := m["targetShape"].(string); ok && targetShape != "" {
		return targetShape
	}
	return nil
}

func flattenDistributionPolicyZones(v any) []string {
	m := toMap(v)
	if m == nil {
		return nil
	}
	zonesList := toSlice(m["zones"])
	if len(zonesList) == 0 {
		return nil
	}

	var zones []string
	for _, rawZone := range zonesList {
		var z string
		if zm := toMap(rawZone); zm != nil {
			z, _ = zm["zone"].(string)
		} else {
			z, _ = rawZone.(string)
		}
		if z != "" {
			zones = append(zones, tpgresource.GetResourceNameFromSelfLink(z))
		}
	}
	return zones
}

func flattenUpdatePolicy(v any, hasFlexibility bool) []map[string]any {
	policyMap := toMap(v)
	if policyMap == nil {
		return nil
	}
	item := make(map[string]any)
	if typeVal, ok := policyMap["type"].(string); ok && typeVal != "" {
		item["type"] = typeVal
	} else if hasFlexibility {
		item["type"] = "OPPORTUNISTIC"
	} else {
		item["type"] = "PROACTIVE"
	}
	if minimalAction, ok := policyMap["minimalAction"].(string); ok && minimalAction != "" {
		item["minimal_action"] = minimalAction
	} else {
		item["minimal_action"] = "REPLACE"
	}
	if mostDisruptive, ok := policyMap["mostDisruptiveAllowedAction"].(string); ok && mostDisruptive != "" {
		item["most_disruptive_allowed_action"] = mostDisruptive
	}
	if minReadySec, ok := policyMap["minReadySec"].(float64); ok {
		item["min_ready_sec"] = int64(minReadySec)
	}
	if instanceRedistributionType, ok := policyMap["instanceRedistributionType"].(string); ok && instanceRedistributionType != "" {
		item["instance_redistribution_type"] = instanceRedistributionType
	}
	if maxSurge, ok := policyMap["maxSurge"].(map[string]any); ok {
		if fixed, ok := maxSurge["fixed"].(float64); ok {
			item["max_surge_fixed"] = int64(fixed)
		} else if percent, ok := maxSurge["percent"].(float64); ok {
			item["max_surge_percent"] = int64(percent)
		}
	}
	if maxUnavailable, ok := policyMap["maxUnavailable"].(map[string]any); ok {
		if fixed, ok := maxUnavailable["fixed"].(float64); ok {
			item["max_unavailable_fixed"] = int64(fixed)
		} else if percent, ok := maxUnavailable["percent"].(float64); ok {
			item["max_unavailable_percent"] = int64(percent)
		}
	}
	if item["max_surge_fixed"] == nil && item["max_surge_percent"] == nil &&
		item["max_unavailable_fixed"] == nil && item["max_unavailable_percent"] == nil {
		item["max_surge_fixed"] = int64(3)
	}
	if replacementMethod, ok := policyMap["replacementMethod"].(string); ok && replacementMethod != "" {
		item["replacement_method"] = replacementMethod
	}
	if len(item) == 0 {
		return nil
	}
	return []map[string]any{item}
}

func flattenNamedPorts(v any) []map[string]any {
	if v == nil {
		return nil
	}
	portsList, ok := v.([]any)
	if !ok || len(portsList) == 0 {
		return nil
	}
	var result []map[string]any
	for _, rawPort := range portsList {
		portMap, ok := rawPort.(map[string]any)
		if !ok {
			continue
		}
		item := make(map[string]any)
		if name, ok := portMap["name"].(string); ok && name != "" {
			item["name"] = name
		}
		if port, ok := portMap["port"].(float64); ok {
			item["port"] = int64(port)
		}
		if len(item) > 0 {
			result = append(result, item)
		}
	}
	return result
}

func flattenStatefulDisks(v any) []map[string]any {
	if v == nil {
		return nil
	}
	policyMap, ok := v.(map[string]any)
	if !ok {
		return nil
	}
	preservedState, ok := policyMap["preservedState"].(map[string]any)
	if !ok {
		return nil
	}
	disksMap, ok := preservedState["disks"].(map[string]any)
	if !ok || len(disksMap) == 0 {
		return nil
	}
	var result []map[string]any
	for deviceName, rawConfig := range disksMap {
		item := make(map[string]any)
		item["device_name"] = deviceName
		if configMap, ok := rawConfig.(map[string]any); ok {
			if autoDelete, ok := configMap["autoDelete"].(string); ok && autoDelete != "" {
				item["delete_rule"] = autoDelete
			}
		}
		result = append(result, item)
	}
	return result
}

func flattenStatefulIPs(v any, key string) []map[string]any {
	if v == nil {
		return nil
	}
	policyMap, ok := v.(map[string]any)
	if !ok {
		return nil
	}
	preservedState, ok := policyMap["preservedState"].(map[string]any)
	if !ok {
		return nil
	}
	ipsMap, ok := preservedState[key].(map[string]any)
	if !ok || len(ipsMap) == 0 {
		return nil
	}
	var result []map[string]any
	for interfaceName, rawConfig := range ipsMap {
		item := make(map[string]any)
		item["interface_name"] = interfaceName
		if configMap, ok := rawConfig.(map[string]any); ok {
			if autoDelete, ok := configMap["autoDelete"].(string); ok && autoDelete != "" {
				item["delete_rule"] = autoDelete
			}
		}
		result = append(result, item)
	}
	return result
}

func flattenStandbyPolicy(v any) []map[string]any {
	if v == nil {
		return nil
	}
	policyMap, ok := v.(map[string]any)
	if !ok {
		return nil
	}
	item := make(map[string]any)
	if initialDelaySec, ok := policyMap["initialDelaySec"].(float64); ok {
		item["initial_delay_sec"] = int64(initialDelaySec)
	}
	if mode, ok := policyMap["mode"].(string); ok && mode != "" {
		item["mode"] = mode
	}
	if len(item) == 0 {
		return nil
	}
	return []map[string]any{item}
}

func flattenInstanceLifecyclePolicy(v any) []map[string]any {
	if v == nil {
		return nil
	}
	policyMap, ok := v.(map[string]any)
	if !ok {
		return nil
	}
	item := make(map[string]any)
	if defaultAction, ok := policyMap["defaultActionOnFailure"].(string); ok && defaultAction != "" {
		item["default_action_on_failure"] = defaultAction
	}
	if onFailedHealthCheck, ok := policyMap["onFailedHealthCheck"].(string); ok && onFailedHealthCheck != "" {
		item["on_failed_health_check"] = onFailedHealthCheck
	}
	if forceUpdate, ok := policyMap["forceUpdateOnRepair"].(string); ok && forceUpdate != "" {
		item["force_update_on_repair"] = forceUpdate
	}
	if onRepair, ok := policyMap["onRepair"].(map[string]any); ok {
		onRepairItem := make(map[string]any)
		if allowChangingZone, ok := onRepair["allowChangingZone"].(string); ok && allowChangingZone != "" {
			onRepairItem["allow_changing_zone"] = allowChangingZone
		}
		if len(onRepairItem) > 0 {
			item["on_repair"] = []map[string]any{onRepairItem}
		}
	}
	if len(item) == 0 {
		return nil
	}
	return []map[string]any{item}
}

func flattenParams(v any) []map[string]any {
	if v == nil {
		return nil
	}
	paramsMap, ok := v.(map[string]any)
	if !ok {
		return nil
	}
	item := make(map[string]any)
	if tags, ok := paramsMap["resourceManagerTags"].(map[string]any); ok && len(tags) > 0 {
		item["resource_manager_tags"] = tags
	}
	if len(item) == 0 {
		return nil
	}
	return []map[string]any{item}
}

func flattenResourcePolicies(v any) []map[string]any {
	if v == nil {
		return nil
	}
	if policyStr, ok := v.(string); ok && policyStr != "" {
		return []map[string]any{{"workload_policy": policyStr}}
	}
	if policyList, ok := v.([]any); ok && len(policyList) > 0 {
		var strList []string
		for _, item := range policyList {
			if s, ok := item.(string); ok && s != "" {
				strList = append(strList, s)
			}
		}
		if len(strList) > 0 {
			return []map[string]any{{"workload_policy": strList[0]}}
		}
	}
	if policyMap, ok := v.(map[string]any); ok {
		item := make(map[string]any)
		if workloadPolicy, ok := policyMap["workloadPolicy"].(string); ok && workloadPolicy != "" {
			item["workload_policy"] = workloadPolicy
		}
		if len(item) > 0 {
			return []map[string]any{item}
		}
	}
	return nil
}

func flattenTargetSizePolicy(v any) []map[string]any {
	if v == nil {
		return nil
	}
	policyMap, ok := v.(map[string]any)
	if !ok {
		return nil
	}
	item := make(map[string]any)
	if mode, ok := policyMap["mode"].(string); ok && mode != "" {
		item["mode"] = mode
	}
	if len(item) == 0 {
		return nil
	}
	return []map[string]any{item}
}

func flattenInstanceSelectionDisks(rawDisks any) []map[string]any {
	var rawDiskList []any
	if dSlice := toSlice(rawDisks); len(dSlice) > 0 {
		rawDiskList = dSlice
	} else if dMap := toMap(rawDisks); len(dMap) > 0 {
		for _, v := range dMap {
			rawDiskList = append(rawDiskList, v)
		}
	}

	if len(rawDiskList) == 0 {
		return nil
	}

	var disks []map[string]any
	for _, rawDisk := range rawDiskList {
		dMap := toMap(rawDisk)
		if len(dMap) == 0 {
			continue
		}

		diskData := make(map[string]any)
		if autoDelete, ok := dMap["autoDelete"].(bool); ok {
			diskData["auto_delete"] = autoDelete
		} else if autoDelete, ok := dMap["auto_delete"].(bool); ok {
			diskData["auto_delete"] = autoDelete
		}
		if boot, ok := dMap["boot"].(bool); ok {
			diskData["boot"] = boot
		}
		if deviceName, ok := dMap["deviceName"].(string); ok && deviceName != "" {
			diskData["device_name"] = deviceName
		} else if deviceName, ok := dMap["device_name"].(string); ok && deviceName != "" {
			diskData["device_name"] = deviceName
		}
		initParams := toMap(dMap["initializeParams"])
		if initParams == nil {
			initParams = toMap(dMap["initialize_params"])
		}
		if diskSize, ok := dMap["diskSizeGb"]; ok {
			if sizeFloat, isFloat := diskSize.(float64); isFloat {
				diskData["disk_size_gb"] = int(sizeFloat)
			} else {
				diskData["disk_size_gb"] = diskSize
			}
		} else if initParams != nil {
			if diskSize, ok := initParams["diskSizeGb"]; ok {
				if sizeFloat, isFloat := diskSize.(float64); isFloat {
					diskData["disk_size_gb"] = int(sizeFloat)
				} else {
					diskData["disk_size_gb"] = diskSize
				}
			} else if diskSize, ok := initParams["disk_size_gb"]; ok {
				if sizeFloat, isFloat := initParams["disk_size_gb"].(float64); isFloat {
					diskData["disk_size_gb"] = int(sizeFloat)
				} else {
					diskData["disk_size_gb"] = diskSize
				}
			}
		}
		if diskType, ok := dMap["diskType"].(string); ok && diskType != "" {
			diskData["disk_type"] = diskType
		} else if initParams != nil {
			if dt, ok := initParams["diskType"].(string); ok && dt != "" {
				diskData["disk_type"] = dt
			} else if dt, ok := initParams["disk_type"].(string); ok && dt != "" {
				diskData["disk_type"] = dt
			}
		}
		if srcImg, ok := dMap["sourceImage"].(string); ok && srcImg != "" {
			diskData["source_image"] = srcImg
		} else if initParams != nil {
			if si, ok := initParams["sourceImage"].(string); ok && si != "" {
				diskData["source_image"] = si
			} else if si, ok := initParams["source_image"].(string); ok && si != "" {
				diskData["source_image"] = si
			}
		}
		if len(diskData) > 0 {
			disks = append(disks, diskData)
		}
	}
	return disks
}

func flattenInstanceSelection(itemMap map[string]any, defaultName string) map[string]any {
	if itemMap == nil {
		return nil
	}
	selData := make(map[string]any)
	name, _ := itemMap["name"].(string)
	if name == "" {
		name = defaultName
	}
	if name != "" {
		selData["name"] = name
	}
	if rank, ok := itemMap["rank"]; ok {
		if rFloat, isFloat := rank.(float64); isFloat {
			selData["rank"] = int(rFloat)
		} else {
			selData["rank"] = rank
		}
	}
	rawMts := itemMap["machineTypes"]
	if rawMts == nil {
		rawMts = itemMap["machine_types"]
	}
	if mts := toSlice(rawMts); len(mts) > 0 {
		var strMts []string
		for _, mt := range mts {
			if strMt, ok := mt.(string); ok && strMt != "" {
				strMts = append(strMts, strMt)
			}
		}
		if len(strMts) > 0 {
			selData["machine_types"] = strMts
		}
	}
	if minCpu, ok := itemMap["minCpuPlatform"].(string); ok && minCpu != "" {
		selData["min_cpu_platform"] = minCpu
	} else if minCpu, ok := itemMap["min_cpu_platform"].(string); ok && minCpu != "" {
		selData["min_cpu_platform"] = minCpu
	}

	if disks := flattenInstanceSelectionDisks(itemMap["disks"]); len(disks) > 0 {
		selData["disks"] = disks
	}
	return selData
}

func flattenInstanceFlexibilityPolicy(v any) []map[string]any {
	m := toMap(v)
	if m == nil {
		return nil
	}

	rawSelections := m["instanceSelections"]
	if rawSelections == nil {
		rawSelections = m["instance_selections"]
	}
	if rawSelections == nil {
		return nil
	}

	var selections []map[string]any
	if selMap := toMap(rawSelections); selMap != nil {
		for name, item := range selMap {
			if selData := flattenInstanceSelection(toMap(item), name); len(selData) > 0 {
				selections = append(selections, selData)
			}
		}
	} else if selSlice := toSlice(rawSelections); len(selSlice) > 0 {
		for _, item := range selSlice {
			if selData := flattenInstanceSelection(toMap(item), ""); len(selData) > 0 {
				selections = append(selections, selData)
			}
		}
	}

	if len(selections) == 0 {
		return nil
	}

	sort.Slice(selections, func(i, j int) bool {
		nameI, _ := selections[i]["name"].(string)
		nameJ, _ := selections[j]["name"].(string)
		return nameI < nameJ
	})

	return []map[string]any{
		{
			"instance_selections": selections,
		},
	}
}
