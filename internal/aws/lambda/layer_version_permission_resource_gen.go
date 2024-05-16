// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_lambda_layer_version_permission", layerVersionPermissionResource)
}

// layerVersionPermissionResource returns the Terraform awscc_lambda_layer_version_permission resource.
// This Terraform resource corresponds to the CloudFormation AWS::Lambda::LayerVersionPermission resource.
func layerVersionPermissionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Action
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The API action that grants access to the layer.",
		//	  "type": "string"
		//	}
		"action": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The API action that grants access to the layer.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ID generated by service",
		//	  "type": "string"
		//	}
		"layer_version_permission_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ID generated by service",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LayerVersionArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name or Amazon Resource Name (ARN) of the layer.",
		//	  "type": "string"
		//	}
		"layer_version_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name or Amazon Resource Name (ARN) of the layer.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: OrganizationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "With the principal set to *, grant permission to all accounts in the specified organization.",
		//	  "type": "string"
		//	}
		"organization_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "With the principal set to *, grant permission to all accounts in the specified organization.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Principal
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An account ID, or * to grant layer usage permission to all accounts in an organization, or all AWS accounts (if organizationId is not specified).",
		//	  "type": "string"
		//	}
		"principal": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An account ID, or * to grant layer usage permission to all accounts in an organization, or all AWS accounts (if organizationId is not specified).",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
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
		Description: "Schema for Lambda LayerVersionPermission",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lambda::LayerVersionPermission").WithTerraformTypeName("awscc_lambda_layer_version_permission")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action":                      "Action",
		"layer_version_arn":           "LayerVersionArn",
		"layer_version_permission_id": "Id",
		"organization_id":             "OrganizationId",
		"principal":                   "Principal",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
