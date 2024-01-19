package app

import "time"

type ServiceConfig struct {
	Http struct {
		Address                 string
		GracefulShutdownTimeout time.Duration
	}
	DB struct {
		Host     string
		Port     int
		User     string
		Password string
		DbName   string
	}
	Broker KafkaBroker `mapstructure:"broker"`
}

type KafkaBroker struct {
	Server      string   `mapstructure:"server"`
	GroupId     string   `mapstructure:"groupId"`
	OffsetReset string   `mapstructure:"offsetReset"`
	Topics      []string `mapstructure:"topics"`
	AgentTopic  string   `mapstructure:"agentTopic"`
}
