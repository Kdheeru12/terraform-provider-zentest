---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "Zenduty: Roles"
subcategory: ""
description: |- 
     Provides a Zenduty Role Resource. This allows Roles to be created, updated, and deleted.
  
---

# Resource : zenduty_roles 
Provides a Zenduty Role Resource. This allows Roles to be created, updated, and deleted.    
## Example Usage
```hcl
    resource "zenduty_roles" "examole_role" {
        team = ""  
        title = "Example Role"
        description = "Role Description"
    } 

```


## Argument Reference

* `team` - (Required) The unique_id of team to create the role in.
* `title` - (Required) The title of the role.
* `description` - (Required) The description of the role.
* `rank` - (Optional) The rank value of the role. ranges from  1 to 10.




<!-- schema generated by tfplugindocs -->
## Data Types

 Required fields:
- **description** (String)
- **team** (String)
- **title** (String)

 Optional fields:

- **rank** (Number) 1 to 10




