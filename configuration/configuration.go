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
func ParseFromEnv(filePath string) (*Configuration, error) {

  err := godotenv.Load(filePath)
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
func GenerateEnv(config *Configuration) {
  godotenv.Marshal(map[string]string {
    "INPUT" : config.Input,
    "OUTPUT" : config.Output,
    "NAME" : config.Name,
    "READABLE" : fmt.Sprintf("%t", config.Readable),
  })
}

// Print the configuration in standard info
func Print(config *Configuration) {
  println(fmt.Sprintf("⤷ Input:    %s", config.Input))
  println(fmt.Sprintf("⤷ Ouput:    %s", config.Output))
  println(fmt.Sprintf("⤷ Name:     %s", config.Name))
  println(fmt.Sprintf("⤷ Readable: %t", config.Readable))
}


