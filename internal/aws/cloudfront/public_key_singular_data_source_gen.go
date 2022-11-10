// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_cloudfront_public_key", publicKeyDataSource)
}

// publicKeyDataSource returns the Terraform awscc_cloudfront_public_key data source.
// This Terraform data source corresponds to the CloudFormation AWS::CloudFront::PublicKey resource.
func publicKeyDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"created_time": {
			// Property: CreatedTime
			// CloudFormation resource type schema:
			//
			//	{
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
		"public_key_config": {
			// Property: PublicKeyConfig
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "properties": {
			//	    "CallerReference": {
			//	      "type": "string"
			//	    },
			//	    "Comment": {
			//	      "type": "string"
			//	    },
			//	    "EncodedKey": {
			//	      "type": "string"
			//	    },
			//	    "Name": {
			//	      "type": "string"
			//	    }
			//	  },
			//	  "required": [
			//	    "CallerReference",
			//	    "Name",
			//	    "EncodedKey"
			//	  ],
			//	  "type": "object"
			//	}
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"caller_reference": {
						// Property: CallerReference
						Type:     types.StringType,
						Computed: true,
					},
					"comment": {
						// Property: Comment
						Type:     types.StringType,
						Computed: true,
					},
					"encoded_key": {
						// Property: EncodedKey
						Type:     types.StringType,
						Computed: true,
					},
					"name": {
						// Property: Name
						Type:     types.StringType,
						Computed: true,
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
		Description: "Data Source schema for AWS::CloudFront::PublicKey",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFront::PublicKey").WithTerraformTypeName("awscc_cloudfront_public_key")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"caller_reference":  "CallerReference",
		"comment":           "Comment",
		"created_time":      "CreatedTime",
		"encoded_key":       "EncodedKey",
		"id":                "Id",
		"name":              "Name",
		"public_key_config": "PublicKeyConfig",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
