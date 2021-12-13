package counters

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func semanticVersionResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: genericCreate,
		ReadContext:   genericRead,
		UpdateContext: genericUpdate,
		DeleteContext: genericDelete,
		CustomizeDiff: semanticVersionDiff,
		Description:   "A semantic version number whose components increment according to the configured triggers.",
		Schema: map[string]*schema.Schema{
			"major_value": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The current major version number.",
			},
			"minor_value": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The current minor version number.",
			},
			"patch_value": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The current patch version number.",
			},
			"value": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The semantic version number as a string in <major>.<minor>.<patch> form.",
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
							Type:     schema.TypeString,
							Computed: true,
						},
						"major_value": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"minor_value": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"patch_value": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"major_triggers": {
							Type: schema.TypeMap,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed: true,
						},
						"minor_triggers": {
							Type: schema.TypeMap,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed: true,
						},
						"patch_triggers": {
							Type: schema.TypeMap,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed: true,
						},
					},
				},
				Computed:    true,
				Description: "A list of semantic versions that this resource has produced.",
			},
			"major_initial_value": {
				Type:        schema.TypeInt,
				Default:     1,
				Computed:    false,
				Optional:    true,
				ForceNew:    true,
				Description: "The initial major version value.",
			},
			"minor_initial_value": {
				Type:        schema.TypeInt,
				Default:     0,
				Computed:    false,
				Optional:    true,
				ForceNew:    true,
				Description: "The initial minor version value.",
			},
			"patch_initial_value": {
				Type:        schema.TypeInt,
				Default:     0,
				Computed:    false,
				Optional:    true,
				ForceNew:    true,
				Description: "The initial patch version value.",
			},
			"major_triggers": {
				Type:        schema.TypeMap,
				Optional:    true,
				Default:     map[string]interface{}{},
				Elem:        schema.TypeString,
				Description: "A map of strings that will cause the major version number to increment when any of the values change.",
			},
			"minor_triggers": {
				Type:        schema.TypeMap,
				Optional:    true,
				Default:     map[string]interface{}{},
				Elem:        schema.TypeString,
				Description: "A map of strings that will cause the minor version number to increment when any of the values change.",
			},
			"patch_triggers": {
				Type:        schema.TypeMap,
				Optional:    true,
				Default:     map[string]interface{}{},
				Elem:        schema.TypeString,
				Description: "A map of strings that will cause the patch version number to increment when any of the values change.",
			},
		},
	}
}

func semanticVersionDiff(context context.Context, diff *schema.ResourceDiff, something interface{}) error {

	major_triggers := diff.Get("major_triggers").(map[string]interface{})
	minor_triggers := diff.Get("minor_triggers").(map[string]interface{})
	patch_triggers := diff.Get("patch_triggers").(map[string]interface{})

	_, exists := diff.GetOkExists("value")

	if !exists {

		major := diff.Get("major_initial_value").(int)
		diff.SetNew("major_value", major)
		minor := diff.Get("minor_initial_value").(int)
		diff.SetNew("minor_value", minor)
		patch := diff.Get("patch_initial_value").(int)
		diff.SetNew("patch_value", patch)
		version := fmt.Sprintf("%d.%d.%d", major, minor, patch)
		diff.SetNew("value", version)

		diff.SetNew("history", appendAndTruncate([]interface{}{}, map[string]interface{}{
			"value":          diff.Get("value"),
			"major_value":    diff.Get("major_value").(int),
			"minor_value":    diff.Get("minor_value").(int),
			"patch_value":    diff.Get("patch_value").(int),
			"major_triggers": major_triggers,
			"minor_triggers": minor_triggers,
			"patch_triggers": patch_triggers,
		}, diff.Get("max_history").(int)))

		return nil
	}

	if diff.HasChange("major_triggers") {
		newMajor := diff.Get("major_value").(int) + 1
		diff.SetNew("major_value", newMajor)
		diff.SetNew("minor_value", 0)
		diff.SetNew("patch_value", 0)
		newVersion := fmt.Sprintf("%d.%d.%d", newMajor, 0, 0)
		diff.SetNew("value", newVersion)

		curHistory := diff.Get("history").([]interface{})
		diff.SetNew("history", appendAndTruncate(curHistory, map[string]interface{}{
			"value":          diff.Get("value"),
			"major_value":    diff.Get("major_value").(int),
			"minor_value":    diff.Get("minor_value").(int),
			"patch_value":    diff.Get("patch_value").(int),
			"major_triggers": major_triggers,
			"minor_triggers": minor_triggers,
			"patch_triggers": patch_triggers,
		}, diff.Get("max_history").(int)))

	} else if diff.HasChange("minor_triggers") {
		curMajor := diff.Get("major_value").(int)
		newMinor := diff.Get("minor_value").(int) + 1
		diff.SetNew("minor_value", newMinor)
		diff.SetNew("patch_value", 0)
		newVersion := fmt.Sprintf("%d.%d.%d", curMajor, newMinor, 0)
		diff.SetNew("value", newVersion)

		curHistory := diff.Get("history").([]interface{})
		diff.SetNew("history", appendAndTruncate(curHistory, map[string]interface{}{
			"value":          diff.Get("value"),
			"major_value":    diff.Get("major_value").(int),
			"minor_value":    diff.Get("minor_value").(int),
			"patch_value":    diff.Get("patch_value").(int),
			"major_triggers": major_triggers,
			"minor_triggers": minor_triggers,
			"patch_triggers": patch_triggers,
		}, diff.Get("max_history").(int)))

	} else if diff.HasChange("patch_triggers") {
		curMajor := diff.Get("major_value").(int)
		curMinor := diff.Get("minor_value").(int)
		newPatch := diff.Get("patch_value").(int) + 1
		diff.SetNew("patch_value", newPatch)
		newVersion := fmt.Sprintf("%d.%d.%d", curMajor, curMinor, newPatch)
		diff.SetNew("value", newVersion)

		curHistory := diff.Get("history").([]interface{})
		diff.SetNew("history", appendAndTruncate(curHistory, map[string]interface{}{
			"value":          diff.Get("value"),
			"major_value":    diff.Get("major_value").(int),
			"minor_value":    diff.Get("minor_value").(int),
			"patch_value":    diff.Get("patch_value").(int),
			"major_triggers": major_triggers,
			"minor_triggers": minor_triggers,
			"patch_triggers": patch_triggers,
		}, diff.Get("max_history").(int)))

	}

	return nil
}
