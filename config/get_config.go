package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config is the singleton for Configuration
var Config *Configuration

// initialize creates an Configuration object from a config jsojson file
func initialize() {
	Config = new(Configuration)
	HydrateGameConfig(Config, "./config/example_game.json")
	HydrateGameConfig(Config, "./config/example_board.json")
}

func HydrateGameConfig(c *Configuration, filepath string) {
	raw, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if err := json.Unmarshal(raw, c); err != nil {
		fmt.Printf("Error initializing configuration: %s\n", err)
		os.Exit(1)
	}
}

// GetConfiguration retrieves Config or initialize it if the value is nil
func GetConfiguration() *Configuration {
	if Config == nil {
		initialize()
	}
	return Config
}
