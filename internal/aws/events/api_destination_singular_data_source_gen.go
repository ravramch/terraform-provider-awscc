// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package events

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_events_api_destination", apiDestinationDataSource)
}

// apiDestinationDataSource returns the Terraform awscc_events_api_destination data source.
// This Terraform data source corresponds to the CloudFormation AWS::Events::ApiDestination resource.
func apiDestinationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The arn of the api destination.",
			//	  "type": "string"
			//	}
			Description: "The arn of the api destination.",
			Type:        types.StringType,
			Computed:    true,
		},
		"connection_arn": {
			// Property: ConnectionArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The arn of the connection.",
			//	  "type": "string"
			//	}
			Description: "The arn of the connection.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 512,
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"http_method": {
			// Property: HttpMethod
			// CloudFormation resource type schema:
			//
			//	{
			//	  "enum": [
			//	    "GET",
			//	    "HEAD",
			//	    "POST",
			//	    "OPTIONS",
			//	    "PUT",
			//	    "DELETE",
			//	    "PATCH"
			//	  ],
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"invocation_endpoint": {
			// Property: InvocationEndpoint
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Url endpoint to invoke.",
			//	  "type": "string"
			//	}
			Description: "Url endpoint to invoke.",
			Type:        types.StringType,
			Computed:    true,
		},
		"invocation_rate_limit_per_second": {
			// Property: InvocationRateLimitPerSecond
			// CloudFormation resource type schema:
			//
			//	{
			//	  "minimum": 1,
			//	  "type": "integer"
			//	}
			Type:     types.Int64Type,
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Name of the apiDestination.",
			//	  "maxLength": 64,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "Name of the apiDestination.",
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
		Description: "Data Source schema for AWS::Events::ApiDestination",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Events::ApiDestination").WithTerraformTypeName("awscc_events_api_destination")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                              "Arn",
		"connection_arn":                   "ConnectionArn",
		"description":                      "Description",
		"http_method":                      "HttpMethod",
		"invocation_endpoint":              "InvocationEndpoint",
		"invocation_rate_limit_per_second": "InvocationRateLimitPerSecond",
		"name":                             "Name",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
