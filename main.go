package main

import (
	"github.com/fishman/terraform-provider-freenas/freenas"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: freenas.Provider})
}
