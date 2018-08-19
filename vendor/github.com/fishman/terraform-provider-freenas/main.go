package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-freenas/freenas"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: freenas.Provider})
}
