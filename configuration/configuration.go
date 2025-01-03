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
	Trace    bool   `json:"trace"`
}

// Set default to configuration
func Default() *Configuration {
	return &Configuration{
		Input:    "./",
		Output:   "./",
		Name:     "master.css",
		Readable: false,
		Trace:    false,
	}
}

// Parse configuration from file
func ParseFromEnv(filePath string) (*Configuration, error) {

	err := godotenv.Load(filePath)
	if nil != err {
		return nil, err
	}

	configuration := Configuration{
		Input:    os.Getenv("INPUT"),
		Output:   os.Getenv("OUTPUT"),
		Name:     os.Getenv("NAME"),
		Readable: os.Getenv("READABLE") == "true",
		Trace:    false,
	}

	return &configuration, nil
}

// Generate an env file with current configuration
func GenerateEnv(config *Configuration) {

	lines, err := godotenv.Marshal(map[string]string{
		"INPUT":    config.Input,
		"OUTPUT":   config.Output,
		"NAME":     config.Name,
		"READABLE": fmt.Sprintf("%t", config.Readable),
	})

	// If unmarshall fails return
	if nil != err {
		println("ERROR: Cannot generate env file.")
		println(err.Error())
		return
	}

	// Open env file
	file, err := os.Create("styles.env")
	if nil != err {
		println("ERROR: Cannot generate env file.")
		println(err.Error())
		return
	}

	file.WriteString(lines)
}

// Print the configuration in standard info
func Print(config *Configuration) {
	println(fmt.Sprintf("⤷ Input:    %s", config.Input))
	println(fmt.Sprintf("⤷ Ouput:    %s", config.Output))
	println(fmt.Sprintf("⤷ Name:     %s", config.Name))
	println(fmt.Sprintf("⤷ Readable: %t", config.Readable))
	println(fmt.Sprintf("⤷ Trace: %t", config.Trace))
}
