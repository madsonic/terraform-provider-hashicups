package hashicups

import (
	"context"

	"github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("HASHICUPS_USERNAME", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("HASHICUPS_PASSWORD", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"hashicups_coffees": dataSourceCoffees(),
			"hashicups_order":   dataSourceOrder(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	var diags diag.Diagnostics

	// authenticated flow
	if (username != "") && (password != "") {
		c, err := hashicups.NewClient(nil, &username, &password)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create Hashicups client",
				Detail:   "Unable to auth user for authenticated HashiCups client",
			})

			return nil, diags
		}

		return c, diags
	}

	// unauthenticated flow
	c, err := hashicups.NewClient(nil, nil, nil)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Hashicups client",
			Detail:   "No authentication credentials provided",
		})

		return nil, diags
	}

	return c, diags
}
