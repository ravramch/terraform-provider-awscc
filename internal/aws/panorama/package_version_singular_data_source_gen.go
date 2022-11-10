// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package panorama

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_panorama_package_version", packageVersionDataSource)
}

// packageVersionDataSource returns the Terraform awscc_panorama_package_version data source.
// This Terraform data source corresponds to the CloudFormation AWS::Panorama::PackageVersion resource.
func packageVersionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"is_latest_patch": {
			// Property: IsLatestPatch
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "boolean"
			//	}
			Type:     types.BoolType,
			Computed: true,
		},
		"mark_latest": {
			// Property: MarkLatest
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "boolean"
			//	}
			Type:     types.BoolType,
			Computed: true,
		},
		"owner_account": {
			// Property: OwnerAccount
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 12,
			//	  "minLength": 1,
			//	  "pattern": "^[0-9a-z\\_]+$",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"package_arn": {
			// Property: PackageArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 255,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"package_id": {
			// Property: PackageId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 255,
			//	  "minLength": 1,
			//	  "pattern": "^[a-zA-Z0-9\\-\\_\\/]+$",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"package_name": {
			// Property: PackageName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 128,
			//	  "minLength": 1,
			//	  "pattern": "^[a-zA-Z0-9\\-\\_]+$",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"package_version": {
			// Property: PackageVersion
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 255,
			//	  "minLength": 1,
			//	  "pattern": "^([0-9]+)\\.([0-9]+)$",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"patch_version": {
			// Property: PatchVersion
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 255,
			//	  "minLength": 1,
			//	  "pattern": "^[a-z0-9]+$",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"registered_time": {
			// Property: RegisteredTime
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "integer"
			//	}
			Type:     types.Int64Type,
			Computed: true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			//
			//	{
			//	  "enum": [
			//	    "REGISTER_PENDING",
			//	    "REGISTER_COMPLETED",
			//	    "FAILED",
			//	    "DELETING"
			//	  ],
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"status_description": {
			// Property: StatusDescription
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 255,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"updated_latest_patch_version": {
			// Property: UpdatedLatestPatchVersion
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 255,
			//	  "minLength": 1,
			//	  "pattern": "^[a-z0-9]+$",
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
		Description: "Data Source schema for AWS::Panorama::PackageVersion",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Panorama::PackageVersion").WithTerraformTypeName("awscc_panorama_package_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"is_latest_patch":              "IsLatestPatch",
		"mark_latest":                  "MarkLatest",
		"owner_account":                "OwnerAccount",
		"package_arn":                  "PackageArn",
		"package_id":                   "PackageId",
		"package_name":                 "PackageName",
		"package_version":              "PackageVersion",
		"patch_version":                "PatchVersion",
		"registered_time":              "RegisteredTime",
		"status":                       "Status",
		"status_description":           "StatusDescription",
		"updated_latest_patch_version": "UpdatedLatestPatchVersion",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
