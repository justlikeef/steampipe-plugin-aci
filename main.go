package main

import (
	"steampipe-plugin-aci/aci"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: aci.Plugin})
}
