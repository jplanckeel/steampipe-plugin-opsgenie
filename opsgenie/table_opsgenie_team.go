package opsgenie

import (
	"context"

	"github.com/opsgenie/opsgenie-go-sdk-v2/team"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableOpsgenieTeam() *plugin.Table {
	return &plugin.Table{
		Name: "opsgenie_team",
		Description: "Opsgenie Team",
		List: &plugin.ListConfig{
			Hydrate: listTeams,
		},
		Columns: []*plugin.Column{
			{Name: "team_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Id"), Description: "The Id of the team."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the team."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "The description of team."},
		},
	}
}

func listTeams(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connectTeam(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("opsgenie_team.listTeams", "connection_error", err)
		return nil, err
	}

	opts := &team.ListTeamRequest{}

	teams, err := conn.List(ctx, opts)
	if err != nil {
		plugin.Logger(ctx).Error("opsgenie_team.listTeams", "api_error", err)
		return nil, err
	}
	for _, t := range teams.Teams {
		d.StreamListItem(ctx, t)
	}
	return nil, nil
}
