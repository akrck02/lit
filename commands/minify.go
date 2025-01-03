package command

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/akrck02/littlestyles/configuration"
	"github.com/akrck02/littlestyles/data"
)

// Minify the css file.
func Minify(config *configuration.Configuration) {

	startTime := time.Now()

	println("Minifying the CSS files with the following configuration.")
	configuration.Print(config)

	// If input file does not exists, raise an error
	if !data.PathExists(config.Input) {
		println("ERROR: input file does not exist.")
		return
	}

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

	// Open the input file
	println(fmt.Sprintf("\nFile tree from %s", config.Input))
	inputFile, err := os.Open(config.Input)
	if nil != err {
		log.Fatal(err)
	}
	defer inputFile.Close()

	// Get file info
	info, err := inputFile.Stat()
	if nil != err {
		log.Fatal(err)
	}

	// If the file is a directory ignore
	if info.IsDir() {
		return
	}

	// Get file extension, ignore if is not a css file.
	extension := filepath.Ext(inputFile.Name())
	if ".css" != strings.ToLower(extension) {
		return
	}

	// Get the lines
	lines := addFile(config, &[]string{}, inputFile, 0)

	// Replace all the comments
	contentString := strings.Join(*lines, "")
	commentRegex := regexp.MustCompile(`/\*([^*]|[\r\n])*\*/`)
	contentString = commentRegex.ReplaceAllString(contentString, "")

	commentRegex = regexp.MustCompile(`/\*(.)*\*/`)
	contentString = strings.TrimSpace(commentRegex.ReplaceAllString(contentString, ""))

	// Write to file
	outputFile.WriteString(contentString)

	// Get file size
	stats, err := outputFile.Stat()
	if nil != err {
		log.Fatal(err)
	}
	size := stats.Size()

	// Log the processing time
	println(fmt.Sprintf("File processed (%db) in %dµs.", size, time.Now().Sub(startTime).Microseconds()))

}

// Add a css file to output file
func addFile(config *configuration.Configuration, lines *[]string, inputFile *os.File, treeLevel int) *[]string {

	if config.Trace {
		println(fmt.Sprintf("%s⤷ %s/%s", strings.Repeat("    ", treeLevel), path.Base(path.Dir(inputFile.Name())), path.Base(inputFile.Name())))
	}

	// Read the file line by line using scanner
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		newLines := processCssLine(config, lines, inputFile, scanner.Text(), treeLevel)
		lines = &newLines
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	return lines
}

// Process a css line
func processCssLine(config *configuration.Configuration, lines *[]string, inputFile *os.File, line string, treeLevel int) []string {

	// If it is a local import statement, try to access the file
	if strings.Contains(line, "@import") && !strings.Contains(line, "http") {

		// Get the referenced url inside @import statements, ignore if not valid
		line = getImportUrlFromLine(line)
		if "" == line {
			return *lines
		}

		// Get local url inside statement as local path
		referencedUrl := path.Clean(fmt.Sprintf("%s/%s", filepath.Dir(inputFile.Name()), line))
		referencedFile, err := os.Open(referencedUrl)
		if nil != err {
			log.Fatal(err)
		}
		defer referencedFile.Close()
		lines = addFile(config, lines, referencedFile, treeLevel+1)
		return *lines
	}

	// Return if the line is empty
	line = trimLine(line)
	if "" == line {
		return *lines
	}

	line = strings.ReplaceAll(line, "  ", "")

	// Write the line to the file
	line = line + " "
	if true == config.Readable {
		line += "\n"
	}
	return append(*lines, line)
}

// Get import url line
func getImportUrlFromLine(line string) string {

	standardImportRegex, _ := regexp.Compile(`(?mi)("(.*?)")`)
	matches := standardImportRegex.FindAllStringSubmatch(line, -1)

	// If an url is present return.
	if 0 != len(matches) {
		return matches[0][2]
	}

	// if not found, return empty.
	return ""
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
