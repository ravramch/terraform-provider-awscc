// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_ipam_pool", iPAMPoolDataSource)
}

// iPAMPoolDataSource returns the Terraform awscc_ec2_ipam_pool data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::IPAMPool resource.
func iPAMPoolDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"address_family": {
			// Property: AddressFamily
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The address family of the address space in this pool. Either IPv4 or IPv6.",
			//	  "type": "string"
			//	}
			Description: "The address family of the address space in this pool. Either IPv4 or IPv6.",
			Type:        types.StringType,
			Computed:    true,
		},
		"allocation_default_netmask_length": {
			// Property: AllocationDefaultNetmaskLength
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The default netmask length for allocations made from this pool. This value is used when the netmask length of an allocation isn't specified.",
			//	  "type": "integer"
			//	}
			Description: "The default netmask length for allocations made from this pool. This value is used when the netmask length of an allocation isn't specified.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"allocation_max_netmask_length": {
			// Property: AllocationMaxNetmaskLength
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The maximum allowed netmask length for allocations made from this pool.",
			//	  "type": "integer"
			//	}
			Description: "The maximum allowed netmask length for allocations made from this pool.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"allocation_min_netmask_length": {
			// Property: AllocationMinNetmaskLength
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The minimum allowed netmask length for allocations made from this pool.",
			//	  "type": "integer"
			//	}
			Description: "The minimum allowed netmask length for allocations made from this pool.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"allocation_resource_tags": {
			// Property: AllocationResourceTags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "When specified, an allocation will not be allowed unless a resource has a matching set of tags.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "A key-value pair to associate with a resource.",
			//	    "properties": {
			//	      "Key": {
			//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//	        "maxLength": 128,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//	        "maxLength": 256,
			//	        "minLength": 0,
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Key",
			//	      "Value"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Description: "When specified, an allocation will not be allowed unless a resource has a matching set of tags.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) of the IPAM Pool.",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) of the IPAM Pool.",
			Type:        types.StringType,
			Computed:    true,
		},
		"auto_import": {
			// Property: AutoImport
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Determines what to do if IPAM discovers resources that haven't been assigned an allocation. If set to true, an allocation will be made automatically.",
			//	  "type": "boolean"
			//	}
			Description: "Determines what to do if IPAM discovers resources that haven't been assigned an allocation. If set to true, an allocation will be made automatically.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"aws_service": {
			// Property: AwsService
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Limits which service in Amazon Web Services that the pool can be used in.",
			//	  "enum": [
			//	    "ec2"
			//	  ],
			//	  "type": "string"
			//	}
			Description: "Limits which service in Amazon Web Services that the pool can be used in.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"ipam_arn": {
			// Property: IpamArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) of the IPAM this pool is a part of.",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) of the IPAM this pool is a part of.",
			Type:        types.StringType,
			Computed:    true,
		},
		"ipam_pool_id": {
			// Property: IpamPoolId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Id of the IPAM Pool.",
			//	  "type": "string"
			//	}
			Description: "Id of the IPAM Pool.",
			Type:        types.StringType,
			Computed:    true,
		},
		"ipam_scope_arn": {
			// Property: IpamScopeArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) of the scope this pool is a part of.",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) of the scope this pool is a part of.",
			Type:        types.StringType,
			Computed:    true,
		},
		"ipam_scope_id": {
			// Property: IpamScopeId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Id of the scope this pool is a part of.",
			//	  "type": "string"
			//	}
			Description: "The Id of the scope this pool is a part of.",
			Type:        types.StringType,
			Computed:    true,
		},
		"ipam_scope_type": {
			// Property: IpamScopeType
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Determines whether this scope contains publicly routable space or space for a private network",
			//	  "enum": [
			//	    "public",
			//	    "private"
			//	  ],
			//	  "type": "string"
			//	}
			Description: "Determines whether this scope contains publicly routable space or space for a private network",
			Type:        types.StringType,
			Computed:    true,
		},
		"locale": {
			// Property: Locale
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The region of this pool. If not set, this will default to \"None\" which will disable non-custom allocations. If the locale has been specified for the source pool, this value must match.",
			//	  "type": "string"
			//	}
			Description: "The region of this pool. If not set, this will default to \"None\" which will disable non-custom allocations. If the locale has been specified for the source pool, this value must match.",
			Type:        types.StringType,
			Computed:    true,
		},
		"pool_depth": {
			// Property: PoolDepth
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The depth of this pool in the source pool hierarchy.",
			//	  "type": "integer"
			//	}
			Description: "The depth of this pool in the source pool hierarchy.",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"provisioned_cidrs": {
			// Property: ProvisionedCidrs
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A list of cidrs representing the address space available for allocation in this pool.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "An address space to be inserted into this pool. All allocations must be made from this address space.",
			//	    "properties": {
			//	      "Cidr": {
			//	        "description": "Represents a single IPv4 or IPv6 CIDR",
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Cidr"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Description: "A list of cidrs representing the address space available for allocation in this pool.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"cidr": {
						// Property: Cidr
						Description: "Represents a single IPv4 or IPv6 CIDR",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"publicly_advertisable": {
			// Property: PubliclyAdvertisable
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Determines whether or not address space from this pool is publicly advertised. Must be set if and only if the pool is IPv6.",
			//	  "type": "boolean"
			//	}
			Description: "Determines whether or not address space from this pool is publicly advertised. Must be set if and only if the pool is IPv6.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"source_ipam_pool_id": {
			// Property: SourceIpamPoolId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Id of this pool's source. If set, all space provisioned in this pool must be free space provisioned in the parent pool.",
			//	  "type": "string"
			//	}
			Description: "The Id of this pool's source. If set, all space provisioned in this pool must be free space provisioned in the parent pool.",
			Type:        types.StringType,
			Computed:    true,
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The state of this pool. This can be one of the following values: \"create-in-progress\", \"create-complete\", \"modify-in-progress\", \"modify-complete\", \"delete-in-progress\", or \"delete-complete\"",
			//	  "enum": [
			//	    "create-in-progress",
			//	    "create-complete",
			//	    "modify-in-progress",
			//	    "modify-complete",
			//	    "delete-in-progress",
			//	    "delete-complete"
			//	  ],
			//	  "type": "string"
			//	}
			Description: "The state of this pool. This can be one of the following values: \"create-in-progress\", \"create-complete\", \"modify-in-progress\", \"modify-complete\", \"delete-in-progress\", or \"delete-complete\"",
			Type:        types.StringType,
			Computed:    true,
		},
		"state_message": {
			// Property: StateMessage
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "An explanation of how the pool arrived at it current state.",
			//	  "type": "string"
			//	}
			Description: "An explanation of how the pool arrived at it current state.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
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
			//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//	        "maxLength": 128,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//	        "maxLength": 256,
			//	        "minLength": 0,
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Key",
			//	      "Value"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::EC2::IPAMPool",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::IPAMPool").WithTerraformTypeName("awscc_ec2_ipam_pool")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"address_family":                    "AddressFamily",
		"allocation_default_netmask_length": "AllocationDefaultNetmaskLength",
		"allocation_max_netmask_length":     "AllocationMaxNetmaskLength",
		"allocation_min_netmask_length":     "AllocationMinNetmaskLength",
		"allocation_resource_tags":          "AllocationResourceTags",
		"arn":                               "Arn",
		"auto_import":                       "AutoImport",
		"aws_service":                       "AwsService",
		"cidr":                              "Cidr",
		"description":                       "Description",
		"ipam_arn":                          "IpamArn",
		"ipam_pool_id":                      "IpamPoolId",
		"ipam_scope_arn":                    "IpamScopeArn",
		"ipam_scope_id":                     "IpamScopeId",
		"ipam_scope_type":                   "IpamScopeType",
		"key":                               "Key",
		"locale":                            "Locale",
		"pool_depth":                        "PoolDepth",
		"provisioned_cidrs":                 "ProvisionedCidrs",
		"publicly_advertisable":             "PubliclyAdvertisable",
		"source_ipam_pool_id":               "SourceIpamPoolId",
		"state":                             "State",
		"state_message":                     "StateMessage",
		"tags":                              "Tags",
		"value":                             "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
