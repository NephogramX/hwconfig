package conf

import (
	"gitee.com/dfrobotcd/chirpstack-api/go/as/external/api"
	"github.com/NephogramX/hwconfig/internal/band"
)

/*
// Config defines the configuration structure.
type NsConfigFull struct {
	GwLog struct {
		RemoteIP   string `mapstructure:"remote_ip"`
		RemotePort int    `mapstructure:"remote_port"`
		Module     string `mapstructure:"module"`
	} `mapstructure:"gwlog"`

	Statistics struct {
		URL     string `mapstructure:"url"`
		Enabled bool   `mapstructure:"enabled"`
	} `mapstructure:"statistics"`

	General struct {
		LogLevel                  int    `mapstructure:"log_level"`
		LogToSyslog               bool   `mapstructure:"log_to_syslog"`
		GRPCDefaultResolverScheme string `mapstructure:"grpc_default_resolver_scheme"`
	} `mapstructure:"general"`

	PostgreSQL struct {
		DSN                string `mapstructure:"dsn"`
		Automigrate        bool   `mapstructure:"automigrate"`
		MaxOpenConnections int    `mapstructure:"max_open_connections"`
		MaxIdleConnections int    `mapstructure:"max_idle_connections"`
	} `mapstructure:"postgresql"`

	Redis struct {
		URL        string   `mapstructure:"url"` // deprecated
		Servers    []string `mapstructure:"servers"`
		Cluster    bool     `mapstructure:"cluster"`
		MasterName string   `mapstructure:"master_name"`
		PoolSize   int      `mapstructure:"pool_size"`
		Password   string   `mapstructure:"password"`
		Database   int      `mapstructure:"database"`
		TLSEnabled bool     `mapstructure:"tls_enabled"`
	} `mapstructure:"redis"`

	NetworkServer struct {
		NetID                lorawan.NetID
		NetIDString          string        `mapstructure:"net_id"`
		DeduplicationDelay   time.Duration `mapstructure:"deduplication_delay"`
		DeviceSessionTTL     time.Duration `mapstructure:"device_session_ttl"`
		GetDownlinkDataDelay time.Duration `mapstructure:"get_downlink_data_delay"`

		Band struct {
			Name                   band.Name `mapstructure:"name"`
			UplinkDwellTime400ms   bool      `mapstructure:"uplink_dwell_time_400ms"`
			DownlinkDwellTime400ms bool      `mapstructure:"downlink_dwell_time_400ms"`
			UplinkMaxEIRP          float32   `mapstructure:"uplink_max_eirp"`
			RepeaterCompatible     bool      `mapstructure:"repeater_compatible"`
		} `mapstructure:"band"`

		NetworkSettings struct {
			InstallationMargin      float64 `mapstructure:"installation_margin"`
			RXWindow                int     `mapstructure:"rx_window"`
			RX1Delay                int     `mapstructure:"rx1_delay"`
			RX1DROffset             int     `mapstructure:"rx1_dr_offset"`
			RX2DR                   int     `mapstructure:"rx2_dr"`
			RX2Frequency            int     `mapstructure:"rx2_frequency"`
			RX2PreferOnRX1DRLt      int     `mapstructure:"rx2_prefer_on_rx1_dr_lt"`
			RX2PreferOnLinkBudget   bool    `mapstructure:"rx2_prefer_on_link_budget"`
			GatewayPreferMinMargin  float64 `mapstructure:"gateway_prefer_min_margin"`
			DownlinkTXPower         int     `mapstructure:"downlink_tx_power"`
			EnabledUplinkChannels   []int   `mapstructure:"enabled_uplink_channels"`
			DisableMACCommands      bool    `mapstructure:"disable_mac_commands"`
			DisableADR              bool    `mapstructure:"disable_adr"`
			MaxMACCommandErrorCount int     `mapstructure:"max_mac_command_error_count"`

			ExtraChannels []struct {
				Frequency int `mapstructure:"frequency"`
				MinDR     int `mapstructure:"min_dr"`
				MaxDR     int `mapstructure:"max_dr"`
			} `mapstructure:"extra_channels"`

			ClassB struct {
				PingSlotDR        int `mapstructure:"ping_slot_dr"`
				PingSlotFrequency int `mapstructure:"ping_slot_frequency"`
			} `mapstructure:"class_b"`

			RejoinRequest struct {
				Enabled   bool `mapstructure:"enabled"`
				MaxCountN int  `mapstructure:"max_count_n"`
				MaxTimeN  int  `mapstructure:"max_time_n"`
			} `mapstructure:"rejoin_request"`
		} `mapstructure:"network_settings"`

		Scheduler struct {
			SchedulerInterval time.Duration `mapstructure:"scheduler_interval"`

			ClassC struct {
				DownlinkLockDuration  time.Duration `mapstructure:"downlink_lock_duration"`
				MulticastGatewayDelay time.Duration `mapstructure:"multicast_gateway_delay"`
			} `mapstructure:"class_c"`
		} `mapstructure:"scheduler"`

		API struct {
			Bind    string `mapstructure:"bind"`
			CACert  string `mapstructure:"ca_cert"`
			TLSCert string `mapstructure:"tls_cert"`
			TLSKey  string `mapstructure:"tls_key"`
		} `mapstructure:"api"`

		Gateway struct {
			// Deprecated
			Stats struct {
				Timezone string
			}

			CACert             string        `mapstructure:"ca_cert"`
			CAKey              string        `mapstructure:"ca_key"`
			ClientCertLifetime time.Duration `mapstructure:"client_cert_lifetime"`

			Backend struct {
				Type                 string `mapstructure:"type"`
				MultiDownlinkFeature string `mapstructure:"multi_downlink_feature"`

				MQTT struct {
					Server               string        `mapstructure:"server"`
					Username             string        `mapstructure:"username"`
					Password             string        `mapstructure:"password"`
					MaxReconnectInterval time.Duration `mapstructure:"max_reconnect_interval"`
					QOS                  uint8         `mapstructure:"qos"`
					CleanSession         bool          `mapstructure:"clean_session"`
					ClientID             string        `mapstructure:"client_id"`
					CACert               string        `mapstructure:"ca_cert"`
					TLSCert              string        `mapstructure:"tls_cert"`
					TLSKey               string        `mapstructure:"tls_key"`

					EventTopic           string `mapstructure:"event_topic"`
					CommandTopicTemplate string `mapstructure:"command_topic_template"`
				} `mapstructure:"mqtt"`

				AMQP struct {
					URL                       string `mapstructure:"url"`
					EventQueueName            string `mapstructure:"event_queue_name"`
					EventRoutingKey           string `mapstructure:"event_routing_key"`
					CommandRoutingKeyTemplate string `mapstructure:"command_routing_key_template"`
				} `mapstructure:"amqp"`

				GCPPubSub struct {
					CredentialsFile         string        `mapstructure:"credentials_file"`
					ProjectID               string        `mapstructure:"project_id"`
					UplinkTopicName         string        `mapstructure:"uplink_topic_name"`
					DownlinkTopicName       string        `mapstructure:"downlink_topic_name"`
					UplinkRetentionDuration time.Duration `mapstructure:"uplink_retention_duration"`
				} `mapstructure:"gcp_pub_sub"`

				AzureIoTHub struct {
					EventsConnectionString   string `mapstructure:"events_connection_string"`
					CommandsConnectionString string `mapstructure:"commands_connection_string"`
				} `mapstructure:"azure_iot_hub"`
			} `mapstructure:"backend"`
		} `mapstructure:"gateway"`
	} `mapstructure:"network_server"`

	JoinServer struct {
		ResolveJoinEUI      bool   `mapstructure:"resolve_join_eui"`
		ResolveDomainSuffix string `mapstructure:"resolve_domain_suffix"`

		Servers []struct {
			Server  string `mapstructure:"server"`
			JoinEUI string `mapstructure:"join_eui"`
			CACert  string `mapstructure:"ca_cert"`
			TLSCert string `mapstructure:"tls_cert"`
			TLSKey  string `mapstructure:"tls_key"`
		} `mapstructure:"servers"`

		Default struct {
			Server  string `mapstructure:"server"`
			CACert  string `mapstructure:"ca_cert"`
			TLSCert string `mapstructure:"tls_cert"`
			TLSKey  string `mapstructure:"tls_key"`
		} `mapstructure:"default"`

		KEK struct {
			Set []KEK `mapstructure:"set"`
		} `mapstructure:"kek"`
	} `mapstructure:"join_server"`

	Roaming struct {
		ResolveNetIDDomainSuffix string `mapstructure:"resolve_netid_domain_suffix"`

		API struct {
			Bind    string `mapstructure:"bind"`
			CACert  string `mapstructure:"ca_cert"`
			TLSCert string `mapstructure:"tls_cert"`
			TLSKey  string `mapstructure:"tls_key"`
		} `mapstructure:"api"`

		Servers []RoamingServer      `mapstructure:"servers"`
		Default DefaultRoamingServer `mapstructure:"default"`

		KEK struct {
			Set []KEK `mapstructure:"set"`
		} `mapstructure:"kek"`
	} `mapstructure:"roaming"`

	NetworkController struct {
		Client nc.NetworkControllerServiceClient `mapstructure:"client"`

		Server  string `mapstructure:"server"`
		CACert  string `mapstructure:"ca_cert"`
		TLSCert string `mapstructure:"tls_cert"`
		TLSKey  string `mapstructure:"tls_key"`
	} `mapstructure:"network_controller"`

	Metrics struct {
		Timezone string `mapstructure:"timezone"`

		Prometheus struct {
			EndpointEnabled    bool   `mapstructure:"endpoint_enabled"`
			Bind               string `mapstructure:"bind"`
			APITimingHistogram bool   `mapstructure:"api_timing_histogram"`
		} `mapstructure:"prometheus"`
	} `mapstructure:"metrics"`

	Monitoring struct {
		Bind                         string `mapstructure:"bind"`
		PrometheusEndpoint           bool   `mapstructure:"prometheus_endpoint"`
		PrometheusAPITimingHistogram bool   `mapstructure:"prometheus_api_timing_histogram"`
		HealthcheckEndpoint          bool   `mapstructure:"healthcheck_endpoint"`
	} `mapstructure:"monitoring"`
}

type RoamingServer struct {
	NetID                  lorawan.NetID
	NetIDString            string        `mapstructure:"net_id"`
	Async                  bool          `mapstructure:"async"`
	AsyncTimeout           time.Duration `mapstructure:"async_timeout"`
	PassiveRoaming         bool          `mapstructure:"passive_roaming"`
	PassiveRoamingLifetime time.Duration `mapstructure:"passive_roaming_lifetime"`
	PassiveRoamingKEKLabel string        `mapstructure:"passive_roaming_kek_label"`
	Server                 string        `mapstructure:"server"`
	CACert                 string        `mapstructure:"ca_cert"`
	TLSCert                string        `mapstructure:"tls_cert"`
	TLSKey                 string        `mapstructure:"tls_key"`
}

type DefaultRoamingServer struct {
	Enabled                bool          `mapstructure:"enabled"`
	Async                  bool          `mapstructure:"async"`
	AsyncTimeout           time.Duration `mapstructure:"async_timeout"`
	PassiveRoaming         bool          `mapstructure:"passive_roaming"`
	PassiveRoamingLifetime time.Duration `mapstructure:"passive_roaming_lifetime"`
	PassiveRoamingKEKLabel string        `mapstructure:"passive_roaming_kek_label"`
	Server                 string        `mapstructure:"server"`
	CACert                 string        `mapstructure:"ca_cert"`
	TLSCert                string        `mapstructure:"tls_cert"`
	TLSKey                 string        `mapstructure:"tls_key"`
}

type KEK struct {
	Label string `mapstructure:"label"`
	KEK   string `mapstructure:"kek"`
}
*/

