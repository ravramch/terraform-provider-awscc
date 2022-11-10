// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iot_role_alias", roleAliasDataSource)
}

// roleAliasDataSource returns the Terraform awscc_iot_role_alias data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoT::RoleAlias resource.
func roleAliasDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"credential_duration_seconds": {
			// Property: CredentialDurationSeconds
			// CloudFormation resource type schema:
			//
			//	{
			//	  "default": 3600,
			//	  "maximum": 43200,
			//	  "minimum": 900,
			//	  "type": "integer"
			//	}
			Type:     types.Int64Type,
			Computed: true,
		},
		"role_alias": {
			// Property: RoleAlias
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 128,
			//	  "minLength": 1,
			//	  "pattern": "[\\w=,@-]+",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"role_alias_arn": {
			// Property: RoleAliasArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 128,
			//	  "minLength": 1,
			//	  "pattern": "[\\w=,@-]+",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 2048,
			//	  "minLength": 20,
			//	  "pattern": "arn:(aws[a-zA-Z-]*)?:iam::\\d{12}:role/?[a-zA-Z_0-9+=,.@\\-_/]+",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "A key-value pair to associate with a resource.",
			//	    "properties": {
			//	      "Key": {
			//	        "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//	        "maxLength": 127,
			//	        "minLength": 1,
			//	        "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//	        "maxLength": 255,
			//	        "minLength": 1,
			//	        "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Value",
			//	      "Key"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array"
			//	}
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::IoT::RoleAlias",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::RoleAlias").WithTerraformTypeName("awscc_iot_role_alias")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"credential_duration_seconds": "CredentialDurationSeconds",
		"key":                         "Key",
		"role_alias":                  "RoleAlias",
		"role_alias_arn":              "RoleAliasArn",
		"role_arn":                    "RoleArn",
		"tags":                        "Tags",
		"value":                       "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
