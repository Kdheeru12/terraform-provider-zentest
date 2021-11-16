package zenduty

import (
	"context"
	"terraform-provider-zenduty/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIntegrations() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIntegrationCreate,
		UpdateContext: resourceIntegrationUpdate,
		DeleteContext: resourceIntegrationDelete,
		ReadContext:   resourceIntegrationRead,
		Schema: map[string]*schema.Schema{
			"application": {
				Type:     schema.TypeString,
				Required: true,
			},
			"team_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"service_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"summary": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceIntegrationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiclient, _ := m.(*Config).Client()

	newIntegration := &client.Integration{}
	var diags diag.Diagnostics
	team_id := d.Get("team_id").(string)
	service_id := d.Get("service_id").(string)
	summary := d.Get("summary").(string)
	if summary != "" {
		newIntegration.Summary = summary
	}
	if v, ok := d.GetOk("name"); ok {
		newIntegration.Name = v.(string)
	}
	if v, ok := d.GetOk("application"); ok {
		newIntegration.Application = v.(string)
	}

	integration, err := apiclient.Integrations.CreateIntegration(team_id, service_id, newIntegration)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(integration.Unique_Id)
	return diags
}

func resourceIntegrationUpdate(Ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	return diags
}

func resourceIntegrationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	return diags
}

func resourceIntegrationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	id := d.Id()
	team_id := d.Get("team_id").(string)
	service_id := d.Get("service_id").(string)
	apiclient, _ := m.(*Config).Client()

	integration, err := apiclient.Integrations.GetIntegrationByID(team_id, service_id, id)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("name", integration.Name)
	return diags
}
