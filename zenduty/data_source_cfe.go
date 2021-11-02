// package shope

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
// )

// func dataSourceCoffees() *schema.Resource {
// 	return &schema.Resource{
// 		ReadContext: dataSourceCoffeesRead,
// 		Schema: map[string]*schema.Schema{
// 			"coffees": &schema.Schema{
// 				Type:     schema.TypeList,
// 				Computed: true,
// 				Elem: &schema.Resource{
// 					Schema: map[string]*schema.Schema{
// 						"id": &schema.Schema{
// 							Type:     schema.TypeInt,
// 							Computed: true,
// 						},
// 						"name": &schema.Schema{
// 							Type:     schema.TypeString,
// 							Computed: true,
// 						},
// 						"teaser": &schema.Schema{
// 							Type:     schema.TypeString,
// 							Computed: true,
// 						},
// 						"description": &schema.Schema{
// 							Type:     schema.TypeString,
// 							Computed: true,
// 						},
// 						"price": &schema.Schema{
// 							Type:     schema.TypeInt,
// 							Computed: true,
// 						},
// 						"image": &schema.Schema{
// 							Type:     schema.TypeString,
// 							Computed: true,
// 						},
// 						"ingredients": &schema.Schema{
// 							Type:     schema.TypeList,
// 							Computed: true,
// 							Elem: &schema.Resource{
// 								Schema: map[string]*schema.Schema{
// 									"ingredient_id": &schema.Schema{
// 										Type:     schema.TypeInt,
// 										Computed: true,
// 									},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// }

// func dataSourceCoffeesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	client := &http.Client{Timeout: 10 * time.Second}

// 	// Warning or errors can be collected in a slice type
// 	var diags diag.Diagnostics

// 	req, err := http.NewRequest("GET", fmt.Sprintf("%s/coffees", "http://localhost:19090"), nil)
// 	if err != nil {
// 		return diag.FromErr(err)
// 	}

// 	r, err := client.Do(req)
// 	if err != nil {
// 		return diag.FromErr(err)
// 	}
// 	defer r.Body.Close()

// 	coffees := make([]map[string]interface{}, 0)
// 	err = json.NewDecoder(r.Body).Decode(&coffees)
// 	if err != nil {
// 		return diag.FromErr(err)
// 	}

// 	if err := d.Set("coffees", coffees); err != nil {
// 		return diag.FromErr(err)
// 	}

// 	// always run
// 	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

// 	return diags
// }

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

func resourceTeamData() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceTeamReads,
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

func resourceTeamReads(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
