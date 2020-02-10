package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func parseFile(path string) (map[string]int, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()
	metricNames := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for lnum := 1; scanner.Scan(); lnum++ {
		line := scanner.Text()
		if line[0:1] == "#" {
			// fmt.Printf("Skipped comment line: %s\n", line)
			continue
		} else {
			// fmt.Printf("Metric: %s\n", metric)
			metric := getMetricName(line)
			metricNames[metric] = 1
		}
	}
	return metricNames, nil
}

// Get just the metric name
func getMetricName(line string) string {
	fields := strings.Fields(line)
	metric := fields[0]
	reg := regexp.MustCompile(`{`)
	metric = reg.Split(metric, -1)[0]
	return metric
}

func main() {
	merger, err := parseFile("files/exporter-merger.txt")
	if err == nil {
		fmt.Printf("Total metrics found: %d\n", len(merger))
	}

	prometheus, err := parseFile("files/prometheus.txt")

	//Check if there's a key in prometheus.txt that doesn't exist in exporter-merger.txt
	allGood := true
	for k := range prometheus {
		if _, ok := merger[k]; !ok {
			allGood = false
			fmt.Printf("Metric not found: %s\n", k)
		}
	}
	if allGood {
		fmt.Println("All good!")
	}
}
