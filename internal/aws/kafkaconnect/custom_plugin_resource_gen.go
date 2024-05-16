// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package kafkaconnect

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_kafkaconnect_custom_plugin", customPluginResource)
}

// customPluginResource returns the Terraform awscc_kafkaconnect_custom_plugin resource.
// This Terraform resource corresponds to the CloudFormation AWS::KafkaConnect::CustomPlugin resource.
func customPluginResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ContentType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of the plugin file.",
		//	  "enum": [
		//	    "JAR",
		//	    "ZIP"
		//	  ],
		//	  "type": "string"
		//	}
		"content_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of the plugin file.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"JAR",
					"ZIP",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CustomPluginArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the custom plugin to use.",
		//	  "pattern": "arn:(aws|aws-us-gov|aws-cn):kafkaconnect:.*",
		//	  "type": "string"
		//	}
		"custom_plugin_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the custom plugin to use.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A summary description of the custom plugin.",
		//	  "maxLength": 1024,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A summary description of the custom plugin.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(1024),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FileDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Details about the custom plugin file.",
		//	  "properties": {
		//	    "FileMd5": {
		//	      "description": "The hex-encoded MD5 checksum of the custom plugin file. You can use it to validate the file.",
		//	      "type": "string"
		//	    },
		//	    "FileSize": {
		//	      "description": "The size in bytes of the custom plugin file. You can use it to validate the file.",
		//	      "format": "int64",
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"file_description": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: FileMd5
				"file_md_5": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The hex-encoded MD5 checksum of the custom plugin file. You can use it to validate the file.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: FileSize
				"file_size": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The size in bytes of the custom plugin file. You can use it to validate the file.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Details about the custom plugin file.",
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Location
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Information about the location of a custom plugin.",
		//	  "properties": {
		//	    "S3Location": {
		//	      "additionalProperties": false,
		//	      "description": "The S3 bucket Amazon Resource Name (ARN), file key, and object version of the plugin file stored in Amazon S3.",
		//	      "properties": {
		//	        "BucketArn": {
		//	          "description": "The Amazon Resource Name (ARN) of an S3 bucket.",
		//	          "type": "string"
		//	        },
		//	        "FileKey": {
		//	          "description": "The file key for an object in an S3 bucket.",
		//	          "type": "string"
		//	        },
		//	        "ObjectVersion": {
		//	          "description": "The version of an object in an S3 bucket.",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "BucketArn",
		//	        "FileKey"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "S3Location"
		//	  ],
		//	  "type": "object"
		//	}
		"location": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: S3Location
				"s3_location": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: BucketArn
						"bucket_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The Amazon Resource Name (ARN) of an S3 bucket.",
							Required:    true,
						}, /*END ATTRIBUTE*/
						// Property: FileKey
						"file_key": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The file key for an object in an S3 bucket.",
							Required:    true,
						}, /*END ATTRIBUTE*/
						// Property: ObjectVersion
						"object_version": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The version of an object in an S3 bucket.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The S3 bucket Amazon Resource Name (ARN), file key, and object version of the plugin file stored in Amazon S3.",
					Required:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Information about the location of a custom plugin.",
			Required:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the custom plugin.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the custom plugin.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 128),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Revision
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The revision of the custom plugin.",
		//	  "format": "int64",
		//	  "type": "integer"
		//	}
		"revision": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The revision of the custom plugin.",
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
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
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
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
		Description: "An example resource schema demonstrating some basic constructs and validation rules.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::KafkaConnect::CustomPlugin").WithTerraformTypeName("awscc_kafkaconnect_custom_plugin")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bucket_arn":        "BucketArn",
		"content_type":      "ContentType",
		"custom_plugin_arn": "CustomPluginArn",
		"description":       "Description",
		"file_description":  "FileDescription",
		"file_key":          "FileKey",
		"file_md_5":         "FileMd5",
		"file_size":         "FileSize",
		"key":               "Key",
		"location":          "Location",
		"name":              "Name",
		"object_version":    "ObjectVersion",
		"revision":          "Revision",
		"s3_location":       "S3Location",
		"tags":              "Tags",
		"value":             "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
