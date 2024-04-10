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
	registry.AddDataSourceFactory("awscc_iam_instance_profile", instanceProfileDataSource)
}

// instanceProfileDataSource returns the Terraform awscc_iam_instance_profile data source.
// This Terraform data source corresponds to the CloudFormation AWS::IAM::InstanceProfile resource.
func instanceProfileDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceProfileName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the instance profile to create.\n This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-",
		//	  "type": "string"
		//	}
		"instance_profile_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the instance profile to create.\n This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Path
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The path to the instance profile. For more information about paths, see [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the *IAM User Guide*.\n This parameter is optional. If it is not included, it defaults to a slash (/).\n This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of either a forward slash (/) by itself or a string that must begin and end with forward slashes. In addition, it can contain any ASCII character from the ! (``\\u0021``) through the DEL character (``\\u007F``), including most punctuation characters, digits, and upper and lowercased letters.",
		//	  "type": "string"
		//	}
		"path": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The path to the instance profile. For more information about paths, see [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the *IAM User Guide*.\n This parameter is optional. If it is not included, it defaults to a slash (/).\n This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of either a forward slash (/) by itself or a string that must begin and end with forward slashes. In addition, it can contain any ASCII character from the ! (``\\u0021``) through the DEL character (``\\u007F``), including most punctuation characters, digits, and upper and lowercased letters.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Roles
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the role to associate with the instance profile. Only one role can be assigned to an EC2 instance at a time, and all applications on the instance share the same role and permissions.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"roles": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The name of the role to associate with the instance profile. Only one role can be assigned to an EC2 instance at a time, and all applications on the instance share the same role and permissions.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IAM::InstanceProfile",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IAM::InstanceProfile").WithTerraformTypeName("awscc_iam_instance_profile")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                   "Arn",
		"instance_profile_name": "InstanceProfileName",
		"path":                  "Path",
		"roles":                 "Roles",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
