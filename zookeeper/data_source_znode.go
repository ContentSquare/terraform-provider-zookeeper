package zookeeper

import (
	"context"
	"errors"
	"fmt"
	"github.com/contentsquare/terraform-provider-zookeeper/api/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func datasourceZkZnode() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceZnodeRead,
		Schema: map[string]*schema.Schema{
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"data": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"zxid": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"data": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceZnodeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*client.Client)
	path := d.Get("path").(string)

	znode, err := c.ReadZNode(path)
	log.Printf("Znode is %+v", znode)
	if err != nil {
		d.SetId("")
		return diag.FromErr(errors.New(fmt.Sprintf("datasource zookeeper node: can't read path %s. err=%v", path, err)))
	}

	d.SetId(znode.Path)
	data := make([]map[string]interface{}, 0)
	entry := make(map[string]interface{}, 0)
	entry["zxid"] = znode.Zxid
	entry["data"] = string(znode.Data)
	data = append(data, entry)
	d.Set("data", data)
	return diags
}
