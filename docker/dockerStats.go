package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msgf("%v", os.Args)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if len(os.Args) != 2 {
		log.Error().Msg("Usage: ./dockerstats <runLength>")
		os.Exit(1)
	}

	totalRunttime, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Error().Msg("Error: Invalid freq")
		os.Exit(1)
	}

	headerWritten := false
	beforeCMD := time.Now().Unix()
	csvFile, err := os.Create(fmt.Sprintf("docker_stats_%v.csv", beforeCMD))
	if err != nil {
		log.Fatal()
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	for x := 0; x < totalRunttime; x++ {

		beforeCMD = time.Now().Unix()
		cmd := exec.Command("docker", "stats", "--no-stream", "--format", "{{.Container}},{{.Name}},{{.CPUPerc}},{{.MemUsage}},{{.NetIO}},{{.BlockIO}},{{.MemPerc}},{{.PIDs}}")
		afterCMD := time.Now().Unix()

		output, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal()
		}

		if err := cmd.Start(); err != nil {
			log.Fatal()
		}

		scanner := bufio.NewScanner(output)
		for scanner.Scan() {
			line := scanner.Text()
			columns := strings.Split(line, ",")

			columns = append(columns, strconv.FormatInt(beforeCMD, 10))
			columns = append(columns, strconv.FormatInt(afterCMD, 10))

			if !headerWritten {
				header := []string{"Container", "Name", "CPU %", "Mem Usage", "Net I/O", "Block I/O", "Mem %", "PIDs"}
				if err := csvWriter.Write(header); err != nil {
					log.Fatal()
				}
				headerWritten = true
			}

			// Write the stats to the CSV file
			fmt.Println(x)
			fmt.Println(columns)

			if err := csvWriter.Write(columns); err != nil {
				log.Fatal()
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal()
		}

		if err := cmd.Wait(); err != nil {
			log.Fatal()
		}
	}

	fmt.Println("Docker stats have been written to the CSV file.")
}
