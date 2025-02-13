// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package appconfig

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_appconfig_deployment", deploymentDataSource)
}

// deploymentDataSource returns the Terraform awscc_appconfig_deployment data source.
// This Terraform data source corresponds to the CloudFormation AWS::AppConfig::Deployment resource.
func deploymentDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ApplicationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The application ID.",
		//	  "type": "string"
		//	}
		"application_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The application ID.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ConfigurationProfileId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The configuration profile ID.",
		//	  "type": "string"
		//	}
		"configuration_profile_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The configuration profile ID.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ConfigurationVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The configuration version to deploy. If deploying an AWS AppConfig hosted configuration version, you can specify either the version number or version label. For all other configurations, you must specify the version number.",
		//	  "type": "string"
		//	}
		"configuration_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The configuration version to deploy. If deploying an AWS AppConfig hosted configuration version, you can specify either the version number or version label. For all other configurations, you must specify the version number.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DeploymentNumber
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The sequence number of the deployment.",
		//	  "type": "string"
		//	}
		"deployment_number": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The sequence number of the deployment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DeploymentStrategyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The deployment strategy ID.",
		//	  "type": "string"
		//	}
		"deployment_strategy_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The deployment strategy ID.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description of the deployment.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description of the deployment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DynamicExtensionParameters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "ExtensionReference": {
		//	        "type": "string"
		//	      },
		//	      "ParameterName": {
		//	        "type": "string"
		//	      },
		//	      "ParameterValue": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"dynamic_extension_parameters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ExtensionReference
					"extension_reference": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: ParameterName
					"parameter_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: ParameterValue
					"parameter_value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The environment ID.",
		//	  "type": "string"
		//	}
		"environment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The environment ID.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KmsKeyIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AWS Key Management Service key identifier (key ID, key alias, or key ARN) provided when the resource was created or updated.",
		//	  "pattern": "^[\\da-f]{8}-[\\da-f]{4}-[\\da-f]{4}-[\\da-f]{4}-[\\da-f]{12}|alias/[a-zA-Z0-9/_-]{1,250}|arn:aws[a-zA-Z-]*:kms:[a-z]{2}(-gov|-iso(b?))?-[a-z]+-\\d{1}:\\d{12}:(key/[0-9a-f-]{36}|alias/[a-zA-Z0-9/_-]{1,250})$",
		//	  "type": "string"
		//	}
		"kms_key_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AWS Key Management Service key identifier (key ID, key alias, or key ARN) provided when the resource was created or updated.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Metadata to assign to the deployment. Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key-value string map. The valid character set is [a-zA-Z1-9+-=._:/]. The tag key can be up to 128 characters and must not start with aws:.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The tag value can be up to 256 characters.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key-value string map. The valid character set is [a-zA-Z1-9+-=._:/]. The tag key can be up to 128 characters and must not start with aws:.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag value can be up to 256 characters.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::AppConfig::Deployment",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::AppConfig::Deployment").WithTerraformTypeName("awscc_appconfig_deployment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"application_id":               "ApplicationId",
		"configuration_profile_id":     "ConfigurationProfileId",
		"configuration_version":        "ConfigurationVersion",
		"deployment_number":            "DeploymentNumber",
		"deployment_strategy_id":       "DeploymentStrategyId",
		"description":                  "Description",
		"dynamic_extension_parameters": "DynamicExtensionParameters",
		"environment_id":               "EnvironmentId",
		"extension_reference":          "ExtensionReference",
		"key":                          "Key",
		"kms_key_identifier":           "KmsKeyIdentifier",
		"parameter_name":               "ParameterName",
		"parameter_value":              "ParameterValue",
		"tags":                         "Tags",
		"value":                        "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
