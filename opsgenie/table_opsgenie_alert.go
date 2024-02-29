package opsgenie

import (
	"context"

	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableOpsgenieAlert() *plugin.Table {
	return &plugin.Table{
		Name: "opsgenie_alert",
		//TODO: change description
		Description: "Opsgenie alerts.",
		List: &plugin.ListConfig{
			Hydrate: listAlerts,
		},
		Columns: []*plugin.Column{
			{Name: "alert_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Id"), Description: ""},
			{Name: "alias", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "message", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "status", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "acknowledged", Type: proto.ColumnType_BOOL, Description: ""},
			{Name: "count", Type: proto.ColumnType_INT, Description: ""},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "last_occurred_at", Type: proto.ColumnType_TIMESTAMP, Description: ""},
			{Name: "source", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "priority", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "owner", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "owner_team_id", Type: proto.ColumnType_STRING, Description: ""},
		},
	}
}

func listAlerts(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connectAlert(ctx, d)
	if err != nil {
		return nil, err
	}

	// Get query config on opsgenie.spc
	query := GetConfig(d.Connection).Query

	opts := &alert.ListAlertRequest{
		Limit:  100,
		Offset: 0,
		Order:  "desc",
		Query:  *query,
	}

	for {
		alerts, err := conn.List(ctx, opts)
		if err != nil {
			return nil, err
		}
		for _, t := range alerts.Alerts {
			d.StreamListItem(ctx, t)
		}

		// Sum of Offset and limit should be lower than 20K
		// https://docs.opsgenie.com/docs/alert-api#list-alerts
		if opts.Offset == 19800 {
			break
		}
		opts.Offset = opts.Offset + 100
	}
	return nil, nil
}
