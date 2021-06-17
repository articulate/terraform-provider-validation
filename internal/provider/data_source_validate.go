package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceValidate() *schema.Resource {
	return &schema.Resource{
		Description: "Validate data",

		ReadContext: dataSourceValidateRead,

		Schema: map[string]*schema.Schema{
			"condition": {
				Description: "Data validation condition.",
				Type:        schema.TypeBool,
				Required:    true,
			},
			"error_message": {
				Description: "Error message to display if condition is not met.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func dataSourceValidateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId(d.Get("error_message").(string))

	if d.Get("condition").(bool) {
		return nil
	}

	return diag.Errorf(d.Id())
}
