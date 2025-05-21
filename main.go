package main

import (
	"flag"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	// "github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/scribe-security/terraform-valint/internal"
)

//go:generate terraform fmt -recursive ./examples/
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

const version string = "dev"

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/scribe-security/valint",
		Debug:   debug,
	}
	fmt.Println("Starting valint provider...", opts, internal.New(version))

	// Wire up ggcr logs.
	// logs.Warn.SetOutput(os.Stderr)
	// if debug {
	// 	logs.Progress.SetOutput(os.Stderr)
	// 	logs.Debug.SetOutput(os.Stderr)
	// }

	// if err := providerserver.Serve(context.Background(), provider.New(version), opts); err != nil {
	// 	log.Fatal(err.Error())
	// }
}
