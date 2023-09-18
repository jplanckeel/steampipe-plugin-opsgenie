package opsgenie

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type opsgenieConfig struct {
	Url   *string `cty:"url"`
	Token *string `cty:"token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"url": {
		Type: schema.TypeString,
	},
	"token": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &opsgenieConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) opsgenieConfig {
	if connection == nil || connection.Config == nil {
		return opsgenieConfig{}
	}
	config, _ := connection.Config.(opsgenieConfig)
	return config
}
