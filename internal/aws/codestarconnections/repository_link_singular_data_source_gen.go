// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package codestarconnections

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_codestarconnections_repository_link", repositoryLinkDataSource)
}

// repositoryLinkDataSource returns the Terraform awscc_codestarconnections_repository_link data source.
// This Terraform data source corresponds to the CloudFormation AWS::CodeStarConnections::RepositoryLink resource.
func repositoryLinkDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ConnectionArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the CodeStarConnection. The ARN is used as the connection reference when the connection is shared between AWS services.",
		//	  "pattern": "arn:(aws|aws-us-gov|aws-cn):.+:.+:[0-9]{12}:.+",
		//	  "type": "string"
		//	}
		"connection_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the CodeStarConnection. The ARN is used as the connection reference when the connection is shared between AWS services.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EncryptionKeyArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the KMS key that the customer can optionally specify to use to encrypt RepositoryLink properties. If not specified, a default key will be used.",
		//	  "pattern": "arn:(aws|aws-us-gov|aws-cn):.+:.+:[0-9]{12}:.+",
		//	  "type": "string"
		//	}
		"encryption_key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the KMS key that the customer can optionally specify to use to encrypt RepositoryLink properties. If not specified, a default key will be used.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OwnerId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "the ID of the entity that owns the repository.",
		//	  "pattern": "[a-za-z0-9_\\.-]+",
		//	  "type": "string"
		//	}
		"owner_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "the ID of the entity that owns the repository.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProviderType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the external provider where your third-party code repository is configured.",
		//	  "pattern": "^(GitHub|Bitbucket|GitHubEnterprise|GitLab)$",
		//	  "type": "string"
		//	}
		"provider_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the external provider where your third-party code repository is configured.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RepositoryLinkArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique Amazon Resource Name (ARN) to designate the repository link.",
		//	  "pattern": "arn:(aws|aws-us-gov|aws-cn):.+:.+:[0-9]{12}:.+",
		//	  "type": "string"
		//	}
		"repository_link_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique Amazon Resource Name (ARN) to designate the repository link.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RepositoryLinkId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A UUID that uniquely identifies the RepositoryLink.",
		//	  "pattern": "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}",
		//	  "type": "string"
		//	}
		"repository_link_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A UUID that uniquely identifies the RepositoryLink.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RepositoryName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The repository for which the link is being created.",
		//	  "pattern": "[a-za-z0-9_\\.-]+",
		//	  "type": "string"
		//	}
		"repository_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The repository for which the link is being created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the tags applied to a RepositoryLink.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, , ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, , ., /, =, +, and -. ",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
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
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, , ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, , ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Specifies the tags applied to a RepositoryLink.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CodeStarConnections::RepositoryLink",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CodeStarConnections::RepositoryLink").WithTerraformTypeName("awscc_codestarconnections_repository_link")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"connection_arn":      "ConnectionArn",
		"encryption_key_arn":  "EncryptionKeyArn",
		"key":                 "Key",
		"owner_id":            "OwnerId",
		"provider_type":       "ProviderType",
		"repository_link_arn": "RepositoryLinkArn",
		"repository_link_id":  "RepositoryLinkId",
		"repository_name":     "RepositoryName",
		"tags":                "Tags",
		"value":               "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
