// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package signer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_signer_signing_profile", signingProfileDataSource)
}

// signingProfileDataSource returns the Terraform awscc_signer_signing_profile data source.
// This Terraform data source corresponds to the CloudFormation AWS::Signer::SigningProfile resource.
func signingProfileDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) of the specified signing profile.",
			//	  "pattern": "^arn:aws(-(cn|gov))?:[a-z-]+:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) of the specified signing profile.",
			Type:        types.StringType,
			Computed:    true,
		},
		"platform_id": {
			// Property: PlatformId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The ID of the target signing platform.",
			//	  "enum": [
			//	    "AWSLambda-SHA384-ECDSA"
			//	  ],
			//	  "type": "string"
			//	}
			Description: "The ID of the target signing platform.",
			Type:        types.StringType,
			Computed:    true,
		},
		"profile_name": {
			// Property: ProfileName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A name for the signing profile. AWS CloudFormation generates a unique physical ID and uses that ID for the signing profile name. ",
			//	  "type": "string"
			//	}
			Description: "A name for the signing profile. AWS CloudFormation generates a unique physical ID and uses that ID for the signing profile name. ",
			Type:        types.StringType,
			Computed:    true,
		},
		"profile_version": {
			// Property: ProfileVersion
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A version for the signing profile. AWS Signer generates a unique version for each profile of the same profile name.",
			//	  "pattern": "^[0-9a-zA-Z]{10}$",
			//	  "type": "string"
			//	}
			Description: "A version for the signing profile. AWS Signer generates a unique version for each profile of the same profile name.",
			Type:        types.StringType,
			Computed:    true,
		},
		"profile_version_arn": {
			// Property: ProfileVersionArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) of the specified signing profile version.",
			//	  "pattern": "^arn:aws(-(cn|gov))?:[a-z-]+:(([a-z]+-)+[0-9])?:([0-9]{12})?:[^.]+$",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) of the specified signing profile version.",
			Type:        types.StringType,
			Computed:    true,
		},
		"signature_validity_period": {
			// Property: SignatureValidityPeriod
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "Signature validity period of the profile.",
			//	  "properties": {
			//	    "Type": {
			//	      "enum": [
			//	        "DAYS",
			//	        "MONTHS",
			//	        "YEARS"
			//	      ],
			//	      "type": "string"
			//	    },
			//	    "Value": {
			//	      "type": "integer"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "Signature validity period of the profile.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"type": {
						// Property: Type
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.Int64Type,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A list of tags associated with the signing profile.",
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
			//	        "type": "string"
			//	      }
			//	    },
			//	    "type": "object"
			//	  },
			//	  "type": "array"
			//	}
			Description: "A list of tags associated with the signing profile.",
			Attributes: tfsdk.ListNestedAttributes(
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
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Signer::SigningProfile",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Signer::SigningProfile").WithTerraformTypeName("awscc_signer_signing_profile")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                       "Arn",
		"key":                       "Key",
		"platform_id":               "PlatformId",
		"profile_name":              "ProfileName",
		"profile_version":           "ProfileVersion",
		"profile_version_arn":       "ProfileVersionArn",
		"signature_validity_period": "SignatureValidityPeriod",
		"tags":                      "Tags",
		"type":                      "Type",
		"value":                     "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
