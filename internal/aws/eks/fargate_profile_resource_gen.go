// Code generated by generators/resource/main.go; DO NOT EDIT.

package eks

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_eks_fargate_profile", fargateProfileResourceType)
}

// fargateProfileResourceType returns the Terraform awscc_eks_fargate_profile resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EKS::FargateProfile resource type.
func fargateProfileResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"cluster_name": {
			// Property: ClusterName
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of the Cluster",
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Name of the Cluster",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtLeast(1),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"fargate_profile_name": {
			// Property: FargateProfileName
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of FargateProfile",
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Name of FargateProfile",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtLeast(1),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"pod_execution_role_arn": {
			// Property: PodExecutionRoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The IAM policy arn for pods",
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The IAM policy arn for pods",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtLeast(1),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"selectors": {
			// Property: Selectors
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Labels": {
			//         "items": {
			//           "additionalProperties": false,
			//           "description": "A key-value pair to associate with a pod.",
			//           "properties": {
			//             "Key": {
			//               "description": "The key name of the label.",
			//               "maxLength": 127,
			//               "minLength": 1,
			//               "type": "string"
			//             },
			//             "Value": {
			//               "description": "The value for the label. ",
			//               "maxLength": 255,
			//               "minLength": 1,
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "Key",
			//             "Value"
			//           ],
			//           "type": "object"
			//         },
			//         "type": "array"
			//       },
			//       "Namespace": {
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Namespace"
			//     ],
			//     "type": "object"
			//   },
			//   "minItems": 1,
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"labels": {
						// Property: Labels
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"key": {
									// Property: Key
									Description: "The key name of the label.",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 127),
									},
								},
								"value": {
									// Property: Value
									Description: "The value for the label. ",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 255),
									},
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Optional: true,
					},
					"namespace": {
						// Property: Namespace
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtLeast(1),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtLeast(1),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"subnets": {
			// Property: Subnets
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 127),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 255),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.UniqueItems(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource Schema for AWS::EKS::FargateProfile",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EKS::FargateProfile").WithTerraformTypeName("awscc_eks_fargate_profile")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                    "Arn",
		"cluster_name":           "ClusterName",
		"fargate_profile_name":   "FargateProfileName",
		"key":                    "Key",
		"labels":                 "Labels",
		"namespace":              "Namespace",
		"pod_execution_role_arn": "PodExecutionRoleArn",
		"selectors":              "Selectors",
		"subnets":                "Subnets",
		"tags":                   "Tags",
		"value":                  "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