type NsConfig struct {
	NetworkServer struct {
		NetID string `mapstructure:"net_id"`

		Band struct {
			Name string `mapstructure:"name"`
		} `mapstructure:"band"`

		NetworkSettings struct {
			InstallationMargin float64 `mapstructure:"installation_margin"`
			RX1Delay           int     `mapstructure:"rx1_delay"`
			RX1DROffset        int     `mapstructure:"rx1_dr_offset"`
			RX2DR              int     `mapstructure:"rx2_dr"`
			RX2Frequency       int     `mapstructure:"rx2_frequency"`
			// DownlinkTXPower       int                 `mapstructure:"downlink_tx_power"`
			EnabledUplinkChannels []int               `mapstructure:"enabled_uplink_channels"`
			DisableADR            bool                `mapstructure:"disable_adr"`
			ExtraChannels         []band.ExtraChannel `mapstructure:"extra_channels"`
			// ExtraChannels         []struct {
			// 	Frequency int `mapstructure:"frequency" toml:"frequency"`
			// 	MinDR     int `mapstructure:"min_dr" toml:"min_dr"`
			// 	MaxDR     int `mapstructure:"max_dr" toml:"max_dr"`
			// } `mapstructure:"extra_channels"`
		} `mapstructure:"network_settings"`
	} `mapstructure:"network_server"`
}

