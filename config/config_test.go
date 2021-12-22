package config

import (
	"fmt"
	"testing"
)

func TestLoadConfig(*testing.T) {

	config := "../bin/config.toml"
	err := LoadConfigFile(config)
	fmt.Println(err)
	fmt.Println(GlobalConfig)
}
