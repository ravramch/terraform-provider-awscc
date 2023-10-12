// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_apigatewayv2_route", routeResource)
}

// routeResource returns the Terraform awscc_apigatewayv2_route resource.
// This Terraform resource corresponds to the CloudFormation AWS::ApiGatewayV2::Route resource.
func routeResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ApiId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The API identifier.",
		//	  "type": "string"
		//	}
		"api_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The API identifier.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ApiKeyRequired
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies whether an API key is required for the route. Supported only for WebSocket APIs.",
		//	  "type": "boolean"
		//	}
		"api_key_required": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether an API key is required for the route. Supported only for WebSocket APIs.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AuthorizationScopes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The authorization scopes supported by this route.",
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"authorization_scopes": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The authorization scopes supported by this route.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AuthorizationType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The authorization type for the route. For WebSocket APIs, valid values are ``NONE`` for open access, ``AWS_IAM`` for using AWS IAM permissions, and ``CUSTOM`` for using a Lambda authorizer. For HTTP APIs, valid values are ``NONE`` for open access, ``JWT`` for using JSON Web Tokens, ``AWS_IAM`` for using AWS IAM permissions, and ``CUSTOM`` for using a Lambda authorizer.",
		//	  "type": "string"
		//	}
		"authorization_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The authorization type for the route. For WebSocket APIs, valid values are ``NONE`` for open access, ``AWS_IAM`` for using AWS IAM permissions, and ``CUSTOM`` for using a Lambda authorizer. For HTTP APIs, valid values are ``NONE`` for open access, ``JWT`` for using JSON Web Tokens, ``AWS_IAM`` for using AWS IAM permissions, and ``CUSTOM`` for using a Lambda authorizer.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AuthorizerId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the ``Authorizer`` resource to be associated with this route. The authorizer identifier is generated by API Gateway when you created the authorizer.",
		//	  "type": "string"
		//	}
		"authorizer_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the ``Authorizer`` resource to be associated with this route. The authorizer identifier is generated by API Gateway when you created the authorizer.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// AuthorizerId is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: ModelSelectionExpression
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The model selection expression for the route. Supported only for WebSocket APIs.",
		//	  "type": "string"
		//	}
		"model_selection_expression": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The model selection expression for the route. Supported only for WebSocket APIs.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: OperationName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The operation name for the route.",
		//	  "type": "string"
		//	}
		"operation_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The operation name for the route.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RequestModels
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The request models for the route. Supported only for WebSocket APIs.",
		//	  "type": "object"
		//	}
		"request_models": schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The request models for the route. Supported only for WebSocket APIs.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RequestParameters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The request parameters for the route. Supported only for WebSocket APIs.",
		//	  "items": {
		//	    "$ref": "#/definitions/ParameterConstraints"
		//	  },
		//	  "type": "object"
		//	}
		"request_parameters": schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The request parameters for the route. Supported only for WebSocket APIs.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// RequestParameters is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: RouteId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"route_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RouteKey
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The route key for the route. For HTTP APIs, the route key can be either ``$default``, or a combination of an HTTP method and resource path, for example, ``GET /pets``.",
		//	  "type": "string"
		//	}
		"route_key": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The route key for the route. For HTTP APIs, the route key can be either ``$default``, or a combination of an HTTP method and resource path, for example, ``GET /pets``.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: RouteResponseSelectionExpression
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The route response selection expression for the route. Supported only for WebSocket APIs.",
		//	  "type": "string"
		//	}
		"route_response_selection_expression": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The route response selection expression for the route. Supported only for WebSocket APIs.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Target
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The target for the route.",
		//	  "type": "string"
		//	}
		"target": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The target for the route.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "The ``AWS::ApiGatewayV2::Route`` resource creates a route for an API.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGatewayV2::Route").WithTerraformTypeName("awscc_apigatewayv2_route")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"api_id":                              "ApiId",
		"api_key_required":                    "ApiKeyRequired",
		"authorization_scopes":                "AuthorizationScopes",
		"authorization_type":                  "AuthorizationType",
		"authorizer_id":                       "AuthorizerId",
		"model_selection_expression":          "ModelSelectionExpression",
		"operation_name":                      "OperationName",
		"request_models":                      "RequestModels",
		"request_parameters":                  "RequestParameters",
		"route_id":                            "RouteId",
		"route_key":                           "RouteKey",
		"route_response_selection_expression": "RouteResponseSelectionExpression",
		"target":                              "Target",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/AuthorizerId",
		"/properties/RequestParameters",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
