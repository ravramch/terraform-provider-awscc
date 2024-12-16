// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package s3tables

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_s3tables_table_bucket", tableBucketResource)
}

// tableBucketResource returns the Terraform awscc_s3tables_table_bucket resource.
// This Terraform resource corresponds to the CloudFormation AWS::S3Tables::TableBucket resource.
func tableBucketResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: TableBucketARN
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the specified table bucket.",
		//	  "examples": [
		//	    "arn:aws:s3tables:us-west-2:123456789012:bucket/mytablebucket"
		//	  ],
		//	  "type": "string"
		//	}
		"table_bucket_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the specified table bucket.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TableBucketName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name for the table bucket.",
		//	  "maxLength": 63,
		//	  "minLength": 3,
		//	  "type": "string"
		//	}
		"table_bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name for the table bucket.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(3, 63),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UnreferencedFileRemoval
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Settings governing the Unreferenced File Removal maintenance action. Unreferenced file removal identifies and deletes all objects that are not referenced by any table snapshots.",
		//	  "properties": {
		//	    "NoncurrentDays": {
		//	      "description": "S3 permanently deletes noncurrent objects after the number of days specified by the NoncurrentDays property.",
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    },
		//	    "Status": {
		//	      "description": "Indicates whether the Unreferenced File Removal maintenance action is enabled.",
		//	      "enum": [
		//	        "Enabled",
		//	        "Disabled"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "UnreferencedDays": {
		//	      "description": "For any object not referenced by your table and older than the UnreferencedDays property, S3 creates a delete marker and marks the object version as noncurrent.",
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"unreferenced_file_removal": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: NoncurrentDays
				"noncurrent_days": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "S3 permanently deletes noncurrent objects after the number of days specified by the NoncurrentDays property.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.AtLeast(1),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Status
				"status": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Indicates whether the Unreferenced File Removal maintenance action is enabled.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"Enabled",
							"Disabled",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: UnreferencedDays
				"unreferenced_days": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "For any object not referenced by your table and older than the UnreferencedDays property, S3 creates a delete marker and marks the object version as noncurrent.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.AtLeast(1),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Settings governing the Unreferenced File Removal maintenance action. Unreferenced file removal identifies and deletes all objects that are not referenced by any table snapshots.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Creates an Amazon S3 Tables table bucket in the same AWS Region where you create the AWS CloudFormation stack.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3Tables::TableBucket").WithTerraformTypeName("awscc_s3tables_table_bucket")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"noncurrent_days":           "NoncurrentDays",
		"status":                    "Status",
		"table_bucket_arn":          "TableBucketARN",
		"table_bucket_name":         "TableBucketName",
		"unreferenced_days":         "UnreferencedDays",
		"unreferenced_file_removal": "UnreferencedFileRemoval",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
