package zenduty

import (
	"context"
	"terraform-provider-zenduty/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
// Provider -
// func Provider() *schema.Provider {
// 	return &schema.Provider{
// 		Schema: map[string]*schema.Schema{
// 			"token": &schema.Schema{
// 				Type:        schema.TypeString,
// 				Optional:    true,
// 				DefaultFunc: schema.EnvDefaultFunc("ddd", nil),
// 			},
// 		},
// 		ResourcesMap: map[string]*schema.Resource{
// 			"team": resourceTeam(),
// 		},
// 		ConfigureFunc: providerConfigure(),
// 	}
// }

// // re
// func providerConfigure(d *schema.ResourceData) (interface{}, error) {
// 	token := d.Get("token").(string)

// 	// Warning or errors can be collected in a slice type

// 	c := client.NewClient(token)

// 	return c, nil
// }

// func Provider() terraform.ResourceProvider {
// 	return &schema.Provider{
// 		Schema: map[string]*schema.Schema{
// 			"token": &schema.Schema{
// 				Type:        schema.TypeString,
// 				Description: "Your Todoist API key",
// 				Required:    true,
// 				DefaultFunc: schema.EnvDefaultFunc("TODOIST_API_KEY", nil),
// 			},
// 		},
// 		ResourcesMap: map[string]*schema.Resource{
// 			"team": resourceTeam(),
// 		},
// 		DataSourcesMap: map[string]*schema.Resource{},
// 		ConfigureFunc:  configureFunc(),
// 	}
// }

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
			"zenduty_team":  resourceTeam(),
			"zenduty_roles": resourceRoles(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"zenduty_teams": resourceTeamData(),
			"zenduty_roles": dataSourceRoles(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

// func configureFunc() func(*schema.ResourceData) (interface{}, error) {
// 	return func(d *schema.ResourceData) (interface{}, error) {
// 		client := client.NewClient(d.Get("token").(string))
// 		return client, nil
// 	}
// }

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	token := d.Get("token").(string)
	// Warning or errors can be collected in a slice type
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
