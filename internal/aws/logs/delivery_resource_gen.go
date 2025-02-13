// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package logs

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_logs_delivery", deliveryResource)
}

// deliveryResource returns the Terraform awscc_logs_delivery resource.
// This Terraform resource corresponds to the CloudFormation AWS::Logs::Delivery resource.
func deliveryResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) that uniquely identifies this delivery.",
		//	  "maxLength": 2048,
		//	  "minLength": 16,
		//	  "pattern": "[\\w#+=/:,.@-]*\\*?",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) that uniquely identifies this delivery.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DeliveryDestinationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the delivery destination that is associated with this delivery.",
		//	  "maxLength": 2048,
		//	  "minLength": 16,
		//	  "pattern": "[\\w#+=/:,.@-]*\\*?",
		//	  "type": "string"
		//	}
		"delivery_destination_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the delivery destination that is associated with this delivery.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(16, 2048),
				stringvalidator.RegexMatches(regexp.MustCompile("[\\w#+=/:,.@-]*\\*?"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DeliveryDestinationType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Displays whether the delivery destination associated with this delivery is CloudWatch Logs, Amazon S3, or Kinesis Data Firehose.",
		//	  "maxLength": 12,
		//	  "minLength": 1,
		//	  "pattern": "^[0-9A-Za-z]+$",
		//	  "type": "string"
		//	}
		"delivery_destination_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Displays whether the delivery destination associated with this delivery is CloudWatch Logs, Amazon S3, or Kinesis Data Firehose.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DeliveryId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique ID that identifies this delivery in your account.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[0-9A-Za-z]+$",
		//	  "type": "string"
		//	}
		"delivery_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique ID that identifies this delivery in your account.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DeliverySourceName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the delivery source that is associated with this delivery.",
		//	  "maxLength": 60,
		//	  "minLength": 1,
		//	  "pattern": "[\\w-]*$",
		//	  "type": "string"
		//	}
		"delivery_source_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the delivery source that is associated with this delivery.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 60),
				stringvalidator.RegexMatches(regexp.MustCompile("[\\w-]*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FieldDelimiter
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The field delimiter to use between record fields when the final output format of a delivery is in Plain , W3C , or Raw format.",
		//	  "maxLength": 5,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"field_delimiter": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The field delimiter to use between record fields when the final output format of a delivery is in Plain , W3C , or Raw format.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 5),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RecordFields
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The list of record fields to be delivered to the destination, in order. If the delivery's log source has mandatory fields, they must be included in this list.",
		//	  "items": {
		//	    "description": "A single record field to be delivered to the destination.",
		//	    "maxLength": 50,
		//	    "minLength": 1,
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"record_fields": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The list of record fields to be delivered to the destination, in order. If the delivery's log source has mandatory fields, they must be included in this list.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.ValueStringsAre(
					stringvalidator.LengthBetween(1, 50),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: S3EnableHiveCompatiblePath
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "This parameter causes the S3 objects that contain delivered logs to use a prefix structure that allows for integration with Apache Hive.",
		//	  "type": "boolean"
		//	}
		"s3_enable_hive_compatible_path": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "This parameter causes the S3 objects that contain delivered logs to use a prefix structure that allows for integration with Apache Hive.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: S3SuffixPath
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "This string allows re-configuring the S3 object prefix to contain either static or variable sections. The valid variables to use in the suffix path will vary by each log source. See ConfigurationTemplate$allowedSuffixPathFields for more info on what values are supported in the suffix path for each log source.",
		//	  "maxLength": 256,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"s3_suffix_path": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "This string allows re-configuring the S3 object prefix to contain either static or variable sections. The valid variables to use in the suffix path will vary by each log source. See ConfigurationTemplate$allowedSuffixPathFields for more info on what values are supported in the suffix path for each log source.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 256),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags that have been assigned to this delivery.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode",
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
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tags that have been assigned to this delivery.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
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
		Description: "This structure contains information about one delivery in your account.\n\nA delivery is a connection between a logical delivery source and a logical delivery destination.\n\nFor more information, see [CreateDelivery](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_CreateDelivery.html).",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Logs::Delivery").WithTerraformTypeName("awscc_logs_delivery")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                            "Arn",
		"delivery_destination_arn":       "DeliveryDestinationArn",
		"delivery_destination_type":      "DeliveryDestinationType",
		"delivery_id":                    "DeliveryId",
		"delivery_source_name":           "DeliverySourceName",
		"field_delimiter":                "FieldDelimiter",
		"key":                            "Key",
		"record_fields":                  "RecordFields",
		"s3_enable_hive_compatible_path": "S3EnableHiveCompatiblePath",
		"s3_suffix_path":                 "S3SuffixPath",
		"tags":                           "Tags",
		"value":                          "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
