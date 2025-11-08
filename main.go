package main

import (
	"fmt"
	"os"

	"go-reloaded/internal/io"
	"go-reloaded/internal/pipeline"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: goreloaded <input> <output>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	text, err := io.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	if err := io.CheckOverwrite(outputFile); err != nil {
		fmt.Printf("Cancelled: %v\n", err)
		os.Exit(1)
	}

	result := pipeline.ProcessText(text)

	if err := io.WriteFile(outputFile, result); err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		os.Exit(1)
	}
}
