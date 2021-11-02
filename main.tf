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

resource "zenduty_roles" "test5" {
  team = "510fca1c-4d29-430a-a164-adcbf1e455f1"
  title = "updateds"
  description = "hello"
}







# data "zenduty_teams" "all" {
#     name = "ff84fab4-dc13-4b89-a00e-5c19ce82d943"
# }


# output "all_teams" {
#   value = data.zenduty_teams.all
# }
