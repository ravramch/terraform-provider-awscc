// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package s3outposts

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_s3outposts_endpoint", endpointResource)
}

// endpointResource returns the Terraform awscc_s3outposts_endpoint resource.
// This Terraform resource corresponds to the CloudFormation AWS::S3Outposts::Endpoint resource.
func endpointResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccessType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "Private",
		//	  "description": "The type of access for the on-premise network connectivity for the Outpost endpoint. To access endpoint from an on-premises network, you must specify the access type and provide the customer owned Ipv4 pool.",
		//	  "enum": [
		//	    "CustomerOwnedIp",
		//	    "Private"
		//	  ],
		//	  "type": "string"
		//	}
		"access_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of access for the on-premise network connectivity for the Outpost endpoint. To access endpoint from an on-premises network, you must specify the access type and provide the customer owned Ipv4 pool.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"CustomerOwnedIp",
					"Private",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				generic.StringDefaultValue("Private"),
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the endpoint.",
		//	  "maxLength": 500,
		//	  "minLength": 5,
		//	  "pattern": "^arn:[^:]+:s3-outposts:[a-zA-Z0-9\\-]+:\\d{12}:outpost\\/[^:]+\\/endpoint/[a-zA-Z0-9]{19}$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the endpoint.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CidrBlock
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The VPC CIDR committed by this endpoint.",
		//	  "maxLength": 20,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"cidr_block": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The VPC CIDR committed by this endpoint.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time the endpoint was created.",
		//	  "pattern": "^([0-2]\\d{3})-(0[0-9]|1[0-2])-([0-2]\\d|3[01])T([01]\\d|2[0-4]):([0-5]\\d):([0-6]\\d)((\\.\\d{3})?)Z$",
		//	  "type": "string"
		//	}
		"creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time the endpoint was created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CustomerOwnedIpv4Pool
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the customer-owned IPv4 pool for the Endpoint. IP addresses will be allocated from this pool for the endpoint.",
		//	  "pattern": "^ipv4pool-coip-([0-9a-f]{17})$",
		//	  "type": "string"
		//	}
		"customer_owned_ipv_4_pool": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the customer-owned IPv4 pool for the Endpoint. IP addresses will be allocated from this pool for the endpoint.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^ipv4pool-coip-([0-9a-f]{17})$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FailedReason
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The failure reason, if any, for a create or delete endpoint operation.",
		//	  "properties": {
		//	    "ErrorCode": {
		//	      "description": "The failure code, if any, for a create or delete endpoint operation.",
		//	      "type": "string"
		//	    },
		//	    "Message": {
		//	      "description": "Additional error details describing the endpoint failure and recommended action.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"failed_reason": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ErrorCode
				"error_code": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The failure code, if any, for a create or delete endpoint operation.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Message
				"message": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Additional error details describing the endpoint failure and recommended action.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The failure reason, if any, for a create or delete endpoint operation.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the endpoint.",
		//	  "maxLength": 500,
		//	  "minLength": 5,
		//	  "pattern": "^[a-zA-Z0-9]{19}$",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the endpoint.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NetworkInterfaces
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The network interfaces of the endpoint.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The container for the network interface.",
		//	    "properties": {
		//	      "NetworkInterfaceId": {
		//	        "maxLength": 100,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "NetworkInterfaceId"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"network_interfaces": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: NetworkInterfaceId
					"network_interface_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The network interfaces of the endpoint.",
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: OutpostId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The id of the customer outpost on which the bucket resides.",
		//	  "pattern": "^(op-[a-f0-9]{17}|\\d{12}|ec2)$",
		//	  "type": "string"
		//	}
		"outpost_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The id of the customer outpost on which the bucket resides.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^(op-[a-f0-9]{17}|\\d{12}|ec2)$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SecurityGroupId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the security group to use with the endpoint.",
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "pattern": "^sg-([0-9a-f]{8}|[0-9a-f]{17})$",
		//	  "type": "string"
		//	}
		"security_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the security group to use with the endpoint.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 100),
				stringvalidator.RegexMatches(regexp.MustCompile("^sg-([0-9a-f]{8}|[0-9a-f]{17})$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "Available",
		//	    "Pending",
		//	    "Deleting",
		//	    "Create_Failed",
		//	    "Delete_Failed"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SubnetId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the subnet in the selected VPC. The subnet must belong to the Outpost.",
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "pattern": "^subnet-([0-9a-f]{8}|[0-9a-f]{17})$",
		//	  "type": "string"
		//	}
		"subnet_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the subnet in the selected VPC. The subnet must belong to the Outpost.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 100),
				stringvalidator.RegexMatches(regexp.MustCompile("^subnet-([0-9a-f]{8}|[0-9a-f]{17})$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource Type Definition for AWS::S3Outposts::Endpoint",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3Outposts::Endpoint").WithTerraformTypeName("awscc_s3outposts_endpoint")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_type":               "AccessType",
		"arn":                       "Arn",
		"cidr_block":                "CidrBlock",
		"creation_time":             "CreationTime",
		"customer_owned_ipv_4_pool": "CustomerOwnedIpv4Pool",
		"endpoint_id":               "Id",
		"error_code":                "ErrorCode",
		"failed_reason":             "FailedReason",
		"message":                   "Message",
		"network_interface_id":      "NetworkInterfaceId",
		"network_interfaces":        "NetworkInterfaces",
		"outpost_id":                "OutpostId",
		"security_group_id":         "SecurityGroupId",
		"status":                    "Status",
		"subnet_id":                 "SubnetId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