func (c *NsConfig) ApiMode(adrDrMin int, adrDrMax int) *api.GateWayMode {
	return &api.GateWayMode{
		Mode: "NS",
		ModeConfig: &api.GateWayMode_Ns{
			Ns: &api.BuiltInNetworkServer{
				NetworkId: c.NetworkServer.NetID,
				Adr: &api.NSADR{
					Enable: !c.NetworkServer.NetworkSettings.DisableADR,
					Margin: int32(c.NetworkServer.NetworkSettings.InstallationMargin),

					DrIdMin: int32(adrDrMin),
					DrIdMax: int32(adrDrMax),
				},
				Rx1: &api.NSRX1{
					DrOffset: int32(c.NetworkServer.NetworkSettings.RX1DROffset),
					Delay:    int32(c.NetworkServer.NetworkSettings.RX1Delay),
				},
				Rx2: &api.NSRX2{
					DrIndex: int32(c.NetworkServer.NetworkSettings.RX2DR),
					Freq:    int32(c.NetworkServer.NetworkSettings.RX2Frequency),
				},
				DwellTimeLimit:                &api.NSDwellTimeLimit{Uplink: 0, Downlink: 0},
				DownlinkTxPower:               0,
				DisableFrameCounterValidation: false,
				EndDeviceStatusQueryInterval:  0,
				StatisticsPeriod:              0,
			},
		},
	}
}

