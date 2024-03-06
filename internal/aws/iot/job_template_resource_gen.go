// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package iot

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddResourceFactory("awscc_iot_job_template", jobTemplateResource)
}

// jobTemplateResource returns the Terraform awscc_iot_job_template resource.
// This Terraform resource corresponds to the CloudFormation AWS::IoT::JobTemplate resource.
func jobTemplateResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AbortConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The criteria that determine when and how a job abort takes place.",
		//	  "properties": {
		//	    "CriteriaList": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "The criteria that determine when and how a job abort takes place.",
		//	        "properties": {
		//	          "Action": {
		//	            "description": "The type of job action to take to initiate the job abort.",
		//	            "enum": [
		//	              "CANCEL"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "FailureType": {
		//	            "description": "The type of job execution failures that can initiate a job abort.",
		//	            "enum": [
		//	              "FAILED",
		//	              "REJECTED",
		//	              "TIMED_OUT",
		//	              "ALL"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "MinNumberOfExecutedThings": {
		//	            "description": "The minimum number of things which must receive job execution notifications before the job can be aborted.",
		//	            "minimum": 1,
		//	            "type": "integer"
		//	          },
		//	          "ThresholdPercentage": {
		//	            "description": "The minimum percentage of job execution failures that must occur to initiate the job abort.",
		//	            "maximum": 100,
		//	            "type": "number"
		//	          }
		//	        },
		//	        "required": [
		//	          "Action",
		//	          "FailureType",
		//	          "MinNumberOfExecutedThings",
		//	          "ThresholdPercentage"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "minItems": 1,
		//	      "type": "array"
		//	    }
		//	  },
		//	  "required": [
		//	    "CriteriaList"
		//	  ],
		//	  "type": "object"
		//	}
		"abort_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CriteriaList
				"criteria_list": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Action
							"action": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The type of job action to take to initiate the job abort.",
								Required:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"CANCEL",
									),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: FailureType
							"failure_type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The type of job execution failures that can initiate a job abort.",
								Required:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"FAILED",
										"REJECTED",
										"TIMED_OUT",
										"ALL",
									),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: MinNumberOfExecutedThings
							"min_number_of_executed_things": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "The minimum number of things which must receive job execution notifications before the job can be aborted.",
								Required:    true,
								Validators: []validator.Int64{ /*START VALIDATORS*/
									int64validator.AtLeast(1),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: ThresholdPercentage
							"threshold_percentage": schema.Float64Attribute{ /*START ATTRIBUTE*/
								Description: "The minimum percentage of job execution failures that must occur to initiate the job abort.",
								Required:    true,
								Validators: []validator.Float64{ /*START VALIDATORS*/
									float64validator.AtMost(100.000000),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					CustomType: cctypes.MultisetType,
					Required:   true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.SizeAtLeast(1),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The criteria that determine when and how a job abort takes place.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// AbortConfig is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description of the Job Template.",
		//	  "maxLength": 2028,
		//	  "pattern": "[^\\p{C}]+",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description of the Job Template.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(2028),
				stringvalidator.RegexMatches(regexp.MustCompile("[^\\p{C}]+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DestinationPackageVersions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "description": "Specifies target package version ARNs for a software update job.",
		//	    "maxLength": 1600,
		//	    "minLength": 1,
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"destination_package_versions": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			CustomType:  cctypes.MultisetType,
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.ValueStringsAre(
					stringvalidator.LengthBetween(1, 1600),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// DestinationPackageVersions is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Document
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The job document. Required if you don't specify a value for documentSource.",
		//	  "maxLength": 32768,
		//	  "type": "string"
		//	}
		"document": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The job document. Required if you don't specify a value for documentSource.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(32768),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// Document is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: DocumentSource
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An S3 link to the job document to use in the template. Required if you don't specify a value for document.",
		//	  "maxLength": 1350,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"document_source": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An S3 link to the job document to use in the template. Required if you don't specify a value for document.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 1350),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// DocumentSource is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: JobArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Optional for copying a JobTemplate from a pre-existing Job configuration.",
		//	  "type": "string"
		//	}
		"job_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Optional for copying a JobTemplate from a pre-existing Job configuration.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// JobArn is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: JobExecutionsRetryConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "RetryCriteriaList": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Specifies how many times a failure type should be retried.",
		//	        "properties": {
		//	          "FailureType": {
		//	            "enum": [
		//	              "FAILED",
		//	              "TIMED_OUT",
		//	              "ALL"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "NumberOfRetries": {
		//	            "maximum": 10,
		//	            "minimum": 0,
		//	            "type": "integer"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "maxItems": 2,
		//	      "minItems": 1,
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"job_executions_retry_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: RetryCriteriaList
				"retry_criteria_list": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: FailureType
							"failure_type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"FAILED",
										"TIMED_OUT",
										"ALL",
									),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: NumberOfRetries
							"number_of_retries": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.Int64{ /*START VALIDATORS*/
									int64validator.Between(0, 10),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
									int64planmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					CustomType: cctypes.MultisetType,
					Optional:   true,
					Computed:   true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.SizeBetween(1, 2),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// JobExecutionsRetryConfig is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: JobExecutionsRolloutConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Allows you to create a staged rollout of a job.",
		//	  "properties": {
		//	    "ExponentialRolloutRate": {
		//	      "additionalProperties": false,
		//	      "description": "The rate of increase for a job rollout. This parameter allows you to define an exponential rate for a job rollout.",
		//	      "properties": {
		//	        "BaseRatePerMinute": {
		//	          "description": "The minimum number of things that will be notified of a pending job, per minute at the start of job rollout. This parameter allows you to define the initial rate of rollout.",
		//	          "minimum": 1,
		//	          "type": "integer"
		//	        },
		//	        "IncrementFactor": {
		//	          "description": "The exponential factor to increase the rate of rollout for a job.",
		//	          "maximum": 5,
		//	          "minimum": 1,
		//	          "type": "number"
		//	        },
		//	        "RateIncreaseCriteria": {
		//	          "additionalProperties": false,
		//	          "description": "The criteria to initiate the increase in rate of rollout for a job.",
		//	          "properties": {
		//	            "NumberOfNotifiedThings": {
		//	              "minimum": 1,
		//	              "type": "integer"
		//	            },
		//	            "NumberOfSucceededThings": {
		//	              "minimum": 1,
		//	              "type": "integer"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "required": [
		//	        "BaseRatePerMinute",
		//	        "IncrementFactor",
		//	        "RateIncreaseCriteria"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "MaximumPerMinute": {
		//	      "description": "The maximum number of things that will be notified of a pending job, per minute. This parameter allows you to create a staged rollout.",
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"job_executions_rollout_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ExponentialRolloutRate
				"exponential_rollout_rate": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: BaseRatePerMinute
						"base_rate_per_minute": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Description: "The minimum number of things that will be notified of a pending job, per minute at the start of job rollout. This parameter allows you to define the initial rate of rollout.",
							Required:    true,
							Validators: []validator.Int64{ /*START VALIDATORS*/
								int64validator.AtLeast(1),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: IncrementFactor
						"increment_factor": schema.Float64Attribute{ /*START ATTRIBUTE*/
							Description: "The exponential factor to increase the rate of rollout for a job.",
							Required:    true,
							Validators: []validator.Float64{ /*START VALIDATORS*/
								float64validator.Between(1.000000, 5.000000),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: RateIncreaseCriteria
						"rate_increase_criteria": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: NumberOfNotifiedThings
								"number_of_notified_things": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									Validators: []validator.Int64{ /*START VALIDATORS*/
										int64validator.AtLeast(1),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
										int64planmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: NumberOfSucceededThings
								"number_of_succeeded_things": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									Validators: []validator.Int64{ /*START VALIDATORS*/
										int64validator.AtLeast(1),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
										int64planmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "The criteria to initiate the increase in rate of rollout for a job.",
							Required:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The rate of increase for a job rollout. This parameter allows you to define an exponential rate for a job rollout.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: MaximumPerMinute
				"maximum_per_minute": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The maximum number of things that will be notified of a pending job, per minute. This parameter allows you to create a staged rollout.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.AtLeast(1),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Allows you to create a staged rollout of a job.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// JobExecutionsRolloutConfig is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: JobTemplateId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z0-9_-]+",
		//	  "type": "string"
		//	}
		"job_template_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 64),
				stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9_-]+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MaintenanceWindows
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Specifies a start time and duration for a scheduled Job.",
		//	    "properties": {
		//	      "DurationInMinutes": {
		//	        "maximum": 1430,
		//	        "minimum": 1,
		//	        "type": "integer"
		//	      },
		//	      "StartTime": {
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"maintenance_windows": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: DurationInMinutes
					"duration_in_minutes": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.Int64{ /*START VALIDATORS*/
							int64validator.Between(1, 1430),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
							int64planmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: StartTime
					"start_time": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType: cctypes.MultisetType,
			Optional:   true,
			Computed:   true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// MaintenanceWindows is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: PresignedUrlConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Configuration for pre-signed S3 URLs.",
		//	  "properties": {
		//	    "ExpiresInSec": {
		//	      "description": "How number (in seconds) pre-signed URLs are valid.",
		//	      "maximum": 3600,
		//	      "minimum": 60,
		//	      "type": "integer"
		//	    },
		//	    "RoleArn": {
		//	      "description": "The ARN of an IAM role that grants grants permission to download files from the S3 bucket where the job data/updates are stored. The role must also grant permission for IoT to download the files.",
		//	      "maxLength": 2048,
		//	      "minLength": 20,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "RoleArn"
		//	  ],
		//	  "type": "object"
		//	}
		"presigned_url_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ExpiresInSec
				"expires_in_sec": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "How number (in seconds) pre-signed URLs are valid.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.Between(60, 3600),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: RoleArn
				"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of an IAM role that grants grants permission to download files from the S3 bucket where the job data/updates are stored. The role must also grant permission for IoT to download the files.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(20, 2048),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Configuration for pre-signed S3 URLs.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// PresignedUrlConfig is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Metadata that can be used to manage the JobTemplate.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The tag's key.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The tag's value.",
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
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag's key.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag's value.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Metadata that can be used to manage the JobTemplate.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeAtMost(50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
				setplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// Tags is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: TimeoutConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Specifies the amount of time each device has to finish its execution of the job.",
		//	  "properties": {
		//	    "InProgressTimeoutInMinutes": {
		//	      "description": "Specifies the amount of time, in minutes, this device has to finish execution of this job.",
		//	      "maximum": 10080,
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "required": [
		//	    "InProgressTimeoutInMinutes"
		//	  ],
		//	  "type": "object"
		//	}
		"timeout_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: InProgressTimeoutInMinutes
				"in_progress_timeout_in_minutes": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "Specifies the amount of time, in minutes, this device has to finish execution of this job.",
					Required:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.Between(1, 10080),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Specifies the amount of time each device has to finish its execution of the job.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// TimeoutConfig is a write-only property.
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Job templates enable you to preconfigure jobs so that you can deploy them to multiple sets of target devices.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::JobTemplate").WithTerraformTypeName("awscc_iot_job_template")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"abort_config":                   "AbortConfig",
		"action":                         "Action",
		"arn":                            "Arn",
		"base_rate_per_minute":           "BaseRatePerMinute",
		"criteria_list":                  "CriteriaList",
		"description":                    "Description",
		"destination_package_versions":   "DestinationPackageVersions",
		"document":                       "Document",
		"document_source":                "DocumentSource",
		"duration_in_minutes":            "DurationInMinutes",
		"expires_in_sec":                 "ExpiresInSec",
		"exponential_rollout_rate":       "ExponentialRolloutRate",
		"failure_type":                   "FailureType",
		"in_progress_timeout_in_minutes": "InProgressTimeoutInMinutes",
		"increment_factor":               "IncrementFactor",
		"job_arn":                        "JobArn",
		"job_executions_retry_config":    "JobExecutionsRetryConfig",
		"job_executions_rollout_config":  "JobExecutionsRolloutConfig",
		"job_template_id":                "JobTemplateId",
		"key":                            "Key",
		"maintenance_windows":            "MaintenanceWindows",
		"maximum_per_minute":             "MaximumPerMinute",
		"min_number_of_executed_things":  "MinNumberOfExecutedThings",
		"number_of_notified_things":      "NumberOfNotifiedThings",
		"number_of_retries":              "NumberOfRetries",
		"number_of_succeeded_things":     "NumberOfSucceededThings",
		"presigned_url_config":           "PresignedUrlConfig",
		"rate_increase_criteria":         "RateIncreaseCriteria",
		"retry_criteria_list":            "RetryCriteriaList",
		"role_arn":                       "RoleArn",
		"start_time":                     "StartTime",
		"tags":                           "Tags",
		"threshold_percentage":           "ThresholdPercentage",
		"timeout_config":                 "TimeoutConfig",
		"value":                          "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/JobArn",
		"/properties/Tags",
		"/properties/Document",
		"/properties/DocumentSource",
		"/properties/TimeoutConfig",
		"/properties/JobExecutionsRolloutConfig",
		"/properties/AbortConfig",
		"/properties/PresignedUrlConfig",
		"/properties/DestinationPackageVersions",
		"/properties/JobExecutionsRetryConfig",
		"/properties/MaintenanceWindows",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
