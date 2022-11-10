// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lakeformation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_lakeformation_principal_permissions", principalPermissionsDataSource)
}

// principalPermissionsDataSource returns the Terraform awscc_lakeformation_principal_permissions data source.
// This Terraform data source corresponds to the CloudFormation AWS::LakeFormation::PrincipalPermissions resource.
func principalPermissionsDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"catalog": {
			// Property: Catalog
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 12,
			//	  "minLength": 12,
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"permissions": {
			// Property: Permissions
			// CloudFormation resource type schema:
			//
			//	{
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "enum": [
			//	      "ALL",
			//	      "SELECT",
			//	      "ALTER",
			//	      "DROP",
			//	      "DELETE",
			//	      "INSERT",
			//	      "DESCRIBE",
			//	      "CREATE_DATABASE",
			//	      "CREATE_TABLE",
			//	      "DATA_LOCATION_ACCESS",
			//	      "CREATE_TAG",
			//	      "ASSOCIATE"
			//	    ],
			//	    "type": "string"
			//	  },
			//	  "type": "array"
			//	}
			Type:     types.ListType{ElemType: types.StringType},
			Computed: true,
		},
		"permissions_with_grant_option": {
			// Property: PermissionsWithGrantOption
			// CloudFormation resource type schema:
			//
			//	{
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "enum": [
			//	      "ALL",
			//	      "SELECT",
			//	      "ALTER",
			//	      "DROP",
			//	      "DELETE",
			//	      "INSERT",
			//	      "DESCRIBE",
			//	      "CREATE_DATABASE",
			//	      "CREATE_TABLE",
			//	      "DATA_LOCATION_ACCESS",
			//	      "CREATE_TAG",
			//	      "ASSOCIATE"
			//	    ],
			//	    "type": "string"
			//	  },
			//	  "type": "array"
			//	}
			Type:     types.ListType{ElemType: types.StringType},
			Computed: true,
		},
		"principal": {
			// Property: Principal
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "properties": {
			//	    "DataLakePrincipalIdentifier": {
			//	      "maxLength": 255,
			//	      "minLength": 1,
			//	      "type": "string"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"data_lake_principal_identifier": {
						// Property: DataLakePrincipalIdentifier
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"principal_identifier": {
			// Property: PrincipalIdentifier
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"resource": {
			// Property: Resource
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "properties": {
			//	    "Catalog": {
			//	      "additionalProperties": false,
			//	      "type": "object"
			//	    },
			//	    "DataCellsFilter": {
			//	      "additionalProperties": false,
			//	      "properties": {
			//	        "DatabaseName": {
			//	          "maxLength": 255,
			//	          "minLength": 1,
			//	          "type": "string"
			//	        },
			//	        "Name": {
			//	          "maxLength": 255,
			//	          "minLength": 1,
			//	          "type": "string"
			//	        },
			//	        "TableCatalogId": {
			//	          "maxLength": 12,
			//	          "minLength": 12,
			//	          "type": "string"
			//	        },
			//	        "TableName": {
			//	          "maxLength": 255,
			//	          "minLength": 1,
			//	          "type": "string"
			//	        }
			//	      },
			//	      "required": [
			//	        "TableCatalogId",
			//	        "DatabaseName",
			//	        "TableName",
			//	        "Name"
			//	      ],
			//	      "type": "object"
			//	    },
			//	    "DataLocation": {
			//	      "additionalProperties": false,
			//	      "properties": {
			//	        "CatalogId": {
			//	          "maxLength": 12,
			//	          "minLength": 12,
			//	          "type": "string"
			//	        },
			//	        "ResourceArn": {
			//	          "type": "string"
			//	        }
			//	      },
			//	      "required": [
			//	        "CatalogId",
			//	        "ResourceArn"
			//	      ],
			//	      "type": "object"
			//	    },
			//	    "Database": {
			//	      "additionalProperties": false,
			//	      "properties": {
			//	        "CatalogId": {
			//	          "maxLength": 12,
			//	          "minLength": 12,
			//	          "type": "string"
			//	        },
			//	        "Name": {
			//	          "maxLength": 255,
			//	          "minLength": 1,
			//	          "type": "string"
			//	        }
			//	      },
			//	      "required": [
			//	        "CatalogId",
			//	        "Name"
			//	      ],
			//	      "type": "object"
			//	    },
			//	    "LFTag": {
			//	      "additionalProperties": false,
			//	      "properties": {
			//	        "CatalogId": {
			//	          "maxLength": 12,
			//	          "minLength": 12,
			//	          "type": "string"
			//	        },
			//	        "TagKey": {
			//	          "maxLength": 255,
			//	          "minLength": 1,
			//	          "type": "string"
			//	        },
			//	        "TagValues": {
			//	          "insertionOrder": false,
			//	          "items": {
			//	            "maxLength": 256,
			//	            "minLength": 0,
			//	            "type": "string"
			//	          },
			//	          "maxItems": 50,
			//	          "minItems": 1,
			//	          "type": "array"
			//	        }
			//	      },
			//	      "required": [
			//	        "CatalogId",
			//	        "TagKey",
			//	        "TagValues"
			//	      ],
			//	      "type": "object"
			//	    },
			//	    "LFTagPolicy": {
			//	      "additionalProperties": false,
			//	      "properties": {
			//	        "CatalogId": {
			//	          "maxLength": 12,
			//	          "minLength": 12,
			//	          "type": "string"
			//	        },
			//	        "Expression": {
			//	          "insertionOrder": false,
			//	          "items": {
			//	            "additionalProperties": false,
			//	            "properties": {
			//	              "TagKey": {
			//	                "maxLength": 128,
			//	                "minLength": 1,
			//	                "type": "string"
			//	              },
			//	              "TagValues": {
			//	                "insertionOrder": false,
			//	                "items": {
			//	                  "maxLength": 256,
			//	                  "minLength": 0,
			//	                  "type": "string"
			//	                },
			//	                "maxItems": 50,
			//	                "minItems": 1,
			//	                "type": "array"
			//	              }
			//	            },
			//	            "type": "object"
			//	          },
			//	          "maxItems": 5,
			//	          "minItems": 1,
			//	          "type": "array"
			//	        },
			//	        "ResourceType": {
			//	          "enum": [
			//	            "DATABASE",
			//	            "TABLE"
			//	          ],
			//	          "type": "string"
			//	        }
			//	      },
			//	      "required": [
			//	        "CatalogId",
			//	        "ResourceType",
			//	        "Expression"
			//	      ],
			//	      "type": "object"
			//	    },
			//	    "Table": {
			//	      "additionalProperties": false,
			//	      "properties": {
			//	        "CatalogId": {
			//	          "maxLength": 12,
			//	          "minLength": 12,
			//	          "type": "string"
			//	        },
			//	        "DatabaseName": {
			//	          "maxLength": 255,
			//	          "minLength": 1,
			//	          "type": "string"
			//	        },
			//	        "Name": {
			//	          "maxLength": 255,
			//	          "minLength": 1,
			//	          "type": "string"
			//	        },
			//	        "TableWildcard": {
			//	          "additionalProperties": false,
			//	          "type": "object"
			//	        }
			//	      },
			//	      "required": [
			//	        "CatalogId",
			//	        "DatabaseName"
			//	      ],
			//	      "type": "object"
			//	    },
			//	    "TableWithColumns": {
			//	      "additionalProperties": false,
			//	      "properties": {
			//	        "CatalogId": {
			//	          "maxLength": 12,
			//	          "minLength": 12,
			//	          "type": "string"
			//	        },
			//	        "ColumnNames": {
			//	          "insertionOrder": false,
			//	          "items": {
			//	            "maxLength": 255,
			//	            "minLength": 1,
			//	            "type": "string"
			//	          },
			//	          "type": "array"
			//	        },
			//	        "ColumnWildcard": {
			//	          "additionalProperties": false,
			//	          "properties": {
			//	            "ExcludedColumnNames": {
			//	              "insertionOrder": false,
			//	              "items": {
			//	                "maxLength": 255,
			//	                "minLength": 1,
			//	                "type": "string"
			//	              },
			//	              "type": "array"
			//	            }
			//	          },
			//	          "type": "object"
			//	        },
			//	        "DatabaseName": {
			//	          "maxLength": 255,
			//	          "minLength": 1,
			//	          "type": "string"
			//	        },
			//	        "Name": {
			//	          "maxLength": 255,
			//	          "minLength": 1,
			//	          "type": "string"
			//	        }
			//	      },
			//	      "required": [
			//	        "CatalogId",
			//	        "DatabaseName",
			//	        "Name"
			//	      ],
			//	      "type": "object"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"catalog": {
						// Property: Catalog
						Type:     types.MapType{ElemType: types.StringType},
						Computed: true,
					},
					"data_cells_filter": {
						// Property: DataCellsFilter
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"database_name": {
									// Property: DatabaseName
									Type:     types.StringType,
									Computed: true,
								},
								"name": {
									// Property: Name
									Type:     types.StringType,
									Computed: true,
								},
								"table_catalog_id": {
									// Property: TableCatalogId
									Type:     types.StringType,
									Computed: true,
								},
								"table_name": {
									// Property: TableName
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"data_location": {
						// Property: DataLocation
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"catalog_id": {
									// Property: CatalogId
									Type:     types.StringType,
									Computed: true,
								},
								"resource_arn": {
									// Property: ResourceArn
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"database": {
						// Property: Database
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"catalog_id": {
									// Property: CatalogId
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
					"lf_tag": {
						// Property: LFTag
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"catalog_id": {
									// Property: CatalogId
									Type:     types.StringType,
									Computed: true,
								},
								"tag_key": {
									// Property: TagKey
									Type:     types.StringType,
									Computed: true,
								},
								"tag_values": {
									// Property: TagValues
									Type:     types.ListType{ElemType: types.StringType},
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"lf_tag_policy": {
						// Property: LFTagPolicy
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"catalog_id": {
									// Property: CatalogId
									Type:     types.StringType,
									Computed: true,
								},
								"expression": {
									// Property: Expression
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"tag_key": {
												// Property: TagKey
												Type:     types.StringType,
												Computed: true,
											},
											"tag_values": {
												// Property: TagValues
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"resource_type": {
									// Property: ResourceType
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"table": {
						// Property: Table
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"catalog_id": {
									// Property: CatalogId
									Type:     types.StringType,
									Computed: true,
								},
								"database_name": {
									// Property: DatabaseName
									Type:     types.StringType,
									Computed: true,
								},
								"name": {
									// Property: Name
									Type:     types.StringType,
									Computed: true,
								},
								"table_wildcard": {
									// Property: TableWildcard
									Type:     types.MapType{ElemType: types.StringType},
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"table_with_columns": {
						// Property: TableWithColumns
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"catalog_id": {
									// Property: CatalogId
									Type:     types.StringType,
									Computed: true,
								},
								"column_names": {
									// Property: ColumnNames
									Type:     types.ListType{ElemType: types.StringType},
									Computed: true,
								},
								"column_wildcard": {
									// Property: ColumnWildcard
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"excluded_column_names": {
												// Property: ExcludedColumnNames
												Type:     types.ListType{ElemType: types.StringType},
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"database_name": {
									// Property: DatabaseName
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
				},
			),
			Computed: true,
		},
		"resource_identifier": {
			// Property: ResourceIdentifier
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
		Description: "Data Source schema for AWS::LakeFormation::PrincipalPermissions",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::LakeFormation::PrincipalPermissions").WithTerraformTypeName("awscc_lakeformation_principal_permissions")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"catalog":                        "Catalog",
		"catalog_id":                     "CatalogId",
		"column_names":                   "ColumnNames",
		"column_wildcard":                "ColumnWildcard",
		"data_cells_filter":              "DataCellsFilter",
		"data_lake_principal_identifier": "DataLakePrincipalIdentifier",
		"data_location":                  "DataLocation",
		"database":                       "Database",
		"database_name":                  "DatabaseName",
		"excluded_column_names":          "ExcludedColumnNames",
		"expression":                     "Expression",
		"lf_tag":                         "LFTag",
		"lf_tag_policy":                  "LFTagPolicy",
		"name":                           "Name",
		"permissions":                    "Permissions",
		"permissions_with_grant_option":  "PermissionsWithGrantOption",
		"principal":                      "Principal",
		"principal_identifier":           "PrincipalIdentifier",
		"resource":                       "Resource",
		"resource_arn":                   "ResourceArn",
		"resource_identifier":            "ResourceIdentifier",
		"resource_type":                  "ResourceType",
		"table":                          "Table",
		"table_catalog_id":               "TableCatalogId",
		"table_name":                     "TableName",
		"table_wildcard":                 "TableWildcard",
		"table_with_columns":             "TableWithColumns",
		"tag_key":                        "TagKey",
		"tag_values":                     "TagValues",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
