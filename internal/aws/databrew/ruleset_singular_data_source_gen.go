// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package databrew

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_databrew_ruleset", rulesetDataSource)
}

// rulesetDataSource returns the Terraform awscc_databrew_ruleset data source.
// This Terraform data source corresponds to the CloudFormation AWS::DataBrew::Ruleset resource.
func rulesetDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Description of the Ruleset",
			//	  "maxLength": 1024,
			//	  "type": "string"
			//	}
			Description: "Description of the Ruleset",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Name of the Ruleset",
			//	  "maxLength": 255,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "Name of the Ruleset",
			Type:        types.StringType,
			Computed:    true,
		},
		"rules": {
			// Property: Rules
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "List of the data quality rules in the ruleset",
			//	  "insertionOrder": true,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "Data quality rule for a target resource (dataset)",
			//	    "properties": {
			//	      "CheckExpression": {
			//	        "description": "Expression with rule conditions",
			//	        "maxLength": 1024,
			//	        "minLength": 4,
			//	        "pattern": "^[\u003e\u003c0-9A-Za-z_.,:)(!= ]+$",
			//	        "type": "string"
			//	      },
			//	      "ColumnSelectors": {
			//	        "insertionOrder": true,
			//	        "items": {
			//	          "additionalProperties": false,
			//	          "description": "Selector of a column from a dataset for profile job configuration. One selector includes either a column name or a regular expression",
			//	          "properties": {
			//	            "Name": {
			//	              "description": "The name of a column from a dataset",
			//	              "maxLength": 255,
			//	              "minLength": 1,
			//	              "type": "string"
			//	            },
			//	            "Regex": {
			//	              "description": "A regular expression for selecting a column from a dataset",
			//	              "maxLength": 255,
			//	              "minLength": 1,
			//	              "type": "string"
			//	            }
			//	          },
			//	          "type": "object"
			//	        },
			//	        "minItems": 1,
			//	        "type": "array"
			//	      },
			//	      "Disabled": {
			//	        "description": "Boolean value to disable/enable a rule",
			//	        "type": "boolean"
			//	      },
			//	      "Name": {
			//	        "description": "Name of the rule",
			//	        "maxLength": 128,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      },
			//	      "SubstitutionMap": {
			//	        "insertionOrder": true,
			//	        "items": {
			//	          "additionalProperties": false,
			//	          "description": "A key-value pair to associate expression's substitution variable names with their values",
			//	          "properties": {
			//	            "Value": {
			//	              "description": "Value or column name",
			//	              "maxLength": 1024,
			//	              "minLength": 0,
			//	              "type": "string"
			//	            },
			//	            "ValueReference": {
			//	              "description": "Variable name",
			//	              "maxLength": 128,
			//	              "minLength": 2,
			//	              "pattern": "^:[A-Za-z0-9_]+$",
			//	              "type": "string"
			//	            }
			//	          },
			//	          "required": [
			//	            "ValueReference",
			//	            "Value"
			//	          ],
			//	          "type": "object"
			//	        },
			//	        "type": "array"
			//	      },
			//	      "Threshold": {
			//	        "additionalProperties": false,
			//	        "properties": {
			//	          "Type": {
			//	            "description": "Threshold type for a rule",
			//	            "enum": [
			//	              "GREATER_THAN_OR_EQUAL",
			//	              "LESS_THAN_OR_EQUAL",
			//	              "GREATER_THAN",
			//	              "LESS_THAN"
			//	            ],
			//	            "type": "string"
			//	          },
			//	          "Unit": {
			//	            "description": "Threshold unit for a rule",
			//	            "enum": [
			//	              "COUNT",
			//	              "PERCENTAGE"
			//	            ],
			//	            "type": "string"
			//	          },
			//	          "Value": {
			//	            "description": "Threshold value for a rule",
			//	            "type": "number"
			//	          }
			//	        },
			//	        "required": [
			//	          "Value"
			//	        ],
			//	        "type": "object"
			//	      }
			//	    },
			//	    "required": [
			//	      "Name",
			//	      "CheckExpression"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "minItems": 1,
			//	  "type": "array"
			//	}
			Description: "List of the data quality rules in the ruleset",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"check_expression": {
						// Property: CheckExpression
						Description: "Expression with rule conditions",
						Type:        types.StringType,
						Computed:    true,
					},
					"column_selectors": {
						// Property: ColumnSelectors
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"name": {
									// Property: Name
									Description: "The name of a column from a dataset",
									Type:        types.StringType,
									Computed:    true,
								},
								"regex": {
									// Property: Regex
									Description: "A regular expression for selecting a column from a dataset",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"disabled": {
						// Property: Disabled
						Description: "Boolean value to disable/enable a rule",
						Type:        types.BoolType,
						Computed:    true,
					},
					"name": {
						// Property: Name
						Description: "Name of the rule",
						Type:        types.StringType,
						Computed:    true,
					},
					"substitution_map": {
						// Property: SubstitutionMap
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"value": {
									// Property: Value
									Description: "Value or column name",
									Type:        types.StringType,
									Computed:    true,
								},
								"value_reference": {
									// Property: ValueReference
									Description: "Variable name",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"threshold": {
						// Property: Threshold
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"type": {
									// Property: Type
									Description: "Threshold type for a rule",
									Type:        types.StringType,
									Computed:    true,
								},
								"unit": {
									// Property: Unit
									Description: "Threshold unit for a rule",
									Type:        types.StringType,
									Computed:    true,
								},
								"value": {
									// Property: Value
									Description: "Threshold value for a rule",
									Type:        types.Float64Type,
									Computed:    true,
								},
							},
						),
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
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "A key-value pair to associate with a resource",
			//	    "properties": {
			//	      "Key": {
			//	        "maxLength": 128,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "maxLength": 256,
			//	        "minLength": 0,
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Value",
			//	      "Key"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": false
			//	}
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
		"target_arn": {
			// Property: TargetArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Arn of the target resource (dataset) to apply the ruleset to",
			//	  "maxLength": 2048,
			//	  "minLength": 20,
			//	  "type": "string"
			//	}
			Description: "Arn of the target resource (dataset) to apply the ruleset to",
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
		Description: "Data Source schema for AWS::DataBrew::Ruleset",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataBrew::Ruleset").WithTerraformTypeName("awscc_databrew_ruleset")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"check_expression": "CheckExpression",
		"column_selectors": "ColumnSelectors",
		"description":      "Description",
		"disabled":         "Disabled",
		"key":              "Key",
		"name":             "Name",
		"regex":            "Regex",
		"rules":            "Rules",
		"substitution_map": "SubstitutionMap",
		"tags":             "Tags",
		"target_arn":       "TargetArn",
		"threshold":        "Threshold",
		"type":             "Type",
		"unit":             "Unit",
		"value":            "Value",
		"value_reference":  "ValueReference",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
