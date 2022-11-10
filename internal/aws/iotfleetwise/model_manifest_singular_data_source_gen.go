// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iotfleetwise

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iotfleetwise_model_manifest", modelManifestDataSource)
}

// modelManifestDataSource returns the Terraform awscc_iotfleetwise_model_manifest data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoTFleetWise::ModelManifest resource.
func modelManifestDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"creation_time": {
			// Property: CreationTime
			// CloudFormation resource type schema:
			//
			//	{
			//	  "format": "date-time",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 2048,
			//	  "minLength": 1,
			//	  "pattern": "",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"last_modification_time": {
			// Property: LastModificationTime
			// CloudFormation resource type schema:
			//
			//	{
			//	  "format": "date-time",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 100,
			//	  "minLength": 1,
			//	  "pattern": "^[a-zA-Z\\d\\-_:]+$",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"nodes": {
			// Property: Nodes
			// CloudFormation resource type schema:
			//
			//	{
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "type": "string"
			//	  },
			//	  "minItems": 1,
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Type:     types.SetType{ElemType: types.StringType},
			Computed: true,
		},
		"signal_catalog_arn": {
			// Property: SignalCatalogArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			//
			//	{
			//	  "default": "DRAFT",
			//	  "enum": [
			//	    "ACTIVE",
			//	    "DRAFT"
			//	  ],
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
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "Key": {
			//	        "maxLength": 128,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "maxLength": 256,
			//	        "minLength": 0,
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Key",
			//	      "Value"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "maxItems": 50,
			//	  "minItems": 0,
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
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
		Description: "Data Source schema for AWS::IoTFleetWise::ModelManifest",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTFleetWise::ModelManifest").WithTerraformTypeName("awscc_iotfleetwise_model_manifest")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                    "Arn",
		"creation_time":          "CreationTime",
		"description":            "Description",
		"key":                    "Key",
		"last_modification_time": "LastModificationTime",
		"name":                   "Name",
		"nodes":                  "Nodes",
		"signal_catalog_arn":     "SignalCatalogArn",
		"status":                 "Status",
		"tags":                   "Tags",
		"value":                  "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
