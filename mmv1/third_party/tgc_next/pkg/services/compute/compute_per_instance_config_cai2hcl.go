package compute

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/cai2hcl/converters/utils"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/cai2hcl/models"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/caiasset"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/tgcresource"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/tpgresource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ComputePerInstanceConfigCai2hclConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewComputePerInstanceConfigCai2hclConverter(provider *schema.Provider) models.Cai2hclConverter {
	schema := provider.ResourcesMap[ComputePerInstanceConfigSchemaName].Schema

	return &ComputePerInstanceConfigCai2hclConverter{
		name:   ComputePerInstanceConfigSchemaName,
		schema: schema,
	}
}

func (c *ComputePerInstanceConfigCai2hclConverter) Convert(assets []caiasset.Asset, options *models.ResourceConverterOptions) ([]*models.TerraformResourceBlock, error) {
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

		hclData := make(map[string]any)
		outputFields := map[string]struct{}{"project": {}}
		utils.ParseUrlParamValuesFromAssetName(asset.Name, "//compute.googleapis.com/projects/{{project}}/zones/{{zone}}/instanceGroupManagers/{{instance_group_manager}}/perInstanceConfigs/{{name}}", outputFields, hclData)

		if name != "" {
			hclData["name"] = name
		}
		if igm, ok := data["instanceGroupManager"].(string); ok && igm != "" {
			hclData["instance_group_manager"] = tpgresource.GetResourceNameFromSelfLink(igm)
		} else if igmName := tgcresource.ParseFieldValue(asset.Name, "instanceGroupManagers"); igmName != "" {
			hclData["instance_group_manager"] = igmName
		}
		if zone, ok := data["zone"].(string); ok && zone != "" {
			hclData["zone"] = zone
		}
		if preservedState := flattenPerInstanceConfigPreservedState(data["preservedState"]); len(preservedState) > 0 {
			hclData["preserved_state"] = preservedState
		}

		ctyVal, err := utils.MapToCtyValWithSchema(hclData, c.schema)
		if err != nil {
			return nil, fmt.Errorf("error converting map to cty.Value: %w", err)
		}

		block := &models.TerraformResourceBlock{
			Labels: []string{c.name, name},
			Value:  ctyVal,
		}
		blocks = append(blocks, block)
	}

	return blocks, nil
}

func flattenPerInstanceConfigPreservedState(v any) []map[string]any {
	if v == nil {
		return nil
	}
	original, ok := v.(map[string]any)
	if !ok || len(original) == 0 {
		return nil
	}
	transformed := make(map[string]any)
	if metadata, ok := original["metadata"].(map[string]any); ok && len(metadata) > 0 {
		m := make(map[string]string)
		for k, val := range metadata {
			if s, ok := val.(string); ok {
				m[k] = s
			}
		}
		if len(m) > 0 {
			transformed["metadata"] = m
		}
	}
	if disks, ok := original["disks"].(map[string]any); ok && len(disks) > 0 {
		var diskList []map[string]any
		for devName, diskRaw := range disks {
			if diskObj, ok := diskRaw.(map[string]any); ok {
				d := map[string]any{
					"device_name": devName,
				}
				if autoDelete, ok := diskObj["autoDelete"].(string); ok && autoDelete != "" {
					d["delete_rule"] = autoDelete
				}
				if source, ok := diskObj["source"].(string); ok && source != "" {
					d["source"] = source
				}
				if mode, ok := diskObj["mode"].(string); ok && mode != "" {
					d["mode"] = mode
				}
				diskList = append(diskList, d)
			}
		}
		if len(diskList) > 0 {
			transformed["disk"] = diskList
		}
	}
	if internalIPs, ok := original["internalIPs"].(map[string]any); ok && len(internalIPs) > 0 {
		var ipList []map[string]any
		for ifaceName, raw := range internalIPs {
			if ipObj, ok := raw.(map[string]any); ok {
				item := map[string]any{
					"interface_name": ifaceName,
				}
				if autoDelete, ok := ipObj["autoDelete"].(string); ok && autoDelete != "" {
					item["auto_delete"] = autoDelete
				}
				if ipAddr, ok := ipObj["ipAddress"].(map[string]any); ok {
					if addr, ok := ipAddr["address"].(string); ok && addr != "" {
						item["ip_address"] = []map[string]any{{"address": addr}}
					}
				}
				ipList = append(ipList, item)
			}
		}
		if len(ipList) > 0 {
			transformed["internal_ip"] = ipList
		}
	}
	if externalIPs, ok := original["externalIPs"].(map[string]any); ok && len(externalIPs) > 0 {
		var ipList []map[string]any
		for ifaceName, raw := range externalIPs {
			if ipObj, ok := raw.(map[string]any); ok {
				item := map[string]any{
					"interface_name": ifaceName,
				}
				if autoDelete, ok := ipObj["autoDelete"].(string); ok && autoDelete != "" {
					item["auto_delete"] = autoDelete
				}
				if ipAddr, ok := ipObj["ipAddress"].(map[string]any); ok {
					if addr, ok := ipAddr["address"].(string); ok && addr != "" {
						item["ip_address"] = []map[string]any{{"address": addr}}
					}
				}
				ipList = append(ipList, item)
			}
		}
		if len(ipList) > 0 {
			transformed["external_ip"] = ipList
		}
	}
	if len(transformed) == 0 {
		return nil
	}
	return []map[string]any{transformed}
}
