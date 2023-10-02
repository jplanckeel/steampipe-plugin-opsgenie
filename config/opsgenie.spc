  connection "steampipe-plugin-opsgenie" {
    plugin = "opsgenie"


    # username / Email address of agent user who have permission to access the API (required).
    #url      = "opsgenie_api_url"

    # API token for your opsgenie instance (required).
    # See https://docs.opsgenie.com/docs/api-access-management
    #token      = "YOUR_opsgenie_API_KEY"

    # To filter request you can add opsgenie query
    #query= "status: open AND responders: 'My_Team'"
  }
