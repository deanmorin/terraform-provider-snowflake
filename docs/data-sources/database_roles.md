---
page_title: "snowflake_database_roles Data Source - terraform-provider-snowflake"
subcategory: ""
description: |-
  
---

!> **V1 release candidate** This data source was reworked and is a release candidate for the V1. We do not expect significant changes in it before the V1. We will welcome any feedback and adjust the data source if needed. Any errors reported will be resolved with a higher priority. We encourage checking this data source out before the V1 release. Please follow the [migration guide](https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/main/MIGRATION_GUIDE.md#v0920--v0930) to use it.

# snowflake_database_roles (Data Source)



## Example Usage

```terraform
data "snowflake_database_roles" "db_roles" {
  database = "MYDB"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `in_database` (String) The database from which to return the database roles from.

### Optional

- `like` (String) Filters the output with **case-insensitive** pattern, with support for SQL wildcard characters (`%` and `_`).
- `limit` (Block List, Max: 1) Limits the number of rows returned. If the `limit.from` is set, then the limit wll start from the first element matched by the expression. The expression is only used to match with the first element, later on the elements are not matched by the prefix, but you can enforce a certain pattern with `starts_with` or `like`. (see [below for nested schema](#nestedblock--limit))

### Read-Only

- `database_roles` (List of Object) Holds the aggregated output of all database role details queries. (see [below for nested schema](#nestedatt--database_roles))
- `id` (String) The ID of this resource.

<a id="nestedblock--limit"></a>
### Nested Schema for `limit`

Required:

- `rows` (Number) The maximum number of rows to return.

Optional:

- `from` (String) Specifies a **case-sensitive** pattern that is used to match object name. After the first match, the limit on the number of rows will be applied.


<a id="nestedatt--database_roles"></a>
### Nested Schema for `database_roles`

Read-Only:

- `show_output` (List of Object) (see [below for nested schema](#nestedobjatt--database_roles--show_output))

<a id="nestedobjatt--database_roles--show_output"></a>
### Nested Schema for `database_roles.show_output`

Read-Only:

- `comment` (String)
- `created_on` (String)
- `database_name` (String)
- `granted_database_roles` (Number)
- `granted_to_database_roles` (Number)
- `granted_to_roles` (Number)
- `is_current` (Boolean)
- `is_default` (Boolean)
- `is_inherited` (Boolean)
- `name` (String)
- `owner` (String)
- `owner_role_type` (String)
