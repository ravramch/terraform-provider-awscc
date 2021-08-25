// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotwireless

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_iotwireless_task_definition", taskDefinitionResourceType)
}

// taskDefinitionResourceType returns the Terraform awscc_iotwireless_task_definition resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoTWireless::TaskDefinition resource type.
func taskDefinitionResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "TaskDefinition arn. Returned after successful create.",
			//   "type": "string"
			// }
			Description: "TaskDefinition arn. Returned after successful create.",
			Type:        types.StringType,
			Computed:    true,
		},
		"auto_create_tasks": {
			// Property: AutoCreateTasks
			// CloudFormation resource type schema:
			// {
			//   "description": "Whether to automatically create tasks using this task definition for all gateways with the specified current version. If false, the task must me created by calling CreateWirelessGatewayTask.",
			//   "type": "boolean"
			// }
			Description: "Whether to automatically create tasks using this task definition for all gateways with the specified current version. If false, the task must me created by calling CreateWirelessGatewayTask.",
			Type:        types.BoolType,
			Required:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the new wireless gateway task definition",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ID of the new wireless gateway task definition",
			Type:        types.StringType,
			Computed:    true,
		},
		"lo_ra_wan_update_gateway_task_entry": {
			// Property: LoRaWANUpdateGatewayTaskEntry
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "CurrentVersion": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "Model": {
			//           "maxLength": 4096,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "PackageVersion": {
			//           "maxLength": 32,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "Station": {
			//           "maxLength": 4096,
			//           "minLength": 1,
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "UpdateVersion": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "Model": {
			//           "maxLength": 4096,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "PackageVersion": {
			//           "maxLength": 32,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "Station": {
			//           "maxLength": 4096,
			//           "minLength": 1,
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"current_version": {
						// Property: CurrentVersion
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"model": {
									// Property: Model
									Type:     types.StringType,
									Optional: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 4096),
									},
								},
								"package_version": {
									// Property: PackageVersion
									Type:     types.StringType,
									Optional: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 32),
									},
								},
								"station": {
									// Property: Station
									Type:     types.StringType,
									Optional: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 4096),
									},
								},
							},
						),
						Optional: true,
					},
					"update_version": {
						// Property: UpdateVersion
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"model": {
									// Property: Model
									Type:     types.StringType,
									Optional: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 4096),
									},
								},
								"package_version": {
									// Property: PackageVersion
									Type:     types.StringType,
									Optional: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 32),
									},
								},
								"station": {
									// Property: Station
									Type:     types.StringType,
									Optional: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 4096),
									},
								},
							},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the new resource.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The name of the new resource.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of key-value pairs that contain metadata for the destination.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of key-value pairs that contain metadata for the destination.",
			Attributes: providertypes.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 127),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 255),
						},
					},
				},
				providertypes.SetNestedAttributesOptions{
					MaxItems: 50,
				},
			),
			Optional: true,
		},
		"task_definition_type": {
			// Property: TaskDefinitionType
			// CloudFormation resource type schema:
			// {
			//   "description": "A filter to list only the wireless gateway task definitions that use this task definition type",
			//   "enum": [
			//     "UPDATE"
			//   ],
			//   "type": "string"
			// }
			Description: "A filter to list only the wireless gateway task definitions that use this task definition type",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"UPDATE",
				}),
			},
		},
		"update": {
			// Property: Update
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "LoRaWAN": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "CurrentVersion": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Model": {
			//               "maxLength": 4096,
			//               "minLength": 1,
			//               "type": "string"
			//             },
			//             "PackageVersion": {
			//               "maxLength": 32,
			//               "minLength": 1,
			//               "type": "string"
			//             },
			//             "Station": {
			//               "maxLength": 4096,
			//               "minLength": 1,
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "SigKeyCrc": {
			//           "format": "int64",
			//           "type": "integer"
			//         },
			//         "UpdateSignature": {
			//           "maxLength": 4096,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "UpdateVersion": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "Model": {
			//               "maxLength": 4096,
			//               "minLength": 1,
			//               "type": "string"
			//             },
			//             "PackageVersion": {
			//               "maxLength": 32,
			//               "minLength": 1,
			//               "type": "string"
			//             },
			//             "Station": {
			//               "maxLength": 4096,
			//               "minLength": 1,
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "UpdateDataRole": {
			//       "maxLength": 2048,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "UpdateDataSource": {
			//       "maxLength": 4096,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"lo_ra_wan": {
						// Property: LoRaWAN
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"current_version": {
									// Property: CurrentVersion
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"model": {
												// Property: Model
												Type:     types.StringType,
												Optional: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 4096),
												},
											},
											"package_version": {
												// Property: PackageVersion
												Type:     types.StringType,
												Optional: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 32),
												},
											},
											"station": {
												// Property: Station
												Type:     types.StringType,
												Optional: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 4096),
												},
											},
										},
									),
									Optional: true,
								},
								"sig_key_crc": {
									// Property: SigKeyCrc
									Type:     types.NumberType,
									Optional: true,
								},
								"update_signature": {
									// Property: UpdateSignature
									Type:     types.StringType,
									Optional: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 4096),
									},
								},
								"update_version": {
									// Property: UpdateVersion
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"model": {
												// Property: Model
												Type:     types.StringType,
												Optional: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 4096),
												},
											},
											"package_version": {
												// Property: PackageVersion
												Type:     types.StringType,
												Optional: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 32),
												},
											},
											"station": {
												// Property: Station
												Type:     types.StringType,
												Optional: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 4096),
												},
											},
										},
									),
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"update_data_role": {
						// Property: UpdateDataRole
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 2048),
						},
					},
					"update_data_source": {
						// Property: UpdateDataSource
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 4096),
						},
					},
				},
			),
			Optional: true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Creates a gateway task definition.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTWireless::TaskDefinition").WithTerraformTypeName("awscc_iotwireless_task_definition")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                                 "Arn",
		"auto_create_tasks":                   "AutoCreateTasks",
		"current_version":                     "CurrentVersion",
		"id":                                  "Id",
		"key":                                 "Key",
		"lo_ra_wan":                           "LoRaWAN",
		"lo_ra_wan_update_gateway_task_entry": "LoRaWANUpdateGatewayTaskEntry",
		"model":                               "Model",
		"name":                                "Name",
		"package_version":                     "PackageVersion",
		"sig_key_crc":                         "SigKeyCrc",
		"station":                             "Station",
		"tags":                                "Tags",
		"task_definition_type":                "TaskDefinitionType",
		"update":                              "Update",
		"update_data_role":                    "UpdateDataRole",
		"update_data_source":                  "UpdateDataSource",
		"update_signature":                    "UpdateSignature",
		"update_version":                      "UpdateVersion",
		"value":                               "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_iotwireless_task_definition", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
