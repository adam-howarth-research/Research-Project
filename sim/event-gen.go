package main

import (
	"fmt"
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
	ip    net.IP
	edges []*Node
}

func main() {
	log.Info().Msgf("%v", os.Args)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if len(os.Args) != 4 {
		log.Error().Msg("Usage: ./event-gen <duration> <event_per_second> <remote_site_count>")
		os.Exit(1)
	}

	durationInt, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Error().Msg("Error: Invalid duration")
		os.Exit(1)
	}

	eventsPerSecond, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Error().Msg("Error: Invalid events per second")
		os.Exit(1)
	}

	storeCount, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Error().Msg("Error: Invalid remote_site_count")
		os.Exit(1)
	}

	rand.Seed(888)
	officeIP := net.ParseIP("192.168.0.1")
	officeNode := &Node{ip: officeIP}

	storeNodes := make([]*Node, storeCount)
	for i := range storeNodes {
		storeIP := net.ParseIP(fmt.Sprintf("192.168.%d.1", i+1))
		storeNodes[i] = &Node{ip: storeIP}
		officeNode.edges = append(officeNode.edges, storeNodes[i])
		storeNodes[i].edges = append(storeNodes[i].edges, officeNode)
	}

	storeDevices := createDevicesForStores(storeNodes)
	officeDevices := createDevicesForOffice(officeNode)

	duration := time.Duration(durationInt) * time.Second

	flows := generateFlows(storeDevices, officeDevices, duration, eventsPerSecond)

	observedNodes := make(map[string]int)
	observedEdges := make(map[string]int)

	for _, flow := range flows {
		observedNodes[flow.SourceIP.String()] = 1
		observedNodes[flow.DestinationIP.String()] = 1

		observedEdges[flow.SourceIP.String()+flow.DestinationIP.String()] = 1
	}

	fmt.Printf("Total nodes: %d\n", len(observedNodes))
	fmt.Printf("Total edges: %d\n", len(flows))
	fmt.Printf("Total Unique edges: %d\n", len(observedEdges))

}

func generateFlows(storeDevices [][]*Node, officeDevices []*Node, duration time.Duration, eventsPerSecond int) []Flow {
	var flows []Flow
	totalEvents := int(duration.Seconds()) * eventsPerSecond

	for i := 0; i < totalEvents; i++ {
		storeIndex := rand.Intn(len(storeDevices))
		deviceIndex := rand.Intn(len(storeDevices[storeIndex]))
		officeDevice := officeDevices[rand.Intn(len(officeDevices))]

		flow := Flow{
			SourceIP:        storeDevices[storeIndex][deviceIndex].ip,
			DestinationIP:   officeDevice.ip,
			SourcePort:      randomPort(),
			DestinationPort: randomPort(),
			Protocol:        randomProtocol(),
			ByteCount:       randomByteCount(),
		}
		flows = append(flows, flow)
	}

	return flows
}
func randomPort() uint16 {
	return uint16(rand.Intn(65535-1024) + 1024)
}

func randomProtocol() uint8 {
	protocols := []uint8{6, 17} // TCP (6) and UDP (17)
	return protocols[rand.Intn(len(protocols))]
}

func randomByteCount() uint32 {
	return uint32(rand.Intn(10000-100) + 100)
}


func createDevicesForStores(storeNodes []*Node) [][]*Node {
	storeDevices := make([][]*Node, len(storeNodes))

	for i, storeNode := range storeNodes {
		numPosDevices := rand.Intn(8) + 5 // Between 5 and 12 POS devices
		storeDevices[i] = make([]*Node, numPosDevices+2) // POS devices + inventory + office computer

		for j := 0; j < numPosDevices+2; j++ {
			deviceIP := net.ParseIP(fmt.Sprintf("192.168.%d.%d", i+1, j+2))
			storeDevices[i][j] = &Node{ip: deviceIP}
			storeDevices[i][j].edges = append(storeDevices[i][j].edges, storeNode)
			storeNode.edges = append(storeNode.edges, storeDevices[i][j])
		}
	}

	return storeDevices
}

func createDevicesForOffice(officeNode *Node) []*Node {
	officeDevices := make([]*Node, 60)

	for i := 0; i < 60; i++ {
		deviceIP := net.ParseIP(fmt.Sprintf("192.168.0.%d", i+2))
		officeDevices[i] = &Node{ip: deviceIP}
		officeDevices[i].edges = append(officeDevices[i].edges, officeNode)
		officeNode.edges = append(officeNode.edges, officeDevices[i])
	}

	return officeDevices
}