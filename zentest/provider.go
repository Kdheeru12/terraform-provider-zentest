package zentest

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Your Zenduty API key",
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ZENDUTY_API_KEY", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"zenduty_teams":        resourceTeam(),
			"zenduty_roles":        resourceRoles(),
			"zenduty_services":     resourceServices(),
			"zenduty_integrations": resourceIntegrations(),
			"zenduty_schedules":    resourceSchedules(),
			"zenduty_esp":          resourceEsp(),
			"zenduty_incidents":    resourceIncidents(),
			"zenduty_invite":       resourceInvite(),
			"zenduty_member":       resourceMembers(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"zenduty_teams":        dataSourceTeams(),
			"zenduty_roles":        dataSourceRoles(),
			"zenduty_incidents":    dataSourceIncidents(),
			"zenduty_services":     dataSourceServices(),
			"zenduty_integrations": dataSourceIntegrations(),
			"zenduty_schedules":    dataSourceSchedules(),
			"zenduty_esp":          dataSourceEsp(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	token := d.Get("token").(string)
	var diags diag.Diagnostics
	if token != "" {
		client := Config{
			Token: token,
		}
		return &client, diags
	}
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Unable to create zenduty client provider",
		Detail:   "Unable to auth user for authenticated zenduty client",
	})

	return nil, diags
}
