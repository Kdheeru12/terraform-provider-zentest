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

resource "zenduty_team" "team" {
    name = "testprovider"
}

# data "zenduty_teams" "all" {
#     name = "ff84fab4-dc13-4b89-a00e-5c19ce82d943"
# }


# output "all_teams" {
#   value = data.zenduty_teams.all
# }
