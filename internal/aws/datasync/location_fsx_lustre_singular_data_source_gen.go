// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package datasync

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_datasync_location_fsx_lustre", locationFSxLustreDataSource)
}

// locationFSxLustreDataSource returns the Terraform awscc_datasync_location_fsx_lustre data source.
// This Terraform data source corresponds to the CloudFormation AWS::DataSync::LocationFSxLustre resource.
func locationFSxLustreDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"fsx_filesystem_arn": {
			// Property: FsxFilesystemArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) for the FSx for Lustre file system.",
			//	  "maxLength": 128,
			//	  "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):fsx:[a-z\\-0-9]+:[0-9]{12}:file-system/fs-[0-9a-f]+$",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) for the FSx for Lustre file system.",
			Type:        types.StringType,
			Computed:    true,
		},
		"location_arn": {
			// Property: LocationArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) of the Amazon FSx for Lustre file system location that is created.",
			//	  "maxLength": 128,
			//	  "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:location/loc-[0-9a-z]{17}$",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) of the Amazon FSx for Lustre file system location that is created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"location_uri": {
			// Property: LocationUri
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The URL of the FSx for Lustre location that was described.",
			//	  "maxLength": 4356,
			//	  "pattern": "^(efs|nfs|s3|smb|fsxw|hdfs|fsxl)://[a-zA-Z0-9.:/\\-]+$",
			//	  "type": "string"
			//	}
			Description: "The URL of the FSx for Lustre location that was described.",
			Type:        types.StringType,
			Computed:    true,
		},
		"security_group_arns": {
			// Property: SecurityGroupArns
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The ARNs of the security groups that are to use to configure the FSx for Lustre file system.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "maxLength": 128,
			//	    "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):ec2:[a-z\\-0-9]*:[0-9]{12}:security-group/sg-[a-f0-9]+$",
			//	    "type": "string"
			//	  },
			//	  "maxItems": 5,
			//	  "minItems": 1,
			//	  "type": "array"
			//	}
			Description: "The ARNs of the security groups that are to use to configure the FSx for Lustre file system.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"subdirectory": {
			// Property: Subdirectory
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A subdirectory in the location's path.",
			//	  "maxLength": 4096,
			//	  "pattern": "^[a-zA-Z0-9_\\-\\+\\./\\(\\)\\$\\p{Zs}]+$",
			//	  "type": "string"
			//	}
			Description: "A subdirectory in the location's path.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "An array of key-value pairs to apply to this resource.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "A key-value pair to associate with a resource.",
			//	    "properties": {
			//	      "Key": {
			//	        "description": "The key for an AWS resource tag.",
			//	        "maxLength": 256,
			//	        "minLength": 1,
			//	        "pattern": "^[a-zA-Z0-9\\s+=._:/-]+$",
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "description": "The value for an AWS resource tag.",
			//	        "maxLength": 256,
			//	        "minLength": 1,
			//	        "pattern": "^[a-zA-Z0-9\\s+=._:@/-]+$",
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
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key for an AWS resource tag.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for an AWS resource tag.",
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
		Description: "Data Source schema for AWS::DataSync::LocationFSxLustre",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataSync::LocationFSxLustre").WithTerraformTypeName("awscc_datasync_location_fsx_lustre")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"fsx_filesystem_arn":  "FsxFilesystemArn",
		"key":                 "Key",
		"location_arn":        "LocationArn",
		"location_uri":        "LocationUri",
		"security_group_arns": "SecurityGroupArns",
		"subdirectory":        "Subdirectory",
		"tags":                "Tags",
		"value":               "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