func NewNsConfig(b band.Band, a *api.BuiltInNetworkServer) *NsConfig {
	c := &NsConfig{}

	c.NetworkServer.NetID = a.NetworkId
	c.NetworkServer.NetworkSettings.DisableADR = !a.Adr.Enable
	c.NetworkServer.NetworkSettings.InstallationMargin = float64(a.Adr.Margin)
	c.NetworkServer.NetworkSettings.RX1DROffset = int(a.Rx1.DrOffset)
	c.NetworkServer.NetworkSettings.RX1Delay = int(a.Rx1.Delay)
	c.NetworkServer.NetworkSettings.RX2DR = int(a.Rx2.DrIndex)
	c.NetworkServer.NetworkSettings.RX2Frequency = int(a.Rx2.Freq)
	c.NetworkServer.Band.Name = b.String()

	c.NetworkServer.NetworkSettings.ExtraChannels = b.GetExtraChannels()
	// for _, ch := range chs {
	// 	c.NetworkServer.NetworkSettings.ExtraChannels = append(c.NetworkServer.NetworkSettings.ExtraChannels, struct {
	// 		Frequency int `mapstructure:"frequency" toml:"frequency"`
	// 		MinDR     int `mapstructure:"min_dr" toml:"min_dr"`
	// 		MaxDR     int `mapstructure:"max_dr" toml:"max_dr"`
	// 	}{ch.Frequency, ch.MinDR, ch.MaxDR})
	// }

	c.NetworkServer.NetworkSettings.EnabledUplinkChannels = b.GetUplinkChannels()

	return c
}
