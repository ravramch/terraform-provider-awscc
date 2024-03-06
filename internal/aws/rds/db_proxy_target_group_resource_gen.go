// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package rds

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddResourceFactory("awscc_rds_db_proxy_target_group", dBProxyTargetGroupResource)
}

// dBProxyTargetGroupResource returns the Terraform awscc_rds_db_proxy_target_group resource.
// This Terraform resource corresponds to the CloudFormation AWS::RDS::DBProxyTargetGroup resource.
func dBProxyTargetGroupResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ConnectionPoolConfigurationInfo
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "ConnectionBorrowTimeout": {
		//	      "description": "The number of seconds for a proxy to wait for a connection to become available in the connection pool.",
		//	      "type": "integer"
		//	    },
		//	    "InitQuery": {
		//	      "description": "One or more SQL statements for the proxy to run when opening each new database connection.",
		//	      "type": "string"
		//	    },
		//	    "MaxConnectionsPercent": {
		//	      "description": "The maximum size of the connection pool for each target in a target group.",
		//	      "maximum": 100,
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "MaxIdleConnectionsPercent": {
		//	      "description": "Controls how actively the proxy closes idle database connections in the connection pool.",
		//	      "maximum": 100,
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "SessionPinningFilters": {
		//	      "description": "Each item in the list represents a class of SQL operations that normally cause all later statements in a session using a proxy to be pinned to the same underlying database connection.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"connection_pool_configuration_info": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ConnectionBorrowTimeout
				"connection_borrow_timeout": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The number of seconds for a proxy to wait for a connection to become available in the connection pool.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: InitQuery
				"init_query": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "One or more SQL statements for the proxy to run when opening each new database connection.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: MaxConnectionsPercent
				"max_connections_percent": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The maximum size of the connection pool for each target in a target group.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.Between(0, 100),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: MaxIdleConnectionsPercent
				"max_idle_connections_percent": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "Controls how actively the proxy closes idle database connections in the connection pool.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.Between(0, 100),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: SessionPinningFilters
				"session_pinning_filters": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					CustomType:  cctypes.MultisetType,
					Description: "Each item in the list represents a class of SQL operations that normally cause all later statements in a session using a proxy to be pinned to the same underlying database connection.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DBClusterIdentifiers
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"db_cluster_identifiers": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			CustomType:  cctypes.MultisetType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DBInstanceIdentifiers
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"db_instance_identifiers": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			CustomType:  cctypes.MultisetType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DBProxyName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier for the proxy.",
		//	  "maxLength": 64,
		//	  "pattern": "[A-z][0-z]*",
		//	  "type": "string"
		//	}
		"db_proxy_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier for the proxy.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(64),
				stringvalidator.RegexMatches(regexp.MustCompile("[A-z][0-z]*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TargetGroupArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) representing the target group.",
		//	  "type": "string"
		//	}
		"target_group_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) representing the target group.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TargetGroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier for the DBProxyTargetGroup",
		//	  "enum": [
		//	    "default"
		//	  ],
		//	  "type": "string"
		//	}
		"target_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier for the DBProxyTargetGroup",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"default",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
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
		Description: "Resource schema for AWS::RDS::DBProxyTargetGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RDS::DBProxyTargetGroup").WithTerraformTypeName("awscc_rds_db_proxy_target_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"connection_borrow_timeout":          "ConnectionBorrowTimeout",
		"connection_pool_configuration_info": "ConnectionPoolConfigurationInfo",
		"db_cluster_identifiers":             "DBClusterIdentifiers",
		"db_instance_identifiers":            "DBInstanceIdentifiers",
		"db_proxy_name":                      "DBProxyName",
		"init_query":                         "InitQuery",
		"max_connections_percent":            "MaxConnectionsPercent",
		"max_idle_connections_percent":       "MaxIdleConnectionsPercent",
		"session_pinning_filters":            "SessionPinningFilters",
		"target_group_arn":                   "TargetGroupArn",
		"target_group_name":                  "TargetGroupName",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
