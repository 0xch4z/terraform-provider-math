package addition

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"addend": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"augend": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"sum": {
			Type:     schema.TypeInt,
			Computed: true,
		},
	}
}
