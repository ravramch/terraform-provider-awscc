// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_personalize_schema", schemaDataSource)
}

// schemaDataSource returns the Terraform awscc_personalize_schema data source.
// This Terraform data source corresponds to the CloudFormation AWS::Personalize::Schema resource.
func schemaDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"domain": {
			// Property: Domain
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The domain of a Domain dataset group.",
			//	  "enum": [
			//	    "ECOMMERCE",
			//	    "VIDEO_ON_DEMAND"
			//	  ],
			//	  "type": "string"
			//	}
			Description: "The domain of a Domain dataset group.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Name for the schema.",
			//	  "maxLength": 63,
			//	  "minLength": 1,
			//	  "pattern": "^[a-zA-Z0-9][a-zA-Z0-9\\-_]*",
			//	  "type": "string"
			//	}
			Description: "Name for the schema.",
			Type:        types.StringType,
			Computed:    true,
		},
		"schema": {
			// Property: Schema
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A schema in Avro JSON format.",
			//	  "maxLength": 10000,
			//	  "type": "string"
			//	}
			Description: "A schema in Avro JSON format.",
			Type:        types.StringType,
			Computed:    true,
		},
		"schema_arn": {
			// Property: SchemaArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Arn for the schema.",
			//	  "maxLength": 256,
			//	  "pattern": "arn:([a-z\\d-]+):personalize:.*:.*:.+",
			//	  "type": "string"
			//	}
			Description: "Arn for the schema.",
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
		Description: "Data Source schema for AWS::Personalize::Schema",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Personalize::Schema").WithTerraformTypeName("awscc_personalize_schema")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"domain":     "Domain",
		"name":       "Name",
		"schema":     "Schema",
		"schema_arn": "SchemaArn",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
