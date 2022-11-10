// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_lex_bot_version", botVersionDataSource)
}

// botVersionDataSource returns the Terraform awscc_lex_bot_version data source.
// This Terraform data source corresponds to the CloudFormation AWS::Lex::BotVersion resource.
func botVersionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"bot_id": {
			// Property: BotId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Unique ID of resource",
			//	  "maxLength": 10,
			//	  "minLength": 10,
			//	  "pattern": "^[0-9a-zA-Z]+$",
			//	  "type": "string"
			//	}
			Description: "Unique ID of resource",
			Type:        types.StringType,
			Computed:    true,
		},
		"bot_version": {
			// Property: BotVersion
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The version of a bot.",
			//	  "maxLength": 5,
			//	  "minLength": 1,
			//	  "pattern": "^(DRAFT|[0-9]+)$",
			//	  "type": "string"
			//	}
			Description: "The version of a bot.",
			Type:        types.StringType,
			Computed:    true,
		},
		"bot_version_locale_specification": {
			// Property: BotVersionLocaleSpecification
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Specifies the locales that Amazon Lex adds to this version. You can choose the Draft version or any other previously published version for each locale.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "BotVersionLocaleDetails": {
			//	        "additionalProperties": false,
			//	        "description": "The version of a bot used for a bot locale.",
			//	        "properties": {
			//	          "SourceBotVersion": {
			//	            "description": "The version of a bot.",
			//	            "maxLength": 5,
			//	            "minLength": 1,
			//	            "pattern": "^(DRAFT|[0-9]+)$",
			//	            "type": "string"
			//	          }
			//	        },
			//	        "required": [
			//	          "SourceBotVersion"
			//	        ],
			//	        "type": "object"
			//	      },
			//	      "LocaleId": {
			//	        "description": "The identifier of the language and locale that the bot will be used in.",
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "LocaleId",
			//	      "BotVersionLocaleDetails"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "minItems": 1,
			//	  "type": "array"
			//	}
			Description: "Specifies the locales that Amazon Lex adds to this version. You can choose the Draft version or any other previously published version for each locale.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"bot_version_locale_details": {
						// Property: BotVersionLocaleDetails
						Description: "The version of a bot used for a bot locale.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"source_bot_version": {
									// Property: SourceBotVersion
									Description: "The version of a bot.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"locale_id": {
						// Property: LocaleId
						Description: "The identifier of the language and locale that the bot will be used in.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A description of the version. Use the description to help identify the version in lists.",
			//	  "maxLength": 200,
			//	  "type": "string"
			//	}
			Description: "A description of the version. Use the description to help identify the version in lists.",
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
		Description: "Data Source schema for AWS::Lex::BotVersion",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lex::BotVersion").WithTerraformTypeName("awscc_lex_bot_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bot_id":                           "BotId",
		"bot_version":                      "BotVersion",
		"bot_version_locale_details":       "BotVersionLocaleDetails",
		"bot_version_locale_specification": "BotVersionLocaleSpecification",
		"description":                      "Description",
		"locale_id":                        "LocaleId",
		"source_bot_version":               "SourceBotVersion",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
