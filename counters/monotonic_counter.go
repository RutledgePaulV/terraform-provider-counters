package counters

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func monotonicResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: genericCreate,
		ReadContext:   genericRead,
		UpdateContext: genericUpdate,
		DeleteContext: genericDelete,
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
			"max_history": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1000,
				ForceNew:    false,
				Description: "How many versions (max) should this resource store?",
			},
			"history": {
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"triggers": {
							Type: schema.TypeMap,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed: true,
						},
					},
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

	_, exists := diff.GetOkExists("value")

	if !exists {
		start := diff.Get("initial_value")
		diff.SetNew("value", start)

		diff.SetNew("history", appendAndTruncate([]interface{}{}, map[string]interface{}{
			"value":    start,
			"triggers": diff.Get("triggers"),
		}, diff.Get("max_history").(int)))

	} else if diff.HasChange("triggers") {
		step := diff.Get("step").(int)
		newValue := diff.Get("value").(int) + step
		curHistory := diff.Get("history").([]interface{})
		diff.SetNew("value", newValue)
		diff.SetNew("history", appendAndTruncate(curHistory, map[string]interface{}{
			"value":    newValue,
			"triggers": diff.Get("triggers"),
		}, diff.Get("max_history").(int)))
	}

	return nil
}
