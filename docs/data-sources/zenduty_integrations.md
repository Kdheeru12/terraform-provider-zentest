---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "zenduty_integrations Data Source - terraform-provider-zentest"
subcategory: ""
description: |-
  
---

# zenduty_integrations (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **service_id** (String)
- **team_id** (String)

### Optional

- **id** (String) The ID of this resource.
- **integration_id** (String)

### Read-Only

- **results** (List of Object) (see [below for nested schema](#nestedatt--results))

<a id="nestedatt--results"></a>
### Nested Schema for `results`

Read-Only:

- **application** (String)
- **application_reference** (Map of String)
- **create_incidents_for** (Number)
- **created_by** (String)
- **creation_date** (String)
- **default_urgency** (Number)
- **integration_key** (String)
- **integration_type** (Number)
- **name** (String)
- **service** (String)
- **summary** (String)
- **unique_id** (String)


