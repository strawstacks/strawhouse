package config

import (
	"flag"
	"github.com/bsthun/gut"
	"github.com/caarlos0/env"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Config struct {
	WebListen   [2]*string `env:"STRAWHOUSE_WEB_LISTEN" yaml:"webListen" validate:"required"`
	ProtoListen [2]*string `env:"STRAWHOUSE_PROTO_LISTEN" yaml:"protoListen" validate:"required"`
	DataRoot    *string    `env:"STRAWHOUSE_DATA_ROOT" yaml:"dataRoot" validate:"dirpath"`
	PogrebPath  *string    `env:"STRAWHOUSE_POGREB_PATH" yaml:"pogrebPath" validate:"dirpath"`
	Key         *string    `env:"STRAWHOUSE_KEY" yaml:"key" validate:"required"`
}

func Init() *Config {
	// * Parse arguments
	path := flag.String("config", "", "Path to config file")
	flag.Parse()

	// * Declare struct
	config := new(Config)

	// * Fallback to env
	if *path == "" {
		// * Parse env
		if err := env.Parse(config); err != nil {
			gut.Fatal("Unable to parse environment variables", err)
		}
	} else {
		// * Read config
		yml, err := os.ReadFile(*path)
		if err != nil {
			gut.Fatal("Unable to read configuration file", err)
		}

		// * Parse config
		if err := yaml.Unmarshal(yml, config); err != nil {
			gut.Fatal("Unable to parse configuration file", err)
		}
	}

	// * Validate config
	if err := gut.Validate(config); err != nil {
		gut.Fatal("Invalid configuration", err)
	}

	// * Normalize config
	*config.DataRoot, _ = filepath.Abs(*config.DataRoot)

	return config
}
