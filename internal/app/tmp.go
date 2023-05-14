package app

import (
	"fmt"
	"log"
	"os"
)

func initTempDir() {
	tempDir := os.TempDir()

	if err := os.MkdirAll(tempDir, 1777); err != nil {
		log.Fatalf("failed to create temporary directory %s: %s", tempDir, err)
	}

	tempFile, err := os.CreateTemp("", "tmpInit_")
	if err != nil {
		log.Fatalf("failed to create tempFile: %s", err)
	}

	if _, err = fmt.Fprintf(tempFile, "Hello, World!"); err != nil {
		log.Fatalf("failed to write to tempFile: %s", err)
	}

	if err = tempFile.Close(); err != nil {
		log.Fatalf("failed to close tempFile: %s", err)
	}

	if err = os.Remove(tempFile.Name()); err != nil {
		log.Fatalf("failed to delete tempFile: %s", err)
	}

	log.Printf("u temporary directory %s", tempDir)
}
