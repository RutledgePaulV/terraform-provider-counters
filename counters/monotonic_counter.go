package counters

import (
	"context"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func monotonicResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: createCounter,
		ReadContext:   readCounter,
		UpdateContext: updateCounter,
		DeleteContext: deleteCounter,
		CustomizeDiff: customMonotonicDiff,
		Description:   "A monotonic counter which increments according to the configured triggers.",
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
				ForceNew:    true,
				Description: "The amount used to increment / decrement the counter on each revision.",
			},
			"history": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Computed:    true,
				Description: "A list of counter values that this resource has produced.",
			},
			"initial_value": {
				Type:        schema.TypeInt,
				Default:     0,
				Computed:    false,
				Optional:    true,
				ForceNew:    true,
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

func customMonotonicDiff(context context.Context, diff *schema.ResourceDiff, something interface{}) error {
	if diff.HasChange("triggers") {
		step := diff.Get("step").(int)
		newValue := diff.Get("value").(int) + step
		curHistory := diff.Get("history").([]interface{})
		diff.SetNew("value", newValue)
		diff.SetNew("history", append(curHistory, newValue))
	}
	return nil
}

func createCounter(context context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	diagnostics := make(diag.Diagnostics, 0)
	data.SetId(uuid.New().String())
	start := data.Get("initial_value")
	data.Set("value", start)
	data.Set("history", []interface{}{start})
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
