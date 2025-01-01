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
  helpPathFlag := flag.String("h", "-", "-h")
	flag.Parse()

  // Open help if help flag is present.
	if "-" != *helpPathFlag {
	  command.Help()
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
    
    command.Minify(config)
    return 
  } 

  // Minify the CSS files using default configuration.
  config := configuration.Default() 
  command.Minify(config)

}
