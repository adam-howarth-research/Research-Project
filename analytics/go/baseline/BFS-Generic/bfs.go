package main

import (
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

type Node struct {
	ip   net.IP
	edges []*Node
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
	edgeCount := edgeSampleSize
	flows := generateFlows(nodeCount, edgeCount)

	nodeList := createNetworkFromFlows(flows)


	ticker := time.NewTicker(time.Duration(freq) * time.Second)

	for {
		select {
		case <-ticker.C:
			startNode := nodeList[rand.Intn(len(nodeList))]
			start := time.Now()
			bfs(nodeList, startNode)
			elapsed := time.Since(start)
			elapsedMS := elapsed.Microseconds()

			log.Info().Time("start", start).Int("nodes", nodes).Int("edgesamplesize", edgeSampleSize).Int64("elapsed", elapsedMS).Msgf("Computation with node count %d and edge sample %d took %s\n", nodes, edgeSampleSize, elapsed)

		}
	}
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

func createNetworkFromFlows(flows []Flow) []*Node {
	nodes := make(map[string]*Node)

	for _, flow := range flows {
		srcIPStr := flow.SourceIP.String()
		dstIPStr := flow.DestinationIP.String()

		srcNode, srcExists := nodes[srcIPStr]
		if !srcExists {
			srcNode = &Node{ip: flow.SourceIP}
			nodes[srcIPStr] = srcNode
		}

		dstNode, dstExists := nodes[dstIPStr]
		if !dstExists {
			dstNode = &Node{ip: flow.DestinationIP}
			nodes[dstIPStr] = dstNode
		}

		srcNode.edges = append(srcNode.edges, dstNode)
	}

	return nodesList(nodes)
}

func nodesList(nodesMap map[string]*Node) []*Node {
	nodes := make([]*Node, 0, len(nodesMap))
	for _, node := range nodesMap {
		nodes = append(nodes, node)
	}
	return nodes
}

func bfs(nodes []*Node, startNode *Node) {
	visited := make(map[string]bool)
	queue := make([]*Node, 0)

	queue = append(queue, startNode)
	visited[startNode.ip.String()] = true

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		for _, neighbor := range currentNode.edges {
			neighborIP := neighbor.ip.String()
			if !visited[neighborIP] {
				queue = append(queue, neighbor)
				visited[neighborIP] = true
			}
		}
	}
}


