package zenduty

import (
	"context"
	"terraform-provider-zenduty/client"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEsp() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceEspsRead,
		Schema: map[string]*schema.Schema{
			"team_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"esp_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"escalation_policies": {
				Type:        schema.TypeList,
				Description: "List of Escalation policies",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"summary": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"team": {
							Type:     schema.TypeString,
							Required: true,
						},
						"unique_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"rules": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"delay": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"position": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"unique_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"targets": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"target_type": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"target_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"repeat_policy": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"move_to_next": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"global_ep": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceEspsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiclient := m.(*client.Client)

	team_id := d.Get("team_id").(string)
	esp_id := d.Get("esp_id").(string)

	var diags diag.Diagnostics
	if esp_id != "" {
		esp, err := apiclient.GetEscalationPolicyById(team_id, esp_id)
		if err != nil {
			return diag.FromErr(err)
		}
		items := make([]map[string]interface{}, 1)
		item := make(map[string]interface{})
		item["unique_id"] = esp.Unique_Id
		item["name"] = esp.Name
		item["summary"] = esp.Summary
		item["description"] = esp.Description
		item["team"] = esp.Team
		rules := make([]map[string]interface{}, len(esp.Rules))
		for j, rule := range esp.Rules {
			rules[j] = map[string]interface{}{
				"delay":     rule.Delay,
				"position":  rule.Position,
				"unique_id": rule.Unique_Id,
			}
			if rule.Targets != nil {
				targets := make([]map[string]interface{}, len(rule.Targets))
				for k, target := range rule.Targets {
					targets[k] = map[string]interface{}{
						"target_type": target.Target_type,
						"target_id":   target.Target_id,
					}
				}
				rules[j]["targets"] = targets
			}

		}
		item["rules"] = rules

		items[0] = item

		if err := d.Set("escalation_policies", items); err != nil {
			return diag.FromErr(err)
		}
		d.SetId(time.Now().String())

		return diags
	} else {

		esps, err := apiclient.GetEscalationPolicy(team_id)
		if err != nil {
			return diag.FromErr(err)
		}
		items := make([]map[string]interface{}, len(esps))
		for i, esp := range esps {
			item := make(map[string]interface{})
			item["unique_id"] = esp.Unique_Id
			item["name"] = esp.Name
			item["summary"] = esp.Summary
			item["description"] = esp.Description
			item["team"] = esp.Team
			rules := make([]map[string]interface{}, len(esp.Rules))
			for j, rule := range esp.Rules {
				rules[j] = map[string]interface{}{
					"delay":     rule.Delay,
					"position":  rule.Position,
					"unique_id": rule.Unique_Id,
				}
				if rule.Targets != nil {
					targets := make([]map[string]interface{}, len(rule.Targets))
					for k, target := range rule.Targets {
						targets[k] = map[string]interface{}{
							"target_type": target.Target_type,
							"target_id":   target.Target_id,
						}
					}
					rules[j]["targets"] = targets
				}

			}
			item["rules"] = rules
			items[i] = item
		}
		if err := d.Set("escalation_policies", items); err != nil {
			return diag.FromErr(err)
		}
		d.SetId(time.Now().String())

		return diags
	}
}
