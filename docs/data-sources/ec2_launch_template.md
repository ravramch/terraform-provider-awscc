---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_launch_template Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::EC2::LaunchTemplate
---

# awscc_ec2_launch_template (Data Source)

Data Source schema for AWS::EC2::LaunchTemplate



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `default_version_number` (String) The default version of the launch template
- `latest_version_number` (String) The latest version of the launch template
- `launch_template_data` (Attributes) The information for the launch template. (see [below for nested schema](#nestedatt--launch_template_data))
- `launch_template_id` (String) LaunchTemplate ID generated by service
- `launch_template_name` (String) A name for the launch template.
- `tag_specifications` (Attributes List) The tags to apply to the launch template on creation. (see [below for nested schema](#nestedatt--tag_specifications))
- `version_description` (String) A description for the first version of the launch template.

<a id="nestedatt--launch_template_data"></a>
### Nested Schema for `launch_template_data`

Read-Only:

- `block_device_mappings` (Attributes List) The block device mapping. (see [below for nested schema](#nestedatt--launch_template_data--block_device_mappings))
- `capacity_reservation_specification` (Attributes) Specifies an instance's Capacity Reservation targeting option. (see [below for nested schema](#nestedatt--launch_template_data--capacity_reservation_specification))
- `cpu_options` (Attributes) specifies the CPU options for an instance. (see [below for nested schema](#nestedatt--launch_template_data--cpu_options))
- `credit_specification` (Attributes) The user data to make available to the instance. (see [below for nested schema](#nestedatt--launch_template_data--credit_specification))
- `disable_api_stop` (Boolean) Indicates whether to enable the instance for stop protection.
- `disable_api_termination` (Boolean) If you set this parameter to true, you can't terminate the instance using the Amazon EC2 console, CLI, or API.
- `ebs_optimized` (Boolean) Indicates whether the instance is optimized for Amazon EBS I/O.
- `elastic_gpu_specifications` (Attributes List) An elastic GPU to associate with the instance. (see [below for nested schema](#nestedatt--launch_template_data--elastic_gpu_specifications))
- `elastic_inference_accelerators` (Attributes List) The elastic inference accelerator for the instance. (see [below for nested schema](#nestedatt--launch_template_data--elastic_inference_accelerators))
- `enclave_options` (Attributes) Indicates whether the instance is enabled for AWS Nitro Enclaves. (see [below for nested schema](#nestedatt--launch_template_data--enclave_options))
- `hibernation_options` (Attributes) Specifies whether your instance is configured for hibernation. (see [below for nested schema](#nestedatt--launch_template_data--hibernation_options))
- `iam_instance_profile` (Attributes) Specifies an IAM instance profile, which is a container for an IAM role for your instance. (see [below for nested schema](#nestedatt--launch_template_data--iam_instance_profile))
- `image_id` (String) The ID of the AMI. Alternatively, you can specify a Systems Manager parameter, which will resolve to an AMI ID on launch.
- `instance_initiated_shutdown_behavior` (String) Indicates whether an instance stops or terminates when you initiate shutdown from the instance (using the operating system command for system shutdown).
- `instance_market_options` (Attributes) The market (purchasing) option for the instances. (see [below for nested schema](#nestedatt--launch_template_data--instance_market_options))
- `instance_requirements` (Attributes) The attributes for the instance types. (see [below for nested schema](#nestedatt--launch_template_data--instance_requirements))
- `instance_type` (String)
- `kernel_id` (String) The ID of the kernel.
- `key_name` (String) The name of the EC2 key pair
- `license_specifications` (Attributes List) The license configurations. (see [below for nested schema](#nestedatt--launch_template_data--license_specifications))
- `maintenance_options` (Attributes) The maintenance options of your instance. (see [below for nested schema](#nestedatt--launch_template_data--maintenance_options))
- `metadata_options` (Attributes) The metadata options for the instance. (see [below for nested schema](#nestedatt--launch_template_data--metadata_options))
- `monitoring` (Attributes) Specifies whether detailed monitoring is enabled for an instance. (see [below for nested schema](#nestedatt--launch_template_data--monitoring))
- `network_interfaces` (Attributes List) If you specify a network interface, you must specify any security groups and subnets as part of the network interface. (see [below for nested schema](#nestedatt--launch_template_data--network_interfaces))
- `placement` (Attributes) Specifies the placement of an instance. (see [below for nested schema](#nestedatt--launch_template_data--placement))
- `private_dns_name_options` (Attributes) Describes the options for instance hostnames. (see [below for nested schema](#nestedatt--launch_template_data--private_dns_name_options))
- `ram_disk_id` (String)
- `security_group_ids` (List of String) One or more security group IDs.
- `security_groups` (List of String) One or more security group names.
- `tag_specifications` (Attributes List) The tags to apply to the resources that are created during instance launch. (see [below for nested schema](#nestedatt--launch_template_data--tag_specifications))
- `user_data` (String) The user data to make available to the instance.

<a id="nestedatt--launch_template_data--block_device_mappings"></a>
### Nested Schema for `launch_template_data.block_device_mappings`

Read-Only:

- `device_name` (String) The user data to make available to the instance.
- `ebs` (Attributes) Parameters for a block device for an EBS volume in an Amazon EC2 launch template. (see [below for nested schema](#nestedatt--launch_template_data--block_device_mappings--ebs))
- `no_device` (String) To omit the device from the block device mapping, specify an empty string.
- `virtual_name` (String) The virtual device name (ephemeralN).

<a id="nestedatt--launch_template_data--block_device_mappings--ebs"></a>
### Nested Schema for `launch_template_data.block_device_mappings.ebs`

Read-Only:

- `delete_on_termination` (Boolean) Indicates whether the EBS volume is deleted on instance termination.
- `encrypted` (Boolean) Indicates whether the EBS volume is encrypted. Encrypted volumes can only be attached to instances that support Amazon EBS encryption. If you are creating a volume from a snapshot, you can't specify an encryption value.
- `iops` (Number) The number of I/O operations per second (IOPS).
- `kms_key_id` (String) The ARN of the symmetric AWS Key Management Service (AWS KMS) CMK used for encryption.
- `snapshot_id` (String) The ID of the snapshot.
- `throughput` (Number) The throughput to provision for a gp3 volume, with a maximum of 1,000 MiB/s.
- `volume_size` (Number) The size of the volume, in GiBs. You must specify either a snapshot ID or a volume size.
- `volume_type` (String) The volume type.



<a id="nestedatt--launch_template_data--capacity_reservation_specification"></a>
### Nested Schema for `launch_template_data.capacity_reservation_specification`

Read-Only:

- `capacity_reservation_preference` (String) Indicates the instance's Capacity Reservation preferences.
- `capacity_reservation_target` (Attributes) Specifies a target Capacity Reservation. (see [below for nested schema](#nestedatt--launch_template_data--capacity_reservation_specification--capacity_reservation_target))

<a id="nestedatt--launch_template_data--capacity_reservation_specification--capacity_reservation_target"></a>
### Nested Schema for `launch_template_data.capacity_reservation_specification.capacity_reservation_target`

Read-Only:

- `capacity_reservation_id` (String) The ID of the Capacity Reservation in which to run the instance.
- `capacity_reservation_resource_group_arn` (String) The ARN of the Capacity Reservation resource group in which to run the instance.



<a id="nestedatt--launch_template_data--cpu_options"></a>
### Nested Schema for `launch_template_data.cpu_options`

Read-Only:

- `amd_sev_snp` (String) Indicates whether to enable the instance for AMD SEV-SNP. AMD SEV-SNP is supported with M6a, R6a, and C6a instance types only.
- `core_count` (Number) The number of CPU cores for the instance.
- `threads_per_core` (Number) The number of threads per CPU core. To disable multithreading for the instance, specify a value of 1. Otherwise, specify the default value of 2.


<a id="nestedatt--launch_template_data--credit_specification"></a>
### Nested Schema for `launch_template_data.credit_specification`

Read-Only:

- `cpu_credits` (String) The user data to make available to the instance.


<a id="nestedatt--launch_template_data--elastic_gpu_specifications"></a>
### Nested Schema for `launch_template_data.elastic_gpu_specifications`

Read-Only:

- `type` (String) The type of Elastic Graphics accelerator.


<a id="nestedatt--launch_template_data--elastic_inference_accelerators"></a>
### Nested Schema for `launch_template_data.elastic_inference_accelerators`

Read-Only:

- `count` (Number) The number of elastic inference accelerators to attach to the instance.
- `type` (String) The type of elastic inference accelerator.


<a id="nestedatt--launch_template_data--enclave_options"></a>
### Nested Schema for `launch_template_data.enclave_options`

Read-Only:

- `enabled` (Boolean) If this parameter is set to true, the instance is enabled for AWS Nitro Enclaves; otherwise, it is not enabled for AWS Nitro Enclaves.


<a id="nestedatt--launch_template_data--hibernation_options"></a>
### Nested Schema for `launch_template_data.hibernation_options`

Read-Only:

- `configured` (Boolean) TIf you set this parameter to true, the instance is enabled for hibernation.


<a id="nestedatt--launch_template_data--iam_instance_profile"></a>
### Nested Schema for `launch_template_data.iam_instance_profile`

Read-Only:

- `arn` (String) The Amazon Resource Name (ARN) of the instance profile.
- `name` (String) The name of the instance profile.


<a id="nestedatt--launch_template_data--instance_market_options"></a>
### Nested Schema for `launch_template_data.instance_market_options`

Read-Only:

- `market_type` (String) The market type.
- `spot_options` (Attributes) Specifies options for Spot Instances. (see [below for nested schema](#nestedatt--launch_template_data--instance_market_options--spot_options))

<a id="nestedatt--launch_template_data--instance_market_options--spot_options"></a>
### Nested Schema for `launch_template_data.instance_market_options.spot_options`

Read-Only:

- `block_duration_minutes` (Number) Deprecated
- `instance_interruption_behavior` (String) The behavior when a Spot Instance is interrupted. The default is terminate.
- `max_price` (String) The maximum hourly price you're willing to pay for the Spot Instances.
- `spot_instance_type` (String) The Spot Instance request type.
- `valid_until` (String) The end date of the request, in UTC format (YYYY-MM-DDTHH:MM:SSZ). Supported only for persistent requests.



<a id="nestedatt--launch_template_data--instance_requirements"></a>
### Nested Schema for `launch_template_data.instance_requirements`

Read-Only:

- `accelerator_count` (Attributes) The minimum and maximum number of accelerators (GPUs, FPGAs, or AWS Inferential chips) on an instance. (see [below for nested schema](#nestedatt--launch_template_data--instance_requirements--accelerator_count))
- `accelerator_manufacturers` (List of String) Indicates whether instance types must have accelerators by specific manufacturers.
- `accelerator_names` (List of String) The accelerators that must be on the instance type.
- `accelerator_total_memory_mi_b` (Attributes) The minimum and maximum amount of total accelerator memory, in MiB. (see [below for nested schema](#nestedatt--launch_template_data--instance_requirements--accelerator_total_memory_mi_b))
- `accelerator_types` (List of String) The accelerator types that must be on the instance type.
- `allowed_instance_types` (List of String) The instance types to apply your specified attributes against.
- `bare_metal` (String) Indicates whether bare metal instance types must be included, excluded, or required.
- `baseline_ebs_bandwidth_mbps` (Attributes) The minimum and maximum baseline bandwidth to Amazon EBS, in Mbps. (see [below for nested schema](#nestedatt--launch_template_data--instance_requirements--baseline_ebs_bandwidth_mbps))
- `burstable_performance` (String)
- `cpu_manufacturers` (List of String) The CPU manufacturers to include.
- `excluded_instance_types` (List of String) The instance types to exclude.
- `instance_generations` (List of String) Indicates whether current or previous generation instance types are included.
- `local_storage` (String) The user data to make available to the instance.
- `local_storage_types` (List of String) The type of local storage that is required.
- `memory_gi_b_per_v_cpu` (Attributes) The minimum and maximum amount of memory per vCPU, in GiB. (see [below for nested schema](#nestedatt--launch_template_data--instance_requirements--memory_gi_b_per_v_cpu))
- `memory_mi_b` (Attributes) The minimum and maximum amount of memory, in MiB. (see [below for nested schema](#nestedatt--launch_template_data--instance_requirements--memory_mi_b))
- `network_bandwidth_gbps` (Attributes) The minimum and maximum amount of network bandwidth, in gigabits per second (Gbps). (see [below for nested schema](#nestedatt--launch_template_data--instance_requirements--network_bandwidth_gbps))
- `network_interface_count` (Attributes) TThe minimum and maximum number of network interfaces. (see [below for nested schema](#nestedatt--launch_template_data--instance_requirements--network_interface_count))
- `on_demand_max_price_percentage_over_lowest_price` (Number) The price protection threshold for On-Demand Instances.
- `require_hibernate_support` (Boolean) Indicates whether instance types must support hibernation for On-Demand Instances.
- `spot_max_price_percentage_over_lowest_price` (Number) The price protection threshold for Spot Instances.
- `total_local_storage_gb` (Attributes) The minimum and maximum amount of total local storage, in GB. (see [below for nested schema](#nestedatt--launch_template_data--instance_requirements--total_local_storage_gb))
- `v_cpu_count` (Attributes) The minimum and maximum number of vCPUs. (see [below for nested schema](#nestedatt--launch_template_data--instance_requirements--v_cpu_count))

<a id="nestedatt--launch_template_data--instance_requirements--accelerator_count"></a>
### Nested Schema for `launch_template_data.instance_requirements.accelerator_count`

Read-Only:

- `max` (Number) The maximum number of accelerators.
- `min` (Number) The minimum number of accelerators.


<a id="nestedatt--launch_template_data--instance_requirements--accelerator_total_memory_mi_b"></a>
### Nested Schema for `launch_template_data.instance_requirements.accelerator_total_memory_mi_b`

Read-Only:

- `max` (Number) The maximum amount of accelerator memory, in MiB.
- `min` (Number) The minimum amount of accelerator memory, in MiB.


<a id="nestedatt--launch_template_data--instance_requirements--baseline_ebs_bandwidth_mbps"></a>
### Nested Schema for `launch_template_data.instance_requirements.baseline_ebs_bandwidth_mbps`

Read-Only:

- `max` (Number) The maximum baseline bandwidth, in Mbps.
- `min` (Number) The minimum baseline bandwidth, in Mbps.


<a id="nestedatt--launch_template_data--instance_requirements--memory_gi_b_per_v_cpu"></a>
### Nested Schema for `launch_template_data.instance_requirements.memory_gi_b_per_v_cpu`

Read-Only:

- `max` (Number) The maximum amount of memory per vCPU, in GiB.
- `min` (Number) TThe minimum amount of memory per vCPU, in GiB.


<a id="nestedatt--launch_template_data--instance_requirements--memory_mi_b"></a>
### Nested Schema for `launch_template_data.instance_requirements.memory_mi_b`

Read-Only:

- `max` (Number) The maximum amount of memory, in MiB.
- `min` (Number) The minimum amount of memory, in MiB.


<a id="nestedatt--launch_template_data--instance_requirements--network_bandwidth_gbps"></a>
### Nested Schema for `launch_template_data.instance_requirements.network_bandwidth_gbps`

Read-Only:

- `max` (Number) The maximum amount of network bandwidth, in Gbps.
- `min` (Number) The minimum amount of network bandwidth, in Gbps.


<a id="nestedatt--launch_template_data--instance_requirements--network_interface_count"></a>
### Nested Schema for `launch_template_data.instance_requirements.network_interface_count`

Read-Only:

- `max` (Number)
- `min` (Number)


<a id="nestedatt--launch_template_data--instance_requirements--total_local_storage_gb"></a>
### Nested Schema for `launch_template_data.instance_requirements.total_local_storage_gb`

Read-Only:

- `max` (Number)
- `min` (Number)


<a id="nestedatt--launch_template_data--instance_requirements--v_cpu_count"></a>
### Nested Schema for `launch_template_data.instance_requirements.v_cpu_count`

Read-Only:

- `max` (Number) The maximum number of vCPUs.
- `min` (Number) The minimum number of vCPUs.



<a id="nestedatt--launch_template_data--license_specifications"></a>
### Nested Schema for `launch_template_data.license_specifications`

Read-Only:

- `license_configuration_arn` (String) The Amazon Resource Name (ARN) of the license configuration.


<a id="nestedatt--launch_template_data--maintenance_options"></a>
### Nested Schema for `launch_template_data.maintenance_options`

Read-Only:

- `auto_recovery` (String) Disables the automatic recovery behavior of your instance or sets it to default.


<a id="nestedatt--launch_template_data--metadata_options"></a>
### Nested Schema for `launch_template_data.metadata_options`

Read-Only:

- `http_endpoint` (String) Enables or disables the HTTP metadata endpoint on your instances. If the parameter is not specified, the default state is enabled.
- `http_protocol_ipv_6` (String) Enables or disables the IPv6 endpoint for the instance metadata service.
- `http_put_response_hop_limit` (Number) The desired HTTP PUT response hop limit for instance metadata requests. The larger the number, the further instance metadata requests can travel.
- `http_tokens` (String) IMDSv2 uses token-backed sessions.
- `instance_metadata_tags` (String) Set to enabled to allow access to instance tags from the instance metadata.


<a id="nestedatt--launch_template_data--monitoring"></a>
### Nested Schema for `launch_template_data.monitoring`

Read-Only:

- `enabled` (Boolean) Specify true to enable detailed monitoring.


<a id="nestedatt--launch_template_data--network_interfaces"></a>
### Nested Schema for `launch_template_data.network_interfaces`

Read-Only:

- `associate_carrier_ip_address` (Boolean) Indicates whether to associate a Carrier IP address with eth0 for a new network interface.
- `associate_public_ip_address` (Boolean) Associates a public IPv4 address with eth0 for a new network interface.
- `delete_on_termination` (Boolean) Indicates whether the network interface is deleted when the instance is terminated.
- `description` (String) A description for the network interface.
- `device_index` (Number) The device index for the network interface attachment.
- `groups` (List of String) The IDs of one or more security groups.
- `interface_type` (String) The type of network interface.
- `ipv_4_prefix_count` (Number) The number of IPv4 prefixes to be automatically assigned to the network interface.
- `ipv_4_prefixes` (Attributes List) One or more IPv4 prefixes to be assigned to the network interface. (see [below for nested schema](#nestedatt--launch_template_data--network_interfaces--ipv_4_prefixes))
- `ipv_6_address_count` (Number) The number of IPv6 addresses to assign to a network interface.
- `ipv_6_addresses` (Attributes List) One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. (see [below for nested schema](#nestedatt--launch_template_data--network_interfaces--ipv_6_addresses))
- `ipv_6_prefix_count` (Number) The number of IPv6 prefixes to be automatically assigned to the network interface.
- `ipv_6_prefixes` (Attributes List) One or more IPv6 prefixes to be assigned to the network interface. (see [below for nested schema](#nestedatt--launch_template_data--network_interfaces--ipv_6_prefixes))
- `network_card_index` (Number) The index of the network card.
- `network_interface_id` (String) The ID of the network interface.
- `private_ip_address` (String) The primary private IPv4 address of the network interface.
- `private_ip_addresses` (Attributes List) One or more private IPv4 addresses. (see [below for nested schema](#nestedatt--launch_template_data--network_interfaces--private_ip_addresses))
- `secondary_private_ip_address_count` (Number) The number of secondary private IPv4 addresses to assign to a network interface.
- `subnet_id` (String) The ID of the subnet for the network interface.

<a id="nestedatt--launch_template_data--network_interfaces--ipv_4_prefixes"></a>
### Nested Schema for `launch_template_data.network_interfaces.ipv_4_prefixes`

Read-Only:

- `ipv_4_prefix` (String) The IPv4 prefix.


<a id="nestedatt--launch_template_data--network_interfaces--ipv_6_addresses"></a>
### Nested Schema for `launch_template_data.network_interfaces.ipv_6_addresses`

Read-Only:

- `ipv_6_address` (String)


<a id="nestedatt--launch_template_data--network_interfaces--ipv_6_prefixes"></a>
### Nested Schema for `launch_template_data.network_interfaces.ipv_6_prefixes`

Read-Only:

- `ipv_6_prefix` (String)


<a id="nestedatt--launch_template_data--network_interfaces--private_ip_addresses"></a>
### Nested Schema for `launch_template_data.network_interfaces.private_ip_addresses`

Read-Only:

- `primary` (Boolean) Indicates whether the private IPv4 address is the primary private IPv4 address. Only one IPv4 address can be designated as primary.
- `private_ip_address` (String) The private IPv4 address.



<a id="nestedatt--launch_template_data--placement"></a>
### Nested Schema for `launch_template_data.placement`

Read-Only:

- `affinity` (String) The affinity setting for an instance on a Dedicated Host.
- `availability_zone` (String) The Availability Zone for the instance.
- `group_id` (String) The Group Id of a placement group. You must specify the Placement Group Group Id to launch an instance in a shared placement group.
- `group_name` (String) The name of the placement group for the instance.
- `host_id` (String) The ID of the Dedicated Host for the instance.
- `host_resource_group_arn` (String) The ARN of the host resource group in which to launch the instances. If you specify a host resource group ARN, omit the Tenancy parameter or set it to host.
- `partition_number` (Number) The number of the partition the instance should launch in. Valid only if the placement group strategy is set to partition.
- `spread_domain` (String) Reserved for future use.
- `tenancy` (String) The tenancy of the instance (if the instance is running in a VPC). An instance with a tenancy of dedicated runs on single-tenant hardware.


<a id="nestedatt--launch_template_data--private_dns_name_options"></a>
### Nested Schema for `launch_template_data.private_dns_name_options`

Read-Only:

- `enable_resource_name_dns_a_record` (Boolean) Indicates whether to respond to DNS queries for instance hostnames with DNS A records.
- `enable_resource_name_dns_aaaa_record` (Boolean) Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records.
- `hostname_type` (String) The type of hostname for EC2 instances.


<a id="nestedatt--launch_template_data--tag_specifications"></a>
### Nested Schema for `launch_template_data.tag_specifications`

Read-Only:

- `resource_type` (String) The type of resource to tag.
- `tags` (Attributes List) The tags for the resource. (see [below for nested schema](#nestedatt--launch_template_data--tag_specifications--tags))

<a id="nestedatt--launch_template_data--tag_specifications--tags"></a>
### Nested Schema for `launch_template_data.tag_specifications.tags`

Read-Only:

- `key` (String)
- `value` (String)




<a id="nestedatt--tag_specifications"></a>
### Nested Schema for `tag_specifications`

Read-Only:

- `resource_type` (String) The type of resource to tag.
- `tags` (Attributes List) The tags for the resource. (see [below for nested schema](#nestedatt--tag_specifications--tags))

<a id="nestedatt--tag_specifications--tags"></a>
### Nested Schema for `tag_specifications.tags`

Read-Only:

- `key` (String)
- `value` (String)