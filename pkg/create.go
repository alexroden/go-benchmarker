package pkg

import (
	"fmt"
	"os"
	"os/exec"
)

func Create(fileName string) error {
	cmd := exec.Command("sh", "-c", fmt.Sprintf("go test ./... -bench . -benchmem > %s", fileName))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}

	return nil
}

