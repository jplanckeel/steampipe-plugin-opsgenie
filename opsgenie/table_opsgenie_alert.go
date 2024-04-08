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
		Description: "Opsgenie Alerts",
		List: &plugin.ListConfig{
			Hydrate: listAlerts,
		},
		Columns: []*plugin.Column{
			{Name: "alert_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Id"), Description: "The Id of the alert."},
			{Name: "alias", Type: proto.ColumnType_STRING, Description: "Client-defined identifier of the alert, that is also the key element of Alert De-Duplication."},
			{Name: "message", Type: proto.ColumnType_STRING, Description: "Message of the alert."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The current status of the alert (e.g., open, closed)."},
			{Name: "acknowledged", Type: proto.ColumnType_BOOL, Description: "A boolean indicating whether the alert has been acknowledged."},
			{Name: "count", Type: proto.ColumnType_INT, Description: "This field indicates the number of times this alert has been triggered or reported."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of when the alert was created."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp when the alert was last updated.."},
			{Name: "last_occurred_at", Type: proto.ColumnType_TIMESTAMP, Description: "This timestamp shows when the alert was last triggered."},
			{Name: "source", Type: proto.ColumnType_STRING, Description: "Display name of the request source."},
			{Name: "priority", Type: proto.ColumnType_STRING, Description: "Priority level of the alert. Possible values are P1, P2, P3, P4 and P5. Default value is P3."},
			{Name: "owner", Type: proto.ColumnType_STRING, Description: "The user or entity that is currently responsible for the alert."},
			{Name: "owner_team_id", Type: proto.ColumnType_STRING, Description: "The Id of the team that owns the alert."},
		},
	}
}

func listAlerts(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connectAlert(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("opsgenie_alert.listAlerts", "connection_error", err)
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
			plugin.Logger(ctx).Error("opsgenie_alert.listAlerts", "api_error", err)
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
