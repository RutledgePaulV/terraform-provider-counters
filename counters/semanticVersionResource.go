package counters

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/diag"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func semanticVersionResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: createSemanticVersion,
		ReadContext:   readSemanticVersion,
		UpdateContext: updateSemanticVersion,
		DeleteContext: deleteSemanticVersion,
		Schema: map[string]*schema.Schema{
			"major": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The current major version number.",
			},
			"minor": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The current minor version number.",
			},
			"patch": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The current patch version number.",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The semantic version number as a string in <major>.<minor>.<patch> form.",
			},
			"start_major": {
				Type:        schema.TypeInt,
				Default:     0,
				Computed:    false,
				Description: "The initial major version value.",
			},
			"start_minor": {
				Type:        schema.TypeInt,
				Default:     0,
				Computed:    false,
				Description: "The initial minor version value.",
			},
			"start_patch": {
				Type:        schema.TypeInt,
				Default:     0,
				Computed:    false,
				Description: "The initial patch version value.",
			},
			"triggers_major": {
				Type:        schema.TypeMap,
				Default:     map[string]string{},
				Description: "A map of strings that will cause the major version number to increment when any of the values change.",
			},
			"triggers_minor": {
				Type:        schema.TypeMap,
				Default:     map[string]string{},
				Description: "A map of strings that will cause the minor version number to increment when any of the values change.",
			},
			"triggers_patch": {
				Type:        schema.TypeMap,
				Default:     map[string]string{},
				Description: "A map of strings that will cause the patch version number to increment when any of the values change.",
			},
		},
	}
}

func createSemanticVersion(context context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	diagnostics := make(diag.Diagnostics, 0)

	return diagnostics
}

func readSemanticVersion(context context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	diagnostics := make(diag.Diagnostics, 0)

	return diagnostics
}

func updateSemanticVersion(context context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	diagnostics := make(diag.Diagnostics, 0)

	return diagnostics
}

func deleteSemanticVersion(context context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	diagnostics := make(diag.Diagnostics, 0)

	return diagnostics
}
