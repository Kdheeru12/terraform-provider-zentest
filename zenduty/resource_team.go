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
		},
	}
}

func resourceTeamCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiclient := m.(*client.Client)
	newteam := &client.Team{}

	var diags diag.Diagnostics
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
	apiclient := m.(*client.Client)
	newteam := &client.Team{}
	id := d.Id()
	newteam.Unique_Id = id
	var diags diag.Diagnostics
	if v, ok := d.GetOk("name"); ok {
		newteam.Name = v.(string)

	}
	_, err := apiclient.UpdateTeam(id, newteam)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags

}

func resourceTeamDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiclient := m.(*client.Client)
	id := d.Id()
	var diags diag.Diagnostics
	err := apiclient.DeleteTeam(id)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceTeamRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiclient := m.(*client.Client)
	id := d.Id()
	var diags diag.Diagnostics

	t, err := apiclient.GetTeamById(id)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("name", t.Name)

	return diags
}
