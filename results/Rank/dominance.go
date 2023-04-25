package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Metric struct {
	Container string
	Name      string
	Value     float64
	Ranking   int
}

func main() {
	// Read the input from a file or use the given input as a string
	input := `Average CPU usage of pcr - dataset 1: 0.01%
Average memory usage of pcr - dataset 1: 0.06%
Average CPU usage of pcr - dataset 2: 0.02%
Average memory usage of pcr - dataset 2: 0.11%
Average CPU usage of pcr - dataset 3: 0.03%
Average memory usage of pcr - dataset 3: 0.11%
Average CPU usage of pcr - dataset 4: 0.09%
Average memory usage of pcr - dataset 4: 0.54%
Average CPU usage of klddos - dataset 1: 0.03%
Average memory usage of klddos - dataset 1: 0.06%
Average CPU usage of klddos - dataset 2: 0.17%
Average memory usage of klddos - dataset 2: 0.16%
Average CPU usage of klddos - dataset 3: 0.85%
Average memory usage of klddos - dataset 3: 0.89%
Average CPU usage of klddos - dataset 4: 8.40%
Average memory usage of klddos - dataset 4: 8.11%
Average CPU usage of baseline - dataset 1: 0.05%
Average memory usage of baseline - dataset 1: 0.05%
Average CPU usage of baseline - dataset 2: 0.05%
Average memory usage of baseline - dataset 2: 0.06%
Average CPU usage of baseline - dataset 3: 0.13%
Average memory usage of baseline - dataset 3: 0.05%
Average CPU usage of baseline - dataset 4: 0.40%
Average memory usage of baseline - dataset 4: 0.05%
Average CPU usage of bfs - dataset 1: 0.09%
Average memory usage of bfs - dataset 1: 0.11%
Average CPU usage of bfs - dataset 2: 0.70%
Average memory usage of bfs - dataset 2: 0.17%
Average CPU usage of bfs - dataset 3: 3.12%
Average memory usage of bfs - dataset 3: 0.44%
Average CPU usage of bfs - dataset 4: 52.43%
Average memory usage of bfs - dataset 4: 3.26%` 

	metrics := parseMetrics(input)
	containerMetrics := groupByContainer(metrics)
	rankDominancePerContainer(containerMetrics)

	for container, metrics := range containerMetrics {
		fmt.Printf("Container: %s\n", container)
		for _, metric := range metrics {
			fmt.Printf("%s: %.2f%% (Rank: %d)\n", metric.Name, metric.Value, metric.Ranking)
		}
		fmt.Println()
	}
}

func parseMetrics(input string) []Metric {
	lines := strings.Split(input, "\n")
	metrics := make([]Metric, len(lines))

	for i, line := range lines {
		fields := strings.Split(line, ": ")
		parts := strings.Split(fields[0], " - ")
		container, name := parts[0], parts[1]
		value, _ := strconv.ParseFloat(strings.TrimSuffix(fields[1], "%"), 64)
		metrics[i] = Metric{Container: container, Name: name, Value: value}
	}

	return metrics
}

func groupByContainer(metrics []Metric) map[string][]Metric {
	containerMetrics := make(map[string][]Metric)
	for _, metric := range metrics {
		containerMetrics[metric.Container] = append(containerMetrics[metric.Container], metric)
	}
	return containerMetrics
}

func rankDominancePerContainer(containerMetrics map[string][]Metric) {
	for _, metrics := range containerMetrics {
		sort.SliceStable(metrics, func(i, j int) bool {
			return metrics[i].Value < metrics[j].Value
		})

		rank := 1
		for i := range metrics {
			if i > 0 && metrics[i].Value != metrics[i-1].Value {
				rank++
			}
			metrics[i].Ranking = rank
		}
	}
}