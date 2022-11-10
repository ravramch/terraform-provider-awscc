// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package servicecatalogappregistry

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_servicecatalogappregistry_attribute_group_association", attributeGroupAssociationDataSource)
}

// attributeGroupAssociationDataSource returns the Terraform awscc_servicecatalogappregistry_attribute_group_association data source.
// This Terraform data source corresponds to the CloudFormation AWS::ServiceCatalogAppRegistry::AttributeGroupAssociation resource.
func attributeGroupAssociationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"application": {
			// Property: Application
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name or the Id of the Application.",
			//	  "maxLength": 256,
			//	  "minLength": 1,
			//	  "pattern": "\\w+|[a-z0-9]{12}",
			//	  "type": "string"
			//	}
			Description: "The name or the Id of the Application.",
			Type:        types.StringType,
			Computed:    true,
		},
		"application_arn": {
			// Property: ApplicationArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "pattern": "arn:aws[-a-z]*:servicecatalog:[a-z]{2}(-gov)?-[a-z]+-\\d:\\d{12}:/applications/[a-z0-9]+",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"attribute_group": {
			// Property: AttributeGroup
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name or the Id of the AttributeGroup.",
			//	  "maxLength": 256,
			//	  "minLength": 1,
			//	  "pattern": "\\w+|[a-z0-9]{12}",
			//	  "type": "string"
			//	}
			Description: "The name or the Id of the AttributeGroup.",
			Type:        types.StringType,
			Computed:    true,
		},
		"attribute_group_arn": {
			// Property: AttributeGroupArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "pattern": "arn:aws[-a-z]*:servicecatalog:[a-z]{2}(-gov)?-[a-z]+-\\d:\\d{12}:/attribute-groups/[a-z0-9]+",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::ServiceCatalogAppRegistry::AttributeGroupAssociation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ServiceCatalogAppRegistry::AttributeGroupAssociation").WithTerraformTypeName("awscc_servicecatalogappregistry_attribute_group_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"application":         "Application",
		"application_arn":     "ApplicationArn",
		"attribute_group":     "AttributeGroup",
		"attribute_group_arn": "AttributeGroupArn",
		"id":                  "Id",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
