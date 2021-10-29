package zenduty

import (
	"terraform-provider-zenduty/client"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceTeam() *schema.Resource {
	return &schema.Resource{
		Create: resourceTeamCreate,
		Read:   resourceTeamRead,
		Update: resourceTeamUpdate,
		Delete: resourceTeamDelete,

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

func resourceTeamCreate(d *schema.ResourceData, m interface{}) error {
	apiclient := m.(*client.Client)
	newteam := &client.Team{}
	if v, ok := d.GetOk("name"); ok {
		newteam.Name = v.(string)
	}
	task, err := apiclient.CreateTeam(newteam)
	if err != nil {
		return err
	}
	d.SetId(task.Unique_Id)
	return resourceTeamRead(d, m)
}

func resourceTeamUpdate(d *schema.ResourceData, m interface{}) error {

	return nil
}

func resourceTeamDelete(d *schema.ResourceData, m interface{}) error {

	return nil
}

func resourceTeamRead(d *schema.ResourceData, m interface{}) error {
	apiclient := m.(*client.Client)
	id := d.Id()
	t, err := apiclient.GetTeam(id)
	if err != nil {
		return err
	}
	d.Set("name", t.Name)
	d.Set("unique_id", t.Unique_Id)
	d.Set("creation_date", t.Creation_Date)
	d.Set("account", t.Account)
	d.Set("creation_date", t.Creation_Date)
	return nil
}
