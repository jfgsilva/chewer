# chewer Package

The `chewer` package is a Go library designed to reduce boilerplate code when building command-line tools (CLIs). It provides a simple and flexible framework for ingesting input, processing it through custom logic, and then outputting the result. The package is structured around three main functions:

1. **Ingest**: Reads input from stdin or from files provided as command-line arguments.
2. **Chew**: Processes the input using custom logic defined by implementing a `Processor` interface.
3. **Spit**: Outputs the processed data to stdout, with any errors directed to stderr.

## How `chewer` Works

### 1. Ingesting Data

The `Ingest` function reads input data from either stdin or files passed as command-line arguments. If no arguments are provided, it reads from stdin. The data is returned as a slice of strings (`[]string`).

```go
input, err := chewer.Ingest(os.Args[1:])
```

### 2. Processing Data with `Processor` Interface

The core of the `chewer` package's flexibility comes from the `Processor` interface. This interface defines a single method, `Process`, which takes the ingested data as input and returns the processed data along with any potential errors.

```go
type Processor interface {
    Process(input []string) ([]string, error)
}
```

To use `chewer` in your CLI, you create a struct that implements the `Processor` interface by defining the `Process` method with your custom logic.

### 3. Spitting Out Data

The `Spit` function takes the processed output and prints it to stdout. If there are errors, they are printed to stderr.

```go
chewer.Spit(output, err)
```

## Example: simple `gosort` CLI

The `gosort` CLI is an example of how to use the `chewer` package to build a simple tool that sorts lines of text. The core logic for sorting is implemented in a struct that satisfies the `Processor` interface.

### Implementing the `SortProcessor`

Here's how you can implement a sorting processor using the `Processor` interface:

```go
package main

import (
    "os"
    "sort"
    "github.com/jfgsilva/chewer"
)

// SortProcessor is an implementation of the Processor interface for sorting.
type SortProcessor struct{}

// Process sorts the input lines.
func (p SortProcessor) Process(input []string) ([]string, error) {
    sort.Strings(input) // Sort the lines alphabetically
    return input, nil
}

func main() {
    // Step 1: Ingest input
    input, err := chewer.Ingest(os.Args[1:])
    
    // Step 2: Chew the input with the SortProcessor
    processor := SortProcessor{}
    output, chewErr := chewer.Chew(input, processor)

    // Step 3: Spit the output
    chewer.Spit(output, chewErr)
}
```

### Building and Running `gosort`

1. **Build the executable**:
    ```bash
    go build -o gosort
    ```

2. **Run the `gosort` CLI**:

    - **From stdin**:
        ```bash
        echo -e "banana\napple\ncherry" | ./gosort
        ```
        **Output**:
        ```bash
        apple
        banana
        cherry
        ```

    - **From files**:
        ```bash
        ./gosort file1.txt file2.txt
        ```

    This will read the lines from the provided files, sort them, and output the sorted lines.

