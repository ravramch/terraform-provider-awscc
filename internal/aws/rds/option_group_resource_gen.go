// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package rds

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
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
	registry.AddResourceFactory("awscc_rds_option_group", optionGroupResource)
}

// optionGroupResource returns the Terraform awscc_rds_option_group resource.
// This Terraform resource corresponds to the CloudFormation AWS::RDS::OptionGroup resource.
func optionGroupResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: EngineName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates the name of the engine that this option group can be applied to.",
		//	  "type": "string"
		//	}
		"engine_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates the name of the engine that this option group can be applied to.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MajorEngineVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates the major engine version associated with this option group.",
		//	  "type": "string"
		//	}
		"major_engine_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates the major engine version associated with this option group.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: OptionConfigurations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "arrayType": "AttributeList",
		//	  "description": "Indicates what options are available in the option group.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The OptionConfiguration property type specifies an individual option, and its settings, within an AWS::RDS::OptionGroup resource.",
		//	    "properties": {
		//	      "DBSecurityGroupMemberships": {
		//	        "description": "A list of DBSecurityGroupMembership name strings used for this option.",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "type": "string"
		//	        },
		//	        "type": "array",
		//	        "uniqueItems": true
		//	      },
		//	      "OptionName": {
		//	        "description": "The configuration of options to include in a group.",
		//	        "type": "string"
		//	      },
		//	      "OptionSettings": {
		//	        "description": "The option settings to include in an option group.",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "description": "The OptionSetting property type specifies the value for an option within an OptionSetting property.",
		//	          "properties": {
		//	            "Name": {
		//	              "description": "The name of the option that has settings that you can set.",
		//	              "type": "string"
		//	            },
		//	            "Value": {
		//	              "description": "The current value of the option setting.",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "type": "array"
		//	      },
		//	      "OptionVersion": {
		//	        "description": "The version for the option.",
		//	        "type": "string"
		//	      },
		//	      "Port": {
		//	        "description": "The optional port for the option.",
		//	        "type": "integer"
		//	      },
		//	      "VpcSecurityGroupMemberships": {
		//	        "description": "A list of VpcSecurityGroupMembership name strings used for this option.",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "type": "string"
		//	        },
		//	        "type": "array",
		//	        "uniqueItems": true
		//	      }
		//	    },
		//	    "required": [
		//	      "OptionName"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"option_configurations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: DBSecurityGroupMemberships
					"db_security_group_memberships": schema.SetAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Description: "A list of DBSecurityGroupMembership name strings used for this option.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
							setplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: OptionName
					"option_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The configuration of options to include in a group.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: OptionSettings
					"option_settings": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Name
								"name": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The name of the option that has settings that you can set.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: Value
								"value": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The current value of the option setting.",
									Optional:    true,
									Computed:    true,
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						CustomType:  cctypes.MultisetType,
						Description: "The option settings to include in an option group.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
							listplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: OptionVersion
					"option_version": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The version for the option.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Port
					"port": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "The optional port for the option.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
							int64planmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: VpcSecurityGroupMemberships
					"vpc_security_group_memberships": schema.SetAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Description: "A list of VpcSecurityGroupMembership name strings used for this option.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
							setplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType:  cctypes.MultisetType,
			Description: "Indicates what options are available in the option group.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: OptionGroupDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Provides a description of the option group.",
		//	  "type": "string"
		//	}
		"option_group_description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Provides a description of the option group.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: OptionGroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the name of the option group.",
		//	  "type": "string"
		//	}
		"option_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the name of the option group.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
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
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType:  cctypes.MultisetType,
			Description: "An array of key-value pairs to apply to this resource.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
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
		Description: "The AWS::RDS::OptionGroup resource creates an option group, to enable and configure features that are specific to a particular DB engine.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RDS::OptionGroup").WithTerraformTypeName("awscc_rds_option_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"db_security_group_memberships":  "DBSecurityGroupMemberships",
		"engine_name":                    "EngineName",
		"key":                            "Key",
		"major_engine_version":           "MajorEngineVersion",
		"name":                           "Name",
		"option_configurations":          "OptionConfigurations",
		"option_group_description":       "OptionGroupDescription",
		"option_group_name":              "OptionGroupName",
		"option_name":                    "OptionName",
		"option_settings":                "OptionSettings",
		"option_version":                 "OptionVersion",
		"port":                           "Port",
		"tags":                           "Tags",
		"value":                          "Value",
		"vpc_security_group_memberships": "VpcSecurityGroupMemberships",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
