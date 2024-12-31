package configuration

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

// This struct represents the needed configuration
// to run the minifier tool.
type Configuration struct {
	Input    string `json:"input"`
	Output   string `json:"output"`
	Name     string `json:"name"`
	Readable bool   `json:"readable"`
}

// Set default to configuration
func Default() *Configuration {
  return &Configuration {
    Input: "./",
    Output: "./",
    Name: "master.css",
    Readable: false,
  }
}

// Parse configuration from file
func ParseFromEnv(file_path string) (*Configuration, error) {

  err := godotenv.Load(file_path)
  if nil != err {
    return nil, err
  }

  configuration := Configuration {
		Input:      os.Getenv("INPUT"),
		Output:     os.Getenv("OUTPUT"),
		Name:       os.Getenv("NAME"),
		Readable:   os.Getenv("READABLE") == "true",
	}

  return &configuration, nil
}

// Generate an env file with current configuration
func GenerateEnv(configuration *Configuration) {
  godotenv.Marshal(map[string]string {
    "INPUT" : configuration.Input,
    "OUTPUT" : configuration.Output,
    "NAME" : configuration.Name,
    "READABLE" : fmt.Sprintf("%t", configuration.Readable),
  })
}

// Print the configuration in standard info
func Print(configuration *Configuration) {
  println(fmt.Sprintf("Input: %s", configuration.Input))
  println(fmt.Sprintf("Ouput: %s", configuration.Output))
  println(fmt.Sprintf("Name: %s", configuration.Name))
  println(fmt.Sprintf("Readable: %t", configuration.Readable))
}

