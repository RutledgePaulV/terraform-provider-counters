package counters

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"testing"
)

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func step1() string {
	return `
		provider counters {
		}
		resource counters_monotonic this {
			initial_value = 35
			triggers = {
				hash = "potatoes"
			}
		}
	`
}

func step2() string {
	return `
		provider counters {
		}
		resource counters_monotonic this {
			initial_value = 35
			triggers = {
				hash = "eggs"
			}
		}
	`
}

func TestAccItem_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"counters": func() (*schema.Provider, error) {
				return Provider(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: step1(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("counters_monotonic.this", "value", "35"),
					resource.TestCheckResourceAttr("counters_monotonic.this", "history.0", "35"),
				),
			},
			{
				Config: step2(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("counters_monotonic.this", "value", "36"),
					resource.TestCheckResourceAttr("counters_monotonic.this", "history.0", "35"),
					resource.TestCheckResourceAttr("counters_monotonic.this", "history.1", "36"),
				),
			},
		},
	})
}
