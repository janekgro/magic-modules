package compute

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/cai2hcl/converters/utils"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/cai2hcl/models"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/caiasset"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/tpgresource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ComputeInstanceGroupCai2hclConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewComputeInstanceGroupCai2hclConverter(provider *schema.Provider) models.Cai2hclConverter {
	schema := provider.ResourcesMap[ComputeInstanceGroupSchemaName].Schema

	return &ComputeInstanceGroupCai2hclConverter{
		name:   ComputeInstanceGroupSchemaName,
		schema: schema,
	}
}

func (c *ComputeInstanceGroupCai2hclConverter) Convert(assets []caiasset.Asset, options *models.ResourceConverterOptions) ([]*models.TerraformResourceBlock, error) {
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
		utils.ParseUrlParamValuesFromAssetName(asset.Name, "//compute.googleapis.com/projects/{{project}}/zones/{{zone}}/instanceGroups/{{name}}", outputFields, hclData)

		if name != "" {
			hclData["name"] = name
		}
		if desc, ok := data["description"]; ok && desc != nil {
			hclData["description"] = desc
		}
		if zone, ok := data["zone"].(string); ok && zone != "" {
			hclData["zone"] = tpgresource.GetResourceNameFromSelfLink(zone)
		}
		hclData["network"] = data["network"]

		if instances, ok := data["instances"].([]any); ok && len(instances) > 0 {
			hclData["instances"] = instances
		}
		hclData["named_port"] = flattenNamedPorts(data["namedPorts"])

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
