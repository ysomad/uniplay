package app

import (
	"fmt"
	"log"
	"os"
)

func initTempDir() {
	tempDir := os.TempDir()
	if err := os.MkdirAll(tempDir, 1777); err != nil {
		log.Fatalf("Failed to create temporary directory %s: %s", tempDir, err)
	}

	tempFile, err := os.CreateTemp("", "tmpInit_")
	if err != nil {
		log.Fatalf("Failed to create tempFile: %s", err)
	}

	if _, err = fmt.Fprintf(tempFile, "Hello, World!"); err != nil {
		log.Fatalf("Failed to write to tempFile: %s", err)
	}

	if err = tempFile.Close(); err != nil {
		log.Fatalf("Failed to close tempFile: %s", err)
	}

	if err = os.Remove(tempFile.Name()); err != nil {
		log.Fatalf("Failed to delete tempFile: %s", err)
	}

	log.Printf("Using temporary directory %s", tempDir)
}
