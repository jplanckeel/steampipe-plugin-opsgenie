package opsgenie

import (
	"context"
	"errors"
	"os"

	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/team"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connectAlert(ctx context.Context, d *plugin.QueryData) (*alert.Client, error) {

	url := os.Getenv("OPSGENIE_URL")
	token := os.Getenv("OPSGENIE_TOKEN")

	opsgenieConfig := GetConfig(d.Connection)
	if opsgenieConfig.Url != nil {
		url = *opsgenieConfig.Url
	}
	if opsgenieConfig.Token != nil {
		token = *opsgenieConfig.Token
	}

	if url == "" {
		return nil, errors.New("'url' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	if token == "" {
		return nil, errors.New("'token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	// You can set custom *http.Client here
	client, err := alert.NewClient(&client.Config{
		ApiKey:         token,
		OpsGenieAPIURL: client.ApiUrl(url),
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}

func connectTeam(ctx context.Context, d *plugin.QueryData) (*team.Client, error) {

	url := os.Getenv("OPSGENIE_URL")
	token := os.Getenv("OPSGENIE_TOKEN")

	opsgenieConfig := GetConfig(d.Connection)
	if opsgenieConfig.Url != nil {
		url = *opsgenieConfig.Url
	}
	if opsgenieConfig.Token != nil {
		token = *opsgenieConfig.Token
	}

	if url == "" {
		return nil, errors.New("'url' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	if token == "" {
		return nil, errors.New("'token' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	// You can set custom *http.Client here
	client, err := team.NewClient(&client.Config{
		ApiKey:         token,
		OpsGenieAPIURL: client.ApiUrl(url),
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}
