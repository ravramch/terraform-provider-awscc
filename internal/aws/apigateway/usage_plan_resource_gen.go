// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddResourceFactory("awscc_apigateway_usage_plan", usagePlanResource)
}

// usagePlanResource returns the Terraform awscc_apigateway_usage_plan resource.
// This Terraform resource corresponds to the CloudFormation AWS::ApiGateway::UsagePlan resource.
func usagePlanResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ApiStages
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The associated API stages of a usage plan.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "API stage name of the associated API stage in a usage plan.",
		//	    "properties": {
		//	      "ApiId": {
		//	        "description": "API Id of the associated API stage in a usage plan.",
		//	        "type": "string"
		//	      },
		//	      "Stage": {
		//	        "description": "API stage name of the associated API stage in a usage plan.",
		//	        "type": "string"
		//	      },
		//	      "Throttle": {
		//	        "additionalProperties": false,
		//	        "description": "Map containing method level throttling information for API stage in a usage plan.",
		//	        "patternProperties": {
		//	          "": {
		//	            "additionalProperties": false,
		//	            "description": "``ThrottleSettings`` is a property of the [AWS::ApiGateway::UsagePlan](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html) resource that specifies the overall request rate (average requests per second) and burst capacity when users call your REST APIs.",
		//	            "properties": {
		//	              "BurstLimit": {
		//	                "description": "The API target request burst rate limit. This allows more requests through for a period of time than the target rate limit.",
		//	                "minimum": 0,
		//	                "type": "integer"
		//	              },
		//	              "RateLimit": {
		//	                "description": "The API target request rate limit.",
		//	                "minimum": 0,
		//	                "type": "number"
		//	              }
		//	            },
		//	            "type": "object"
		//	          }
		//	        },
		//	        "type": "object"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"api_stages": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ApiId
					"api_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "API Id of the associated API stage in a usage plan.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Stage
					"stage": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "API stage name of the associated API stage in a usage plan.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Throttle
					"throttle":                // Pattern: ""
					schema.MapNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: BurstLimit
								"burst_limit": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Description: "The API target request burst rate limit. This allows more requests through for a period of time than the target rate limit.",
									Optional:    true,
									Computed:    true,
									Validators: []validator.Int64{ /*START VALIDATORS*/
										int64validator.AtLeast(0),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
										int64planmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: RateLimit
								"rate_limit": schema.Float64Attribute{ /*START ATTRIBUTE*/
									Description: "The API target request rate limit.",
									Optional:    true,
									Computed:    true,
									Validators: []validator.Float64{ /*START VALIDATORS*/
										float64validator.AtLeast(0.000000),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
										float64planmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Description: "Map containing method level throttling information for API stage in a usage plan.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
							mapplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The associated API stages of a usage plan.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.UniqueValues(),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of a usage plan.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of a usage plan.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Quota
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The target maximum number of permitted requests per a given unit time interval.",
		//	  "properties": {
		//	    "Limit": {
		//	      "description": "The target maximum number of requests that can be made in a given time period.",
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "Offset": {
		//	      "description": "The number of requests subtracted from the given limit in the initial time period.",
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "Period": {
		//	      "description": "The time period in which the limit applies. Valid values are \"DAY\", \"WEEK\" or \"MONTH\".",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"quota": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Limit
				"limit": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The target maximum number of requests that can be made in a given time period.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.AtLeast(0),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Offset
				"offset": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The number of requests subtracted from the given limit in the initial time period.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.AtLeast(0),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Period
				"period": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The time period in which the limit applies. Valid values are \"DAY\", \"WEEK\" or \"MONTH\".",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The target maximum number of permitted requests per a given unit time interval.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The collection of tags. Each tag element is associated with a given resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType:  cctypes.MultisetType,
			Description: "The collection of tags. Each tag element is associated with a given resource.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Throttle
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A map containing method level throttling information for API stage in a usage plan.",
		//	  "properties": {
		//	    "BurstLimit": {
		//	      "description": "The API target request burst rate limit. This allows more requests through for a period of time than the target rate limit.",
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "RateLimit": {
		//	      "description": "The API target request rate limit.",
		//	      "minimum": 0,
		//	      "type": "number"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"throttle": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: BurstLimit
				"burst_limit": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The API target request burst rate limit. This allows more requests through for a period of time than the target rate limit.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.AtLeast(0),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: RateLimit
				"rate_limit": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "The API target request rate limit.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Float64{ /*START VALIDATORS*/
						float64validator.AtLeast(0.000000),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
						float64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "A map containing method level throttling information for API stage in a usage plan.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UsagePlanName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of a usage plan.",
		//	  "type": "string"
		//	}
		"usage_plan_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of a usage plan.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	schema := schema.Schema{
		Description: "The ``AWS::ApiGateway::UsagePlan`` resource creates a usage plan for deployed APIs. A usage plan sets a target for the throttling and quota limits on individual client API keys. For more information, see [Creating and Using API Usage Plans in Amazon API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-usage-plans.html) in the *API Gateway Developer Guide*.\n In some cases clients can exceed the targets that you set. Don?t rely on usage plans to control costs. Consider using [](https://docs.aws.amazon.com/cost-management/latest/userguide/budgets-managing-costs.html) to monitor costs and [](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) to manage API requests.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::UsagePlan").WithTerraformTypeName("awscc_apigateway_usage_plan")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"api_id":          "ApiId",
		"api_stages":      "ApiStages",
		"burst_limit":     "BurstLimit",
		"description":     "Description",
		"id":              "Id",
		"key":             "Key",
		"limit":           "Limit",
		"offset":          "Offset",
		"period":          "Period",
		"quota":           "Quota",
		"rate_limit":      "RateLimit",
		"stage":           "Stage",
		"tags":            "Tags",
		"throttle":        "Throttle",
		"usage_plan_name": "UsagePlanName",
		"value":           "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
