package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/desotech-it/terraform-provider-guacamole/guacamole"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "Run the provider in debug mode")
	flag.Parse()

	opts := &plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return guacamole.Provider()
		},
	}

	if debugMode {
		err := plugin.Debug(context.Background(), "registry.terraform.io/desotech-it/guacamole", opts)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	plugin.Serve(opts)
}
