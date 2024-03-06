// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package datasync

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddResourceFactory("awscc_datasync_location_efs", locationEFSResource)
}

// locationEFSResource returns the Terraform awscc_datasync_location_efs resource.
// This Terraform resource corresponds to the CloudFormation AWS::DataSync::LocationEFS resource.
func locationEFSResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccessPointArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) for the Amazon EFS Access point that DataSync uses when accessing the EFS file system.",
		//	  "maxLength": 128,
		//	  "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):elasticfilesystem:[a-z\\-0-9]+:[0-9]{12}:access-point/fsap-[0-9a-f]{8,40}$",
		//	  "type": "string"
		//	}
		"access_point_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) for the Amazon EFS Access point that DataSync uses when accessing the EFS file system.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(128),
				stringvalidator.RegexMatches(regexp.MustCompile("^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):elasticfilesystem:[a-z\\-0-9]+:[0-9]{12}:access-point/fsap-[0-9a-f]{8,40}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Ec2Config
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The subnet and security group that DataSync uses to access target EFS file system.",
		//	  "properties": {
		//	    "SecurityGroupArns": {
		//	      "description": "The Amazon Resource Names (ARNs) of the security groups that are configured for the Amazon EC2 resource.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "maxLength": 128,
		//	        "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):ec2:[a-z\\-0-9]*:[0-9]{12}:security-group/.*$",
		//	        "type": "string"
		//	      },
		//	      "maxItems": 5,
		//	      "minItems": 1,
		//	      "type": "array"
		//	    },
		//	    "SubnetArn": {
		//	      "description": "The ARN of the subnet that DataSync uses to access the target EFS file system.",
		//	      "maxLength": 128,
		//	      "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):ec2:[a-z\\-0-9]*:[0-9]{12}:subnet/.*$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "SecurityGroupArns",
		//	    "SubnetArn"
		//	  ],
		//	  "type": "object"
		//	}
		"ec_2_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: SecurityGroupArns
				"security_group_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					CustomType:  cctypes.MultisetType,
					Description: "The Amazon Resource Names (ARNs) of the security groups that are configured for the Amazon EC2 resource.",
					Required:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.SizeBetween(1, 5),
						listvalidator.ValueStringsAre(
							stringvalidator.LengthAtMost(128),
							stringvalidator.RegexMatches(regexp.MustCompile("^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):ec2:[a-z\\-0-9]*:[0-9]{12}:security-group/.*$"), ""),
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: SubnetArn
				"subnet_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the subnet that DataSync uses to access the target EFS file system.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(128),
						stringvalidator.RegexMatches(regexp.MustCompile("^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):ec2:[a-z\\-0-9]*:[0-9]{12}:subnet/.*$"), ""),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The subnet and security group that DataSync uses to access target EFS file system.",
			Required:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EfsFilesystemArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) for the Amazon EFS file system.",
		//	  "maxLength": 128,
		//	  "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):elasticfilesystem:[a-z\\-0-9]*:[0-9]{12}:file-system/fs-.*$",
		//	  "type": "string"
		//	}
		"efs_filesystem_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) for the Amazon EFS file system.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(128),
				stringvalidator.RegexMatches(regexp.MustCompile("^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):elasticfilesystem:[a-z\\-0-9]*:[0-9]{12}:file-system/fs-.*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// EfsFilesystemArn is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: FileSystemAccessRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the AWS IAM role that the DataSync will assume when mounting the EFS file system.",
		//	  "maxLength": 128,
		//	  "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):iam::[0-9]{12}:role/.*$",
		//	  "type": "string"
		//	}
		"file_system_access_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the AWS IAM role that the DataSync will assume when mounting the EFS file system.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(128),
				stringvalidator.RegexMatches(regexp.MustCompile("^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):iam::[0-9]{12}:role/.*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InTransitEncryption
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Protocol that is used for encrypting the traffic exchanged between the DataSync Agent and the EFS file system.",
		//	  "enum": [
		//	    "NONE",
		//	    "TLS1_2"
		//	  ],
		//	  "type": "string"
		//	}
		"in_transit_encryption": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Protocol that is used for encrypting the traffic exchanged between the DataSync Agent and the EFS file system.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"NONE",
					"TLS1_2",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LocationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the Amazon EFS file system location that is created.",
		//	  "maxLength": 128,
		//	  "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:location/loc-[0-9a-z]{17}$",
		//	  "type": "string"
		//	}
		"location_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the Amazon EFS file system location that is created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LocationUri
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The URL of the EFS location that was described.",
		//	  "maxLength": 4356,
		//	  "pattern": "^(efs|nfs|s3|smb|fsxw)://[a-zA-Z0-9.\\-/]+$",
		//	  "type": "string"
		//	}
		"location_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The URL of the EFS location that was described.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Subdirectory
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A subdirectory in the location's path. This subdirectory in the EFS file system is used to read data from the EFS source location or write data to the EFS destination.",
		//	  "maxLength": 4096,
		//	  "pattern": "^[a-zA-Z0-9_\\-\\+\\./\\(\\)\\$\\p{Zs}]+$",
		//	  "type": "string"
		//	}
		"subdirectory": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A subdirectory in the location's path. This subdirectory in the EFS file system is used to read data from the EFS source location or write data to the EFS destination.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(4096),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_\\-\\+\\./\\(\\)\\$\\p{Zs}]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// Subdirectory is a write-only property.
		}, /*END ATTRIBUTE*/
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
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key for an AWS resource tag.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
							stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9\\s+=._:/-]+$"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for an AWS resource tag.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
							stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9\\s+=._:@/-]+$"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeAtMost(50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource schema for AWS::DataSync::LocationEFS.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataSync::LocationEFS").WithTerraformTypeName("awscc_datasync_location_efs")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_point_arn":            "AccessPointArn",
		"ec_2_config":                 "Ec2Config",
		"efs_filesystem_arn":          "EfsFilesystemArn",
		"file_system_access_role_arn": "FileSystemAccessRoleArn",
		"in_transit_encryption":       "InTransitEncryption",
		"key":                         "Key",
		"location_arn":                "LocationArn",
		"location_uri":                "LocationUri",
		"security_group_arns":         "SecurityGroupArns",
		"subdirectory":                "Subdirectory",
		"subnet_arn":                  "SubnetArn",
		"tags":                        "Tags",
		"value":                       "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/EfsFilesystemArn",
		"/properties/Subdirectory",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
