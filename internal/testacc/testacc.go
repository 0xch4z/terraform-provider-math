package testacc

import (
	"log"
	"testing"

	"github.com/Charliekenney23/terraform-provider-math/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Providers() map[string]*schema.Provider {
	provider := provider.Provider()
	if err := provider.InternalValidate(); err != nil {
		log.Fatalf("provider validation failed: %s", err)
	}

	return map[string]*schema.Provider{
		"math": provider,
	}
}

func PreCheck(t *testing.T) {
}
