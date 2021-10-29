package zenduty

import (
	"terraform-provider-zenduty/client"

	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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

func Provider() terraform.ResourceProvider {
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
			"team": resourceTeam(),
		},
		ConfigureFunc: configureFunc(),
	}
}

func configureFunc() func(*schema.ResourceData) (interface{}, error) {
	return func(d *schema.ResourceData) (interface{}, error) {
		client := client.NewClient(d.Get("api_key").(string))
		return client, nil
	}
}
