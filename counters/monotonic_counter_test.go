package counters

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

var testAccProviders map[string]func() (*schema.Provider, error)

func init() {
	testAccProviders = map[string]func() (*schema.Provider, error){
		"counters": func() (*schema.Provider, error) {
			return Provider(), nil
		},
	}
}

func testAccPreCheck(t *testing.T) {

}

func testCheckDestroy(state *terraform.State) error {
	return nil
}

func terraformSource() string {
	return `
		provider counters {
		}
		resource counters_monotonic_counter this {
		}
	`
}

func TestAccItem_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviders,
		CheckDestroy:      testCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: terraformSource(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("counters_monotonic_counter.this", "value", "0"),
				),
			},
		},
	})
}
