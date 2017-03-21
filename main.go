package main

import (
	"bufio"
	"encoding/xml"
	"flag"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	// CoverageFileDelimiter - Delimiter used in Coverage Metrics
	CoverageFileDelimiter = ":"
	// CoverageLineAndColumnDelimiter - Delimiter used between line number and column number
	CoverageLineAndColumnDelimiter = "."
	// CoverageStartAndEndLineDelimiter - Delimiter used between start line number and end line number
	CoverageStartAndEndLineDelimiter = ","
	// CoveredAndUnCoveredLinesDelimiter - Delimiter used between covered lines and uncovered lines
	CoveredAndUnCoveredLinesDelimiter = " "
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var outputFile string
var inputFile string

func init() {
	flag.StringVar(&outputFile, "outputfile", "coverage.xml", "File name to write output to.")
	flag.StringVar(&inputFile, "inputfile", "gover.coverprofile", "File name to read coverage profile from.")
}

func main() {

	flag.Parse()

	file, err := os.Open(inputFile)
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	fileArr := processCoverageData(scanner)
	cov := &coverage{}
	cov.File = fileArr
	xmlForm, marshalError := xml.Marshal(cov)
	check(marshalError)
	writeError := ioutil.WriteFile(outputFile, xmlForm, 0644)
	check(writeError)
}

func processCoverageData(scanner *bufio.Scanner) []File {
	fileMap := make(map[string][]LineToCover)

	for scanner.Scan() {
		data := scanner.Text()
		cover := strings.Split(data, CoverageFileDelimiter)
		if cover[0] != "mode" {
			filePath := cover[0]
			if fileMap[filePath] == nil { // New FilePath Entry is found
				fileMap[filePath] = splitStartAndEndLineNumbers(cover[1])
			} else { // Old FilePath Entry...just append covered lines
				fileMap[filePath] = append(fileMap[filePath], splitStartAndEndLineNumbers(cover[1])...)
			}
		}
	}

	returnFileArrays := make([]File, 0)
	// Loop over the map.
	for file, linesToCover := range fileMap {
		returnFileArrays = append(returnFileArrays, File{
			Path:        file,
			LineToCover: linesToCover,
		})
	}

	return returnFileArrays
}

func splitStartAndEndLineNumbers(startAndEndLines string) []LineToCover {
	lToCover := make([]LineToCover, 0)
	covered := false
	startAndEndLineNumbers := strings.Split(startAndEndLines, CoverageStartAndEndLineDelimiter)
	coveredAndUnCoveredLines := strings.Split(startAndEndLineNumbers[1], CoveredAndUnCoveredLinesDelimiter)
	startLine := fetchLineFromLineAndColumn(startAndEndLineNumbers[0])
	endLine := fetchLineFromLineAndColumn(coveredAndUnCoveredLines[0])
	if coveredAndUnCoveredLines[2] == "1" {
		covered = true
	}
	for i := startLine; i < endLine; i++ {
		lToCover = append(lToCover, LineToCover{
			LineNumber: i,
			Covered:    covered,
		})
	}
	return lToCover
}

func fetchLineFromLineAndColumn(lineWithColumn string) (lineNumber int) {
	lineAndColumn := strings.Split(lineWithColumn, CoverageLineAndColumnDelimiter)
	i, err := strconv.Atoi(lineAndColumn[0])
	check(err)
	return i
}
