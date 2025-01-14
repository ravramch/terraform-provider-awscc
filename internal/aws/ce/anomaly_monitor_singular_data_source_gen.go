// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ce

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ce_anomaly_monitor", anomalyMonitorDataSource)
}

// anomalyMonitorDataSource returns the Terraform awscc_ce_anomaly_monitor data source.
// This Terraform data source corresponds to the CloudFormation AWS::CE::AnomalyMonitor resource.
func anomalyMonitorDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreationDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The date when the monitor was created. ",
		//	  "maxLength": 40,
		//	  "minLength": 0,
		//	  "pattern": "(\\d{4}-\\d{2}-\\d{2})(T\\d{2}:\\d{2}:\\d{2}Z)?",
		//	  "type": "string"
		//	}
		"creation_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The date when the monitor was created. ",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DimensionalValueCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The value for evaluated dimensions.",
		//	  "minimum": 0,
		//	  "type": "integer"
		//	}
		"dimensional_value_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The value for evaluated dimensions.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LastEvaluatedDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The date when the monitor last evaluated for anomalies.",
		//	  "maxLength": 40,
		//	  "minLength": 0,
		//	  "pattern": "(\\d{4}-\\d{2}-\\d{2})(T\\d{2}:\\d{2}:\\d{2}Z)?|(NOT_EVALUATED_YET)",
		//	  "type": "string"
		//	}
		"last_evaluated_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The date when the monitor last evaluated for anomalies.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LastUpdatedDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The date when the monitor was last updated.",
		//	  "maxLength": 40,
		//	  "minLength": 0,
		//	  "pattern": "(\\d{4}-\\d{2}-\\d{2})(T\\d{2}:\\d{2}:\\d{2}Z)?",
		//	  "type": "string"
		//	}
		"last_updated_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The date when the monitor was last updated.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MonitorArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Monitor ARN",
		//	  "pattern": "^arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:[-a-zA-Z0-9/:_]+$",
		//	  "type": "string"
		//	}
		"monitor_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Monitor ARN",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MonitorDimension
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The dimensions to evaluate",
		//	  "enum": [
		//	    "SERVICE"
		//	  ],
		//	  "type": "string"
		//	}
		"monitor_dimension": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The dimensions to evaluate",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MonitorName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the monitor.",
		//	  "maxLength": 1024,
		//	  "minLength": 0,
		//	  "pattern": "[\\S\\s]*",
		//	  "type": "string"
		//	}
		"monitor_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the monitor.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MonitorSpecification
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"monitor_specification": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: MonitorType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "DIMENSIONAL",
		//	    "CUSTOM"
		//	  ],
		//	  "type": "string"
		//	}
		"monitor_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceTags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Tags to assign to monitor.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name for the tag.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag.",
		//	        "maxLength": 256,
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
		//	  "maxItems": 200,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"resource_tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name for the tag.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Tags to assign to monitor.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CE::AnomalyMonitor",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CE::AnomalyMonitor").WithTerraformTypeName("awscc_ce_anomaly_monitor")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"creation_date":           "CreationDate",
		"dimensional_value_count": "DimensionalValueCount",
		"key":                     "Key",
		"last_evaluated_date":     "LastEvaluatedDate",
		"last_updated_date":       "LastUpdatedDate",
		"monitor_arn":             "MonitorArn",
		"monitor_dimension":       "MonitorDimension",
		"monitor_name":            "MonitorName",
		"monitor_specification":   "MonitorSpecification",
		"monitor_type":            "MonitorType",
		"resource_tags":           "ResourceTags",
		"value":                   "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
