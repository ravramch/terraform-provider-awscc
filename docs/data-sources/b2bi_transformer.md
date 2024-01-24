---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_b2bi_transformer Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::B2BI::Transformer
---

# awscc_b2bi_transformer (Data Source)

Data Source schema for AWS::B2BI::Transformer



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `created_at` (String)
- `edi_type` (Attributes) (see [below for nested schema](#nestedatt--edi_type))
- `file_format` (String)
- `mapping_template` (String)
- `modified_at` (String)
- `name` (String)
- `sample_document` (String)
- `status` (String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))
- `transformer_arn` (String)
- `transformer_id` (String)

<a id="nestedatt--edi_type"></a>
### Nested Schema for `edi_type`

Read-Only:

- `x12_details` (Attributes) (see [below for nested schema](#nestedatt--edi_type--x12_details))

<a id="nestedatt--edi_type--x12_details"></a>
### Nested Schema for `edi_type.x12_details`

Read-Only:

- `transaction_set` (String)
- `version` (String)



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)