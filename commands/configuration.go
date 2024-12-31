package command

import "github.com/akrck02/littlestyles/configuration"

// Generate a new configuration file.
func Generate() {
  configuration.GenerateEnv(configuration.Default())
}

// Load a configuration from file
func LoadFromFile(file_path string) (*configuration.Configuration, error) {
  return configuration.ParseFromEnv(file_path)
}
