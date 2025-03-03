package pkg

import (
	"fmt"
	"os"
	"os/exec"
)

func Create() error {
	if err := renameFile(); err != nil {
		return err
	}

	cmd := exec.Command("sh", "-c", "go test ./... -bench . -benchmem > benchmark_new.txt")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}

	return nil
}

func renameFile() error {
	if _, err := os.Stat("benchmark_new.txt"); err == nil {
		if err := Accept(); err != nil {
			return err
		}
	}

	return nil
}
