package addition_test

import (
	"testing"

	"github.com/Charliekenney23/terraform-provider-math/internal/testacc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const dataSourceName = "data.math_addition.test"

func TestDataSource_basic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testacc.PreCheck(t) },
		Providers: testacc.Providers(),
		Steps: []resource.TestStep{
			{
				Config: configBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "addend", "2"),
					resource.TestCheckResourceAttr(dataSourceName, "augend", "3"),
					resource.TestCheckResourceAttr(dataSourceName, "sum", "5"),
				),
			},
		},
	})
}

const configBasic = `
data "math_addition" "test" {
	addend = 2
	augend = 3
}`
