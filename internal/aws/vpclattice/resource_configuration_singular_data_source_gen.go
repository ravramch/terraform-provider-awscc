// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package vpclattice

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_vpclattice_resource_configuration", resourceConfigurationDataSource)
}

// resourceConfigurationDataSource returns the Terraform awscc_vpclattice_resource_configuration data source.
// This Terraform data source corresponds to the CloudFormation AWS::VpcLattice::ResourceConfiguration resource.
func resourceConfigurationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AllowAssociationToSharableServiceNetwork
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"allow_association_to_sharable_service_network": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:[a-z0-9f\\-]+:vpc-lattice:[a-zA-Z0-9\\-]+:\\d{12}:resourceconfiguration/rcfg-[0-9a-z]{17}$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 22,
		//	  "minLength": 22,
		//	  "pattern": "^rcfg-[0-9a-z]{17}$",
		//	  "type": "string"
		//	}
		"resource_configuration_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 40,
		//	  "minLength": 3,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PortRanges
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 11,
		//	    "minLength": 1,
		//	    "pattern": "^((\\d{1,5}\\-\\d{1,5})|(\\d+))$",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"port_ranges": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProtocolType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "TCP"
		//	  ],
		//	  "type": "string"
		//	}
		"protocol_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceConfigurationAuthType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "NONE",
		//	    "AWS_IAM"
		//	  ],
		//	  "type": "string"
		//	}
		"resource_configuration_auth_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceConfigurationDefinition
		// CloudFormation resource type schema:
		//
		//	{
		//	  "properties": {
		//	    "ArnResource": {
		//	      "maxLength": 1224,
		//	      "pattern": "",
		//	      "type": "string"
		//	    },
		//	    "DnsResource": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "DomainName": {
		//	          "maxLength": 255,
		//	          "minLength": 3,
		//	          "type": "string"
		//	        },
		//	        "IpAddressType": {
		//	          "enum": [
		//	            "IPV4",
		//	            "IPV6",
		//	            "DUALSTACK"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "DomainName",
		//	        "IpAddressType"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "IpResource": {
		//	      "maxLength": 39,
		//	      "minLength": 4,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"resource_configuration_definition": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ArnResource
				"arn_resource": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: DnsResource
				"dns_resource": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: DomainName
						"domain_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: IpAddressType
						"ip_address_type": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: IpResource
				"ip_resource": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceConfigurationGroupId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 22,
		//	  "minLength": 22,
		//	  "pattern": "^rcfg-[0-9a-z]{17}$",
		//	  "type": "string"
		//	}
		"resource_configuration_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceConfigurationType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "GROUP",
		//	    "CHILD",
		//	    "SINGLE",
		//	    "ARN"
		//	  ],
		//	  "type": "string"
		//	}
		"resource_configuration_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceGatewayId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"resource_gateway_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "minItems": 0,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::VpcLattice::ResourceConfiguration",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::VpcLattice::ResourceConfiguration").WithTerraformTypeName("awscc_vpclattice_resource_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allow_association_to_sharable_service_network": "AllowAssociationToSharableServiceNetwork",
		"arn":                               "Arn",
		"arn_resource":                      "ArnResource",
		"dns_resource":                      "DnsResource",
		"domain_name":                       "DomainName",
		"ip_address_type":                   "IpAddressType",
		"ip_resource":                       "IpResource",
		"key":                               "Key",
		"name":                              "Name",
		"port_ranges":                       "PortRanges",
		"protocol_type":                     "ProtocolType",
		"resource_configuration_auth_type":  "ResourceConfigurationAuthType",
		"resource_configuration_definition": "ResourceConfigurationDefinition",
		"resource_configuration_group_id":   "ResourceConfigurationGroupId",
		"resource_configuration_id":         "Id",
		"resource_configuration_type":       "ResourceConfigurationType",
		"resource_gateway_id":               "ResourceGatewayId",
		"tags":                              "Tags",
		"value":                             "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
