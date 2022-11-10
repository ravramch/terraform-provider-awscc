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
	registry.AddDataSourceFactory("awscc_ec2_egress_only_internet_gateway", egressOnlyInternetGatewayDataSource)
}

// egressOnlyInternetGatewayDataSource returns the Terraform awscc_ec2_egress_only_internet_gateway data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::EgressOnlyInternetGateway resource.
func egressOnlyInternetGatewayDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Service Generated ID of the EgressOnlyInternetGateway",
			//	  "type": "string"
			//	}
			Description: "Service Generated ID of the EgressOnlyInternetGateway",
			Type:        types.StringType,
			Computed:    true,
		},
		"vpc_id": {
			// Property: VpcId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The ID of the VPC for which to create the egress-only internet gateway.",
			//	  "type": "string"
			//	}
			Description: "The ID of the VPC for which to create the egress-only internet gateway.",
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
		Description: "Data Source schema for AWS::EC2::EgressOnlyInternetGateway",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::EgressOnlyInternetGateway").WithTerraformTypeName("awscc_ec2_egress_only_internet_gateway")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"id":     "Id",
		"vpc_id": "VpcId",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
