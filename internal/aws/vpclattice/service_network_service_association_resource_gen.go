// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package vpclattice

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
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
	registry.AddResourceFactory("awscc_vpclattice_service_network_service_association", serviceNetworkServiceAssociationResource)
}

// serviceNetworkServiceAssociationResource returns the Terraform awscc_vpclattice_service_network_service_association resource.
// This Terraform resource corresponds to the CloudFormation AWS::VpcLattice::ServiceNetworkServiceAssociation resource.
func serviceNetworkServiceAssociationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:[a-z0-9\\-]+:vpc-lattice:[a-zA-Z0-9\\-]+:\\d{12}:servicenetworkserviceassociation/snsa-[0-9a-z]{17}$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DnsEntry
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "DomainName": {
		//	      "type": "string"
		//	    },
		//	    "HostedZoneId": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"dns_entry": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DomainName
				"domain_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: HostedZoneId
				"hosted_zone_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 17,
		//	  "pattern": "^snsa-[0-9a-z]{17}$",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ServiceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:[a-z0-9\\-]+:vpc-lattice:[a-zA-Z0-9\\-]+:\\d{12}:service/svc-[0-9a-z]{17}$",
		//	  "type": "string"
		//	}
		"service_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ServiceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 21,
		//	  "minLength": 21,
		//	  "pattern": "^svc-[0-9a-z]{17}$",
		//	  "type": "string"
		//	}
		"service_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ServiceIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^((svc-[0-9a-z]{17})|(arn:[a-z0-9\\-]+:vpc-lattice:[a-zA-Z0-9\\-]+:\\d{12}:service/svc-[0-9a-z]{17}))$",
		//	  "type": "string"
		//	}
		"service_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(20, 2048),
				stringvalidator.RegexMatches(regexp.MustCompile("^((svc-[0-9a-z]{17})|(arn:[a-z0-9\\-]+:vpc-lattice:[a-zA-Z0-9\\-]+:\\d{12}:service/svc-[0-9a-z]{17}))$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// ServiceIdentifier is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: ServiceName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 40,
		//	  "minLength": 3,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"service_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ServiceNetworkArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:[a-z0-9\\-]+:vpc-lattice:[a-zA-Z0-9\\-]+:\\d{12}:servicenetwork/sn-[0-9a-z]{17}$",
		//	  "type": "string"
		//	}
		"service_network_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ServiceNetworkId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 20,
		//	  "minLength": 20,
		//	  "pattern": "^sn-[0-9a-z]{17}$",
		//	  "type": "string"
		//	}
		"service_network_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ServiceNetworkIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^((sn-[0-9a-z]{17})|(arn:[a-z0-9\\-]+:vpc-lattice:[a-zA-Z0-9\\-]+:\\d{12}:servicenetwork/sn-[0-9a-z]{17}))$",
		//	  "type": "string"
		//	}
		"service_network_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(20, 2048),
				stringvalidator.RegexMatches(regexp.MustCompile("^((sn-[0-9a-z]{17})|(arn:[a-z0-9\\-]+:vpc-lattice:[a-zA-Z0-9\\-]+:\\d{12}:servicenetwork/sn-[0-9a-z]{17}))$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// ServiceNetworkIdentifier is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: ServiceNetworkName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 63,
		//	  "minLength": 3,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"service_network_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "CREATE_IN_PROGRESS",
		//	    "ACTIVE",
		//	    "DELETE_IN_PROGRESS",
		//	    "CREATE_FAILED",
		//	    "DELETE_FAILED"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
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
		//	  "maxItems": 50,
		//	  "minItems": 0,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeBetween(0, 50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
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
		Description: "Associates a service with a service network.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::VpcLattice::ServiceNetworkServiceAssociation").WithTerraformTypeName("awscc_vpclattice_service_network_service_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                                    "Arn",
		"created_at":                             "CreatedAt",
		"dns_entry":                              "DnsEntry",
		"domain_name":                            "DomainName",
		"hosted_zone_id":                         "HostedZoneId",
		"key":                                    "Key",
		"service_arn":                            "ServiceArn",
		"service_id":                             "ServiceId",
		"service_identifier":                     "ServiceIdentifier",
		"service_name":                           "ServiceName",
		"service_network_arn":                    "ServiceNetworkArn",
		"service_network_id":                     "ServiceNetworkId",
		"service_network_identifier":             "ServiceNetworkIdentifier",
		"service_network_name":                   "ServiceNetworkName",
		"service_network_service_association_id": "Id",
		"status":                                 "Status",
		"tags":                                   "Tags",
		"value":                                  "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/ServiceNetworkIdentifier",
		"/properties/ServiceIdentifier",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
