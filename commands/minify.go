package command

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/akrck02/littlestyles/configuration"
	"github.com/akrck02/littlestyles/data"
)

// Minify the css file.
func Minify(config *configuration.Configuration) {

  println("Minifying the CSS files with the following configuration:")
  configuration.Print(config)

  // If input file does not exists, raise an error
  if !data.PathExists(config.Input) {
    println("ERROR: input file does not exist.")
    return  
  }

  inputFile, err := os.Open(config.Input)

  // If the output directory does not exists, create it.
  if !data.PathExists(config.Output) {
    err := os.MkdirAll(config.Output, os.ModePerm)
    if nil != err {
      println(fmt.Sprintf("ERROR: Cannot create output directory %s.", config.Output))    
      return
    }
  }

  // If output file already exists, remove previous version.
  outputFilePath := fmt.Sprintf("%s/%s", config.Output, config.Name)
  if data.PathExists(outputFilePath) {
    os.Remove(outputFilePath)
  }

  // Create output file. 
  outputFile, err := os.Create(outputFilePath)
  if nil != err {
    log.Fatal(err)
  }
  defer outputFile.Close()

  // Access the file
  println(fmt.Sprintf("\nStarting to access file tree from %s", config.Input))
  access(config, inputFile, outputFile)

}

// Access a file or directory.
func access(config *configuration.Configuration, currentFile *os.File, outputFile *os.File) {

  // Get file info
  currentFileInfo, err := currentFile.Stat()
  if nil != err {
    log.Fatal(err)
  }

  // If the file is a directory ignore
  if currentFileInfo.IsDir() {
    return
  }

  // Get file extension, if it is empty ignore file. 
  extension := filepath.Ext(currentFile.Name())
  if "" == extension {
    return
  }

  // If extension is not css ignore
  if ".css" != strings.ToLower(extension){
    return
  }

  // Add current file to main file
  currentPath := filepath.Dir(config.Input)
  addToFile(config, currentPath, currentFile, outputFile)
}

// Add file contents to master file.
func addToFile(config *configuration.Configuration, currentPath string, currentFile *os.File, outputFile *os.File) {

  println(fmt.Sprintf("Checking out file %s", currentFile.Name()))
  
  // read the file line by line using scanner
  scanner := bufio.NewScanner(currentFile)
  for scanner.Scan() {
    line := scanner.Text()

    if strings.Contains(line, "@import") && !strings.Contains(line, "http") {      

      // Get the referenced url inside @import statements
      line = strings.ReplaceAll(line, "@import", "")
      line = strings.ReplaceAll(line, "\"", "")
      line = strings.ReplaceAll(line, ";","")
      println(fmt.Sprintf("⤷ Import line: %s", line))

      url := strings.ReplaceAll(line, "url(", "")
      url = strings.ReplaceAll(line, "(", "")
      url = strings.ReplaceAll(line, ")", "")
      url = strings.TrimSpace(line)

      println(fmt.Sprintf("⤷ Attempting to open %s", url))

      _, err := currentFile.Stat()
      if nil != err {
        log.Fatal(err)
      }

      // Get local url inside statement as local path
      referencedUrl := fmt.Sprintf("%s/%s", currentPath, url)
      println(fmt.Sprintf("⤷ Referenced url: %s", referencedUrl))

      referencedFile, err := os.Open(referencedUrl)
      if nil != err {
        log.Fatal(err)
      }

      // Add file content to file
      addToFile(config, url, referencedFile, outputFile)

    } else {

      line = trimLine(line)
      if "" == line {
        break
      }

      //outputFile.WriteString(line)
      print(line)
      if true == config.Readable {
        println()
        //outputFile.WriteString("\n")
      }
    }
  }
  if err := scanner.Err(); err != nil {
        log.Fatal(err)
  }
}

// Trim line 
func trimLine(line string) string {
  line = strings.TrimSpace(line)
  return strings.ReplaceAll(line, "\n", "") 
}

// NOTE: this isn't multi-Unicode-codepoint aware, like specifying skintone or
// gender of an emoji: https://unicode.org/emoji/charts/full-emoji-modifiers.html
func substr(input string, start int, length int) string {
    asRunes := []rune(input)
    
    if start >= len(asRunes) {
        return ""
    }
    
    if start+length > len(asRunes) {
        length = len(asRunes) - start
    }
    
    return string(asRunes[start : start+length])
}
