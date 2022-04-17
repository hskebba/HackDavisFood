package dataapi

import "github.com/caarlos0/env/v6"

var config struct {
	CockroachURI string `env:"COCKROACH_URI" envDefault:"postgresql://root@4054eb123985:26257/defaultdb?sslmode=disable"`
}

func parseConfig() {
	// parse environment config
	if err := env.Parse(&config); err != nil {
		panic(err)
	}
}
