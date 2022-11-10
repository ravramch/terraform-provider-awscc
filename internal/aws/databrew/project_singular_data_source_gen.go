// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package databrew

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_databrew_project", projectDataSource)
}

// projectDataSource returns the Terraform awscc_databrew_project data source.
// This Terraform data source corresponds to the CloudFormation AWS::DataBrew::Project resource.
func projectDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"dataset_name": {
			// Property: DatasetName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Dataset name",
			//	  "maxLength": 255,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "Dataset name",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Project name",
			//	  "maxLength": 255,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "Project name",
			Type:        types.StringType,
			Computed:    true,
		},
		"recipe_name": {
			// Property: RecipeName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Recipe name",
			//	  "maxLength": 255,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "Recipe name",
			Type:        types.StringType,
			Computed:    true,
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Role arn",
			//	  "type": "string"
			//	}
			Description: "Role arn",
			Type:        types.StringType,
			Computed:    true,
		},
		"sample": {
			// Property: Sample
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "Sample",
			//	  "properties": {
			//	    "Size": {
			//	      "description": "Sample size",
			//	      "minimum": 1,
			//	      "type": "integer"
			//	    },
			//	    "Type": {
			//	      "description": "Sample type",
			//	      "enum": [
			//	        "FIRST_N",
			//	        "LAST_N",
			//	        "RANDOM"
			//	      ],
			//	      "type": "string"
			//	    }
			//	  },
			//	  "required": [
			//	    "Type"
			//	  ],
			//	  "type": "object"
			//	}
			Description: "Sample",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"size": {
						// Property: Size
						Description: "Sample size",
						Type:        types.Int64Type,
						Computed:    true,
					},
					"type": {
						// Property: Type
						Description: "Sample type",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
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
			//	    "description": "A key-value pair to associate with a resource.",
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
			//	      "Value",
			//	      "Key"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": false
			//	}
			Attributes: tfsdk.ListNestedAttributes(
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
		Description: "Data Source schema for AWS::DataBrew::Project",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataBrew::Project").WithTerraformTypeName("awscc_databrew_project")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"dataset_name": "DatasetName",
		"key":          "Key",
		"name":         "Name",
		"recipe_name":  "RecipeName",
		"role_arn":     "RoleArn",
		"sample":       "Sample",
		"size":         "Size",
		"tags":         "Tags",
		"type":         "Type",
		"value":        "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
