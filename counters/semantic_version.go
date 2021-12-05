package counters

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func semanticVersionResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: createSemanticVersion,
		ReadContext:   readSemanticVersion,
		UpdateContext: updateSemanticVersion,
		DeleteContext: deleteSemanticVersion,
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
				Default:     map[string]string{},
				Description: "A map of strings that will cause the major version number to increment when any of the values change.",
			},
			"minor_triggers": {
				Type:        schema.TypeMap,
				Optional:    true,
				Default:     map[string]string{},
				Description: "A map of strings that will cause the minor version number to increment when any of the values change.",
			},
			"patch_triggers": {
				Type:        schema.TypeMap,
				Optional:    true,
				Default:     map[string]string{},
				Description: "A map of strings that will cause the patch version number to increment when any of the values change.",
			},
		},
	}
}

func createSemanticVersion(context context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	diagnostics := make(diag.Diagnostics, 0)
	data.SetId(uuid.New().String())
	major := data.Get("major_initial_value").(int)
	data.Set("major_value", major)
	minor := data.Get("minor_initial_value").(int)
	data.Set("minor_value", minor)
	patch := data.Get("patch_initial_value").(int)
	data.Set("patch_value", patch)
	data.Set("value", fmt.Sprintf("%d.%d.%d", major, minor, patch))
	return diagnostics
}

func readSemanticVersion(context context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	diagnostics := make(diag.Diagnostics, 0)
	return diagnostics
}

func updateSemanticVersion(context context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	diagnostics := make(diag.Diagnostics, 0)
	if data.HasChanges("patch_triggers") {
		value := data.Get("patch_value").(int)
		data.Set("patch_value", value+1)
	}

	if data.HasChanges("minor_triggers") {
		value := data.Get("minor_value").(int)
		data.Set("minor_value", value+1)
		data.Set("patch_value", 0)
	}

	if data.HasChanges("major_triggers") {
		value := data.Get("major_value").(int)
		data.Set("major_value", value+1)
		data.Set("minor_value", 0)
		data.Set("patch_value", 0)
	}

	major := data.Get("major_value").(int)
	minor := data.Get("minor_value").(int)
	patch := data.Get("patch_value").(int)

	newVersion := fmt.Sprintf("%d.%d.%d", major, minor, patch)

	if newVersion != data.Get("value") {
		data.Set("value", newVersion)
	}

	return diagnostics
}

func deleteSemanticVersion(context context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	diagnostics := make(diag.Diagnostics, 0)
	return diagnostics
}
