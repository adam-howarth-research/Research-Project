package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Run "docker stats" command with the "--no-stream" flag to get a single snapshot
	beforeCMD := time.Now().Unix()
	cmd := exec.Command("docker", "stats", "--no-stream", "--format", "{{.Container}},{{.Name}},{{.CPUPerc}},{{.MemUsage}},{{.NetIO}},{{.BlockIO}},{{.MemPerc}},{{.PIDs}}")
	afterCMD := time.Now().Unix()

	output, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// Create a new CSV file
	csvFile, err := os.Create(fmt.Sprintf("docker_stats_%v.csv", beforeCMD))
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	scanner := bufio.NewScanner(output)
	headerWritten := false

	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Split(line, ",")

		columns = append(columns, strconv.FormatInt(beforeCMD, 10))
		columns = append(columns, strconv.FormatInt(afterCMD, 10))

		// Write header if it's not written yet
		if !headerWritten {
			header := []string{"Container", "Name", "CPU %", "Mem Usage", "Net I/O", "Block I/O", "Mem %", "PIDs"}
			if err := csvWriter.Write(header); err != nil {
				log.Fatal(err)
			}
			headerWritten = true
		}

		// Write the stats to the CSV file
		if err := csvWriter.Write(columns); err != nil {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Docker stats have been written to the CSV file.")
}
