package main

import (
	"fmt"
	"math"
	"math/rand"
	"net"
	"time"
	"os"
	"strconv"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Flow struct {
	SourceIP        net.IP
	DestinationIP   net.IP
	SourcePort      uint16
	DestinationPort uint16
	Protocol        uint8
	ByteCount       uint32
}

func generateIPs(nodeCount int) []net.IP {
	ips := make([]net.IP, nodeCount)
	for i := 0; i < nodeCount; i++ {
		ips[i] = net.IPv4(byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256)))
	}
	return ips
}

func generateFlows(nodeCount, edgeCount int) []Flow {
	ips := generateIPs(nodeCount)
	flows := make([]Flow, edgeCount)

	for i := 0; i < edgeCount; i++ {
		srcIP := ips[rand.Intn(nodeCount)]
		dstIP := ips[rand.Intn(nodeCount)]
		for srcIP.Equal(dstIP) {
			dstIP = ips[rand.Intn(nodeCount)]
		}

		flows[i] = Flow{
			SourceIP:        srcIP,
			DestinationIP:   dstIP,
			SourcePort:      uint16(rand.Uint32() % 65536),
			DestinationPort: uint16(rand.Uint32() % 65536),
			Protocol:        uint8(rand.Intn(3)),
			ByteCount:       rand.Uint32() % 100000,
		}
	}

	return flows
}

func createHistogram(flows []Flow, numBins int) []int {
	histogram := make([]int, numBins)
	maxBytes := uint32(0)

	for _, flow := range flows {
		if flow.ByteCount > maxBytes {
			maxBytes = flow.ByteCount
		}
	}

	binSize := float64(maxBytes) / float64(numBins)
	for _, flow := range flows {
		binIndex := int(math.Floor(float64(flow.ByteCount) / binSize))
		if binIndex >= numBins {
			binIndex = numBins - 1
		}
		histogram[binIndex]++
	}

	return histogram
}

func calculateKLD(p, q []float64) float64 {
	kld := 0.0
	for i := 0; i < len(p); i++ {
		if p[i] == 0 || q[i] == 0 {
			continue
		}
		kld += p[i] * math.Log2(p[i]/q[i])
	}
	return kld
}

func normalizeHistogram(histogram []int) []float64 {
	normalized := make([]float64, len(histogram))
	total := 0

	for _, count := range histogram {
		total += count
	}

	for i, count := range histogram {
		normalized[i] = float64(count) / float64(total)
	}

	return normalized
}

func main() {

	log.Info().Msgf("%v", os.Args)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if len(os.Args) != 4 {
		log.Error().Msg("Usage: ./baseline <node_count> <edgeSampleSize>")
		os.Exit(1)
	}

	nodes, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Error().Msg("Error: Invalid node_count")
		os.Exit(1)
	}

	edgeSampleSize, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Error().Msg("Error: Invalid edgeSampleSize")
		os.Exit(1)
	}

	freq, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Error().Msg("Error: Invalid run frequency")
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())

	nodeCount := nodes
	normalEdgeCount := edgeSampleSize
	attackEdgeCount := edgeSampleSize // Simulating an increase in traffic volume (e.g., due to a DDoS attack)

	normalFlows := generateFlows(nodeCount, normalEdgeCount)
	attackFlows := generateFlows(nodeCount, attackEdgeCount)


	ticker := time.NewTicker(time.Duration(freq) * time.Second)

	for {
		select {
		case <-ticker.C:
			start := time.Now()
			compute(normalFlows, attackFlows)
			elapsed := time.Since(start)
			elapsedMS := elapsed.Microseconds()

			log.Info().Time("start", start).Int("nodes", nodes).Int("edgesamplesize", edgeSampleSize).Int64("elapsed", elapsedMS).Msgf("Computation with node count %d and edge sample %d took %s\n", nodes, edgeSampleSize, elapsed)

		}
	}
}

func compute(flowset1, flowset2 []Flow) {
	numBins := 10

	histNormal := createHistogram(flowset1, numBins)
	histAttack := createHistogram(flowset2, numBins)

	normalizedHistNormal := normalizeHistogram(histNormal)
	normalizedHistAttack := normalizeHistogram(histAttack)

	kld := calculateKLD(normalizedHistNormal, normalizedHistAttack)
	fmt.Printf("Kullback-Leibler Divergence: %f\n", kld)

	threshold := 0.5
	if kld > threshold {
		fmt.Println("DDoS attack detected!")
	} else {
		fmt.Println("No DDoS attack detected.")
	}
}