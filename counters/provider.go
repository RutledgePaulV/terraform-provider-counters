package counters

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"counters_monotonic":        monotonicResource(),
			"counters_semantic_version": semanticVersionResource(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return nil, make(diag.Diagnostics, 0)
}
