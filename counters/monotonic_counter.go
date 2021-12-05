package counters

import (
	"context"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func monotonicCounterResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: createCounter,
		ReadContext:   readCounter,
		UpdateContext: updateCounter,
		DeleteContext: deleteCounter,
		Schema: map[string]*schema.Schema{
			"value": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The current value of the counter.",
			},
			"step": {
				Type:        schema.TypeInt,
				Default:     1,
				Computed:    false,
				Optional:    true,
				Description: "The amount used to increment / decrement the counter on each revision.",
			},
			"start": {
				Type:        schema.TypeInt,
				Default:     0,
				Computed:    false,
				Optional:    true,
				Description: "The initial value to use for the counter.",
			},
			"triggers": {
				Type:        schema.TypeMap,
				Optional:    true,
				Default:     map[string]string{},
				Description: "A map of strings that will cause a change to the counter when any of the values change.",
			},
		},
	}
}

func createCounter(context context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	diagnostics := make(diag.Diagnostics, 0)
	data.SetId(uuid.New().String())
	start, _ := data.GetOk("start")
	err := data.Set("value", start)
	if err != nil {
		diagnostics = append(diagnostics, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Encountered error initializing value.",
		})
	}
	return diagnostics
}

func readCounter(context context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	diagnostics := make(diag.Diagnostics, 0)

	return diagnostics
}

func updateCounter(context context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	diagnostics := make(diag.Diagnostics, 0)

	return diagnostics
}

func deleteCounter(context context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	diagnostics := make(diag.Diagnostics, 0)

	return diagnostics
}
