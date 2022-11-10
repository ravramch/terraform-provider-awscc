// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_appstream_app_block", appBlockDataSource)
}

// appBlockDataSource returns the Terraform awscc_appstream_app_block data source.
// This Terraform data source corresponds to the CloudFormation AWS::AppStream::AppBlock resource.
func appBlockDataSource(ctx context.Context) (datasource.DataSource, error) {
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
		"created_time": {
			// Property: CreatedTime
			// CloudFormation resource type schema:
			//
			//	{
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
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"display_name": {
			// Property: DisplayName
			// CloudFormation resource type schema:
			//
			//	{
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
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"setup_script_details": {
			// Property: SetupScriptDetails
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "properties": {
			//	    "ExecutableParameters": {
			//	      "type": "string"
			//	    },
			//	    "ExecutablePath": {
			//	      "type": "string"
			//	    },
			//	    "ScriptS3Location": {
			//	      "additionalProperties": false,
			//	      "properties": {
			//	        "S3Bucket": {
			//	          "type": "string"
			//	        },
			//	        "S3Key": {
			//	          "type": "string"
			//	        }
			//	      },
			//	      "required": [
			//	        "S3Bucket",
			//	        "S3Key"
			//	      ],
			//	      "type": "object"
			//	    },
			//	    "TimeoutInSeconds": {
			//	      "type": "integer"
			//	    }
			//	  },
			//	  "required": [
			//	    "ScriptS3Location",
			//	    "ExecutablePath",
			//	    "TimeoutInSeconds"
			//	  ],
			//	  "type": "object"
			//	}
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"executable_parameters": {
						// Property: ExecutableParameters
						Type:     types.StringType,
						Computed: true,
					},
					"executable_path": {
						// Property: ExecutablePath
						Type:     types.StringType,
						Computed: true,
					},
					"script_s3_location": {
						// Property: ScriptS3Location
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"s3_bucket": {
									// Property: S3Bucket
									Type:     types.StringType,
									Computed: true,
								},
								"s3_key": {
									// Property: S3Key
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"timeout_in_seconds": {
						// Property: TimeoutInSeconds
						Type:     types.Int64Type,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"source_s3_location": {
			// Property: SourceS3Location
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "properties": {
			//	    "S3Bucket": {
			//	      "type": "string"
			//	    },
			//	    "S3Key": {
			//	      "type": "string"
			//	    }
			//	  },
			//	  "required": [
			//	    "S3Bucket",
			//	    "S3Key"
			//	  ],
			//	  "type": "object"
			//	}
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"s3_bucket": {
						// Property: S3Bucket
						Type:     types.StringType,
						Computed: true,
					},
					"s3_key": {
						// Property: S3Key
						Type:     types.StringType,
						Computed: true,
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
			//	    "properties": {
			//	      "TagKey": {
			//	        "type": "string"
			//	      },
			//	      "TagValue": {
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "TagKey",
			//	      "TagValue"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"tag_key": {
						// Property: TagKey
						Type:     types.StringType,
						Computed: true,
					},
					"tag_value": {
						// Property: TagValue
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
		Description: "Data Source schema for AWS::AppStream::AppBlock",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::AppStream::AppBlock").WithTerraformTypeName("awscc_appstream_app_block")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                   "Arn",
		"created_time":          "CreatedTime",
		"description":           "Description",
		"display_name":          "DisplayName",
		"executable_parameters": "ExecutableParameters",
		"executable_path":       "ExecutablePath",
		"name":                  "Name",
		"s3_bucket":             "S3Bucket",
		"s3_key":                "S3Key",
		"script_s3_location":    "ScriptS3Location",
		"setup_script_details":  "SetupScriptDetails",
		"source_s3_location":    "SourceS3Location",
		"tag_key":               "TagKey",
		"tag_value":             "TagValue",
		"tags":                  "Tags",
		"timeout_in_seconds":    "TimeoutInSeconds",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
