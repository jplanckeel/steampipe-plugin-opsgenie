connection "steampipe-plugin-opsgenie" {
  plugin = "jplanckeel/opsgenie"

  # URL to access the API (required).
  # If using the EU instance of Opsgenie, the URL needs to be api.eu.opsgenie.com for requests to be successful.
  #url      = "api.opsgenie.com"
  

  # API token for your opsgenie instance (required).
  # See https://docs.opsgenie.com/docs/api-access-management
  #token      = "33d0d62a6a163083ba7b3bab31bd6612"
  
  # To filter request you can add opsgenie query
  #query= "status: open AND responders: 'My_Team'"
}
