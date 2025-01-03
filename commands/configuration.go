package command

import (
	"fmt"
	"time"

	"github.com/akrck02/littlestyles/configuration"
)

// Generate a new configuration file.
func Generate() {
  startTime := time.Now()
  configuration.GenerateEnv(configuration.Default())
  println(fmt.Sprintf("Config file styles.env generated in %dµs.", time.Now().Sub(startTime).Microseconds()))
}

// Load a configuration from file
func LoadFromFile(filePath string) (*configuration.Configuration, error) {
  return configuration.ParseFromEnv(filePath)
}
