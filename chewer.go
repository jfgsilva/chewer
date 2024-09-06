package chewer

import (
	"bufio"
	"fmt"
	"os"
)

// Processor defines an interface for processing input data.
type Processor interface {
	Process(input []string) ([]string, error)
}

// Ingest reads from stdin or files provided as arguments.
func Ingest(args []string) ([]string, error) {
	var input []string

	if len(args) == 0 {
		// Read from stdin
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input = append(input, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("error reading from stdin: %v", err)
		}
	} else {
		// Read from files
		for _, file := range args {
			f, err := os.Open(file)
			if err != nil {
				return nil, fmt.Errorf("error opening file %s: %v", file, err)
			}
			defer f.Close()

			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				input = append(input, scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				return nil, fmt.Errorf("error reading file %s: %v", file, err)
			}
		}
	}

	return input, nil
}

// Chew processes the input data using the provided Processor implementation.
func Chew(input []string, p Processor) ([]string, error) {
	return p.Process(input)
}

// Spit writes the processed data to stdout. Errors are written to stderr.
func Spit(output []string, err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}

	for _, line := range output {
		fmt.Println(line)
	}
}
