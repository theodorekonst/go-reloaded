package internal_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"go-reloaded/internal/pipeline"
)

func TestGolden(t *testing.T) {
	testdataDir := "../testdata"

	files, err := os.ReadDir(testdataDir)
	if err != nil {
		t.Fatalf("Failed to read testdata directory: %v", err)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".txt") && !strings.HasSuffix(file.Name(), ".want.txt") {
			testName := strings.TrimSuffix(file.Name(), ".txt")
			t.Run(testName, func(t *testing.T) {
				inputPath := filepath.Join(testdataDir, file.Name())
				wantPath := filepath.Join(testdataDir, testName+".want.txt")

				input, err := os.ReadFile(inputPath)
				if err != nil {
					t.Fatalf("Failed to read input file %s: %v", inputPath, err)
				}

				want, err := os.ReadFile(wantPath)
				if err != nil {
					t.Fatalf("Failed to read want file %s: %v", wantPath, err)
				}

				got := pipeline.ProcessText(string(input))
				wantStr := string(want)

				if got != wantStr {
					t.Errorf("Test %s failed:\nGot:  %q\nWant: %q", testName, got, wantStr)
				}
			})
		}
	}
}
