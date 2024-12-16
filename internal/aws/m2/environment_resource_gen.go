// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package m2

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_m2_environment", environmentResource)
}

// environmentResource returns the Terraform awscc_m2_environment resource.
// This Terraform resource corresponds to the CloudFormation AWS::M2::Environment resource.
func environmentResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the environment.",
		//	  "maxLength": 500,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the environment.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 500),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EngineType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The target platform for the environment.",
		//	  "enum": [
		//	    "microfocus",
		//	    "bluage"
		//	  ],
		//	  "type": "string"
		//	}
		"engine_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The target platform for the environment.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"microfocus",
					"bluage",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EngineVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The version of the runtime engine for the environment.",
		//	  "pattern": "^\\S{1,10}$",
		//	  "type": "string"
		//	}
		"engine_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The version of the runtime engine for the environment.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^\\S{1,10}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the runtime environment.",
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"environment_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the runtime environment.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique identifier of the environment.",
		//	  "pattern": "^\\S{1,80}$",
		//	  "type": "string"
		//	}
		"environment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique identifier of the environment.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: HighAvailabilityConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Defines the details of a high availability configuration.",
		//	  "properties": {
		//	    "DesiredCapacity": {
		//	      "maximum": 100,
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "required": [
		//	    "DesiredCapacity"
		//	  ],
		//	  "type": "object"
		//	}
		"high_availability_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DesiredCapacity
				"desired_capacity": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.Between(1, 100),
						fwvalidators.NotNullInt64(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Defines the details of a high availability configuration.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InstanceType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of instance underlying the environment.",
		//	  "pattern": "^\\S{1,20}$",
		//	  "type": "string"
		//	}
		"instance_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of instance underlying the environment.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^\\S{1,20}$"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: KmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID or the Amazon Resource Name (ARN) of the customer managed KMS Key used for encrypting environment-related resources.",
		//	  "maxLength": 2048,
		//	  "type": "string"
		//	}
		"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID or the Amazon Resource Name (ARN) of the customer managed KMS Key used for encrypting environment-related resources.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(2048),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the environment.",
		//	  "pattern": "^[A-Za-z0-9][A-Za-z0-9_\\-]{1,59}$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the environment.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^[A-Za-z0-9][A-Za-z0-9_\\-]{1,59}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NetworkType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "ipv4",
		//	    "dual"
		//	  ],
		//	  "type": "string"
		//	}
		"network_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"ipv4",
					"dual",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PreferredMaintenanceWindow
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Configures a desired maintenance window for the environment. If you do not provide a value, a random system-generated value will be assigned.",
		//	  "pattern": "^\\S{1,50}$",
		//	  "type": "string"
		//	}
		"preferred_maintenance_window": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Configures a desired maintenance window for the environment. If you do not provide a value, a random system-generated value will be assigned.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^\\S{1,50}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PubliclyAccessible
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies whether the environment is publicly accessible.",
		//	  "type": "boolean"
		//	}
		"publicly_accessible": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether the environment is publicly accessible.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
				boolplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SecurityGroupIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The list of security groups for the VPC associated with this environment.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "pattern": "^\\S{1,50}$",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"security_group_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The list of security groups for the VPC associated with this environment.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.ValueStringsAre(
					stringvalidator.RegexMatches(regexp.MustCompile("^\\S{1,50}$"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StorageConfigurations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The storage configurations defined for the runtime environment.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "description": "Defines the storage configuration for an environment.",
		//	    "properties": {
		//	      "Efs": {
		//	        "additionalProperties": false,
		//	        "description": "Defines the storage configuration for an Amazon EFS file system.",
		//	        "properties": {
		//	          "FileSystemId": {
		//	            "description": "The file system identifier.",
		//	            "pattern": "^\\S{1,200}$",
		//	            "type": "string"
		//	          },
		//	          "MountPoint": {
		//	            "description": "The mount point for the file system.",
		//	            "pattern": "^\\S{1,200}$",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "FileSystemId",
		//	          "MountPoint"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "Fsx": {
		//	        "additionalProperties": false,
		//	        "description": "Defines the storage configuration for an Amazon FSx file system.",
		//	        "properties": {
		//	          "FileSystemId": {
		//	            "description": "The file system identifier.",
		//	            "pattern": "^\\S{1,200}$",
		//	            "type": "string"
		//	          },
		//	          "MountPoint": {
		//	            "description": "The mount point for the file system.",
		//	            "pattern": "^\\S{1,200}$",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "FileSystemId",
		//	          "MountPoint"
		//	        ],
		//	        "type": "object"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"storage_configurations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Efs
					"efs": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: FileSystemId
							"file_system_id": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The file system identifier.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.RegexMatches(regexp.MustCompile("^\\S{1,200}$"), ""),
									fwvalidators.NotNullString(),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: MountPoint
							"mount_point": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The mount point for the file system.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.RegexMatches(regexp.MustCompile("^\\S{1,200}$"), ""),
									fwvalidators.NotNullString(),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "Defines the storage configuration for an Amazon EFS file system.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Fsx
					"fsx": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: FileSystemId
							"file_system_id": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The file system identifier.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.RegexMatches(regexp.MustCompile("^\\S{1,200}$"), ""),
									fwvalidators.NotNullString(),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: MountPoint
							"mount_point": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The mount point for the file system.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.RegexMatches(regexp.MustCompile("^\\S{1,200}$"), ""),
									fwvalidators.NotNullString(),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "Defines the storage configuration for an Amazon FSx file system.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The storage configurations defined for the runtime environment.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SubnetIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique identifiers of the subnets assigned to this runtime environment.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "pattern": "^\\S{1,50}$",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"subnet_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The unique identifiers of the subnets assigned to this runtime environment.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.ValueStringsAre(
					stringvalidator.RegexMatches(regexp.MustCompile("^\\S{1,50}$"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Tags associated to this environment.",
		//	  "patternProperties": {
		//	    "": {
		//	      "maxLength": 256,
		//	      "minLength": 0,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "Tags associated to this environment.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
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
		Description: "Represents a runtime environment that can run migrated mainframe applications.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::M2::Environment").WithTerraformTypeName("awscc_m2_environment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"description":                  "Description",
		"desired_capacity":             "DesiredCapacity",
		"efs":                          "Efs",
		"engine_type":                  "EngineType",
		"engine_version":               "EngineVersion",
		"environment_arn":              "EnvironmentArn",
		"environment_id":               "EnvironmentId",
		"file_system_id":               "FileSystemId",
		"fsx":                          "Fsx",
		"high_availability_config":     "HighAvailabilityConfig",
		"instance_type":                "InstanceType",
		"kms_key_id":                   "KmsKeyId",
		"mount_point":                  "MountPoint",
		"name":                         "Name",
		"network_type":                 "NetworkType",
		"preferred_maintenance_window": "PreferredMaintenanceWindow",
		"publicly_accessible":          "PubliclyAccessible",
		"security_group_ids":           "SecurityGroupIds",
		"storage_configurations":       "StorageConfigurations",
		"subnet_ids":                   "SubnetIds",
		"tags":                         "Tags",
	})

	opts = opts.WithCreateTimeoutInMinutes(120).WithDeleteTimeoutInMinutes(120)

	opts = opts.WithUpdateTimeoutInMinutes(120)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
