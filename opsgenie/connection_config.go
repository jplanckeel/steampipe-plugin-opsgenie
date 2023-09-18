package opsgenie

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type opsgenieConfig struct {
	Url   *string `cty:"url"`
	Token *string `cty:"token"`
	Query *string `cty:"query"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"url": {
		Type: schema.TypeString,
	},
	"token": {
		Type: schema.TypeString,
	},
	"query": {
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

	var query string
	if config.Query == nil {
		config.Query = &query
	}

	return config
}
