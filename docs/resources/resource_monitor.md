---
page_title: "snowflake_resource_monitor Resource - terraform-provider-snowflake"
subcategory: ""
description: |-
  
---

# snowflake_resource_monitor (Resource)



## Example Usage

```terraform
resource "snowflake_resource_monitor" "monitor" {
  name         = "monitor"
  credit_quota = 100

  frequency       = "DAILY"
  start_timestamp = "2020-12-07 00:00"
  end_timestamp   = "2021-12-07 00:00"

  notify_triggers            = [40, 50]
  suspend_triggers           = 50
  suspend_immediate_triggers = 90

  notify_users = ["USERONE", "USERTWO"]
}
```

-> **Note** Instead of using fully_qualified_name, you can reference objects managed outside Terraform by constructing a correct ID, consult [identifiers guide](https://registry.terraform.io/providers/Snowflake-Labs/snowflake/latest/docs/guides/identifiers#new-computed-fully-qualified-name-field-in-resources).
<!-- TODO(SNOW-1634854): include an example showing both methods-->

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Identifier for the resource monitor; must be unique for your account. Due to technical limitations (read more [here](https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/main/docs/technical-documentation/identifiers_rework_design_decisions.md#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `(`, `)`, `"`

### Optional

- `credit_quota` (Number) The number of credits allocated to the resource monitor per frequency interval. When total usage for all warehouses assigned to the monitor reaches this number for the current frequency interval, the resource monitor is considered to be at 100% of quota.
- `end_timestamp` (String) The date and time when the resource monitor suspends the assigned warehouses.
- `frequency` (String) The frequency interval at which the credit usage resets to 0. Valid values are (case-insensitive): `MONTHLY` | `DAILY` | `WEEKLY` | `YEARLY` | `NEVER`. If you set a `frequency` for a resource monitor, you must also set `start_timestamp`. If you specify `NEVER` for the frequency, the credit usage for the warehouse does not reset. After removing this field from the config, the previously set value will be preserved on the Snowflake side, not the default value. That's due to Snowflake limitation and the lack of unset functionality for this parameter.
- `notify_triggers` (Set of Number) Specifies a list of percentages of the credit quota. After reaching any of the values the users passed in the notify_users field will be notified (to receive the notification they should have notifications enabled). Values over 100 are supported.
- `notify_users` (Set of String) Specifies the list of users (their identifiers) to receive email notifications on resource monitors.
- `start_timestamp` (String) The date and time when the resource monitor starts monitoring credit usage for the assigned warehouses. If you set a `start_timestamp` for a resource monitor, you must also set `frequency`.  After removing this field from the config, the previously set value will be preserved on the Snowflake side, not the default value. That's due to Snowflake limitation and the lack of unset functionality for this parameter.
- `suspend_immediate_trigger` (Number) Represents a numeric value specified as a percentage of the credit quota. Values over 100 are supported. After reaching this value, all assigned warehouses immediately cancel any currently running queries or statements. In addition, this action sends a notification to all users who have enabled notifications for themselves.
- `suspend_trigger` (Number) Represents a numeric value specified as a percentage of the credit quota. Values over 100 are supported. After reaching this value, all assigned warehouses while allowing currently running queries to complete will be suspended. No new queries can be executed by the warehouses until the credit quota for the resource monitor is increased. In addition, this action sends a notification to all users who have enabled notifications for themselves.

### Read-Only

- `fully_qualified_name` (String) Fully qualified name of the resource. For more information, see [object name resolution](https://docs.snowflake.com/en/sql-reference/name-resolution).
- `id` (String) The ID of this resource.
- `show_output` (List of Object) Outputs the result of `SHOW RESOURCE MONITORS` for the given resource monitor. (see [below for nested schema](#nestedatt--show_output))

<a id="nestedatt--show_output"></a>
### Nested Schema for `show_output`

Read-Only:

- `comment` (String)
- `created_on` (String)
- `credit_quota` (Number)
- `end_time` (String)
- `frequency` (String)
- `level` (String)
- `name` (String)
- `owner` (String)
- `remaining_credits` (Number)
- `start_time` (String)
- `suspend_at` (Number)
- `suspend_immediate_at` (Number)
- `used_credits` (Number)

## Import

Import is supported using the following syntax:

```shell
# format is the resource monitor name
terraform import snowflake_resource_monitor.example 'resourceMonitorName'
```
