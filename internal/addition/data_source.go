package addition

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSource() *schema.Resource {
	return &schema.Resource{
		Schema:      dataSourceSchema(),
		ReadContext: dataSourceRead,
	}
}

func dataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	addend := d.Get("addend").(int)
	augend := d.Get("augend").(int)
	d.SetId(fmt.Sprintf("%d+%d", addend, augend))
	d.Set("sum", addend+augend)
	return nil
}
