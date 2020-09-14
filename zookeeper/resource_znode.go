package zookeeper

import (
	"context"
	"errors"
	"fmt"
	"github.com/contentsquare/terraform-provider-zookeeper/api/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceZkZnode() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceZnodeCreate,
		ReadContext:   resourceZnodeRead,
		UpdateContext: resourceZnodeUpdate,
		DeleteContext: resourceZnodeDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Required: false,
			},
			"zxid": &schema.Schema{
				Type:             schema.TypeInt,
				Computed:         true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"zxid": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}
func resourceZnodeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*client.Client)
	path := d.Get("path").(string)
	data := d.Get("data").(string)

	err := c.CreateZNode(path, data)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(path)

	return diags
}

func resourceZnodeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	path := d.Id()
	c := m.(*client.Client)
	znode, err := c.ReadZNode(path)
	if err != nil {
		if err.Error() == "zk: node does not exist" {
			d.SetId("")
			return diags
		}
		return diag.FromErr(err)
	}
	d.Set("path", path)
	d.Set("data", string(znode.Data))
	d.Set("zxid", znode.Zxid)
	return diags
}

func resourceZnodeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	path := d.Id()
	c := m.(*client.Client)
	data := d.Get("data").(string)
	err := c.UpdateZnode(path, data)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("path", path)
	d.Set("data", data)
	return diags
}

func resourceZnodeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	path := d.Id()
	c := m.(*client.Client)
	err := c.DeleteZnode(path)
	if err != nil {
		return diag.FromErr(errors.New(fmt.Sprintf("Unable to delete znode %s. err=%s", path, err)))
	}
	return diags
}

//func resourceZnodeExists(ctx context.Context, d *schema.ResourceData, m interface{}) (bool, diag.Diagnostics){
//	var diags diag.Diagnostics
//	path := d.Id()
//	c := m.(*client.Client)
//	exists,  err := c.ZnodeExists(path)
//	if err != nil {
//		return exists, diag.FromErr(err)
//	}
//	return exists, diags
//}
//
//
