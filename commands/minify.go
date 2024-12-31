package command

import "github.com/akrck02/littlestyles/configuration"

// Minify the css file
func Minify(current_configuration *configuration.Configuration) {

  println("Minifying the CSS files with the following configuration:")
  configuration.Print(current_configuration)

}

