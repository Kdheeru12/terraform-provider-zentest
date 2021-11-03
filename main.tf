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

# resource "zenduty_roles" "testinge" {
#   team= "510fca1c-4d29-430a-a164-adcbf1e455f1"
#   title = "News delete"
#   description = "This is the description for the new Role"

# }



