package main

import (
	"github.com/contentsquare/terraform-provider-zookeeper/zookeeper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

var (
	version, commit string
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return zookeeper.Provider()
		},
	})
}
