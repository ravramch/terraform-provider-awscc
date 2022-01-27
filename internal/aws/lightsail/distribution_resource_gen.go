// Code generated by generators/resource/main.go; DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_lightsail_distribution", distributionResourceType)
}

// distributionResourceType returns the Terraform awscc_lightsail_distribution resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Lightsail::Distribution resource type.
func distributionResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"able_to_update_bundle": {
			// Property: AbleToUpdateBundle
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates whether the bundle that is currently applied to your distribution, specified using the distributionName parameter, can be changed to another bundle.",
			//   "type": "boolean"
			// }
			Description: "Indicates whether the bundle that is currently applied to your distribution, specified using the distributionName parameter, can be changed to another bundle.",
			Type:        types.BoolType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"bundle_id": {
			// Property: BundleId
			// CloudFormation resource type schema:
			// {
			//   "description": "The bundle ID to use for the distribution.",
			//   "type": "string"
			// }
			Description: "The bundle ID to use for the distribution.",
			Type:        types.StringType,
			Required:    true,
		},
		"cache_behavior_settings": {
			// Property: CacheBehaviorSettings
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "An object that describes the cache behavior settings for the distribution.",
			//   "properties": {
			//     "AllowedHTTPMethods": {
			//       "description": "The HTTP methods that are processed and forwarded to the distribution's origin.",
			//       "type": "string"
			//     },
			//     "CachedHTTPMethods": {
			//       "description": "The HTTP method responses that are cached by your distribution.",
			//       "type": "string"
			//     },
			//     "DefaultTTL": {
			//       "description": "The default amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the content has been updated.",
			//       "format": "int64",
			//       "type": "integer"
			//     },
			//     "ForwardedCookies": {
			//       "additionalProperties": false,
			//       "description": "An object that describes the cookies that are forwarded to the origin. Your content is cached based on the cookies that are forwarded.",
			//       "properties": {
			//         "CookiesAllowList": {
			//           "description": "The specific cookies to forward to your distribution's origin.",
			//           "insertionOrder": false,
			//           "items": {
			//             "type": "string"
			//           },
			//           "type": "array",
			//           "uniqueItems": true
			//         },
			//         "Option": {
			//           "description": "Specifies which cookies to forward to the distribution's origin for a cache behavior: all, none, or allow-list to forward only the cookies specified in the cookiesAllowList parameter.",
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "ForwardedHeaders": {
			//       "additionalProperties": false,
			//       "description": "An object that describes the headers that are forwarded to the origin. Your content is cached based on the headers that are forwarded.",
			//       "properties": {
			//         "HeadersAllowList": {
			//           "description": "The specific headers to forward to your distribution's origin.",
			//           "insertionOrder": false,
			//           "items": {
			//             "type": "string"
			//           },
			//           "type": "array",
			//           "uniqueItems": true
			//         },
			//         "Option": {
			//           "description": "The headers that you want your distribution to forward to your origin and base caching on.",
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "ForwardedQueryStrings": {
			//       "additionalProperties": false,
			//       "description": "An object that describes the query strings that are forwarded to the origin. Your content is cached based on the query strings that are forwarded.",
			//       "properties": {
			//         "Option": {
			//           "description": "Indicates whether the distribution forwards and caches based on query strings.",
			//           "type": "boolean"
			//         },
			//         "QueryStringsAllowList": {
			//           "description": "The specific query strings that the distribution forwards to the origin.",
			//           "insertionOrder": false,
			//           "items": {
			//             "type": "string"
			//           },
			//           "type": "array",
			//           "uniqueItems": true
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "MaximumTTL": {
			//       "description": "The maximum amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the object has been updated.",
			//       "format": "int64",
			//       "type": "integer"
			//     },
			//     "MinimumTTL": {
			//       "description": "The minimum amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the object has been updated.",
			//       "format": "int64",
			//       "type": "integer"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "An object that describes the cache behavior settings for the distribution.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"allowed_http_methods": {
						// Property: AllowedHTTPMethods
						Description: "The HTTP methods that are processed and forwarded to the distribution's origin.",
						Type:        types.StringType,
						Optional:    true,
					},
					"cached_http_methods": {
						// Property: CachedHTTPMethods
						Description: "The HTTP method responses that are cached by your distribution.",
						Type:        types.StringType,
						Optional:    true,
					},
					"default_ttl": {
						// Property: DefaultTTL
						Description: "The default amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the content has been updated.",
						Type:        types.NumberType,
						Optional:    true,
					},
					"forwarded_cookies": {
						// Property: ForwardedCookies
						Description: "An object that describes the cookies that are forwarded to the origin. Your content is cached based on the cookies that are forwarded.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"cookies_allow_list": {
									// Property: CookiesAllowList
									Description: "The specific cookies to forward to your distribution's origin.",
									Type:        types.SetType{ElemType: types.StringType},
									Optional:    true,
								},
								"option": {
									// Property: Option
									Description: "Specifies which cookies to forward to the distribution's origin for a cache behavior: all, none, or allow-list to forward only the cookies specified in the cookiesAllowList parameter.",
									Type:        types.StringType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"forwarded_headers": {
						// Property: ForwardedHeaders
						Description: "An object that describes the headers that are forwarded to the origin. Your content is cached based on the headers that are forwarded.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"headers_allow_list": {
									// Property: HeadersAllowList
									Description: "The specific headers to forward to your distribution's origin.",
									Type:        types.SetType{ElemType: types.StringType},
									Optional:    true,
								},
								"option": {
									// Property: Option
									Description: "The headers that you want your distribution to forward to your origin and base caching on.",
									Type:        types.StringType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"forwarded_query_strings": {
						// Property: ForwardedQueryStrings
						Description: "An object that describes the query strings that are forwarded to the origin. Your content is cached based on the query strings that are forwarded.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"option": {
									// Property: Option
									Description: "Indicates whether the distribution forwards and caches based on query strings.",
									Type:        types.BoolType,
									Optional:    true,
								},
								"query_strings_allow_list": {
									// Property: QueryStringsAllowList
									Description: "The specific query strings that the distribution forwards to the origin.",
									Type:        types.SetType{ElemType: types.StringType},
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
					"maximum_ttl": {
						// Property: MaximumTTL
						Description: "The maximum amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the object has been updated.",
						Type:        types.NumberType,
						Optional:    true,
					},
					"minimum_ttl": {
						// Property: MinimumTTL
						Description: "The minimum amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the object has been updated.",
						Type:        types.NumberType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"cache_behaviors": {
			// Property: CacheBehaviors
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of objects that describe the per-path cache behavior for the distribution.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "Describes the per-path cache behavior of an Amazon Lightsail content delivery network (CDN) distribution.",
			//     "properties": {
			//       "Behavior": {
			//         "description": "The cache behavior for the specified path.",
			//         "type": "string"
			//       },
			//       "Path": {
			//         "description": "The path to a directory or file to cached, or not cache. Use an asterisk symbol to specify wildcard directories (path/to/assets/*), and file types (*.html, *jpg, *js). Directories and file paths are case-sensitive.",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of objects that describe the per-path cache behavior for the distribution.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"behavior": {
						// Property: Behavior
						Description: "The cache behavior for the specified path.",
						Type:        types.StringType,
						Optional:    true,
					},
					"path": {
						// Property: Path
						Description: "The path to a directory or file to cached, or not cache. Use an asterisk symbol to specify wildcard directories (path/to/assets/*), and file types (*.html, *jpg, *js). Directories and file paths are case-sensitive.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Optional: true,
		},
		"certificate_name": {
			// Property: CertificateName
			// CloudFormation resource type schema:
			// {
			//   "description": "The certificate attached to the Distribution.",
			//   "type": "string"
			// }
			Description: "The certificate attached to the Distribution.",
			Type:        types.StringType,
			Optional:    true,
		},
		"default_cache_behavior": {
			// Property: DefaultCacheBehavior
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "An object that describes the default cache behavior for the distribution.",
			//   "properties": {
			//     "Behavior": {
			//       "description": "The cache behavior of the distribution.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "An object that describes the default cache behavior for the distribution.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"behavior": {
						// Property: Behavior
						Description: "The cache behavior of the distribution.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Required: true,
		},
		"distribution_arn": {
			// Property: DistributionArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"distribution_name": {
			// Property: DistributionName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name for the distribution.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name for the distribution.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"ip_address_type": {
			// Property: IpAddressType
			// CloudFormation resource type schema:
			// {
			//   "description": "The IP address type for the distribution.",
			//   "type": "string"
			// }
			Description: "The IP address type for the distribution.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"is_enabled": {
			// Property: IsEnabled
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates whether the distribution is enabled.",
			//   "type": "boolean"
			// }
			Description: "Indicates whether the distribution is enabled.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"origin": {
			// Property: Origin
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "An object that describes the origin resource for the distribution, such as a Lightsail instance or load balancer.",
			//   "properties": {
			//     "Name": {
			//       "description": "The name of the origin resource.",
			//       "type": "string"
			//     },
			//     "ProtocolPolicy": {
			//       "description": "The protocol that your Amazon Lightsail distribution uses when establishing a connection with your origin to pull content.",
			//       "type": "string"
			//     },
			//     "RegionName": {
			//       "description": "The AWS Region name of the origin resource.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "An object that describes the origin resource for the distribution, such as a Lightsail instance or load balancer.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"name": {
						// Property: Name
						Description: "The name of the origin resource.",
						Type:        types.StringType,
						Optional:    true,
					},
					"protocol_policy": {
						// Property: ProtocolPolicy
						Description: "The protocol that your Amazon Lightsail distribution uses when establishing a connection with your origin to pull content.",
						Type:        types.StringType,
						Optional:    true,
					},
					"region_name": {
						// Property: RegionName
						Description: "The AWS Region name of the origin resource.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Required: true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "The status of the distribution.",
			//   "type": "string"
			// }
			Description: "The status of the distribution.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Optional: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			tfsdk.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::Lightsail::Distribution",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lightsail::Distribution").WithTerraformTypeName("awscc_lightsail_distribution")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"able_to_update_bundle":    "AbleToUpdateBundle",
		"allowed_http_methods":     "AllowedHTTPMethods",
		"behavior":                 "Behavior",
		"bundle_id":                "BundleId",
		"cache_behavior_settings":  "CacheBehaviorSettings",
		"cache_behaviors":          "CacheBehaviors",
		"cached_http_methods":      "CachedHTTPMethods",
		"certificate_name":         "CertificateName",
		"cookies_allow_list":       "CookiesAllowList",
		"default_cache_behavior":   "DefaultCacheBehavior",
		"default_ttl":              "DefaultTTL",
		"distribution_arn":         "DistributionArn",
		"distribution_name":        "DistributionName",
		"forwarded_cookies":        "ForwardedCookies",
		"forwarded_headers":        "ForwardedHeaders",
		"forwarded_query_strings":  "ForwardedQueryStrings",
		"headers_allow_list":       "HeadersAllowList",
		"ip_address_type":          "IpAddressType",
		"is_enabled":               "IsEnabled",
		"key":                      "Key",
		"maximum_ttl":              "MaximumTTL",
		"minimum_ttl":              "MinimumTTL",
		"name":                     "Name",
		"option":                   "Option",
		"origin":                   "Origin",
		"path":                     "Path",
		"protocol_policy":          "ProtocolPolicy",
		"query_strings_allow_list": "QueryStringsAllowList",
		"region_name":              "RegionName",
		"status":                   "Status",
		"tags":                     "Tags",
		"value":                    "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}