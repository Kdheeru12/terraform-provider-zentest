terraform {
  required_providers {
    zenduty = {
      source  = "terraform.zenduty1.com/zenduty1corp/zenduty1"
      version = ">= 1.0"

    }
  }
}

provider "zenduty" {
    token = "0e2038520fca2fbd7f3d9aace062c4fe911be36b"
  
}

# resource "zenduty_roles" "testing" {
#   team= "510fca1c-4d29-430a-a164-adcbf1e455f1"
#   title = "News updated"
#   description = "This is the description for the new Role"

# }

resource "zenduty_roles" "rank_test" {
  team= "510fca1c-4d29-430a-a164-adcbf1e455f1"
  title = "crud"
  description = "T is the description for the new Role"
  rank = 10
}

# data "zenduty_roles" "roles" {
#   team_id = "510fca1c-4d29-430a-a164-adcbf1e455f1"
# }

# output "roles" {
#   value = data.zenduty_roles.roles
  
# }

data "zenduty_incidents" "incidents"{

}
output "roles" {
  value = data.zenduty_incidents.incidents
  
}







