package zenduty

import (
	"context"
	"terraform-provider-zenduty/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// func resourceTeam() *schema.Resource {
// 	return &schema.Resource{
// 		CreateContext: resourceOrderCreate,
// 		ReadContext:   resourceOrderRead,
// 		UpdateContext: resourceOrderUpdate,
// 		DeleteContext: resourceOrderDelete,
// 		Schema:        map[string]*schema.Schema{},
// 	}
// }

func resourceTeam() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTeamCreate,
		ReadContext:   resourceTeamRead,
		UpdateContext: resourceTeamUpdate,
		DeleteContext: resourceTeamDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"unique_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"account": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"creation_date": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"members": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"unique_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"team": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"user": {
							Type:     schema.TypeMap,
							Optional: true,
						},
						"joining_date": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"role": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceTeamCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiclient := m.(*client.Client)
	newteam := &client.Team{}

	var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Unable to create zenduty client",
		Detail:   "Unable to auth user for authenticated zenduty client",
	})
	if v, ok := d.GetOk("name"); ok {
		newteam.Name = v.(string)

	}
	task, err := apiclient.CreateTeam(newteam)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(task.Unique_Id)
	return diags
}

func resourceTeamUpdate(Ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	return diags

}

func resourceTeamDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	return diags
}

func resourceTeamRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiclient := m.(*client.Client)
	id := d.Id()
	var diags diag.Diagnostics

	t, err := apiclient.GetTeam(id)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("name", t.Name)
	d.Set("unique_id", t.Unique_Id)
	d.Set("creation_date", t.Creation_Date)
	d.Set("account", t.Account)
	d.Set("creation_date", t.Creation_Date)
	return diags
}
