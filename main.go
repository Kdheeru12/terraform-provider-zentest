// package main

// import (
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

// 	"terraform-provider-zenduty/zenduty"

// 	"github.com/hashicorp/terraform-plugin-sdk/terraform"
// )

// func main() {
// 	plugin.Serve(&plugin.ServeOpts{
// 		ProviderFunc: func() terraform.ResourceProvider {
// 			return zenduty.Provider()
// 		},
// 	})
// }

package main

import (
	"terraform-provider-zenduty/zenduty"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return zenduty.Provider()
		},
	})
}
