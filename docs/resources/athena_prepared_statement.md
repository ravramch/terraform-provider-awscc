---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_athena_prepared_statement Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::Athena::PreparedStatement
---

# awscc_athena_prepared_statement (Resource)

Resource schema for AWS::Athena::PreparedStatement



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **query_statement** (String) The query string for the prepared statement.
- **statement_name** (String) The name of the prepared statement.
- **work_group** (String) The name of the workgroup to which the prepared statement belongs.

### Optional

- **description** (String) The description of the prepared statement.

### Read-Only

- **id** (String) Uniquely identifies the resource.

