package main

import (
	"flag"

	command "github.com/akrck02/littlestyles/commands"
	"github.com/akrck02/littlestyles/configuration"
)

// Log app title to standard output
func logAppTitle() {
	println("-------------------------------------")
	println("     Little styles by akrck02        ")
	println("-------------------------------------")
	println()
}

// Main command handle
func main() {

	logAppTitle()

	// Set flags for the cli tool
	configPathFlag := flag.String("f", "", "-f ./my/string")
	helpPathFlag := flag.Bool("h", false, "-h")
	generateConfigFlag := flag.Bool("g", false, "-g")
	traceConfigFlag := flag.Bool("t", false, "-t")
	flag.Parse()

	// Open help if help flag is present.
	if true == *helpPathFlag {
		command.Help()
		return
	}

	// Generate configuration if flag is present
	if true == *generateConfigFlag {
		command.Generate()
		return
	}

	// Load configuration from file and minify the files.
	configPath := *configPathFlag
	if "" != configPath {
		config, err := command.LoadFromFile(configPath)

		if nil != err {
			println("ERROR: Could not load configuration from file.")
			return
		}

		// Trace if flag is present
		config.Trace = *traceConfigFlag
		command.Minify(config)
		return
	}

	// Minify the CSS files using default configuration.
	config := configuration.Default()
	config.Trace = *traceConfigFlag
	command.Minify(config)

}
