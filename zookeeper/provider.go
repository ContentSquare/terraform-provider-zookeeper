package zookeeper

import (
	"context"
	"github.com/contentsquare/terraform-provider-zookeeper/api/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ZOOKEEPER_HOST", "localhost"),
			},
			"port": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ZOOKEEPER_PORT", 2181),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"zookeeper_node": resourceZkZnode(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"zookeeper_node": datasourceZkZnode(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	address := d.Get("host").(string)
	port := d.Get("port").(int)
	c, err := client.NewClient(address, port)
	if err != nil {
		return c, diag.FromErr(err)
	}
	return c, diags
}