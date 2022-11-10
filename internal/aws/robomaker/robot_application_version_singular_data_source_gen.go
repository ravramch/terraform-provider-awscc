// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package robomaker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_robomaker_robot_application_version", robotApplicationVersionDataSource)
}

// robotApplicationVersionDataSource returns the Terraform awscc_robomaker_robot_application_version data source.
// This Terraform data source corresponds to the CloudFormation AWS::RoboMaker::RobotApplicationVersion resource.
func robotApplicationVersionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"application": {
			// Property: Application
			// CloudFormation resource type schema:
			//
			//	{
			//	  "pattern": "arn:[\\w+=/,.@-]+:[\\w+=/,.@-]+:[\\w+=/,.@-]*:[0-9]*:[\\w+=,.@-]+(/[\\w+=,.@-]+)*",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"application_version": {
			// Property: ApplicationVersion
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "pattern": "arn:[\\w+=/,.@-]+:[\\w+=/,.@-]+:[\\w+=/,.@-]*:[0-9]*:[\\w+=,.@-]+(/[\\w+=,.@-]+)*",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"current_revision_id": {
			// Property: CurrentRevisionId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The revision ID of robot application.",
			//	  "maxLength": 40,
			//	  "minLength": 1,
			//	  "pattern": "[a-zA-Z0-9_.\\-]*",
			//	  "type": "string"
			//	}
			Description: "The revision ID of robot application.",
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
		Description: "Data Source schema for AWS::RoboMaker::RobotApplicationVersion",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RoboMaker::RobotApplicationVersion").WithTerraformTypeName("awscc_robomaker_robot_application_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"application":         "Application",
		"application_version": "ApplicationVersion",
		"arn":                 "Arn",
		"current_revision_id": "CurrentRevisionId",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
