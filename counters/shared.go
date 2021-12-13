package counters

import (
	"context"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func truncate(list []interface{}, max int) []interface{} {
	if len(list) > max {
		return list[(len(list) - max):]
	} else {
		return list
	}
}

func appendAndTruncate(list []interface{}, item interface{}, max int) []interface{} {
	return truncate(append(list, item), max)
}

func genericCreate(context context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	diagnostics := make(diag.Diagnostics, 0)
	data.SetId(uuid.New().String())
	return diagnostics
}

func genericRead(context context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	diagnostics := make(diag.Diagnostics, 0)
	return diagnostics
}

func genericUpdate(context context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	diagnostics := make(diag.Diagnostics, 0)
	return diagnostics
}

func genericDelete(context context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	diagnostics := make(diag.Diagnostics, 0)
	return diagnostics
}
