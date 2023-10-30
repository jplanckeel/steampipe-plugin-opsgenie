connection "opsgenie" {
  plugin  = "jplanckeel/opsgenie"

  # URL to access the API (required).
  # If using the EU instance of Opsgenie, the URL needs to be api.eu.opsgenie.com for requests to be successful.
  # url = "api.opsgenie.com"

  # API token for your opsgenie instance (required).
  # See https://docs.opsgenie.com/docs/api-access-management
  # token = "5c44f27d-8dd5-4939-aa5f-499d8cssf64a"

  # To filter request you can add opsgenie query
  # query = "status: open AND responders: 'My_Team'"
}
