package main

import (
	"github.com/atsaki/terraform-provider-idcf/idcf"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: idcf.Provider,
	})
}
