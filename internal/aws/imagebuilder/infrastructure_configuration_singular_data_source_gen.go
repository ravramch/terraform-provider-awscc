// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_imagebuilder_infrastructure_configuration", infrastructureConfigurationDataSource)
}

// infrastructureConfigurationDataSource returns the Terraform awscc_imagebuilder_infrastructure_configuration data source.
// This Terraform data source corresponds to the CloudFormation AWS::ImageBuilder::InfrastructureConfiguration resource.
func infrastructureConfigurationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) of the infrastructure configuration.",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) of the infrastructure configuration.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The description of the infrastructure configuration.",
			//	  "type": "string"
			//	}
			Description: "The description of the infrastructure configuration.",
			Type:        types.StringType,
			Computed:    true,
		},
		"instance_metadata_options": {
			// Property: InstanceMetadataOptions
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "The instance metadata option settings for the infrastructure configuration.",
			//	  "properties": {
			//	    "HttpPutResponseHopLimit": {
			//	      "description": "Limit the number of hops that an instance metadata request can traverse to reach its destination.",
			//	      "type": "integer"
			//	    },
			//	    "HttpTokens": {
			//	      "description": "Indicates whether a signed token header is required for instance metadata retrieval requests. The values affect the response as follows: ",
			//	      "enum": [
			//	        "required",
			//	        "optional"
			//	      ],
			//	      "type": "string"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "The instance metadata option settings for the infrastructure configuration.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"http_put_response_hop_limit": {
						// Property: HttpPutResponseHopLimit
						Description: "Limit the number of hops that an instance metadata request can traverse to reach its destination.",
						Type:        types.Int64Type,
						Computed:    true,
					},
					"http_tokens": {
						// Property: HttpTokens
						Description: "Indicates whether a signed token header is required for instance metadata retrieval requests. The values affect the response as follows: ",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"instance_profile_name": {
			// Property: InstanceProfileName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The instance profile of the infrastructure configuration.",
			//	  "type": "string"
			//	}
			Description: "The instance profile of the infrastructure configuration.",
			Type:        types.StringType,
			Computed:    true,
		},
		"instance_types": {
			// Property: InstanceTypes
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The instance types of the infrastructure configuration.",
			//	  "insertionOrder": true,
			//	  "items": {
			//	    "type": "string"
			//	  },
			//	  "type": "array"
			//	}
			Description: "The instance types of the infrastructure configuration.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"key_pair": {
			// Property: KeyPair
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The EC2 key pair of the infrastructure configuration..",
			//	  "type": "string"
			//	}
			Description: "The EC2 key pair of the infrastructure configuration..",
			Type:        types.StringType,
			Computed:    true,
		},
		"logging": {
			// Property: Logging
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "The logging configuration of the infrastructure configuration.",
			//	  "properties": {
			//	    "S3Logs": {
			//	      "additionalProperties": false,
			//	      "description": "The S3 path in which to store the logs.",
			//	      "properties": {
			//	        "S3BucketName": {
			//	          "description": "S3BucketName",
			//	          "type": "string"
			//	        },
			//	        "S3KeyPrefix": {
			//	          "description": "S3KeyPrefix",
			//	          "type": "string"
			//	        }
			//	      },
			//	      "type": "object"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "The logging configuration of the infrastructure configuration.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"s3_logs": {
						// Property: S3Logs
						Description: "The S3 path in which to store the logs.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"s3_bucket_name": {
									// Property: S3BucketName
									Description: "S3BucketName",
									Type:        types.StringType,
									Computed:    true,
								},
								"s3_key_prefix": {
									// Property: S3KeyPrefix
									Description: "S3KeyPrefix",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the infrastructure configuration.",
			//	  "type": "string"
			//	}
			Description: "The name of the infrastructure configuration.",
			Type:        types.StringType,
			Computed:    true,
		},
		"resource_tags": {
			// Property: ResourceTags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "The tags attached to the resource created by Image Builder.",
			//	  "patternProperties": {
			//	    "": {
			//	      "type": "string"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "The tags attached to the resource created by Image Builder.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Computed: true,
		},
		"security_group_ids": {
			// Property: SecurityGroupIds
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The security group IDs of the infrastructure configuration.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "type": "string"
			//	  },
			//	  "type": "array"
			//	}
			Description: "The security group IDs of the infrastructure configuration.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"sns_topic_arn": {
			// Property: SnsTopicArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The SNS Topic Amazon Resource Name (ARN) of the infrastructure configuration.",
			//	  "type": "string"
			//	}
			Description: "The SNS Topic Amazon Resource Name (ARN) of the infrastructure configuration.",
			Type:        types.StringType,
			Computed:    true,
		},
		"subnet_id": {
			// Property: SubnetId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The subnet ID of the infrastructure configuration.",
			//	  "type": "string"
			//	}
			Description: "The subnet ID of the infrastructure configuration.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "The tags associated with the component.",
			//	  "patternProperties": {
			//	    "": {
			//	      "type": "string"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "The tags associated with the component.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Computed: true,
		},
		"terminate_instance_on_failure": {
			// Property: TerminateInstanceOnFailure
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The terminate instance on failure configuration of the infrastructure configuration.",
			//	  "type": "boolean"
			//	}
			Description: "The terminate instance on failure configuration of the infrastructure configuration.",
			Type:        types.BoolType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::ImageBuilder::InfrastructureConfiguration",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ImageBuilder::InfrastructureConfiguration").WithTerraformTypeName("awscc_imagebuilder_infrastructure_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                           "Arn",
		"description":                   "Description",
		"http_put_response_hop_limit":   "HttpPutResponseHopLimit",
		"http_tokens":                   "HttpTokens",
		"instance_metadata_options":     "InstanceMetadataOptions",
		"instance_profile_name":         "InstanceProfileName",
		"instance_types":                "InstanceTypes",
		"key_pair":                      "KeyPair",
		"logging":                       "Logging",
		"name":                          "Name",
		"resource_tags":                 "ResourceTags",
		"s3_bucket_name":                "S3BucketName",
		"s3_key_prefix":                 "S3KeyPrefix",
		"s3_logs":                       "S3Logs",
		"security_group_ids":            "SecurityGroupIds",
		"sns_topic_arn":                 "SnsTopicArn",
		"subnet_id":                     "SubnetId",
		"tags":                          "Tags",
		"terminate_instance_on_failure": "TerminateInstanceOnFailure",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
