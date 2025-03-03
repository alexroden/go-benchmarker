package pkg

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	red   = "\033[31m" // Red for increased time
	green = "\033[32m" // Green for decreased time
	reset = "\033[0m"  // Reset to default color
)

var benchmarkRegex = regexp.MustCompile(`^ok\s+([\S]+)\s+([\d\.]+)s$`)

func Compare(oldFile, newFile string) error {
	oldBenchmarks, err := parseFile(oldFile)
	if err != nil {
		return err
	}

	newBenchmarks, err := parseFile(newFile)
	if err != nil {
		return err
	}

	fmt.Println("Benchmark Time Differences:")
	fmt.Println("---------------------------------")
	for service, oldTime := range oldBenchmarks {
		if newTime, exists := newBenchmarks[service]; exists {
			diff := newTime - oldTime
			changeIndicator := ""
			color := reset

			if diff > 0.05 {
				color = red
				changeIndicator = fmt.Sprintf("▲ Increased by %.3fs", diff)
			} else if diff < -0.05 {
				color = green
				changeIndicator = fmt.Sprintf("▼ Decreased by %.3fs", -diff)
			} else {
				changeIndicator = "No change"
			}

			fmt.Printf("%-50s | %s%s%s\n", service, color, changeIndicator, reset)
		}
	}

	return nil
}

func parseFile(filePath string) (map[string]float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("run create first")
	}
	defer file.Close()

	results := make(map[string]float64)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		matches := benchmarkRegex.FindStringSubmatch(line)
		if len(matches) == 3 {
			service := matches[1]
			time, err := strconv.ParseFloat(matches[2], 64)
			if err == nil {
				results[service] = time
			}
		}
	}

	return results, scanner.Err()
}
