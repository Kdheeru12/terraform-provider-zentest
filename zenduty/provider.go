package zenduty

import (
	"context"
	"terraform-provider-zenduty/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Your Todoist API key",
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TODOIST_API_KEY", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"zenduty_team":         resourceTeam(),
			"zenduty_roles":        resourceRoles(),
			"zenduty_services":     resourceServices(),
			"zenduty_integrations": resourceIntegrations(),
			"zenduty_schedules":    resourceSchedules(),
			"zenduty_esp":          resourceEsp(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"zenduty_teams":        dataSourceTeams(),
			"zenduty_roles":        dataSourceRoles(),
			"zenduty_incidents":    dataSourceIncidents(),
			"zenduty_services":     dataSourceServices(),
			"zenduty_integrations": dataSourceIntegrations(),
			"zenduty_schedules":    dataSourceSchedules(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	token := d.Get("token").(string)
	var diags diag.Diagnostics
	if token != "" {
		client := client.NewClient(token)
		return client, diags
	}
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Unable to create zenduty client provider",
		Detail:   "Unable to auth user for authenticated zenduty client",
	})

	return nil, diags
}
