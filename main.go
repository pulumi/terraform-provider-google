package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/pulumi/terraform-provider-google/google"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: google.Provider})
}
