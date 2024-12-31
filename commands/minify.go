package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/akrck02/littlestyles/configuration"
	"github.com/akrck02/littlestyles/data"
)

// Minify the css file.
func Minify(current_configuration *configuration.Configuration) {

  println("Minifying the CSS files with the following configuration:")
  configuration.Print(current_configuration)

  // If input file does not exists, raise an error
  if !data.PathExists(current_configuration.Input) {
    println("ERROR: input file does not exist.")
    return  
  }

  input_file, err := os.Open(current_configuration.Input)

  // If the output directory does not exists, create it.
  if !data.PathExists(current_configuration.Output) {
    err := os.MkdirAll(current_configuration.Output, os.ModePerm)
    if nil != err {
      println(fmt.Sprintf("Cannot create output directory %s.", current_configuration.Output))    
      return
    }
  }

  // If output file already exists, remove previous version.
  output_file_path := fmt.Sprintf("%s/%s", current_configuration.Output, current_configuration.Name)
  if data.PathExists(output_file_path) {
    os.Remove(output_file_path)
  }

  // Create output file. 
  output_file, err := os.Create(output_file_path)
  if nil != err {
    panic(err)
  }
  defer output_file.Close()

  // Access the file
  access(current_configuration, input_file, output_file)

}


// Access a file or directory.
func access(current_configuration *configuration.Configuration, input_file *os.File, output_file *os.File) {

  // Get file info
  info, err := input_file.Stat()
  if nil != err {
    panic(err)
  }

  // If the file is a directory ignore
  if info.IsDir() {
    return
  }

  // Get file extension is css add it to main file 
  if filepath.Ext(input_file.Name()) == "css" {
    addToFile(current_configuration, input_file, output_file)
  }
}

// Add file contents to master file.
func addToFile(current_configuration *configuration.Configuration, input_file *os.File, output_file *os.File) {

  
  

}
