package main

import (
	"flag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/articulate/terraform-provider-validation/internal/provider"
)

// these will be set by the goreleaser configuration
// to appropriate values for the compiled binary
var version = "dev"

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.New(version),
		ProviderAddr: "registry.terraform.io/articulate/validation",
		Debug:        debugMode,
	})
}
