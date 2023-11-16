// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package s3

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_s3_storage_lens_group", storageLensGroupDataSource)
}

// storageLensGroupDataSource returns the Terraform awscc_s3_storage_lens_group data source.
// This Terraform data source corresponds to the CloudFormation AWS::S3::StorageLensGroup resource.
func storageLensGroupDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Filter
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Sets the Storage Lens Group filter.",
		//	  "properties": {
		//	    "And": {
		//	      "additionalProperties": false,
		//	      "description": "The Storage Lens group will include objects that match all of the specified filter values.",
		//	      "properties": {
		//	        "MatchAnyPrefix": {
		//	          "description": "Filter to match any of the specified prefixes.",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "maxLength": 1024,
		//	            "type": "string"
		//	          },
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        },
		//	        "MatchAnySuffix": {
		//	          "description": "Filter to match any of the specified suffixes.",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "maxLength": 1024,
		//	            "type": "string"
		//	          },
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        },
		//	        "MatchAnyTag": {
		//	          "description": "Filter to match any of the specified object tags.",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "Key": {
		//	                "maxLength": 128,
		//	                "minLength": 1,
		//	                "type": "string"
		//	              },
		//	              "Value": {
		//	                "maxLength": 256,
		//	                "minLength": 0,
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "Key",
		//	              "Value"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        },
		//	        "MatchObjectAge": {
		//	          "additionalProperties": false,
		//	          "description": "Filter to match all of the specified values for the minimum and maximum object age.",
		//	          "properties": {
		//	            "DaysGreaterThan": {
		//	              "description": "Minimum object age to which the rule applies.",
		//	              "minimum": 1,
		//	              "type": "integer"
		//	            },
		//	            "DaysLessThan": {
		//	              "description": "Maximum object age to which the rule applies.",
		//	              "minimum": 1,
		//	              "type": "integer"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "MatchObjectSize": {
		//	          "additionalProperties": false,
		//	          "description": "Filter to match all of the specified values for the minimum and maximum object size.",
		//	          "properties": {
		//	            "BytesGreaterThan": {
		//	              "description": "Minimum object size to which the rule applies.",
		//	              "format": "int64",
		//	              "minimum": 1,
		//	              "type": "integer"
		//	            },
		//	            "BytesLessThan": {
		//	              "description": "Maximum object size to which the rule applies.",
		//	              "format": "int64",
		//	              "minimum": 1,
		//	              "type": "integer"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object",
		//	      "uniqueItems": true
		//	    },
		//	    "MatchAnyPrefix": {
		//	      "description": "Filter to match any of the specified prefixes.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "maxLength": 1024,
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "MatchAnySuffix": {
		//	      "description": "Filter to match any of the specified suffixes.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "maxLength": 1024,
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "MatchAnyTag": {
		//	      "description": "Filter to match any of the specified object tags.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Key": {
		//	            "maxLength": 128,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "Value": {
		//	            "maxLength": 256,
		//	            "minLength": 0,
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Key",
		//	          "Value"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "MatchObjectAge": {
		//	      "additionalProperties": false,
		//	      "description": "Filter to match all of the specified values for the minimum and maximum object age.",
		//	      "properties": {
		//	        "DaysGreaterThan": {
		//	          "description": "Minimum object age to which the rule applies.",
		//	          "minimum": 1,
		//	          "type": "integer"
		//	        },
		//	        "DaysLessThan": {
		//	          "description": "Maximum object age to which the rule applies.",
		//	          "minimum": 1,
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "MatchObjectSize": {
		//	      "additionalProperties": false,
		//	      "description": "Filter to match all of the specified values for the minimum and maximum object size.",
		//	      "properties": {
		//	        "BytesGreaterThan": {
		//	          "description": "Minimum object size to which the rule applies.",
		//	          "format": "int64",
		//	          "minimum": 1,
		//	          "type": "integer"
		//	        },
		//	        "BytesLessThan": {
		//	          "description": "Maximum object size to which the rule applies.",
		//	          "format": "int64",
		//	          "minimum": 1,
		//	          "type": "integer"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "Or": {
		//	      "additionalProperties": false,
		//	      "description": "The Storage Lens group will include objects that match any of the specified filter values.",
		//	      "properties": {
		//	        "MatchAnyPrefix": {
		//	          "description": "Filter to match any of the specified prefixes.",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "maxLength": 1024,
		//	            "type": "string"
		//	          },
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        },
		//	        "MatchAnySuffix": {
		//	          "description": "Filter to match any of the specified suffixes.",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "maxLength": 1024,
		//	            "type": "string"
		//	          },
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        },
		//	        "MatchAnyTag": {
		//	          "description": "Filter to match any of the specified object tags.",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "Key": {
		//	                "maxLength": 128,
		//	                "minLength": 1,
		//	                "type": "string"
		//	              },
		//	              "Value": {
		//	                "maxLength": 256,
		//	                "minLength": 0,
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "Key",
		//	              "Value"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        },
		//	        "MatchObjectAge": {
		//	          "additionalProperties": false,
		//	          "description": "Filter to match all of the specified values for the minimum and maximum object age.",
		//	          "properties": {
		//	            "DaysGreaterThan": {
		//	              "description": "Minimum object age to which the rule applies.",
		//	              "minimum": 1,
		//	              "type": "integer"
		//	            },
		//	            "DaysLessThan": {
		//	              "description": "Maximum object age to which the rule applies.",
		//	              "minimum": 1,
		//	              "type": "integer"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "MatchObjectSize": {
		//	          "additionalProperties": false,
		//	          "description": "Filter to match all of the specified values for the minimum and maximum object size.",
		//	          "properties": {
		//	            "BytesGreaterThan": {
		//	              "description": "Minimum object size to which the rule applies.",
		//	              "format": "int64",
		//	              "minimum": 1,
		//	              "type": "integer"
		//	            },
		//	            "BytesLessThan": {
		//	              "description": "Maximum object size to which the rule applies.",
		//	              "format": "int64",
		//	              "minimum": 1,
		//	              "type": "integer"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object",
		//	      "uniqueItems": true
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"filter": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: And
				"and": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: MatchAnyPrefix
						"match_any_prefix": schema.SetAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "Filter to match any of the specified prefixes.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: MatchAnySuffix
						"match_any_suffix": schema.SetAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "Filter to match any of the specified suffixes.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: MatchAnyTag
						"match_any_tag": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
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
							Description: "Filter to match any of the specified object tags.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: MatchObjectAge
						"match_object_age": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: DaysGreaterThan
								"days_greater_than": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Description: "Minimum object age to which the rule applies.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: DaysLessThan
								"days_less_than": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Description: "Maximum object age to which the rule applies.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Filter to match all of the specified values for the minimum and maximum object age.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: MatchObjectSize
						"match_object_size": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: BytesGreaterThan
								"bytes_greater_than": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Description: "Minimum object size to which the rule applies.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: BytesLessThan
								"bytes_less_than": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Description: "Maximum object size to which the rule applies.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Filter to match all of the specified values for the minimum and maximum object size.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The Storage Lens group will include objects that match all of the specified filter values.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: MatchAnyPrefix
				"match_any_prefix": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Filter to match any of the specified prefixes.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: MatchAnySuffix
				"match_any_suffix": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Filter to match any of the specified suffixes.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: MatchAnyTag
				"match_any_tag": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
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
					Description: "Filter to match any of the specified object tags.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: MatchObjectAge
				"match_object_age": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: DaysGreaterThan
						"days_greater_than": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "Minimum object age to which the rule applies.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: DaysLessThan
						"days_less_than": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "Maximum object age to which the rule applies.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Filter to match all of the specified values for the minimum and maximum object age.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: MatchObjectSize
				"match_object_size": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: BytesGreaterThan
						"bytes_greater_than": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "Minimum object size to which the rule applies.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: BytesLessThan
						"bytes_less_than": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "Maximum object size to which the rule applies.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Filter to match all of the specified values for the minimum and maximum object size.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Or
				"or": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: MatchAnyPrefix
						"match_any_prefix": schema.SetAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "Filter to match any of the specified prefixes.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: MatchAnySuffix
						"match_any_suffix": schema.SetAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "Filter to match any of the specified suffixes.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: MatchAnyTag
						"match_any_tag": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
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
							Description: "Filter to match any of the specified object tags.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: MatchObjectAge
						"match_object_age": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: DaysGreaterThan
								"days_greater_than": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Description: "Minimum object age to which the rule applies.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: DaysLessThan
								"days_less_than": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Description: "Maximum object age to which the rule applies.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Filter to match all of the specified values for the minimum and maximum object age.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: MatchObjectSize
						"match_object_size": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: BytesGreaterThan
								"bytes_greater_than": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Description: "Minimum object size to which the rule applies.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: BytesLessThan
								"bytes_less_than": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Description: "Maximum object size to which the rule applies.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Filter to match all of the specified values for the minimum and maximum object size.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The Storage Lens group will include objects that match any of the specified filter values.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Sets the Storage Lens Group filter.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name that identifies the Amazon S3 Storage Lens Group.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9\\-_]+$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name that identifies the Amazon S3 Storage Lens Group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StorageLensGroupArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN for the Amazon S3 Storage Lens Group.",
		//	  "type": "string"
		//	}
		"storage_lens_group_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN for the Amazon S3 Storage Lens Group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A set of tags (key-value pairs) for this Amazon S3 Storage Lens Group.",
		//	  "insertionOrder": true,
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
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
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
			Description: "A set of tags (key-value pairs) for this Amazon S3 Storage Lens Group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::S3::StorageLensGroup",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3::StorageLensGroup").WithTerraformTypeName("awscc_s3_storage_lens_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"and":                    "And",
		"bytes_greater_than":     "BytesGreaterThan",
		"bytes_less_than":        "BytesLessThan",
		"days_greater_than":      "DaysGreaterThan",
		"days_less_than":         "DaysLessThan",
		"filter":                 "Filter",
		"key":                    "Key",
		"match_any_prefix":       "MatchAnyPrefix",
		"match_any_suffix":       "MatchAnySuffix",
		"match_any_tag":          "MatchAnyTag",
		"match_object_age":       "MatchObjectAge",
		"match_object_size":      "MatchObjectSize",
		"name":                   "Name",
		"or":                     "Or",
		"storage_lens_group_arn": "StorageLensGroupArn",
		"tags":                   "Tags",
		"value":                  "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
