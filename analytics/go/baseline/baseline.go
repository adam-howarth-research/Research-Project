package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"


)

func compute(nodes int, factor int) {
	// Simulate the time complexity by performing operations proportional to nodes * factor
	operations := nodes * factor

	for i := 0; i < operations; i++ {
		// Perform a simple operation (e.g., addition) to simulate the computation
		_ = rand.Float64() + rand.Float64()
	}
}

func main() {

	log.Info().Msgf("%v", os.Args)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if len(os.Args) != 4 {
		log.Error().Msg("Usage: ./baseline <node_count> <time_complexity_factor>")
		os.Exit(1)
	}

	nodes, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Error().Msg("Error: Invalid node_count")
		os.Exit(1)
	}

	factor, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Error().Msg("Error: Invalid time_complexity_factor")
		os.Exit(1)
	}

	freq, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Error().Msg("Error: Invalid run frequency")
		os.Exit(1)
	}

	ticker := time.NewTicker(time.Duration(freq) * time.Second)

	for {
		select {
		case <-ticker.C:
			start := time.Now()
			compute(nodes, factor)
			elapsed := time.Since(start)
			elapsedMS := elapsed.Microseconds()

			log.Info().Time("start", start).Int("nodes", nodes).Int("factor", factor).Int64("elapsed", elapsedMS).Msgf("Computation with node count %d and time complexity factor %d took %s\n", nodes, factor, elapsed)

		
			log.Debug().Msgf("Computation with node count %d and time complexity factor %d took %s\n", nodes, factor, elapsed)
			}
	}

}
