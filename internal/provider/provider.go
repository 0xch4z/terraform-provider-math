package provider

import (
	"github.com/Charliekenney23/terraform-provider-math/internal/addition"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: providerSchema(),
		DataSourcesMap: map[string]*schema.Resource{
			"math_addition": addition.DataSource(),
		},
	}
}
