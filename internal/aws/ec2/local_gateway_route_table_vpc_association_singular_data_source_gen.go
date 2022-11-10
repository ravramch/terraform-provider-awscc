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
	registry.AddDataSourceFactory("awscc_ec2_local_gateway_route_table_vpc_association", localGatewayRouteTableVPCAssociationDataSource)
}

// localGatewayRouteTableVPCAssociationDataSource returns the Terraform awscc_ec2_local_gateway_route_table_vpc_association data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::LocalGatewayRouteTableVPCAssociation resource.
func localGatewayRouteTableVPCAssociationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"local_gateway_id": {
			// Property: LocalGatewayId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The ID of the local gateway.",
			//	  "type": "string"
			//	}
			Description: "The ID of the local gateway.",
			Type:        types.StringType,
			Computed:    true,
		},
		"local_gateway_route_table_id": {
			// Property: LocalGatewayRouteTableId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The ID of the local gateway route table.",
			//	  "type": "string"
			//	}
			Description: "The ID of the local gateway route table.",
			Type:        types.StringType,
			Computed:    true,
		},
		"local_gateway_route_table_vpc_association_id": {
			// Property: LocalGatewayRouteTableVpcAssociationId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The ID of the association.",
			//	  "type": "string"
			//	}
			Description: "The ID of the association.",
			Type:        types.StringType,
			Computed:    true,
		},
		"state": {
			// Property: State
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The state of the association.",
			//	  "type": "string"
			//	}
			Description: "The state of the association.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The tags for the association.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "Key": {
			//	        "maxLength": 127,
			//	        "minLength": 1,
			//	        "pattern": "",
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "maxLength": 255,
			//	        "minLength": 1,
			//	        "pattern": "",
			//	        "type": "string"
			//	      }
			//	    },
			//	    "type": "object"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Description: "The tags for the association.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"vpc_id": {
			// Property: VpcId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The ID of the VPC.",
			//	  "type": "string"
			//	}
			Description: "The ID of the VPC.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::EC2::LocalGatewayRouteTableVPCAssociation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::LocalGatewayRouteTableVPCAssociation").WithTerraformTypeName("awscc_ec2_local_gateway_route_table_vpc_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"key":                          "Key",
		"local_gateway_id":             "LocalGatewayId",
		"local_gateway_route_table_id": "LocalGatewayRouteTableId",
		"local_gateway_route_table_vpc_association_id": "LocalGatewayRouteTableVpcAssociationId",
		"state":  "State",
		"tags":   "Tags",
		"value":  "Value",
		"vpc_id": "VpcId",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
