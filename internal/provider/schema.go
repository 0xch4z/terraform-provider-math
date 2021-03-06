package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"token": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}
