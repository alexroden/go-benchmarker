package main

import (
	"fmt"
	"os"

	"github.com/alexroden/go-benchmarker/pkg"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go-benchmarker [create|compare|accept]")
		return
	}

	switch os.Args[1] {
	case "create":
		if err := pkg.Create("benchmark.txt"); err != nil {
			fmt.Println("Error:", err)
		}
	case "compare":
		if err := pkg.Compare("benchmark.txt", "benchmark_new.txt"); err != nil {
			fmt.Println("Error:", err)
		}
	case "accept":
		if err := pkg.Accept(); err != nil {
			fmt.Println("Error:", err)
		}
	default:
		fmt.Printf("Error: %s not supported", os.Args[1])
	}
}
