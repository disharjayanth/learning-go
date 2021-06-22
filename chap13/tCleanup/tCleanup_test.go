package tcleanup

import (
	"os"
	"testing"
)

func createFile(t *testing.T) (string, error) {
	f, err := os.Create("tempFile.txt")
	if err != nil {
		return "", err
	}

	f.Write([]byte("Hello World!"))

	t.Cleanup(func() {
		os.Remove(f.Name())
	})

	return f.Name(), nil
}

func TestFileProcessing(t *testing.T) {
	fname, err := createFile(t)
	if err != nil {
		t.Fatal(err)
	}

	if fname != "tempFile.txt" {
		t.Errorf("Expected %v got %v", "tempFile.txt", fname)
	}
}
