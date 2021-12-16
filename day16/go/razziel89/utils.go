package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// Function readLine reads one line from stdin via a global reader instance.
func readLine() (string, error) {
	return reader.ReadString('\n')
}

// tag::utils[]

// ReadLinesAsLines reads in some lines.
func ReadLinesAsLines() ([]string, error) {
	result := []string{}
	for {
		line, err := readLine()
		if err == io.EOF {
			// Success case, no more input to read.
			return result, nil
		}
		if err != nil {
			return []string{}, err
		}
		line = strings.TrimSpace(line)
		result = append(result, line)
	}
}

// end::utils[]
