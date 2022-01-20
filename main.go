package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/loen1995/terraform-custom-provider/petstore"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: petstore.Provider})
}
