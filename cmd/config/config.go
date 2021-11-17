package main

import (
	"github.com/spf13/viper"
)

type NATS struct {
	Addresses          string
	SearchEventSubject string
	SearchEventQueue   string
}

type Config struct {
	Nats NATS
}

func newConfig() *Config {
	return &Config{
		Nats: NATS {
			Addresses: viper.GetString("nats.addresses"),
			SearchEventSubject: viper.GetString("nats.searchEventSubject"),
		},
	}
}
