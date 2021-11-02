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

func resourceRoles() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRoleCreate,
		UpdateContext: resourceRoleUpdate,
		DeleteContext: resourceRoleDelete,
		ReadContext:   resourceRoleRead,
		Schema: map[string]*schema.Schema{
			"team": {
				Type:     schema.TypeString,
				Required: true,
			},
			"unique_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"creation_date": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rank": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceRoleCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiclient := m.(*client.Client)
	newrole := &client.Roles{}

	var diags diag.Diagnostics
	// diags = append(diags, diag.Diagnostic{
	// 	Severity: diag.Error,
	// 	Summary:  "Unable to create zenduty client",
	// 	Detail:   "Unable to auth user for authenticated zenduty client",
	// })
	if v, ok := d.GetOk("team"); ok {
		newrole.Team = v.(string)

	}
	if v, ok := d.GetOk("description"); ok {
		newrole.Description = v.(string)

	}
	if v, ok := d.GetOk("title"); ok {
		newrole.Title = v.(string)
	}

	role, err := apiclient.CreateRole(newrole.Team, newrole)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(role.Unique_Id)
	d.Set("team", role.Team)
	return diags
}

func resourceRoleUpdate(Ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiclient := m.(*client.Client)
	newrole := &client.Roles{}
	team_id := d.Get("team").(string)
	id := d.Id()
	newrole.Unique_Id = id
	var diags diag.Diagnostics
	if v, ok := d.GetOk("description"); ok {
		newrole.Description = v.(string)

	}
	if v, ok := d.GetOk("title"); ok {
		newrole.Title = v.(string)
	}
	// diags = append(diags, diag.Diagnostic{
	// 	Severity: diag.Error,
	// 	Summary:  "Unable to update zenduty client",
	// 	Detail:   newrole.Title + " " + newrole.Description + " " + newrole.Unique_Id,
	// })
	_, err := apiclient.UpdateRoles(team_id, newrole)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceRoleDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	apiclient := m.(*client.Client)
	id := d.Id()
	team_id := d.Get("team").(string)
	var diags diag.Diagnostics

	err := apiclient.DeleteRole(team_id, id)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceRoleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil

}