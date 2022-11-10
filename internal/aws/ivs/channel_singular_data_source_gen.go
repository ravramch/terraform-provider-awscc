// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ivs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ivs_channel", channelDataSource)
}

// channelDataSource returns the Terraform awscc_ivs_channel data source.
// This Terraform data source corresponds to the CloudFormation AWS::IVS::Channel resource.
func channelDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Channel ARN is automatically generated on creation and assigned as the unique identifier.",
			//	  "maxLength": 128,
			//	  "minLength": 1,
			//	  "pattern": "^arn:aws:ivs:[a-z0-9-]+:[0-9]+:channel/[a-zA-Z0-9-]+$",
			//	  "type": "string"
			//	}
			Description: "Channel ARN is automatically generated on creation and assigned as the unique identifier.",
			Type:        types.StringType,
			Computed:    true,
		},
		"authorized": {
			// Property: Authorized
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Whether the channel is authorized.",
			//	  "type": "boolean"
			//	}
			Description: "Whether the channel is authorized.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"ingest_endpoint": {
			// Property: IngestEndpoint
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Channel ingest endpoint, part of the definition of an ingest server, used when you set up streaming software.",
			//	  "type": "string"
			//	}
			Description: "Channel ingest endpoint, part of the definition of an ingest server, used when you set up streaming software.",
			Type:        types.StringType,
			Computed:    true,
		},
		"latency_mode": {
			// Property: LatencyMode
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Channel latency mode.",
			//	  "enum": [
			//	    "NORMAL",
			//	    "LOW"
			//	  ],
			//	  "type": "string"
			//	}
			Description: "Channel latency mode.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Channel",
			//	  "maxLength": 128,
			//	  "minLength": 0,
			//	  "pattern": "^[a-zA-Z0-9-_]*$",
			//	  "type": "string"
			//	}
			Description: "Channel",
			Type:        types.StringType,
			Computed:    true,
		},
		"playback_url": {
			// Property: PlaybackUrl
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Channel Playback URL.",
			//	  "type": "string"
			//	}
			Description: "Channel Playback URL.",
			Type:        types.StringType,
			Computed:    true,
		},
		"recording_configuration_arn": {
			// Property: RecordingConfigurationArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "default": "",
			//	  "description": "Recording Configuration ARN. A value other than an empty string indicates that recording is enabled. Default: ?? (recording is disabled).",
			//	  "maxLength": 128,
			//	  "minLength": 0,
			//	  "pattern": "^$|arn:aws:ivs:[a-z0-9-]+:[0-9]+:recording-configuration/[a-zA-Z0-9-]+$",
			//	  "type": "string"
			//	}
			Description: "Recording Configuration ARN. A value other than an empty string indicates that recording is enabled. Default: ?? (recording is disabled).",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A list of key-value pairs that contain metadata for the asset model.",
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
			//	      "Value",
			//	      "Key"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "maxItems": 50,
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Description: "A list of key-value pairs that contain metadata for the asset model.",
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
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Channel type, which determines the allowable resolution and bitrate. If you exceed the allowable resolution or bitrate, the stream probably will disconnect immediately.",
			//	  "enum": [
			//	    "STANDARD",
			//	    "BASIC"
			//	  ],
			//	  "type": "string"
			//	}
			Description: "Channel type, which determines the allowable resolution and bitrate. If you exceed the allowable resolution or bitrate, the stream probably will disconnect immediately.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::IVS::Channel",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IVS::Channel").WithTerraformTypeName("awscc_ivs_channel")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                         "Arn",
		"authorized":                  "Authorized",
		"ingest_endpoint":             "IngestEndpoint",
		"key":                         "Key",
		"latency_mode":                "LatencyMode",
		"name":                        "Name",
		"playback_url":                "PlaybackUrl",
		"recording_configuration_arn": "RecordingConfigurationArn",
		"tags":                        "Tags",
		"type":                        "Type",
		"value":                       "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
