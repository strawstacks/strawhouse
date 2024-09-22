package config

import (
	"flag"
	"github.com/bsthun/goutils"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Client struct {
	Name *string `yaml:"name" validate:"alphanum,lowercase,min=4,max=16"`
	Key  *string `yaml:"key" validate:"base64,len=64"`
}

type Config struct {
	WebListen   [2]*string `yaml:"webListen" validate:"required"`
	ProtoListen [2]*string `yaml:"protoListen" validate:"required"`
	DataRoot    *string    `yaml:"dataRoot" validate:"dirpath"`
	ObserveRoot *string    `yaml:"observeRoot" validate:"dirpath"`
	Clients     []*Client  `yaml:"clients"`
}

func Init() *Config {
	// * Parse arguments
	path := flag.String("config", "/etc/strawhouse/backend/config.yml", "Path to config file")
	flag.Parse()

	// * Declare struct
	config := new(Config)

	// * Read config
	yml, err := os.ReadFile(*path)
	if err != nil {
		uu.Fatal("Unable to read configuration file", err)
	}

	// * Parse config
	if err := yaml.Unmarshal(yml, config); err != nil {
		uu.Fatal("Unable to parse configuration file", err)
	}

	// * Validate config
	if err := uu.Validate(config); err != nil {
		uu.Fatal("Invalid configuration", err)
	}

	// * Normalize config
	*config.DataRoot, _ = filepath.Abs(*config.DataRoot)
	*config.ObserveRoot, _ = filepath.Abs(*config.ObserveRoot)

	return config
}
