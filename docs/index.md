---
organization: jplanckeel
category: ["alert & onc-all management"]
icon_url: "/images/plugins/turbot/opsgenie.svg"
brand_color: "#2684FF"
display_name: "Opsgenie"
short_name: "opsgenie"
description: "Steampipe plugin for querying teams and alerts from Opsgenie."
og_description: "Query Opsgenie with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/opsgenie-social-graphic.png"
---

# Opsgenie + Steampipe

[Opsgenie](https://www.atlassian.com/software/opsgenie) provides on-call and alert management to keep services always one.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List teams in your Opsgenie account:

```sql
select 
  team_id,
  name,
 description
from 
  opsgenie_teams;
```

```
+--------------------------------------+------------------------------------+--------------------------------------------+
| team_id                              | name                               | description                                |
+--------------------------------------+------------------------------------+--------------------------------------------+
| 8cfdd4da-73e9-4526-be90-02111f2f2f1f | Infra_Team                         | Infrastructure Team                        |
| 555d4f34-46d5-41b6-88bd-12df8z1f7104 | Dev_Team                           | Developper Team                            |
+--------------------------------------+------------------------------------+--------------------------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/jplanckeel/opsgenie/tables)**

## Get started

### Install

Download and install the latest Opsgenie plugin:

```bash
steampipe plugin install opsgenie
```

### Credentials

| Item        | Description                                                                                                                            |
| :---------- | :------------------------------------------------------------------------------------------------------------------------------------- |
| Credentials | Opsgenie requires an [API token](https://id.atlassian.com/manage-profile/security/api-tokens), sit base url and token for all requests.|

<!-- | Permissions | Grant the `ReadOnlyAccess` policy to your user or role.                                                                                | -->

### Configuration

Installing the latest opsgenie plugin will create a config file (`~/.steampipe/config/opsgenie.spc`) with a single connection named `opsgenie`:

```hcl
  connection "steampipe-plugin-opsgenie" {
    plugin = "opsgenie"


    # username / Email address of agent user who have permission to access the API (required).
    # default Organization
    #url      = "api.opsgenie.com"
    # for EU Organization
    #url      = "api.eu.opsgenie.com"

    # API token for your opsgenie instance (required).
    # See https://docs.opsgenie.com/docs/api-access-management
    #token      = "YOUR_opsgenie_API_KEY"

    # To filter request you can add opsgenie query
    #query= "status: open AND responders: 'My_Team'"
  }
```

- `url` - The site url of your attlassian opsgenie subscription.
- `token` - [API token](https://id.atlassian.com/manage-profile/security/api-tokens) for user's Atlassian account.
- `query` - Query to filter alert [Query](https://support.atlassian.com/opsgenie/docs/search-queries-for-alerts/).

Alternatively, you can also use the standard Opsgenie environment variables to obtain credentials **only if other arguments (`url`, `token`) are not specified** in the connection:

```sh
export OPSGENIE_URL=https://your-domain.atlassian.net/
export OPSGENIE_TOKEN=8WqcdT0rvIZpCjtDqReF48B1
```

## Get involved

- Open source: https://github.com/jplanckeel/steampipe-plugin-opsgenie
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
