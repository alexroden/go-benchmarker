package pkg

import (
	"fmt"
	"os"
)

func Accept() error {
	if _, err := os.Stat("benchmark.txt"); err == nil {
		os.Remove("benchmark.txt")
	}

	err := os.Rename("benchmark_new.txt", "benchmark.txt")
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}

	return nil
}
