terraform {
  required_providers {
    zenduty = {
      source  = "terraform.zenduty1.com/zenduty1corp/zenduty1"
      version = ">= 1.0"

    }
  }
}

provider "zenduty" {
    token = "3b44da5b6cc076b459c45a6256b2e0e8b03af91c"
  
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

resource "zenduty_roles" "test" {
  team= "dd518f4d-dbce-4ad2-b5be-ceff597c67f8"
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

# data "zenduty_incidents" "incidents"{

# }
# output "roles" {
#   value = data.zenduty_incidents.incidents
  
# }

# data "zenduty_teams" "teams"{
#   team_id = "f212685d-4268-49c5-9009-9069f58cfcbd"
# }
# output "teams" {
#   value = data.zenduty_teams.teams
  
# }

# resource "zenduty_services" "service1" {
#   name = "terraform delete"
#   team_id = "dd518f4d-dbce-4ad2-b5be-ceff597c67f8"
#   description = "This is the description for the new Service"
#   escalation_policy = "86d2d574-0027-4593-acfa-7830c9a69dd6"

# }

# data "zenduty_services" "services"{
#   team_id = "dd518f4d-dbce-4ad2-b5be-ceff597c67f8"
#   id = "e6618036-56a7-4e84-980e-f15ac027d556"
# }

# output "services" {
#   value = data.zenduty_services.services
  
# }

# resource "zenduty_integrations" "integration" {
#   team_id = "dd518f4d-dbce-4ad2-b5be-ceff597c67f8"
#   service_id = "e6618036-56a7-4e84-980e-f15ac027d556"
#   application = "c9acbca3-75e0-44b5-a2c9-891918dd128b"
#   name = "terraformd"
#   summary = "This is the summary for the new Integration"
# }

# data "zenduty_integrations" "integrations"{
#   team_id = "dd518f4d-dbce-4ad2-b5be-ceff597c67f8"
#   service_id = "e6618036-56a7-4e84-980e-f15ac027d556"
#   integration_id = "2c20a676-fba2-40e7-aa83-8235d7972b8d"
# }

# output "integrations" {
#   value = data.zenduty_integrations.integrations
  
# }

resource "zenduty_schedules" "name" {
  name = "terraform delete"
  team_id = "dd518f4d-dbce-4ad2-b5be-ceff597c67f8"
  time_zone = "Asia/Kolkata"
  
}


# data "zenduty_schedules" "schedules"{
#   team_id = "dd518f4d-dbce-4ad2-b5be-ceff597c67f8"
#   schedule_id = "a5707ecf-5768-4123-8f05-1d6cf487b7b0"
# }
# output "schedules" {
#   value = data.zenduty_schedules.schedules
  
# }



# chain 

resource "zenduty_team" "team1"{
  name = "terraform chain"

}

resource "zenduty_services" "service1" {
  name = "terraform chain"
  team_id = zenduty_team.team1.id
  description = "This is the description for the new Service"
  escalation_policy = "5bb8cc73-13f9-4b8d-b96a-7a7d77e8c0e2"

}
resource "zenduty_integrations" "integration" {
  team_id = zenduty_team.team1.id
  service_id = zenduty_services.service1.id
  application = "c9acbca3-75e0-44b5-a2c9-891918dd128b"
  name = "terraform chain"
  summary = "This is the summary for the new Integration"
}


resource "zenduty_schedules" "schedule1" {
  name = "terraform chain"
  team_id = zenduty_team.team1.id
  time_zone = "Asia/Kolkata"
  
}

resource "zenduty_roles" "role1" {
  team= zenduty_team.team1.id
  title = "terraform chain"
  description = "T is the description for the new Role"
  rank = 8
}

# data "zenduty_teams" "team1"{
#   team_id = zenduty_team.team1.id
# }
# output "teams" {
#   value = data.zenduty_teams.team1
# }

# data "zenduty_services" "services"{
#   team_id = zenduty_team.team1.id
#   id = zenduty_services.service1.id
# }

# output "services" {
#   value = data.zenduty_services.services
  
# }

# data "zenduty_integrations" "integrations"{
#   team_id = zenduty_team.team1.id
#   service_id = zenduty_services.service1.id
#   integration_id = zenduty_integrations.integration.id
# }

# output "integrations" {
#   value = data.zenduty_integrations.integrations
  
# }


# data "zenduty_schedules" "schedules"{
#   team_id = zenduty_team.team1.id
#   schedule_id = zenduty_schedules.schedule1.id
# }
# output "schedules" {
#   value = data.zenduty_schedules.schedules
  
# }

# data "zenduty_roles" "roles" {
#   team_id = zenduty_team.team1.id
# }

# output "roles" {
#   value = data.zenduty_roles.roles
  
# }


# resource "zenduty_esp" "esp1" {
#   name = "terraform_delete"
#   team_id = zenduty_team.team1.id
#   summary = "This is the summary for the new ESP"
#   description = "This is the description for the new ESP"
#   rules {
#     delay = 0
#     targets {
#       target_type = 2
#       target_id ="50012040-37db-4594-a268-a"
#     }
#     position = 0
#     unique_id = "295958ad-945b-40b4-abb6-56af4f98f626"
#   }
#   repeat_policy = 3
#   move_to_next = true
  
# }










