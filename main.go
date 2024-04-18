package main

import (
	"github.com/jplanckeel/steampipe-plugin-opsgenie/opsgenie"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: opsgenie.Plugin})
}
