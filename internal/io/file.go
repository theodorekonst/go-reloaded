package io

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func WriteFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

func CheckOverwrite(filename string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil
	}

	fmt.Print("Overwrite? (y/n): ")
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(strings.ToLower(response))

	if response != "y" {
		return fmt.Errorf("file exists")
	}
	return nil
}
