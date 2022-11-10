// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package nimblestudio

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_nimblestudio_streaming_image", streamingImageDataSource)
}

// streamingImageDataSource returns the Terraform awscc_nimblestudio_streaming_image data source.
// This Terraform data source corresponds to the CloudFormation AWS::NimbleStudio::StreamingImage resource.
func streamingImageDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "\u003cp\u003eA human-readable description of the streaming image.\u003c/p\u003e",
			//	  "maxLength": 256,
			//	  "minLength": 0,
			//	  "type": "string"
			//	}
			Description: "<p>A human-readable description of the streaming image.</p>",
			Type:        types.StringType,
			Computed:    true,
		},
		"ec_2_image_id": {
			// Property: Ec2ImageId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "\u003cp\u003eThe ID of an EC2 machine image with which to create this streaming image.\u003c/p\u003e",
			//	  "pattern": "^ami-[0-9A-z]+$",
			//	  "type": "string"
			//	}
			Description: "<p>The ID of an EC2 machine image with which to create this streaming image.</p>",
			Type:        types.StringType,
			Computed:    true,
		},
		"encryption_configuration": {
			// Property: EncryptionConfiguration
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "\u003cp\u003eTODO\u003c/p\u003e",
			//	  "properties": {
			//	    "KeyArn": {
			//	      "description": "\u003cp\u003eThe ARN for a KMS key that is used to encrypt studio data.\u003c/p\u003e",
			//	      "minLength": 4,
			//	      "pattern": "^arn:.*",
			//	      "type": "string"
			//	    },
			//	    "KeyType": {
			//	      "description": "\u003cp/\u003e",
			//	      "enum": [
			//	        "CUSTOMER_MANAGED_KEY"
			//	      ],
			//	      "type": "string"
			//	    }
			//	  },
			//	  "required": [
			//	    "KeyType"
			//	  ],
			//	  "type": "object"
			//	}
			Description: "<p>TODO</p>",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"key_arn": {
						// Property: KeyArn
						Description: "<p>The ARN for a KMS key that is used to encrypt studio data.</p>",
						Type:        types.StringType,
						Computed:    true,
					},
					"key_type": {
						// Property: KeyType
						Description: "<p/>",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"eula_ids": {
			// Property: EulaIds
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "\u003cp\u003eThe list of EULAs that must be accepted before a Streaming Session can be started using this streaming image.\u003c/p\u003e",
			//	  "items": {
			//	    "type": "string"
			//	  },
			//	  "type": "array"
			//	}
			Description: "<p>The list of EULAs that must be accepted before a Streaming Session can be started using this streaming image.</p>",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "\u003cp\u003eA friendly name for a streaming image resource.\u003c/p\u003e",
			//	  "maxLength": 64,
			//	  "minLength": 0,
			//	  "type": "string"
			//	}
			Description: "<p>A friendly name for a streaming image resource.</p>",
			Type:        types.StringType,
			Computed:    true,
		},
		"owner": {
			// Property: Owner
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "\u003cp\u003eThe owner of the streaming image, either the studioId that contains the streaming image, or 'amazon' for images that are provided by Amazon Nimble Studio.\u003c/p\u003e",
			//	  "type": "string"
			//	}
			Description: "<p>The owner of the streaming image, either the studioId that contains the streaming image, or 'amazon' for images that are provided by Amazon Nimble Studio.</p>",
			Type:        types.StringType,
			Computed:    true,
		},
		"platform": {
			// Property: Platform
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "\u003cp\u003eThe platform of the streaming image, either WINDOWS or LINUX.\u003c/p\u003e",
			//	  "pattern": "^[a-zA-Z]*$",
			//	  "type": "string"
			//	}
			Description: "<p>The platform of the streaming image, either WINDOWS or LINUX.</p>",
			Type:        types.StringType,
			Computed:    true,
		},
		"streaming_image_id": {
			// Property: StreamingImageId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"studio_id": {
			// Property: StudioId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "\u003cp\u003eThe studioId. \u003c/p\u003e",
			//	  "type": "string"
			//	}
			Description: "<p>The studioId. </p>",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "",
			//	  "patternProperties": {
			//	    "": {
			//	      "type": "string"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::NimbleStudio::StreamingImage",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::NimbleStudio::StreamingImage").WithTerraformTypeName("awscc_nimblestudio_streaming_image")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"description":              "Description",
		"ec_2_image_id":            "Ec2ImageId",
		"encryption_configuration": "EncryptionConfiguration",
		"eula_ids":                 "EulaIds",
		"key_arn":                  "KeyArn",
		"key_type":                 "KeyType",
		"name":                     "Name",
		"owner":                    "Owner",
		"platform":                 "Platform",
		"streaming_image_id":       "StreamingImageId",
		"studio_id":                "StudioId",
		"tags":                     "Tags",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
