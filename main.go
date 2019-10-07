package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/vi7/terraform-provider-pritunlko/pritunlko"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: pritunlko.Provider
	})
}
