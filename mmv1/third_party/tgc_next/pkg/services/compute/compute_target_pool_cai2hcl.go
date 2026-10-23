package compute

import (
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/cai2hcl/converters/utils"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/cai2hcl/models"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/caiasset"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/tpgresource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ComputeTargetPoolCai2hclConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewComputeTargetPoolCai2hclConverter(provider *schema.Provider) models.Cai2hclConverter {
	schema := provider.ResourcesMap[ComputeTargetPoolSchemaName].Schema

	return &ComputeTargetPoolCai2hclConverter{
		name:   ComputeTargetPoolSchemaName,
		schema: schema,
	}
}

func (c *ComputeTargetPoolCai2hclConverter) Convert(assets []caiasset.Asset, options *models.ResourceConverterOptions) ([]*models.TerraformResourceBlock, error) {
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
		utils.ParseUrlParamValuesFromAssetName(asset.Name, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/targetPools/{{name}}", outputFields, hclData)

		if name != "" {
			hclData["name"] = name
		}
		if desc, ok := data["description"]; ok && desc != nil {
			hclData["description"] = desc
		}
		if region, ok := data["region"].(string); ok && region != "" {
			hclData["region"] = tpgresource.GetResourceNameFromSelfLink(region)
		}
		if backupPool, ok := data["backupPool"].(string); ok && backupPool != "" {
			hclData["backup_pool"] = backupPool
		}
		if failoverRatio, ok := data["failoverRatio"].(float64); ok {
			hclData["failover_ratio"] = failoverRatio
		}
		if healthChecks, ok := data["healthChecks"].([]any); ok && len(healthChecks) > 0 {
			hclData["health_checks"] = healthChecks
		}
		if sessionAffinity, ok := data["sessionAffinity"].(string); ok && sessionAffinity != "" {
			hclData["session_affinity"] = sessionAffinity
		}
		if securityPolicy, ok := data["securityPolicy"].(string); ok && securityPolicy != "" {
			hclData["security_policy"] = securityPolicy
		}
		if rawInstances, ok := data["instances"].([]any); ok && len(rawInstances) > 0 {
			cleanInstances := make([]string, 0, len(rawInstances))
			for _, raw := range rawInstances {
				if inst, ok := raw.(string); ok && inst != "" {
					cleanInstances = append(cleanInstances, strings.TrimPrefix(inst, "//compute.googleapis.com/"))
				}
			}
			if len(cleanInstances) > 0 {
				hclData["instances"] = cleanInstances
			}
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
