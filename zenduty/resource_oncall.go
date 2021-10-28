package zenduty

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTodoistTask() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"escalation_policy": &schema.Schema{
				Type: schema.TypeMap,
			},
			"team": &schema.Schema{
				Type: schema.TypeMap,
			},
			"users": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": &schema.Schema{
							Type: schema.TypeString,
						},
						"first_name": &schema.Schema{
							Type: schema.TypeString,
						},
						"last_name": &schema.Schema{
							Type: schema.TypeString,
						},
					},
				},
			},
		},
	}
}
