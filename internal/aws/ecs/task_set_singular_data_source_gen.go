// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ecs_task_set", taskSetDataSource)
}

// taskSetDataSource returns the Terraform awscc_ecs_task_set data source.
// This Terraform data source corresponds to the CloudFormation AWS::ECS::TaskSet resource.
func taskSetDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"cluster": {
			// Property: Cluster
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.",
			//	  "type": "string"
			//	}
			Description: "The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.",
			Type:        types.StringType,
			Computed:    true,
		},
		"external_id": {
			// Property: ExternalId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "An optional non-unique tag that identifies this task set in external systems. If the task set is associated with a service discovery registry, the tasks in this task set will have the ECS_TASK_SET_EXTERNAL_ID AWS Cloud Map attribute set to the provided value. ",
			//	  "type": "string"
			//	}
			Description: "An optional non-unique tag that identifies this task set in external systems. If the task set is associated with a service discovery registry, the tasks in this task set will have the ECS_TASK_SET_EXTERNAL_ID AWS Cloud Map attribute set to the provided value. ",
			Type:        types.StringType,
			Computed:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The ID of the task set.",
			//	  "type": "string"
			//	}
			Description: "The ID of the task set.",
			Type:        types.StringType,
			Computed:    true,
		},
		"launch_type": {
			// Property: LaunchType
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The launch type that new tasks in the task set will use. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html in the Amazon Elastic Container Service Developer Guide. ",
			//	  "enum": [
			//	    "EC2",
			//	    "FARGATE"
			//	  ],
			//	  "type": "string"
			//	}
			Description: "The launch type that new tasks in the task set will use. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html in the Amazon Elastic Container Service Developer Guide. ",
			Type:        types.StringType,
			Computed:    true,
		},
		"load_balancers": {
			// Property: LoadBalancers
			// CloudFormation resource type schema:
			//
			//	{
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "A load balancer object representing the load balancer to use with the task set. The supported load balancer types are either an Application Load Balancer or a Network Load Balancer. ",
			//	    "properties": {
			//	      "ContainerName": {
			//	        "description": "The name of the container (as it appears in a container definition) to associate with the load balancer.",
			//	        "type": "string"
			//	      },
			//	      "ContainerPort": {
			//	        "description": "The port on the container to associate with the load balancer. This port must correspond to a containerPort in the task definition the tasks in the service are using. For tasks that use the EC2 launch type, the container instance they are launched on must allow ingress traffic on the hostPort of the port mapping.",
			//	        "type": "integer"
			//	      },
			//	      "LoadBalancerName": {
			//	        "description": "The name of the load balancer to associate with the Amazon ECS service or task set. A load balancer name is only specified when using a Classic Load Balancer. If you are using an Application Load Balancer or a Network Load Balancer this should be omitted.",
			//	        "type": "string"
			//	      },
			//	      "TargetGroupArn": {
			//	        "description": "The full Amazon Resource Name (ARN) of the Elastic Load Balancing target group or groups associated with a service or task set. A target group ARN is only specified when using an Application Load Balancer or Network Load Balancer. If you are using a Classic Load Balancer this should be omitted. For services using the ECS deployment controller, you can specify one or multiple target groups. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/register-multiple-targetgroups.html in the Amazon Elastic Container Service Developer Guide. For services using the CODE_DEPLOY deployment controller, you are required to define two target groups for the load balancer. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-bluegreen.html in the Amazon Elastic Container Service Developer Guide. If your service's task definition uses the awsvpc network mode (which is required for the Fargate launch type), you must choose ip as the target type, not instance, when creating your target groups because tasks that use the awsvpc network mode are associated with an elastic network interface, not an Amazon EC2 instance.",
			//	        "type": "string"
			//	      }
			//	    },
			//	    "type": "object"
			//	  },
			//	  "type": "array"
			//	}
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"container_name": {
						// Property: ContainerName
						Description: "The name of the container (as it appears in a container definition) to associate with the load balancer.",
						Type:        types.StringType,
						Computed:    true,
					},
					"container_port": {
						// Property: ContainerPort
						Description: "The port on the container to associate with the load balancer. This port must correspond to a containerPort in the task definition the tasks in the service are using. For tasks that use the EC2 launch type, the container instance they are launched on must allow ingress traffic on the hostPort of the port mapping.",
						Type:        types.Int64Type,
						Computed:    true,
					},
					"load_balancer_name": {
						// Property: LoadBalancerName
						Description: "The name of the load balancer to associate with the Amazon ECS service or task set. A load balancer name is only specified when using a Classic Load Balancer. If you are using an Application Load Balancer or a Network Load Balancer this should be omitted.",
						Type:        types.StringType,
						Computed:    true,
					},
					"target_group_arn": {
						// Property: TargetGroupArn
						Description: "The full Amazon Resource Name (ARN) of the Elastic Load Balancing target group or groups associated with a service or task set. A target group ARN is only specified when using an Application Load Balancer or Network Load Balancer. If you are using a Classic Load Balancer this should be omitted. For services using the ECS deployment controller, you can specify one or multiple target groups. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/register-multiple-targetgroups.html in the Amazon Elastic Container Service Developer Guide. For services using the CODE_DEPLOY deployment controller, you are required to define two target groups for the load balancer. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-bluegreen.html in the Amazon Elastic Container Service Developer Guide. If your service's task definition uses the awsvpc network mode (which is required for the Fargate launch type), you must choose ip as the target type, not instance, when creating your target groups because tasks that use the awsvpc network mode are associated with an elastic network interface, not an Amazon EC2 instance.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"network_configuration": {
			// Property: NetworkConfiguration
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "An object representing the network configuration for a task or service.",
			//	  "properties": {
			//	    "AwsVpcConfiguration": {
			//	      "additionalProperties": false,
			//	      "description": "The VPC subnets and security groups associated with a task. All specified subnets and security groups must be from the same VPC.",
			//	      "properties": {
			//	        "AssignPublicIp": {
			//	          "description": "Whether the task's elastic network interface receives a public IP address. The default value is DISABLED.",
			//	          "enum": [
			//	            "DISABLED",
			//	            "ENABLED"
			//	          ],
			//	          "type": "string"
			//	        },
			//	        "SecurityGroups": {
			//	          "description": "The security groups associated with the task or service. If you do not specify a security group, the default security group for the VPC is used. There is a limit of 5 security groups that can be specified per AwsVpcConfiguration.",
			//	          "items": {
			//	            "type": "string"
			//	          },
			//	          "maxItems": 5,
			//	          "type": "array"
			//	        },
			//	        "Subnets": {
			//	          "description": "The subnets associated with the task or service. There is a limit of 16 subnets that can be specified per AwsVpcConfiguration.",
			//	          "items": {
			//	            "type": "string"
			//	          },
			//	          "maxItems": 16,
			//	          "type": "array"
			//	        }
			//	      },
			//	      "required": [
			//	        "Subnets"
			//	      ],
			//	      "type": "object"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "An object representing the network configuration for a task or service.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"aws_vpc_configuration": {
						// Property: AwsVpcConfiguration
						Description: "The VPC subnets and security groups associated with a task. All specified subnets and security groups must be from the same VPC.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"assign_public_ip": {
									// Property: AssignPublicIp
									Description: "Whether the task's elastic network interface receives a public IP address. The default value is DISABLED.",
									Type:        types.StringType,
									Computed:    true,
								},
								"security_groups": {
									// Property: SecurityGroups
									Description: "The security groups associated with the task or service. If you do not specify a security group, the default security group for the VPC is used. There is a limit of 5 security groups that can be specified per AwsVpcConfiguration.",
									Type:        types.ListType{ElemType: types.StringType},
									Computed:    true,
								},
								"subnets": {
									// Property: Subnets
									Description: "The subnets associated with the task or service. There is a limit of 16 subnets that can be specified per AwsVpcConfiguration.",
									Type:        types.ListType{ElemType: types.StringType},
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"platform_version": {
			// Property: PlatformVersion
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The platform version that the tasks in the task set should use. A platform version is specified only for tasks using the Fargate launch type. If one isn't specified, the LATEST platform version is used by default.",
			//	  "type": "string"
			//	}
			Description: "The platform version that the tasks in the task set should use. A platform version is specified only for tasks using the Fargate launch type. If one isn't specified, the LATEST platform version is used by default.",
			Type:        types.StringType,
			Computed:    true,
		},
		"scale": {
			// Property: Scale
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "A floating-point percentage of the desired number of tasks to place and keep running in the task set.",
			//	  "properties": {
			//	    "Unit": {
			//	      "description": "The unit of measure for the scale value.",
			//	      "enum": [
			//	        "PERCENT"
			//	      ],
			//	      "type": "string"
			//	    },
			//	    "Value": {
			//	      "description": "The value, specified as a percent total of a service's desiredCount, to scale the task set. Accepted values are numbers between 0 and 100.",
			//	      "maximum": 100,
			//	      "minimum": 0,
			//	      "type": "number"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "A floating-point percentage of the desired number of tasks to place and keep running in the task set.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"unit": {
						// Property: Unit
						Description: "The unit of measure for the scale value.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value, specified as a percent total of a service's desiredCount, to scale the task set. Accepted values are numbers between 0 and 100.",
						Type:        types.Float64Type,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"service": {
			// Property: Service
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The short name or full Amazon Resource Name (ARN) of the service to create the task set in.",
			//	  "type": "string"
			//	}
			Description: "The short name or full Amazon Resource Name (ARN) of the service to create the task set in.",
			Type:        types.StringType,
			Computed:    true,
		},
		"service_registries": {
			// Property: ServiceRegistries
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The details of the service discovery registries to assign to this task set. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html.",
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "ContainerName": {
			//	        "description": "The container name value, already specified in the task definition, to be used for your service discovery service. If the task definition that your service task specifies uses the bridge or host network mode, you must specify a containerName and containerPort combination from the task definition. If the task definition that your service task specifies uses the awsvpc network mode and a type SRV DNS record is used, you must specify either a containerName and containerPort combination or a port value, but not both.",
			//	        "type": "string"
			//	      },
			//	      "ContainerPort": {
			//	        "description": "The port value, already specified in the task definition, to be used for your service discovery service. If the task definition your service task specifies uses the bridge or host network mode, you must specify a containerName and containerPort combination from the task definition. If the task definition your service task specifies uses the awsvpc network mode and a type SRV DNS record is used, you must specify either a containerName and containerPort combination or a port value, but not both.",
			//	        "type": "integer"
			//	      },
			//	      "Port": {
			//	        "description": "The port value used if your service discovery service specified an SRV record. This field may be used if both the awsvpc network mode and SRV records are used.",
			//	        "type": "integer"
			//	      },
			//	      "RegistryArn": {
			//	        "description": "The Amazon Resource Name (ARN) of the service registry. The currently supported service registry is AWS Cloud Map. For more information, see https://docs.aws.amazon.com/cloud-map/latest/api/API_CreateService.html",
			//	        "type": "string"
			//	      }
			//	    },
			//	    "type": "object"
			//	  },
			//	  "type": "array"
			//	}
			Description: "The details of the service discovery registries to assign to this task set. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"container_name": {
						// Property: ContainerName
						Description: "The container name value, already specified in the task definition, to be used for your service discovery service. If the task definition that your service task specifies uses the bridge or host network mode, you must specify a containerName and containerPort combination from the task definition. If the task definition that your service task specifies uses the awsvpc network mode and a type SRV DNS record is used, you must specify either a containerName and containerPort combination or a port value, but not both.",
						Type:        types.StringType,
						Computed:    true,
					},
					"container_port": {
						// Property: ContainerPort
						Description: "The port value, already specified in the task definition, to be used for your service discovery service. If the task definition your service task specifies uses the bridge or host network mode, you must specify a containerName and containerPort combination from the task definition. If the task definition your service task specifies uses the awsvpc network mode and a type SRV DNS record is used, you must specify either a containerName and containerPort combination or a port value, but not both.",
						Type:        types.Int64Type,
						Computed:    true,
					},
					"port": {
						// Property: Port
						Description: "The port value used if your service discovery service specified an SRV record. This field may be used if both the awsvpc network mode and SRV records are used.",
						Type:        types.Int64Type,
						Computed:    true,
					},
					"registry_arn": {
						// Property: RegistryArn
						Description: "The Amazon Resource Name (ARN) of the service registry. The currently supported service registry is AWS Cloud Map. For more information, see https://docs.aws.amazon.com/cloud-map/latest/api/API_CreateService.html",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"task_definition": {
			// Property: TaskDefinition
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The short name or full Amazon Resource Name (ARN) of the task definition for the tasks in the task set to use.",
			//	  "type": "string"
			//	}
			Description: "The short name or full Amazon Resource Name (ARN) of the task definition for the tasks in the task set to use.",
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
		Description: "Data Source schema for AWS::ECS::TaskSet",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECS::TaskSet").WithTerraformTypeName("awscc_ecs_task_set")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"assign_public_ip":      "AssignPublicIp",
		"aws_vpc_configuration": "AwsVpcConfiguration",
		"cluster":               "Cluster",
		"container_name":        "ContainerName",
		"container_port":        "ContainerPort",
		"external_id":           "ExternalId",
		"id":                    "Id",
		"launch_type":           "LaunchType",
		"load_balancer_name":    "LoadBalancerName",
		"load_balancers":        "LoadBalancers",
		"network_configuration": "NetworkConfiguration",
		"platform_version":      "PlatformVersion",
		"port":                  "Port",
		"registry_arn":          "RegistryArn",
		"scale":                 "Scale",
		"security_groups":       "SecurityGroups",
		"service":               "Service",
		"service_registries":    "ServiceRegistries",
		"subnets":               "Subnets",
		"target_group_arn":      "TargetGroupArn",
		"task_definition":       "TaskDefinition",
		"unit":                  "Unit",
		"value":                 "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
