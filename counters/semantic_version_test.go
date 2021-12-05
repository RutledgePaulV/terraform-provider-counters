package counters

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"testing"
)

func TestProviderSemantic(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func step1Semantic() string {
	return `
		provider counters {
		}
		resource counters_semantic_version this {
			patch_triggers = {
				hash = "potatoes"
			}
		}
	`
}

func step2Semantic() string {
	return `
		provider counters {
		}
		resource counters_semantic_version this {
			patch_triggers = {
				hash = "eggs"
			}
		}
	`
}

func step3Semantic() string {
	return `
		provider counters {
		}
		resource counters_semantic_version this {
			minor_triggers = {
				hash = "potatoes"
			}
			patch_triggers = {
				hash = "eggs"
			}
		}
	`
}

func TestAccItem_Semantic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"counters": func() (*schema.Provider, error) {
				return Provider(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: step1Semantic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("counters_semantic_version.this", "value", "1.0.0"),
				),
			},
			{
				Config: step2Semantic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("counters_semantic_version.this", "value", "1.0.1"),
				),
			},
			{
				Config: step3Semantic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("counters_semantic_version.this", "value", "1.1.0"),
				),
			},
		},
	})
}
