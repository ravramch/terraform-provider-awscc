// Code generated by generators/resource/main.go; DO NOT EDIT.

package synthetics

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"

	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_synthetics_canary", canaryResourceType)
}

// canaryResourceType returns the Terraform awscc_synthetics_canary resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Synthetics::Canary resource type.
func canaryResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"artifact_s3_location": {
			// Property: ArtifactS3Location
			// CloudFormation resource type schema:
			// {
			//   "description": "Provide the s3 bucket output location for test results",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Provide the s3 bucket output location for test results",
			Type:        types.StringType,
			Required:    true,
		},
		"code": {
			// Property: Code
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Handler": {
			//       "type": "string"
			//     },
			//     "S3Bucket": {
			//       "type": "string"
			//     },
			//     "S3Key": {
			//       "type": "string"
			//     },
			//     "S3ObjectVersion": {
			//       "type": "string"
			//     },
			//     "Script": {
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Handler"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"handler": {
						// Property: Handler
						Type:     types.StringType,
						Required: true,
					},
					"s3_bucket": {
						// Property: S3Bucket
						Type:     types.StringType,
						Optional: true,
						// S3Bucket is a write-only attribute.
					},
					"s3_key": {
						// Property: S3Key
						Type:     types.StringType,
						Optional: true,
						// S3Key is a write-only attribute.
					},
					"s3_object_version": {
						// Property: S3ObjectVersion
						Type:     types.StringType,
						Optional: true,
						// S3ObjectVersion is a write-only attribute.
					},
					"script": {
						// Property: Script
						Type:     types.StringType,
						Optional: true,
						// Script is a write-only attribute.
					},
				},
			),
			Required: true,
		},
		"execution_role_arn": {
			// Property: ExecutionRoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Lambda Execution role used to run your canaries",
			//   "type": "string"
			// }
			Description: "Lambda Execution role used to run your canaries",
			Type:        types.StringType,
			Required:    true,
		},
		"failure_retention_period": {
			// Property: FailureRetentionPeriod
			// CloudFormation resource type schema:
			// {
			//   "description": "Retention period of failed canary runs represented in number of days",
			//   "type": "integer"
			// }
			Description: "Retention period of failed canary runs represented in number of days",
			Type:        types.NumberType,
			Optional:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "Id of the canary",
			//   "type": "string"
			// }
			Description: "Id of the canary",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of the canary.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Name of the canary.",
			Type:        types.StringType,
			Required:    true,
			// Name is a force-new attribute.
		},
		"run_config": {
			// Property: RunConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "ActiveTracing": {
			//       "description": "Enable active tracing if set to true",
			//       "type": "boolean"
			//     },
			//     "EnvironmentVariables": {
			//       "additionalProperties": false,
			//       "description": "Environment variable key-value pairs.",
			//       "patternProperties": {
			//         "": {
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "MemoryInMB": {
			//       "description": "Provide maximum memory available for canary in MB",
			//       "type": "integer"
			//     },
			//     "TimeoutInSeconds": {
			//       "description": "Provide maximum canary timeout per run in seconds",
			//       "type": "integer"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"active_tracing": {
						// Property: ActiveTracing
						Description: "Enable active tracing if set to true",
						Type:        types.BoolType,
						Optional:    true,
					},
					"environment_variables": {
						// Property: EnvironmentVariables
						Description: "Environment variable key-value pairs.",
						// Pattern: ""
						Type:     types.MapType{ElemType: types.StringType},
						Optional: true,
					},
					"memory_in_mb": {
						// Property: MemoryInMB
						Description: "Provide maximum memory available for canary in MB",
						Type:        types.NumberType,
						Optional:    true,
					},
					"timeout_in_seconds": {
						// Property: TimeoutInSeconds
						Description: "Provide maximum canary timeout per run in seconds",
						Type:        types.NumberType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"runtime_version": {
			// Property: RuntimeVersion
			// CloudFormation resource type schema:
			// {
			//   "description": "Runtime version of Synthetics Library",
			//   "type": "string"
			// }
			Description: "Runtime version of Synthetics Library",
			Type:        types.StringType,
			Required:    true,
		},
		"schedule": {
			// Property: Schedule
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "DurationInSeconds": {
			//       "type": "string"
			//     },
			//     "Expression": {
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Expression"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"duration_in_seconds": {
						// Property: DurationInSeconds
						Type:     types.StringType,
						Optional: true,
					},
					"expression": {
						// Property: Expression
						Type:     types.StringType,
						Required: true,
					},
				},
			),
			Required: true,
		},
		"start_canary_after_creation": {
			// Property: StartCanaryAfterCreation
			// CloudFormation resource type schema:
			// {
			//   "description": "Runs canary if set to True. Default is False",
			//   "type": "boolean"
			// }
			Description: "Runs canary if set to True. Default is False",
			Type:        types.BoolType,
			Required:    true,
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			// {
			//   "description": "State of the canary",
			//   "type": "string"
			// }
			Description: "State of the canary",
			Type:        types.StringType,
			Computed:    true,
		},
		"success_retention_period": {
			// Property: SuccessRetentionPeriod
			// CloudFormation resource type schema:
			// {
			//   "description": "Retention period of successful canary runs represented in number of days",
			//   "type": "integer"
			// }
			Description: "Retention period of successful canary runs represented in number of days",
			Type:        types.NumberType,
			Optional:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"vpc_config": {
			// Property: VPCConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "SecurityGroupIds": {
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array"
			//     },
			//     "SubnetIds": {
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array"
			//     },
			//     "VpcId": {
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "SubnetIds",
			//     "SecurityGroupIds"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"security_group_ids": {
						// Property: SecurityGroupIds
						Type:     types.ListType{ElemType: types.StringType},
						Required: true,
					},
					"subnet_ids": {
						// Property: SubnetIds
						Type:     types.ListType{ElemType: types.StringType},
						Required: true,
					},
					"vpc_id": {
						// Property: VpcId
						Type:     types.StringType,
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"visual_reference": {
			// Property: VisualReference
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "BaseCanaryRunId": {
			//       "description": "Canary run id to be used as base reference for visual testing",
			//       "type": "string"
			//     },
			//     "BaseScreenshots": {
			//       "description": "List of screenshots used as base reference for visual testing",
			//       "items": {
			//         "properties": {
			//           "IgnoreCoordinates": {
			//             "description": "List of coordinates of rectangles to be ignored during visual testing",
			//             "items": {
			//               "description": "Coordinates of a rectangle to be ignored during visual testing",
			//               "type": "string"
			//             },
			//             "type": "array"
			//           },
			//           "ScreenshotName": {
			//             "description": "Name of the screenshot to be used as base reference for visual testing",
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "ScreenshotName"
			//         ],
			//         "type": "object"
			//       },
			//       "type": "array"
			//     }
			//   },
			//   "required": [
			//     "BaseCanaryRunId"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"base_canary_run_id": {
						// Property: BaseCanaryRunId
						Description: "Canary run id to be used as base reference for visual testing",
						Type:        types.StringType,
						Required:    true,
					},
					"base_screenshots": {
						// Property: BaseScreenshots
						Description: "List of screenshots used as base reference for visual testing",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"ignore_coordinates": {
									// Property: IgnoreCoordinates
									Description: "List of coordinates of rectangles to be ignored during visual testing",
									Type:        types.ListType{ElemType: types.StringType},
									Optional:    true,
								},
								"screenshot_name": {
									// Property: ScreenshotName
									Description: "Name of the screenshot to be used as base reference for visual testing",
									Type:        types.StringType,
									Required:    true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::Synthetics::Canary",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Synthetics::Canary").WithTerraformTypeName("awscc_synthetics_canary")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"active_tracing":              "ActiveTracing",
		"artifact_s3_location":        "ArtifactS3Location",
		"base_canary_run_id":          "BaseCanaryRunId",
		"base_screenshots":            "BaseScreenshots",
		"code":                        "Code",
		"duration_in_seconds":         "DurationInSeconds",
		"environment_variables":       "EnvironmentVariables",
		"execution_role_arn":          "ExecutionRoleArn",
		"expression":                  "Expression",
		"failure_retention_period":    "FailureRetentionPeriod",
		"handler":                     "Handler",
		"id":                          "Id",
		"ignore_coordinates":          "IgnoreCoordinates",
		"key":                         "Key",
		"memory_in_mb":                "MemoryInMB",
		"name":                        "Name",
		"run_config":                  "RunConfig",
		"runtime_version":             "RuntimeVersion",
		"s3_bucket":                   "S3Bucket",
		"s3_key":                      "S3Key",
		"s3_object_version":           "S3ObjectVersion",
		"schedule":                    "Schedule",
		"screenshot_name":             "ScreenshotName",
		"script":                      "Script",
		"security_group_ids":          "SecurityGroupIds",
		"start_canary_after_creation": "StartCanaryAfterCreation",
		"state":                       "State",
		"subnet_ids":                  "SubnetIds",
		"success_retention_period":    "SuccessRetentionPeriod",
		"tags":                        "Tags",
		"timeout_in_seconds":          "TimeoutInSeconds",
		"value":                       "Value",
		"visual_reference":            "VisualReference",
		"vpc_config":                  "VPCConfig",
		"vpc_id":                      "VpcId",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Code/S3Bucket",
		"/properties/Code/S3Key",
		"/properties/Code/S3ObjectVersion",
		"/properties/Code/Script",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_synthetics_canary", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
