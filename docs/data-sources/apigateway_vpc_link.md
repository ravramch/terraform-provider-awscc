---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_apigateway_vpc_link Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::ApiGateway::VpcLink
---

# awscc_apigateway_vpc_link (Data Source)

Data Source schema for AWS::ApiGateway::VpcLink



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `description` (String)
- `name` (String)
- `tags` (Attributes Set) An array of arbitrary tags (key-value pairs) to associate with the VPC link. (see [below for nested schema](#nestedatt--tags))
- `target_arns` (List of String)
- `vpc_link_id` (String)

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)
