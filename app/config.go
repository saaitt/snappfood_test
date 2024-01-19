package app

import "time"

type ServiceConfig struct {
	Http struct {
		Address                 string
		GracefulShutdownTimeout time.Duration
	}
	DB     PostgresDB  `mapstructure:"db"`
	Broker KafkaBroker `mapstructure:"broker"`
}

type PostgresDB struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

type KafkaBroker struct {
	Server      string   `mapstructure:"server"`
	GroupId     string   `mapstructure:"groupId"`
	OffsetReset string   `mapstructure:"offsetReset"`
	Topics      []string `mapstructure:"topics"`
	AgentTopic  string   `mapstructure:"agentTopic"`
}
