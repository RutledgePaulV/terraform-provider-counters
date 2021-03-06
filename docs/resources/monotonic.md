---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "counters_monotonic Resource - terraform-provider-counters"
subcategory: ""
description: |-
  A monotonic counter which increments according to the configured triggers.
---

# counters_monotonic (Resource)

A monotonic counter which increments according to the configured triggers.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.
- **initial_value** (Number) The initial value to use for the counter.
- **max_history** (Number) How many versions (max) should this resource store?
- **step** (Number) The amount used to increment / decrement the counter on each revision.
- **triggers** (Map of String) A map of strings that will cause a change to the counter when any of the values change.

### Read-Only

- **history** (List of Object) A list of counter values that this resource has produced. (see [below for nested schema](#nestedatt--history))
- **value** (Number) The current value of the counter.

<a id="nestedatt--history"></a>
### Nested Schema for `history`

Read-Only:

- **triggers** (Map of String)
- **value** (Number)


