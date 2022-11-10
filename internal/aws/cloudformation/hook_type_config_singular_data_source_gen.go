// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_cloudformation_hook_type_config", hookTypeConfigDataSource)
}

// hookTypeConfigDataSource returns the Terraform awscc_cloudformation_hook_type_config data source.
// This Terraform data source corresponds to the CloudFormation AWS::CloudFormation::HookTypeConfig resource.
func hookTypeConfigDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"configuration": {
			// Property: Configuration
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The configuration data for the extension, in this account and region.",
			//	  "pattern": "[\\s\\S]+",
			//	  "type": "string"
			//	}
			Description: "The configuration data for the extension, in this account and region.",
			Type:        types.StringType,
			Computed:    true,
		},
		"configuration_alias": {
			// Property: ConfigurationAlias
			// CloudFormation resource type schema:
			//
			//	{
			//	  "default": "default",
			//	  "description": "An alias by which to refer to this extension configuration data.",
			//	  "enum": [
			//	    "default"
			//	  ],
			//	  "pattern": "^[a-zA-Z0-9]{1,256}$",
			//	  "type": "string"
			//	}
			Description: "An alias by which to refer to this extension configuration data.",
			Type:        types.StringType,
			Computed:    true,
		},
		"configuration_arn": {
			// Property: ConfigurationArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) for the configuration data, in this account and region.",
			//	  "pattern": "^arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type(-configuration)?/hook/.+$",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) for the configuration data, in this account and region.",
			Type:        types.StringType,
			Computed:    true,
		},
		"type_arn": {
			// Property: TypeArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) of the type without version number.",
			//	  "pattern": "^arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type/hook/.+$",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) of the type without version number.",
			Type:        types.StringType,
			Computed:    true,
		},
		"type_name": {
			// Property: TypeName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the type being registered.\n\nWe recommend that type names adhere to the following pattern: company_or_organization::service::type.",
			//	  "pattern": "^[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}$",
			//	  "type": "string"
			//	}
			Description: "The name of the type being registered.\n\nWe recommend that type names adhere to the following pattern: company_or_organization::service::type.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::CloudFormation::HookTypeConfig",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFormation::HookTypeConfig").WithTerraformTypeName("awscc_cloudformation_hook_type_config")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"configuration":       "Configuration",
		"configuration_alias": "ConfigurationAlias",
		"configuration_arn":   "ConfigurationArn",
		"type_arn":            "TypeArn",
		"type_name":           "TypeName",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
