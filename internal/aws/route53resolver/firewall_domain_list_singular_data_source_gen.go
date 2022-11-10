// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package route53resolver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_route53resolver_firewall_domain_list", firewallDomainListDataSource)
}

// firewallDomainListDataSource returns the Terraform awscc_route53resolver_firewall_domain_list data source.
// This Terraform data source corresponds to the CloudFormation AWS::Route53Resolver::FirewallDomainList resource.
func firewallDomainListDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Arn",
			//	  "maxLength": 600,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "Arn",
			Type:        types.StringType,
			Computed:    true,
		},
		"creation_time": {
			// Property: CreationTime
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Rfc3339TimeString",
			//	  "maxLength": 40,
			//	  "minLength": 20,
			//	  "type": "string"
			//	}
			Description: "Rfc3339TimeString",
			Type:        types.StringType,
			Computed:    true,
		},
		"creator_request_id": {
			// Property: CreatorRequestId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The id of the creator request.",
			//	  "maxLength": 255,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "The id of the creator request.",
			Type:        types.StringType,
			Computed:    true,
		},
		"domain_count": {
			// Property: DomainCount
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Count",
			//	  "minimum": 0,
			//	  "type": "integer"
			//	}
			Description: "Count",
			Type:        types.Int64Type,
			Computed:    true,
		},
		"domain_file_url": {
			// Property: DomainFileUrl
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "S3 URL to import domains from.",
			//	  "maxLength": 1024,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "S3 URL to import domains from.",
			Type:        types.StringType,
			Computed:    true,
		},
		"domains": {
			// Property: Domains
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "An inline list of domains to use for this domain list.",
			//	  "items": {
			//	    "description": "FirewallDomainName",
			//	    "maxLength": 255,
			//	    "minLength": 1,
			//	    "type": "string"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Description: "An inline list of domains to use for this domain list.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "ResourceId",
			//	  "maxLength": 64,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "ResourceId",
			Type:        types.StringType,
			Computed:    true,
		},
		"managed_owner_name": {
			// Property: ManagedOwnerName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "ServicePrincipal",
			//	  "maxLength": 512,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Description: "ServicePrincipal",
			Type:        types.StringType,
			Computed:    true,
		},
		"modification_time": {
			// Property: ModificationTime
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Rfc3339TimeString",
			//	  "maxLength": 40,
			//	  "minLength": 20,
			//	  "type": "string"
			//	}
			Description: "Rfc3339TimeString",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "FirewallDomainListName",
			//	  "maxLength": 64,
			//	  "minLength": 1,
			//	  "pattern": "",
			//	  "type": "string"
			//	}
			Description: "FirewallDomainListName",
			Type:        types.StringType,
			Computed:    true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "ResolverFirewallDomainList, possible values are COMPLETE, DELETING, UPDATING, COMPLETE_IMPORT_FAILED, IMPORTING, and INACTIVE_OWNER_ACCOUNT_CLOSED.",
			//	  "enum": [
			//	    "COMPLETE",
			//	    "DELETING",
			//	    "UPDATING",
			//	    "COMPLETE_IMPORT_FAILED",
			//	    "IMPORTING",
			//	    "INACTIVE_OWNER_ACCOUNT_CLOSED"
			//	  ],
			//	  "type": "string"
			//	}
			Description: "ResolverFirewallDomainList, possible values are COMPLETE, DELETING, UPDATING, COMPLETE_IMPORT_FAILED, IMPORTING, and INACTIVE_OWNER_ACCOUNT_CLOSED.",
			Type:        types.StringType,
			Computed:    true,
		},
		"status_message": {
			// Property: StatusMessage
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "FirewallDomainListAssociationStatus",
			//	  "type": "string"
			//	}
			Description: "FirewallDomainListAssociationStatus",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Tags",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "A key-value pair to associate with a resource.",
			//	    "properties": {
			//	      "Key": {
			//	        "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//	        "maxLength": 127,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//	        "maxLength": 255,
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
			//	  "uniqueItems": true
			//	}
			Description: "Tags",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
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
		Description: "Data Source schema for AWS::Route53Resolver::FirewallDomainList",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53Resolver::FirewallDomainList").WithTerraformTypeName("awscc_route53resolver_firewall_domain_list")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                "Arn",
		"creation_time":      "CreationTime",
		"creator_request_id": "CreatorRequestId",
		"domain_count":       "DomainCount",
		"domain_file_url":    "DomainFileUrl",
		"domains":            "Domains",
		"id":                 "Id",
		"key":                "Key",
		"managed_owner_name": "ManagedOwnerName",
		"modification_time":  "ModificationTime",
		"name":               "Name",
		"status":             "Status",
		"status_message":     "StatusMessage",
		"tags":               "Tags",
		"value":              "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
