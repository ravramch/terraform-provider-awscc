// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package fsx

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	cctypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddResourceFactory("awscc_fsx_data_repository_association", dataRepositoryAssociationResource)
}

// dataRepositoryAssociationResource returns the Terraform awscc_fsx_data_repository_association resource.
// This Terraform resource corresponds to the CloudFormation AWS::FSx::DataRepositoryAssociation resource.
func dataRepositoryAssociationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AssociationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"association_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: BatchImportMetaDataOnCreate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A boolean flag indicating whether an import data repository task to import metadata should run after the data repository association is created. The task runs if this flag is set to ``true``.",
		//	  "type": "boolean"
		//	}
		"batch_import_meta_data_on_create": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "A boolean flag indicating whether an import data repository task to import metadata should run after the data repository association is created. The task runs if this flag is set to ``true``.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
				boolplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DataRepositoryPath
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The path to the Amazon S3 data repository that will be linked to the file system. The path can be an S3 bucket or prefix in the format ``s3://myBucket/myPrefix/``. This path specifies where in the S3 data repository files will be imported from or exported to.",
		//	  "type": "string"
		//	}
		"data_repository_path": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The path to the Amazon S3 data repository that will be linked to the file system. The path can be an S3 bucket or prefix in the format ``s3://myBucket/myPrefix/``. This path specifies where in the S3 data repository files will be imported from or exported to.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FileSystemId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the file system on which the data repository association is configured.",
		//	  "type": "string"
		//	}
		"file_system_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the file system on which the data repository association is configured.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FileSystemPath
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A path on the Amazon FSx for Lustre file system that points to a high-level directory (such as ``/ns1/``) or subdirectory (such as ``/ns1/subdir/``) that will be mapped 1-1 with ``DataRepositoryPath``. The leading forward slash in the name is required. Two data repository associations cannot have overlapping file system paths. For example, if a data repository is associated with file system path ``/ns1/``, then you cannot link another data repository with file system path ``/ns1/ns2``.\n This path specifies where in your file system files will be exported from or imported to. This file system directory can be linked to only one Amazon S3 bucket, and no other S3 bucket can be linked to the directory.\n  If you specify only a forward slash (``/``) as the file system path, you can link only one data repository to the file system. You can only specify \"/\" as the file system path for the first data repository associated with a file system.",
		//	  "type": "string"
		//	}
		"file_system_path": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A path on the Amazon FSx for Lustre file system that points to a high-level directory (such as ``/ns1/``) or subdirectory (such as ``/ns1/subdir/``) that will be mapped 1-1 with ``DataRepositoryPath``. The leading forward slash in the name is required. Two data repository associations cannot have overlapping file system paths. For example, if a data repository is associated with file system path ``/ns1/``, then you cannot link another data repository with file system path ``/ns1/ns2``.\n This path specifies where in your file system files will be exported from or imported to. This file system directory can be linked to only one Amazon S3 bucket, and no other S3 bucket can be linked to the directory.\n  If you specify only a forward slash (``/``) as the file system path, you can link only one data repository to the file system. You can only specify \"/\" as the file system path for the first data repository associated with a file system.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ImportedFileChunkSize
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. The maximum number of disks that a single file can be striped across is limited by the total number of disks that make up the file system or cache.\n The default chunk size is 1,024 MiB (1 GiB) and can go as high as 512,000 MiB (500 GiB). Amazon S3 objects have a maximum size of 5 TB.",
		//	  "type": "integer"
		//	}
		"imported_file_chunk_size": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. The maximum number of disks that a single file can be striped across is limited by the total number of disks that make up the file system or cache.\n The default chunk size is 1,024 MiB (1 GiB) and can go as high as 512,000 MiB (500 GiB). Amazon S3 objects have a maximum size of 5 TB.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ResourceARN
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"resource_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: S3
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The configuration for an Amazon S3 data repository linked to an Amazon FSx Lustre file system with a data repository association. The configuration defines which file events (new, changed, or deleted files or directories) are automatically imported from the linked data repository to the file system or automatically exported from the file system to the data repository.",
		//	  "properties": {
		//	    "AutoExportPolicy": {
		//	      "additionalProperties": false,
		//	      "description": "Describes a data repository association's automatic export policy. The ``AutoExportPolicy`` defines the types of updated objects on the file system that will be automatically exported to the data repository. As you create, modify, or delete files, Amazon FSx for Lustre automatically exports the defined changes asynchronously once your application finishes modifying the file.\n The ``AutoExportPolicy`` is only supported on Amazon FSx for Lustre file systems with a data repository association.",
		//	      "properties": {
		//	        "Events": {
		//	          "description": "The ``AutoExportPolicy`` can have the following event values:\n  +   ``NEW`` - New files and directories are automatically exported to the data repository as they are added to the file system.\n  +   ``CHANGED`` - Changes to files and directories on the file system are automatically exported to the data repository.\n  +   ``DELETED`` - Files and directories are automatically deleted on the data repository when they are deleted on the file system.\n  \n You can define any combination of event types for your ``AutoExportPolicy``.",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "enum": [
		//	              "NEW",
		//	              "CHANGED",
		//	              "DELETED"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "maxItems": 3,
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        }
		//	      },
		//	      "required": [
		//	        "Events"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "AutoImportPolicy": {
		//	      "additionalProperties": false,
		//	      "description": "Describes the data repository association's automatic import policy. The AutoImportPolicy defines how Amazon FSx keeps your file metadata and directory listings up to date by importing changes to your Amazon FSx for Lustre file system as you modify objects in a linked S3 bucket.\n The ``AutoImportPolicy`` is only supported on Amazon FSx for Lustre file systems with a data repository association.",
		//	      "properties": {
		//	        "Events": {
		//	          "description": "The ``AutoImportPolicy`` can have the following event values:\n  +   ``NEW`` - Amazon FSx automatically imports metadata of files added to the linked S3 bucket that do not currently exist in the FSx file system.\n  +   ``CHANGED`` - Amazon FSx automatically updates file metadata and invalidates existing file content on the file system as files change in the data repository.\n  +   ``DELETED`` - Amazon FSx automatically deletes files on the file system as corresponding files are deleted in the data repository.\n  \n You can define any combination of event types for your ``AutoImportPolicy``.",
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "enum": [
		//	              "NEW",
		//	              "CHANGED",
		//	              "DELETED"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "maxItems": 3,
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        }
		//	      },
		//	      "required": [
		//	        "Events"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"s3": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AutoExportPolicy
				"auto_export_policy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Events
						"events": schema.SetAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "The ``AutoExportPolicy`` can have the following event values:\n  +   ``NEW`` - New files and directories are automatically exported to the data repository as they are added to the file system.\n  +   ``CHANGED`` - Changes to files and directories on the file system are automatically exported to the data repository.\n  +   ``DELETED`` - Files and directories are automatically deleted on the data repository when they are deleted on the file system.\n  \n You can define any combination of event types for your ``AutoExportPolicy``.",
							Required:    true,
							Validators: []validator.Set{ /*START VALIDATORS*/
								setvalidator.SizeAtMost(3),
								setvalidator.ValueStringsAre(
									stringvalidator.OneOf(
										"NEW",
										"CHANGED",
										"DELETED",
									),
								),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Describes a data repository association's automatic export policy. The ``AutoExportPolicy`` defines the types of updated objects on the file system that will be automatically exported to the data repository. As you create, modify, or delete files, Amazon FSx for Lustre automatically exports the defined changes asynchronously once your application finishes modifying the file.\n The ``AutoExportPolicy`` is only supported on Amazon FSx for Lustre file systems with a data repository association.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: AutoImportPolicy
				"auto_import_policy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Events
						"events": schema.SetAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "The ``AutoImportPolicy`` can have the following event values:\n  +   ``NEW`` - Amazon FSx automatically imports metadata of files added to the linked S3 bucket that do not currently exist in the FSx file system.\n  +   ``CHANGED`` - Amazon FSx automatically updates file metadata and invalidates existing file content on the file system as files change in the data repository.\n  +   ``DELETED`` - Amazon FSx automatically deletes files on the file system as corresponding files are deleted in the data repository.\n  \n You can define any combination of event types for your ``AutoImportPolicy``.",
							Required:    true,
							Validators: []validator.Set{ /*START VALIDATORS*/
								setvalidator.SizeAtMost(3),
								setvalidator.ValueStringsAre(
									stringvalidator.OneOf(
										"NEW",
										"CHANGED",
										"DELETED",
									),
								),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Describes the data repository association's automatic import policy. The AutoImportPolicy defines how Amazon FSx keeps your file metadata and directory listings up to date by importing changes to your Amazon FSx for Lustre file system as you modify objects in a linked S3 bucket.\n The ``AutoImportPolicy`` is only supported on Amazon FSx for Lustre file systems with a data repository association.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The configuration for an Amazon S3 data repository linked to an Amazon FSx Lustre file system with a data repository association. The configuration defines which file events (new, changed, or deleted files or directories) are automatically imported from the linked data repository to the file system or automatically exported from the file system to the data repository.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.\n For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Specifies a key-value pair for a resource tag.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "A value that specifies the ``TagKey``, the name of the tag. Tag keys must be unique for the resource to which they are attached.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "A value that specifies the ``TagValue``, the value assigned to the corresponding tag key. Tag values can be null and don't have to be unique in a tag set. For example, you can have a key-value pair in a tag set of ``finances : April`` and also of ``payroll : April``.",
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
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A value that specifies the ``TagKey``, the name of the tag. Tag keys must be unique for the resource to which they are attached.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A value that specifies the ``TagValue``, the value assigned to the corresponding tag key. Tag values can be null and don't have to be unique in a tag set. For example, you can have a key-value pair in a tag set of ``finances : April`` and also of ``payroll : April``.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			CustomType:  cctypes.MultisetType,
			Description: "An array of key-value pairs to apply to this resource.\n For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
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
		Description: "Creates an Amazon FSx for Lustre data repository association (DRA). A data repository association is a link between a directory on the file system and an Amazon S3 bucket or prefix. You can have a maximum of 8 data repository associations on a file system. Data repository associations are supported on all FSx for Lustre 2.12 and newer file systems, excluding ``scratch_1`` deployment type. \n Each data repository association must have a unique Amazon FSx file system directory and a unique S3 bucket or prefix associated with it. You can configure a data repository association for automatic import only, for automatic export only, or for both. To learn more about linking a data repository to your file system, see [Linking your file system to an S3 bucket](https://docs.aws.amazon.com/fsx/latest/LustreGuide/create-dra-linked-data-repo.html).",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::FSx::DataRepositoryAssociation").WithTerraformTypeName("awscc_fsx_data_repository_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"association_id":                   "AssociationId",
		"auto_export_policy":               "AutoExportPolicy",
		"auto_import_policy":               "AutoImportPolicy",
		"batch_import_meta_data_on_create": "BatchImportMetaDataOnCreate",
		"data_repository_path":             "DataRepositoryPath",
		"events":                           "Events",
		"file_system_id":                   "FileSystemId",
		"file_system_path":                 "FileSystemPath",
		"imported_file_chunk_size":         "ImportedFileChunkSize",
		"key":                              "Key",
		"resource_arn":                     "ResourceARN",
		"s3":                               "S3",
		"tags":                             "Tags",
		"value":                            "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(120).WithDeleteTimeoutInMinutes(180)

	opts = opts.WithUpdateTimeoutInMinutes(180)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
