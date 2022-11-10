// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package location

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_location_tracker_consumer", trackerConsumerDataSource)
}

// trackerConsumerDataSource returns the Terraform awscc_location_tracker_consumer data source.
// This Terraform data source corresponds to the CloudFormation AWS::Location::TrackerConsumer resource.
func trackerConsumerDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"consumer_arn": {
			// Property: ConsumerArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 1600,
			//	  "pattern": "^arn(:[a-z0-9]+([.-][a-z0-9]+)*){2}(:([a-z0-9]+([.-][a-z0-9]+)*)?){2}:([^/].*)?$",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"tracker_name": {
			// Property: TrackerName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 100,
			//	  "minLength": 1,
			//	  "pattern": "^[-._\\w]+$",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Location::TrackerConsumer",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Location::TrackerConsumer").WithTerraformTypeName("awscc_location_tracker_consumer")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"consumer_arn": "ConsumerArn",
		"tracker_name": "TrackerName",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
