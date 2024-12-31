package main

import (
	"flag"
	command "github.com/akrck02/littlestyles/commands"
	"github.com/akrck02/littlestyles/configuration"
)

// Log app title to standard output
func log_app_title() {
  println("-------------------------------------")
	println("     Little styles by akrck02        ")
	println("-------------------------------------")
	println()
}


// Main command handle
func main() {

  log_app_title()

  // Set flags for the cli tool 
	config_path_flag := flag.String("f", "", "-f ./my/string")
  help_path_flag := flag.String("h", "-", "-h")
	flag.Parse()

  // Open help if help flag is present.
	if "-" != *help_path_flag {
	  command.Help()
		return
	}

  // Load configuration from file and minify the files.
  config_path := *config_path_flag
  if "" != config_path {
    current_configuration, err := command.LoadFromFile(config_path)

    if nil != err {
      println("ERROR: Could not load configuration from file.")
      return
    }
    
    command.Minify(current_configuration)
    return 
  } 

  // Minify the CSS files using default configuration.
  current_configuration := configuration.Default() 
  command.Minify(current_configuration)

}
