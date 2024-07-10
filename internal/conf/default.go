package conf

const GbDefault = `
[backend]
  type = "semtech_udp"
  [backend.semtech_udp]
    udp_bind = "0.0.0.0:1700"

[filters]
  net_ids = []
  join_euis = []

# Integration configuration.
[integration]
# Payload marshaler.
#
# This defines how the MQTT payloads are encoded. Valid options are:
# * protobuf:  Protobuf encoding (this will become the ChirpStack Gateway Bridge v3 default)
# * json:      JSON encoding (easier for debugging, but less compact than 'protobuf')
marshaler="protobuf"

  # MQTT integration configuration.
  [integration.mqtt]
  # Event topic template.
  event_topic_template="gateway/{{ .GatewayID }}/event/{{ .EventType }}"

  # Command topic template.
  command_topic_template="gateway/{{ .GatewayID }}/command/#"

  # MQTT authentication.
  [integration.mqtt.auth]
  # Type defines the MQTT authentication type to use.
  #
  # Set this to the name of one of the sections below.
  type="generic"

    # Generic MQTT authentication.
    [integration.mqtt.auth.generic]
    # MQTT server (e.g. scheme://host:port where scheme is tcp, ssl or ws)
    server="tcp://127.0.0.1:1884"

    # Connect with the given username (optional)
    username=""

    # Connect with the given password (optional)
    password=""
`
