
# Opsgenie Plugin for Steampipe

Use SQL to query infrastructure including servers, networks, facilities and more from Opsgenie.

- **[Get started →](https://hub.steampipe.io/plugins/jplanckeel/opsgenie)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/jplanckeel/opsgenie/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/jplanckeel/opsgenie/issues)

## Quick start

### Install

Download and install the latest Opsgenie plugin:

```bash
steampipe plugin install jplanckeel/opsgenie
```

Configure your [credentials](https://hub.steampipe.io/plugins/jplanckeel/opsgenie#credentials) and [config file](https://hub.steampipe.io/plugins/jplanckeel/opsgenie#configuration).

Configure your account details in `~/.steampipe/config/opsgenie.spc`:

```hcl
connection "steampipe-plugin-opsgenie" {
  plugin  = "jplanckeel/opsgenie"

  # URL to access the API (required).
  # If using the EU instance of Opsgenie, the URL needs to be api.eu.opsgenie.com for requests to be successful.
  #url = "api.opsgenie.com"

  # API token for your opsgenie instance (required).
  # See https://docs.opsgenie.com/docs/api-access-management
  #token = "5c44f27d-8dd5-4939-aa5f-499d8cssf64a"

  # To filter request you can add opsgenie query
  #query = "status: open AND responders: 'My_Team'"
}
```

Run steampipe:

```shell
steampipe query
```

List teams in your Opsgenie account:

```sql
select
   team_id,
   name,
   description 
from
   opsgenie_team;
```

```
+--------------------------------------+------------------------------------+--------------------------------------------+
| team_id                              | name                               | description                                |
+--------------------------------------+------------------------------------+--------------------------------------------+
| 8cfdd4da-73e9-4526-be90-02111f2f2f1f | Infra_Team                         | Infrastructure Team                        |
| 555d4f34-46d5-41b6-88bd-12df8z1f7104 | Dev_Team                           | Developper Team                            |
+--------------------------------------+------------------------------------+--------------------------------------------+
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/jplanckeel/steampipe-plugin-opsgenie.git
cd steampipe-plugin-opsgenie
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/opsgenie.spc
```

Try it!

```
steampipe query
> .inspect opsgenie
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/jplanckeel/steampipe-plugin-opsgenie/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Opsgenie Plugin](https://github.com/jplanckeel/steampipe-plugin-opsgenie/labels/help%20wanted)
