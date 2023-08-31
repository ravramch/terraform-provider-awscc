// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iam_managed_policy", managedPolicyDataSource)
}

// managedPolicyDataSource returns the Terraform awscc_iam_managed_policy data source.
// This Terraform data source corresponds to the CloudFormation AWS::IAM::ManagedPolicy resource.
func managedPolicyDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AttachmentCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of entities (users, groups, and roles) that the policy is attached to.",
		//	  "type": "integer"
		//	}
		"attachment_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of entities (users, groups, and roles) that the policy is attached to.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreateDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The date and time, in ISO 8601 date-time format, when the policy was created.",
		//	  "type": "string"
		//	}
		"create_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The date and time, in ISO 8601 date-time format, when the policy was created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DefaultVersionId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier for the version of the policy that is set as the default version.",
		//	  "type": "string"
		//	}
		"default_version_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier for the version of the policy that is set as the default version.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A friendly description of the policy.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A friendly description of the policy.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Groups
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name (friendly name, not ARN) of the group to attach the policy to.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"groups": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The name (friendly name, not ARN) of the group to attach the policy to.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IsAttachable
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies whether the policy can be attached to an IAM user, group, or role.",
		//	  "type": "boolean"
		//	}
		"is_attachable": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether the policy can be attached to an IAM user, group, or role.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ManagedPolicyName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The friendly name of the policy.",
		//	  "type": "string"
		//	}
		"managed_policy_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The friendly name of the policy.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Path
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The path for the policy.",
		//	  "type": "string"
		//	}
		"path": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The path for the policy.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PermissionsBoundaryUsageCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of entities (users and roles) for which the policy is used to set the permissions boundary.",
		//	  "type": "integer"
		//	}
		"permissions_boundary_usage_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of entities (users and roles) for which the policy is used to set the permissions boundary.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PolicyArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amazon Resource Name (ARN) of the managed policy",
		//	  "type": "string"
		//	}
		"policy_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Amazon Resource Name (ARN) of the managed policy",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PolicyDocument
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The JSON policy document that you want to use as the content for the new policy.",
		//	  "type": "string"
		//	}
		"policy_document": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The JSON policy document that you want to use as the content for the new policy.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PolicyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The stable and unique string identifying the policy.",
		//	  "type": "string"
		//	}
		"policy_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The stable and unique string identifying the policy.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Roles
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name (friendly name, not ARN) of the role to attach the policy to.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"roles": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The name (friendly name, not ARN) of the role to attach the policy to.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UpdateDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The date and time, in ISO 8601 date-time format, when the policy was last updated.",
		//	  "type": "string"
		//	}
		"update_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The date and time, in ISO 8601 date-time format, when the policy was last updated.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Users
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name (friendly name, not ARN) of the IAM user to attach the policy to.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"users": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The name (friendly name, not ARN) of the IAM user to attach the policy to.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IAM::ManagedPolicy",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IAM::ManagedPolicy").WithTerraformTypeName("awscc_iam_managed_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"attachment_count":                 "AttachmentCount",
		"create_date":                      "CreateDate",
		"default_version_id":               "DefaultVersionId",
		"description":                      "Description",
		"groups":                           "Groups",
		"is_attachable":                    "IsAttachable",
		"managed_policy_name":              "ManagedPolicyName",
		"path":                             "Path",
		"permissions_boundary_usage_count": "PermissionsBoundaryUsageCount",
		"policy_arn":                       "PolicyArn",
		"policy_document":                  "PolicyDocument",
		"policy_id":                        "PolicyId",
		"roles":                            "Roles",
		"update_date":                      "UpdateDate",
		"users":                            "Users",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}