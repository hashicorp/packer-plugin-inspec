package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/packer-plugin-sdk/plugin"

	inspecprovisioner "github.com/hashicorp/packer-plugin-inspec/provisioner/inspec"
	"github.com/hashicorp/packer-plugin-inspec/version"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterProvisioner(plugin.DEFAULT_NAME, new(inspecprovisioner.Provisioner))
	pps.SetVersion(version.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
