package main

import (
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

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

type HostTraffic struct {
	BytesSent     uint32
	BytesReceived uint32
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

func calculateProducerConsumerRatio(trafficData map[string]*HostTraffic) {
	for _, data := range trafficData {
		_ = float64(data.BytesSent) / float64(data.BytesReceived)
		//fmt.Printf("Host: %s, Producer-Consumer Ratio: %f\n", host, ratio)
	}
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if len(os.Args) != 4 {
		log.Error().Msg("Usage: ./producer_consumer_ratio <node_count> <edgeSampleSize> <freq>")
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
	edgeCount := edgeSampleSize

	flows := generateFlows(nodeCount, edgeCount)

	trafficData := make(map[string]*HostTraffic)

	for _, flow := range flows {
		srcIP := flow.SourceIP.String()
		dstIP := flow.DestinationIP.String()

		if _, exists := trafficData[srcIP]; !exists {
			trafficData[srcIP] = &HostTraffic{}
		}
		if _, exists := trafficData[dstIP]; !exists {
			trafficData[dstIP] = &HostTraffic{}
		}

		trafficData[srcIP].BytesSent += flow.ByteCount
		trafficData[dstIP].BytesReceived += flow.ByteCount
	}

	ticker := time.NewTicker(time.Duration(freq) * time.Second)

	for {
		select {
		case <-ticker.C:
			start := time.Now()
			calculateProducerConsumerRatio(trafficData)
			elapsed := time.Since(start)
			elapsedMS := elapsed.Microseconds()

			log.Info().Time("start", start).Int("nodes", nodes).Int("edgesamplesize", edgeSampleSize).Int64("elapsed", elapsedMS).Msgf("Computation with node count %d and edge sample %d took %s\n", nodes, edgeSampleSize, elapsed)
		}
	}
}

