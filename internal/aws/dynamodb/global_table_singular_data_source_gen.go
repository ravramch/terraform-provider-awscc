// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package dynamodb

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_dynamodb_global_table", globalTableDataSource)
}

// globalTableDataSource returns the Terraform awscc_dynamodb_global_table data source.
// This Terraform data source corresponds to the CloudFormation AWS::DynamoDB::GlobalTable resource.
func globalTableDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"attribute_definitions": {
			// Property: AttributeDefinitions
			// CloudFormation resource type schema:
			//
			//	{
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "AttributeName": {
			//	        "maxLength": 255,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      },
			//	      "AttributeType": {
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "AttributeName",
			//	      "AttributeType"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "minItems": 1,
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"attribute_name": {
						// Property: AttributeName
						Type:     types.StringType,
						Computed: true,
					},
					"attribute_type": {
						// Property: AttributeType
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"billing_mode": {
			// Property: BillingMode
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"global_secondary_indexes": {
			// Property: GlobalSecondaryIndexes
			// CloudFormation resource type schema:
			//
			//	{
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "IndexName": {
			//	        "maxLength": 255,
			//	        "minLength": 3,
			//	        "type": "string"
			//	      },
			//	      "KeySchema": {
			//	        "items": {
			//	          "additionalProperties": false,
			//	          "properties": {
			//	            "AttributeName": {
			//	              "maxLength": 255,
			//	              "minLength": 1,
			//	              "type": "string"
			//	            },
			//	            "KeyType": {
			//	              "type": "string"
			//	            }
			//	          },
			//	          "required": [
			//	            "KeyType",
			//	            "AttributeName"
			//	          ],
			//	          "type": "object"
			//	        },
			//	        "maxItems": 2,
			//	        "minItems": 1,
			//	        "type": "array",
			//	        "uniqueItems": true
			//	      },
			//	      "Projection": {
			//	        "additionalProperties": false,
			//	        "properties": {
			//	          "NonKeyAttributes": {
			//	            "insertionOrder": false,
			//	            "items": {
			//	              "type": "string"
			//	            },
			//	            "maxItems": 20,
			//	            "type": "array",
			//	            "uniqueItems": true
			//	          },
			//	          "ProjectionType": {
			//	            "type": "string"
			//	          }
			//	        },
			//	        "type": "object"
			//	      },
			//	      "WriteProvisionedThroughputSettings": {
			//	        "additionalProperties": false,
			//	        "properties": {
			//	          "WriteCapacityAutoScalingSettings": {
			//	            "additionalProperties": false,
			//	            "properties": {
			//	              "MaxCapacity": {
			//	                "minimum": 1,
			//	                "type": "integer"
			//	              },
			//	              "MinCapacity": {
			//	                "minimum": 1,
			//	                "type": "integer"
			//	              },
			//	              "SeedCapacity": {
			//	                "minimum": 1,
			//	                "type": "integer"
			//	              },
			//	              "TargetTrackingScalingPolicyConfiguration": {
			//	                "additionalProperties": false,
			//	                "properties": {
			//	                  "DisableScaleIn": {
			//	                    "type": "boolean"
			//	                  },
			//	                  "ScaleInCooldown": {
			//	                    "minimum": 0,
			//	                    "type": "integer"
			//	                  },
			//	                  "ScaleOutCooldown": {
			//	                    "minimum": 0,
			//	                    "type": "integer"
			//	                  },
			//	                  "TargetValue": {
			//	                    "format": "double",
			//	                    "type": "number"
			//	                  }
			//	                },
			//	                "required": [
			//	                  "TargetValue"
			//	                ],
			//	                "type": "object"
			//	              }
			//	            },
			//	            "required": [
			//	              "MinCapacity",
			//	              "MaxCapacity",
			//	              "TargetTrackingScalingPolicyConfiguration"
			//	            ],
			//	            "type": "object"
			//	          }
			//	        },
			//	        "type": "object"
			//	      }
			//	    },
			//	    "required": [
			//	      "IndexName",
			//	      "Projection",
			//	      "KeySchema"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"index_name": {
						// Property: IndexName
						Type:     types.StringType,
						Computed: true,
					},
					"key_schema": {
						// Property: KeySchema
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"attribute_name": {
									// Property: AttributeName
									Type:     types.StringType,
									Computed: true,
								},
								"key_type": {
									// Property: KeyType
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"projection": {
						// Property: Projection
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"non_key_attributes": {
									// Property: NonKeyAttributes
									Type:     types.SetType{ElemType: types.StringType},
									Computed: true,
								},
								"projection_type": {
									// Property: ProjectionType
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"write_provisioned_throughput_settings": {
						// Property: WriteProvisionedThroughputSettings
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"write_capacity_auto_scaling_settings": {
									// Property: WriteCapacityAutoScalingSettings
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"max_capacity": {
												// Property: MaxCapacity
												Type:     types.Int64Type,
												Computed: true,
											},
											"min_capacity": {
												// Property: MinCapacity
												Type:     types.Int64Type,
												Computed: true,
											},
											"seed_capacity": {
												// Property: SeedCapacity
												Type:     types.Int64Type,
												Computed: true,
											},
											"target_tracking_scaling_policy_configuration": {
												// Property: TargetTrackingScalingPolicyConfiguration
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"disable_scale_in": {
															// Property: DisableScaleIn
															Type:     types.BoolType,
															Computed: true,
														},
														"scale_in_cooldown": {
															// Property: ScaleInCooldown
															Type:     types.Int64Type,
															Computed: true,
														},
														"scale_out_cooldown": {
															// Property: ScaleOutCooldown
															Type:     types.Int64Type,
															Computed: true,
														},
														"target_value": {
															// Property: TargetValue
															Type:     types.Float64Type,
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
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"key_schema": {
			// Property: KeySchema
			// CloudFormation resource type schema:
			//
			//	{
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "AttributeName": {
			//	        "maxLength": 255,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      },
			//	      "KeyType": {
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "KeyType",
			//	      "AttributeName"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "maxItems": 2,
			//	  "minItems": 1,
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"attribute_name": {
						// Property: AttributeName
						Type:     types.StringType,
						Computed: true,
					},
					"key_type": {
						// Property: KeyType
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"local_secondary_indexes": {
			// Property: LocalSecondaryIndexes
			// CloudFormation resource type schema:
			//
			//	{
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "IndexName": {
			//	        "maxLength": 255,
			//	        "minLength": 3,
			//	        "type": "string"
			//	      },
			//	      "KeySchema": {
			//	        "items": {
			//	          "additionalProperties": false,
			//	          "properties": {
			//	            "AttributeName": {
			//	              "maxLength": 255,
			//	              "minLength": 1,
			//	              "type": "string"
			//	            },
			//	            "KeyType": {
			//	              "type": "string"
			//	            }
			//	          },
			//	          "required": [
			//	            "KeyType",
			//	            "AttributeName"
			//	          ],
			//	          "type": "object"
			//	        },
			//	        "maxItems": 2,
			//	        "type": "array",
			//	        "uniqueItems": true
			//	      },
			//	      "Projection": {
			//	        "additionalProperties": false,
			//	        "properties": {
			//	          "NonKeyAttributes": {
			//	            "insertionOrder": false,
			//	            "items": {
			//	              "type": "string"
			//	            },
			//	            "maxItems": 20,
			//	            "type": "array",
			//	            "uniqueItems": true
			//	          },
			//	          "ProjectionType": {
			//	            "type": "string"
			//	          }
			//	        },
			//	        "type": "object"
			//	      }
			//	    },
			//	    "required": [
			//	      "IndexName",
			//	      "Projection",
			//	      "KeySchema"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"index_name": {
						// Property: IndexName
						Type:     types.StringType,
						Computed: true,
					},
					"key_schema": {
						// Property: KeySchema
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"attribute_name": {
									// Property: AttributeName
									Type:     types.StringType,
									Computed: true,
								},
								"key_type": {
									// Property: KeyType
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"projection": {
						// Property: Projection
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"non_key_attributes": {
									// Property: NonKeyAttributes
									Type:     types.SetType{ElemType: types.StringType},
									Computed: true,
								},
								"projection_type": {
									// Property: ProjectionType
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
		"replicas": {
			// Property: Replicas
			// CloudFormation resource type schema:
			//
			//	{
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "ContributorInsightsSpecification": {
			//	        "additionalProperties": false,
			//	        "properties": {
			//	          "Enabled": {
			//	            "type": "boolean"
			//	          }
			//	        },
			//	        "required": [
			//	          "Enabled"
			//	        ],
			//	        "type": "object"
			//	      },
			//	      "GlobalSecondaryIndexes": {
			//	        "insertionOrder": false,
			//	        "items": {
			//	          "additionalProperties": false,
			//	          "properties": {
			//	            "ContributorInsightsSpecification": {
			//	              "additionalProperties": false,
			//	              "properties": {
			//	                "Enabled": {
			//	                  "type": "boolean"
			//	                }
			//	              },
			//	              "required": [
			//	                "Enabled"
			//	              ],
			//	              "type": "object"
			//	            },
			//	            "IndexName": {
			//	              "maxLength": 255,
			//	              "minLength": 3,
			//	              "type": "string"
			//	            },
			//	            "ReadProvisionedThroughputSettings": {
			//	              "additionalProperties": false,
			//	              "properties": {
			//	                "ReadCapacityAutoScalingSettings": {
			//	                  "additionalProperties": false,
			//	                  "properties": {
			//	                    "MaxCapacity": {
			//	                      "minimum": 1,
			//	                      "type": "integer"
			//	                    },
			//	                    "MinCapacity": {
			//	                      "minimum": 1,
			//	                      "type": "integer"
			//	                    },
			//	                    "SeedCapacity": {
			//	                      "minimum": 1,
			//	                      "type": "integer"
			//	                    },
			//	                    "TargetTrackingScalingPolicyConfiguration": {
			//	                      "additionalProperties": false,
			//	                      "properties": {
			//	                        "DisableScaleIn": {
			//	                          "type": "boolean"
			//	                        },
			//	                        "ScaleInCooldown": {
			//	                          "minimum": 0,
			//	                          "type": "integer"
			//	                        },
			//	                        "ScaleOutCooldown": {
			//	                          "minimum": 0,
			//	                          "type": "integer"
			//	                        },
			//	                        "TargetValue": {
			//	                          "format": "double",
			//	                          "type": "number"
			//	                        }
			//	                      },
			//	                      "required": [
			//	                        "TargetValue"
			//	                      ],
			//	                      "type": "object"
			//	                    }
			//	                  },
			//	                  "required": [
			//	                    "MinCapacity",
			//	                    "MaxCapacity",
			//	                    "TargetTrackingScalingPolicyConfiguration"
			//	                  ],
			//	                  "type": "object"
			//	                },
			//	                "ReadCapacityUnits": {
			//	                  "minimum": 1,
			//	                  "type": "integer"
			//	                }
			//	              },
			//	              "type": "object"
			//	            }
			//	          },
			//	          "required": [
			//	            "IndexName"
			//	          ],
			//	          "type": "object"
			//	        },
			//	        "type": "array",
			//	        "uniqueItems": true
			//	      },
			//	      "PointInTimeRecoverySpecification": {
			//	        "additionalProperties": false,
			//	        "properties": {
			//	          "PointInTimeRecoveryEnabled": {
			//	            "type": "boolean"
			//	          }
			//	        },
			//	        "type": "object"
			//	      },
			//	      "ReadProvisionedThroughputSettings": {
			//	        "additionalProperties": false,
			//	        "properties": {
			//	          "ReadCapacityAutoScalingSettings": {
			//	            "additionalProperties": false,
			//	            "properties": {
			//	              "MaxCapacity": {
			//	                "minimum": 1,
			//	                "type": "integer"
			//	              },
			//	              "MinCapacity": {
			//	                "minimum": 1,
			//	                "type": "integer"
			//	              },
			//	              "SeedCapacity": {
			//	                "minimum": 1,
			//	                "type": "integer"
			//	              },
			//	              "TargetTrackingScalingPolicyConfiguration": {
			//	                "additionalProperties": false,
			//	                "properties": {
			//	                  "DisableScaleIn": {
			//	                    "type": "boolean"
			//	                  },
			//	                  "ScaleInCooldown": {
			//	                    "minimum": 0,
			//	                    "type": "integer"
			//	                  },
			//	                  "ScaleOutCooldown": {
			//	                    "minimum": 0,
			//	                    "type": "integer"
			//	                  },
			//	                  "TargetValue": {
			//	                    "format": "double",
			//	                    "type": "number"
			//	                  }
			//	                },
			//	                "required": [
			//	                  "TargetValue"
			//	                ],
			//	                "type": "object"
			//	              }
			//	            },
			//	            "required": [
			//	              "MinCapacity",
			//	              "MaxCapacity",
			//	              "TargetTrackingScalingPolicyConfiguration"
			//	            ],
			//	            "type": "object"
			//	          },
			//	          "ReadCapacityUnits": {
			//	            "minimum": 1,
			//	            "type": "integer"
			//	          }
			//	        },
			//	        "type": "object"
			//	      },
			//	      "Region": {
			//	        "type": "string"
			//	      },
			//	      "SSESpecification": {
			//	        "additionalProperties": false,
			//	        "properties": {
			//	          "KMSMasterKeyId": {
			//	            "type": "string"
			//	          }
			//	        },
			//	        "required": [
			//	          "KMSMasterKeyId"
			//	        ],
			//	        "type": "object"
			//	      },
			//	      "TableClass": {
			//	        "type": "string"
			//	      },
			//	      "Tags": {
			//	        "insertionOrder": false,
			//	        "items": {
			//	          "additionalProperties": false,
			//	          "properties": {
			//	            "Key": {
			//	              "type": "string"
			//	            },
			//	            "Value": {
			//	              "type": "string"
			//	            }
			//	          },
			//	          "required": [
			//	            "Value",
			//	            "Key"
			//	          ],
			//	          "type": "object"
			//	        },
			//	        "type": "array",
			//	        "uniqueItems": true
			//	      }
			//	    },
			//	    "required": [
			//	      "Region"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "minItems": 1,
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"contributor_insights_specification": {
						// Property: ContributorInsightsSpecification
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"enabled": {
									// Property: Enabled
									Type:     types.BoolType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"global_secondary_indexes": {
						// Property: GlobalSecondaryIndexes
						Attributes: tfsdk.SetNestedAttributes(
							map[string]tfsdk.Attribute{
								"contributor_insights_specification": {
									// Property: ContributorInsightsSpecification
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"enabled": {
												// Property: Enabled
												Type:     types.BoolType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"index_name": {
									// Property: IndexName
									Type:     types.StringType,
									Computed: true,
								},
								"read_provisioned_throughput_settings": {
									// Property: ReadProvisionedThroughputSettings
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"read_capacity_auto_scaling_settings": {
												// Property: ReadCapacityAutoScalingSettings
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"max_capacity": {
															// Property: MaxCapacity
															Type:     types.Int64Type,
															Computed: true,
														},
														"min_capacity": {
															// Property: MinCapacity
															Type:     types.Int64Type,
															Computed: true,
														},
														"seed_capacity": {
															// Property: SeedCapacity
															Type:     types.Int64Type,
															Computed: true,
														},
														"target_tracking_scaling_policy_configuration": {
															// Property: TargetTrackingScalingPolicyConfiguration
															Attributes: tfsdk.SingleNestedAttributes(
																map[string]tfsdk.Attribute{
																	"disable_scale_in": {
																		// Property: DisableScaleIn
																		Type:     types.BoolType,
																		Computed: true,
																	},
																	"scale_in_cooldown": {
																		// Property: ScaleInCooldown
																		Type:     types.Int64Type,
																		Computed: true,
																	},
																	"scale_out_cooldown": {
																		// Property: ScaleOutCooldown
																		Type:     types.Int64Type,
																		Computed: true,
																	},
																	"target_value": {
																		// Property: TargetValue
																		Type:     types.Float64Type,
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
											"read_capacity_units": {
												// Property: ReadCapacityUnits
												Type:     types.Int64Type,
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
					"point_in_time_recovery_specification": {
						// Property: PointInTimeRecoverySpecification
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"point_in_time_recovery_enabled": {
									// Property: PointInTimeRecoveryEnabled
									Type:     types.BoolType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"read_provisioned_throughput_settings": {
						// Property: ReadProvisionedThroughputSettings
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"read_capacity_auto_scaling_settings": {
									// Property: ReadCapacityAutoScalingSettings
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"max_capacity": {
												// Property: MaxCapacity
												Type:     types.Int64Type,
												Computed: true,
											},
											"min_capacity": {
												// Property: MinCapacity
												Type:     types.Int64Type,
												Computed: true,
											},
											"seed_capacity": {
												// Property: SeedCapacity
												Type:     types.Int64Type,
												Computed: true,
											},
											"target_tracking_scaling_policy_configuration": {
												// Property: TargetTrackingScalingPolicyConfiguration
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"disable_scale_in": {
															// Property: DisableScaleIn
															Type:     types.BoolType,
															Computed: true,
														},
														"scale_in_cooldown": {
															// Property: ScaleInCooldown
															Type:     types.Int64Type,
															Computed: true,
														},
														"scale_out_cooldown": {
															// Property: ScaleOutCooldown
															Type:     types.Int64Type,
															Computed: true,
														},
														"target_value": {
															// Property: TargetValue
															Type:     types.Float64Type,
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
								"read_capacity_units": {
									// Property: ReadCapacityUnits
									Type:     types.Int64Type,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"region": {
						// Property: Region
						Type:     types.StringType,
						Computed: true,
					},
					"sse_specification": {
						// Property: SSESpecification
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"kms_master_key_id": {
									// Property: KMSMasterKeyId
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"table_class": {
						// Property: TableClass
						Type:     types.StringType,
						Computed: true,
					},
					"tags": {
						// Property: Tags
						Attributes: tfsdk.SetNestedAttributes(
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
				},
			),
			Computed: true,
		},
		"sse_specification": {
			// Property: SSESpecification
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "properties": {
			//	    "SSEEnabled": {
			//	      "type": "boolean"
			//	    },
			//	    "SSEType": {
			//	      "type": "string"
			//	    }
			//	  },
			//	  "required": [
			//	    "SSEEnabled"
			//	  ],
			//	  "type": "object"
			//	}
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"sse_enabled": {
						// Property: SSEEnabled
						Type:     types.BoolType,
						Computed: true,
					},
					"sse_type": {
						// Property: SSEType
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"stream_arn": {
			// Property: StreamArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"stream_specification": {
			// Property: StreamSpecification
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "properties": {
			//	    "StreamViewType": {
			//	      "type": "string"
			//	    }
			//	  },
			//	  "required": [
			//	    "StreamViewType"
			//	  ],
			//	  "type": "object"
			//	}
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"stream_view_type": {
						// Property: StreamViewType
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"table_id": {
			// Property: TableId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"table_name": {
			// Property: TableName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"time_to_live_specification": {
			// Property: TimeToLiveSpecification
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "properties": {
			//	    "AttributeName": {
			//	      "type": "string"
			//	    },
			//	    "Enabled": {
			//	      "type": "boolean"
			//	    }
			//	  },
			//	  "required": [
			//	    "Enabled"
			//	  ],
			//	  "type": "object"
			//	}
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"attribute_name": {
						// Property: AttributeName
						Type:     types.StringType,
						Computed: true,
					},
					"enabled": {
						// Property: Enabled
						Type:     types.BoolType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"write_provisioned_throughput_settings": {
			// Property: WriteProvisionedThroughputSettings
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "properties": {
			//	    "WriteCapacityAutoScalingSettings": {
			//	      "additionalProperties": false,
			//	      "properties": {
			//	        "MaxCapacity": {
			//	          "minimum": 1,
			//	          "type": "integer"
			//	        },
			//	        "MinCapacity": {
			//	          "minimum": 1,
			//	          "type": "integer"
			//	        },
			//	        "SeedCapacity": {
			//	          "minimum": 1,
			//	          "type": "integer"
			//	        },
			//	        "TargetTrackingScalingPolicyConfiguration": {
			//	          "additionalProperties": false,
			//	          "properties": {
			//	            "DisableScaleIn": {
			//	              "type": "boolean"
			//	            },
			//	            "ScaleInCooldown": {
			//	              "minimum": 0,
			//	              "type": "integer"
			//	            },
			//	            "ScaleOutCooldown": {
			//	              "minimum": 0,
			//	              "type": "integer"
			//	            },
			//	            "TargetValue": {
			//	              "format": "double",
			//	              "type": "number"
			//	            }
			//	          },
			//	          "required": [
			//	            "TargetValue"
			//	          ],
			//	          "type": "object"
			//	        }
			//	      },
			//	      "required": [
			//	        "MinCapacity",
			//	        "MaxCapacity",
			//	        "TargetTrackingScalingPolicyConfiguration"
			//	      ],
			//	      "type": "object"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"write_capacity_auto_scaling_settings": {
						// Property: WriteCapacityAutoScalingSettings
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"max_capacity": {
									// Property: MaxCapacity
									Type:     types.Int64Type,
									Computed: true,
								},
								"min_capacity": {
									// Property: MinCapacity
									Type:     types.Int64Type,
									Computed: true,
								},
								"seed_capacity": {
									// Property: SeedCapacity
									Type:     types.Int64Type,
									Computed: true,
								},
								"target_tracking_scaling_policy_configuration": {
									// Property: TargetTrackingScalingPolicyConfiguration
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"disable_scale_in": {
												// Property: DisableScaleIn
												Type:     types.BoolType,
												Computed: true,
											},
											"scale_in_cooldown": {
												// Property: ScaleInCooldown
												Type:     types.Int64Type,
												Computed: true,
											},
											"scale_out_cooldown": {
												// Property: ScaleOutCooldown
												Type:     types.Int64Type,
												Computed: true,
											},
											"target_value": {
												// Property: TargetValue
												Type:     types.Float64Type,
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
		Description: "Data Source schema for AWS::DynamoDB::GlobalTable",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DynamoDB::GlobalTable").WithTerraformTypeName("awscc_dynamodb_global_table")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                                  "Arn",
		"attribute_definitions":                "AttributeDefinitions",
		"attribute_name":                       "AttributeName",
		"attribute_type":                       "AttributeType",
		"billing_mode":                         "BillingMode",
		"contributor_insights_specification":   "ContributorInsightsSpecification",
		"disable_scale_in":                     "DisableScaleIn",
		"enabled":                              "Enabled",
		"global_secondary_indexes":             "GlobalSecondaryIndexes",
		"index_name":                           "IndexName",
		"key":                                  "Key",
		"key_schema":                           "KeySchema",
		"key_type":                             "KeyType",
		"kms_master_key_id":                    "KMSMasterKeyId",
		"local_secondary_indexes":              "LocalSecondaryIndexes",
		"max_capacity":                         "MaxCapacity",
		"min_capacity":                         "MinCapacity",
		"non_key_attributes":                   "NonKeyAttributes",
		"point_in_time_recovery_enabled":       "PointInTimeRecoveryEnabled",
		"point_in_time_recovery_specification": "PointInTimeRecoverySpecification",
		"projection":                           "Projection",
		"projection_type":                      "ProjectionType",
		"read_capacity_auto_scaling_settings":  "ReadCapacityAutoScalingSettings",
		"read_capacity_units":                  "ReadCapacityUnits",
		"read_provisioned_throughput_settings": "ReadProvisionedThroughputSettings",
		"region":                               "Region",
		"replicas":                             "Replicas",
		"scale_in_cooldown":                    "ScaleInCooldown",
		"scale_out_cooldown":                   "ScaleOutCooldown",
		"seed_capacity":                        "SeedCapacity",
		"sse_enabled":                          "SSEEnabled",
		"sse_specification":                    "SSESpecification",
		"sse_type":                             "SSEType",
		"stream_arn":                           "StreamArn",
		"stream_specification":                 "StreamSpecification",
		"stream_view_type":                     "StreamViewType",
		"table_class":                          "TableClass",
		"table_id":                             "TableId",
		"table_name":                           "TableName",
		"tags":                                 "Tags",
		"target_tracking_scaling_policy_configuration": "TargetTrackingScalingPolicyConfiguration",
		"target_value":                          "TargetValue",
		"time_to_live_specification":            "TimeToLiveSpecification",
		"value":                                 "Value",
		"write_capacity_auto_scaling_settings":  "WriteCapacityAutoScalingSettings",
		"write_provisioned_throughput_settings": "WriteProvisionedThroughputSettings",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
