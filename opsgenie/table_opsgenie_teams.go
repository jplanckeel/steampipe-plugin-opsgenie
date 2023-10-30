package opsgenie

import (
	"context"

	"github.com/opsgenie/opsgenie-go-sdk-v2/team"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableOpsgenieTeams() *plugin.Table {
	return &plugin.Table{
		Name: "opsgenie_teams",
		//TODO: change description
		Description: "Opsgenie teams.",
		List: &plugin.ListConfig{
			Hydrate: listTeam,
		},
		Columns: []*plugin.Column{
			{Name: "team_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Id"), Description: ""},
			{Name: "name", Type: proto.ColumnType_STRING, Description: ""},
			{Name: "description", Type: proto.ColumnType_STRING, Description: ""},
		},
	}
}

func listTeam(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connectTeam(ctx, d)
	if err != nil {
		return nil, err
	}

	opts := &team.ListTeamRequest{}

	teams, err := conn.List(ctx, opts)
	if err != nil {
		return nil, err
	}
	for _, t := range teams.Teams {
		d.StreamListItem(ctx, t)
	}
	return nil, nil
}
