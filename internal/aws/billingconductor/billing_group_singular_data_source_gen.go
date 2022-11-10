// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package billingconductor

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_billingconductor_billing_group", billingGroupDataSource)
}

// billingGroupDataSource returns the Terraform awscc_billingconductor_billing_group data source.
// This Terraform data source corresponds to the CloudFormation AWS::BillingConductor::BillingGroup resource.
func billingGroupDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"account_grouping": {
			// Property: AccountGrouping
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "properties": {
			//	    "LinkedAccountIds": {
			//	      "insertionOrder": false,
			//	      "items": {
			//	        "pattern": "[0-9]{12}",
			//	        "type": "string"
			//	      },
			//	      "minItems": 1,
			//	      "type": "array",
			//	      "uniqueItems": true
			//	    }
			//	  },
			//	  "required": [
			//	    "LinkedAccountIds"
			//	  ],
			//	  "type": "object"
			//	}
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"linked_account_ids": {
						// Property: LinkedAccountIds
						Type:     types.SetType{ElemType: types.StringType},
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Billing Group ARN",
			//	  "pattern": "arn:aws(-cn)?:billingconductor::[0-9]{12}:billinggroup/?[0-9]{12}",
			//	  "type": "string"
			//	}
			Description: "Billing Group ARN",
			Type:        types.StringType,
			Computed:    true,
		},
		"computation_preference": {
			// Property: ComputationPreference
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "properties": {
			//	    "PricingPlanArn": {
			//	      "description": "ARN of the attached pricing plan",
			//	      "pattern": "arn:aws(-cn)?:billingconductor::[0-9]{12}:pricingplan/[a-zA-Z0-9]{10}",
			//	      "type": "string"
			//	    }
			//	  },
			//	  "required": [
			//	    "PricingPlanArn"
			//	  ],
			//	  "type": "object"
			//	}
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"pricing_plan_arn": {
						// Property: PricingPlanArn
						Description: "ARN of the attached pricing plan",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"creation_time": {
			// Property: CreationTime
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Creation timestamp in UNIX epoch time format",
			//	  "type": "integer"
			//	}
			Description: "Creation timestamp in UNIX epoch time format",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 1024,
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"last_modified_time": {
			// Property: LastModifiedTime
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Latest modified timestamp in UNIX epoch time format",
			//	  "type": "integer"
			//	}
			Description: "Latest modified timestamp in UNIX epoch time format",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 128,
			//	  "minLength": 1,
			//	  "pattern": "[a-zA-Z0-9_\\+=\\.\\-@]+",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"primary_account_id": {
			// Property: PrimaryAccountId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "This account will act as a virtual payer account of the billing group",
			//	  "pattern": "[0-9]{12}",
			//	  "type": "string"
			//	}
			Description: "This account will act as a virtual payer account of the billing group",
			Type:        types.StringType,
			Computed:    true,
		},
		"size": {
			// Property: Size
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Number of accounts in the billing group",
			//	  "type": "integer"
			//	}
			Description: "Number of accounts in the billing group",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			//
			//	{
			//	  "enum": [
			//	    "ACTIVE",
			//	    "PRIMARY_ACCOUNT_MISSING"
			//	  ],
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"status_reason": {
			// Property: StatusReason
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
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
			//	    "properties": {
			//	      "Key": {
			//	        "maxLength": 128,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "maxLength": 256,
			//	        "minLength": 1,
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
			//	  "uniqueItems": true
			//	}
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
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::BillingConductor::BillingGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::BillingConductor::BillingGroup").WithTerraformTypeName("awscc_billingconductor_billing_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_grouping":       "AccountGrouping",
		"arn":                    "Arn",
		"computation_preference": "ComputationPreference",
		"creation_time":          "CreationTime",
		"description":            "Description",
		"key":                    "Key",
		"last_modified_time":     "LastModifiedTime",
		"linked_account_ids":     "LinkedAccountIds",
		"name":                   "Name",
		"pricing_plan_arn":       "PricingPlanArn",
		"primary_account_id":     "PrimaryAccountId",
		"size":                   "Size",
		"status":                 "Status",
		"status_reason":          "StatusReason",
		"tags":                   "Tags",
		"value":                  "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
