package core

import (
	"time"

	cfg_builder "github.com/quack-pot/cfg_builder/pkg"
)

type DbOptions struct {
	MinConnections uint32 `json:"min_connections"`
	MaxConnections uint32 `json:"max_connections"`

	MaxConnectionLifetime cfg_builder.DurationString `json:"max_connection_lifetime"`
	MaxConnectionIdleTime cfg_builder.DurationString `json:"max_connection_idle_time"`

	HealthCheckPeriod          cfg_builder.DurationString `json:"health_check_period"`
	EstablishConnectionTimeout cfg_builder.DurationString `json:"establish_connection_timeout"`
}

func DefaultDbOptions() *DbOptions {
	return &DbOptions{
		MinConnections: 5,
		MaxConnections: 15,

		MaxConnectionLifetime: cfg_builder.DurationString(time.Duration(15) * time.Minute),
		MaxConnectionIdleTime: cfg_builder.DurationString(time.Duration(5) * time.Minute),

		HealthCheckPeriod:          cfg_builder.DurationString(time.Duration(15) * time.Minute),
		EstablishConnectionTimeout: cfg_builder.DurationString(time.Duration(30) * time.Second),
	}
}
