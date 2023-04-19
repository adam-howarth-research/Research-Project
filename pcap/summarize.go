package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"sort"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ConverstationStats struct {
	TCPPortUsage      map[uint16]int
	UDPPortUsage      map[uint16]int
	TotalPacketCount  int
	TotalConvos       int
	TotalInternal     int
	TotalExternal     int
}

func (cs *ConverstationStats) Update(packet gopacket.Packet) {

	srcSubnet, dstSubnet := GetSubnetInfo(packet)

	transportLayer := packet.TransportLayer()
	if transportLayer != nil {
		switch transportLayer.LayerType() {
		case layers.LayerTypeTCP:
			tcp := transportLayer.(*layers.TCP)
			srcPort := uint16(tcp.SrcPort)
			destPort := uint16(tcp.DstPort)

			cs.TCPPortUsage[srcPort]++
			cs.TCPPortUsage[destPort]++

		case layers.LayerTypeUDP:
			udp := transportLayer.(*layers.UDP)

			srcPort := uint16(udp.SrcPort)
			destPort := uint16(udp.DstPort)

			cs.UDPPortUsage[srcPort]++
			cs.UDPPortUsage[destPort]++
		}

		cs.TotalConvos++

		if srcSubnet == dstSubnet {
			cs.TotalInternal++
		}else{
			cs.TotalExternal++
		}
	}
}

func NewConverstationStats() *ConverstationStats {
	return &ConverstationStats {
		TCPPortUsage:     make(map[uint16]int),
		UDPPortUsage:     make(map[uint16]int),
		TotalPacketCount: 0,
		TotalConvos:      0,
		TotalInternal:    0,
		TotalExternal:    0,
	}
}

type SubnetStats struct {
	SrcConverstation *ConverstationStats
	DstConverstation *ConverstationStats
}

func NewSubnetStats() *SubnetStats {
	return &SubnetStats{
		SrcConverstation: NewConverstationStats(), 
		DstConverstation: NewConverstationStats(), 
	}
}

type HostStats struct {
	SrcConverstation *ConverstationStats
	DstConverstation *ConverstationStats
}

func NewHostStats() *HostStats {
	return &HostStats{
		SrcConverstation: NewConverstationStats(), 
		DstConverstation: NewConverstationStats(), 
	}
}


func GetSubnetInfo(packet gopacket.Packet) (string, string){
	
	srcIP, destIP := GetHostInfo(packet)

	srcSubnet := srcIP.Mask(srcIP.DefaultMask()).String()
	dstSubnet := destIP.Mask(destIP.DefaultMask()).String()

	return srcSubnet, dstSubnet
}

func GetHostInfo(packet gopacket.Packet) (net.IP, net.IP){
	networkLayer := packet.NetworkLayer()
	if networkLayer == nil {
		return nil, nil 	
	}

	ipv4, ok := networkLayer.(*layers.IPv4)
	if !ok {
		return nil, nil
	}

	srcIP := ipv4.SrcIP
	destIP := ipv4.DstIP

	return srcIP, destIP
}


func main() {

	rand.Seed(888)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if len(os.Args) != 2 {
		log.Error().Msgf("Usage: ./summarize <path_to_pcap> %d", len(os.Args))
		os.Exit(1)
	}

	path := os.Args[1]
	
	if path == "" {
		log.Error().Msg("No Pcap Path Provided")
		os.Exit(1)
	}

	// Read pcap file
	handle, err := pcap.OpenOffline(path)
	if err != nil {
		log.Error().Msgf("%v", err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	subnetStatsMap := make(map[string]*SubnetStats)
	hostStatsMap   := make(map[string]*HostStats)

	for packet := range packetSource.Packets() {

		srcSubnet, dstSubnet := GetSubnetInfo(packet)

		srcIP, destIP := GetHostInfo(packet)

		statsSrc, ok := subnetStatsMap[srcSubnet]
		if !ok {
			statsSrc = NewSubnetStats()
			subnetStatsMap[srcSubnet] = statsSrc
		}else{
			statsSrc.SrcConverstation.Update(packet)
		}

		statsSrcHost, ok := hostStatsMap[srcIP.String()]
		if !ok {
			statsSrcHost = NewHostStats()
			hostStatsMap[srcIP.String()] = statsSrcHost
		}else{
			statsSrcHost.SrcConverstation.Update(packet)
		}

		statsDst, ok := subnetStatsMap[dstSubnet]
		if !ok {
			statsDst = NewSubnetStats()
			subnetStatsMap[dstSubnet] = statsDst
		}else{
			statsDst.DstConverstation.Update(packet)
		}
		
		statsDstHost, ok := hostStatsMap[destIP.String()]
		if !ok {
			statsDstHost = NewHostStats()
			hostStatsMap[destIP.String()] = statsDstHost
		}else{
			statsDstHost.DstConverstation.Update(packet)
		}
		
	}

	displaySubnetStats(subnetStatsMap)

	simulatedEvents := generateSimulatedEvents(subnetStatsMap, hostStatsMap, 5)
	displaySimulatedEvents(simulatedEvents)
}

func displaySubnetStats(subnetStatsMap map[string]*SubnetStats) {
	for subnet, stats := range subnetStatsMap {
			fmt.Printf("%s: %v\n", subnet, stats)
	}
}

func selectRandomItem(items []string, probabilities []float32) string {
    if len(items) != len(probabilities) {
        return ""
    }

    randVal := rand.Float32()
    selectedIndex := sort.Search(len(probabilities), func(i int) bool { return probabilities[i] >= randVal })

    if selectedIndex >= len(items) {
        selectedIndex = len(items) - 1
    }

    return items[selectedIndex]
}


func generateSimulatedEvents(subnetStatsMap map[string]*SubnetStats, hostStatsMap map[string]*HostStats, recordCount int) []string {
    var simulatedEvents []string

    // Prepare subnet selection
    subnetLabels, subnetProbabilities := prepareSelection(subnetStatsMap)

    // Prepare host selection
    hostLabels, hostProbabilities := prepareSelection(hostStatsMap)

    for i := 0; i < recordCount; i++ {
        // Select source and destination subnets
        srcSubnet := selectRandomItem(subnetLabels, subnetProbabilities)
        dstSubnet := selectRandomItem(subnetLabels, subnetProbabilities)

        // Select source and destination IPs within the selected subnets
        srcIP := selectRandomIP(hostLabels, hostProbabilities, srcSubnet)
        dstIP := selectRandomIP(hostLabels, hostProbabilities, dstSubnet)

        // Select source and destination ports, and protocol
        srcPort, dstPort, protocol := selectPortsAndProtocol(hostStatsMap[srcIP], hostStatsMap[dstIP])

        // Create simulated network event
        event := fmt.Sprintf("%s, %s, %d, %d, %s", srcIP, dstIP, srcPort, dstPort, protocol)
        simulatedEvents = append(simulatedEvents, event)
    }

    return simulatedEvents
}

func selectRandomIP(ips []string, probabilities []float32, subnet string) string {
    var filteredIPs []string
    var filteredProbabilities []float32

    for idx, ip := range ips {
        ipSubnet := net.ParseIP(ip).Mask(net.ParseIP(ip).DefaultMask()).String()
        if ipSubnet == subnet {
            filteredIPs = append(filteredIPs, ip)
            filteredProbabilities = append(filteredProbabilities, probabilities[idx])
        }
    }

    return selectRandomItem(filteredIPs, filteredProbabilities)
}


func selectPortsAndProtocol(srcHostStats, dstHostStats *HostStats) (uint16, uint16, string) {
    var protocol string
    var srcPort, dstPort uint16

    if srcHostStats == nil || dstHostStats == nil {
        srcPort = uint16(rand.Intn(65536))
        dstPort = uint16(rand.Intn(65536))
        protocol = "TCP"

        return srcPort, dstPort, protocol
    }

    // Calculate total conversations for both TCP and UDP for source and destination hosts
	srcTotalTCP := 0
	if srcHostStats != nil {
		for _, v := range(srcHostStats.SrcConverstation.TCPPortUsage) {
			srcTotalTCP += v
		}
	}
	srcTotalUDP := 0
	if srcHostStats != nil {
		for _, v := range(srcHostStats.SrcConverstation.UDPPortUsage) {
			srcTotalUDP += v
		}
	}
	dstTotalTCP := 0
	if dstHostStats != nil {
		for _, v := range(dstHostStats.SrcConverstation.TCPPortUsage) {
			dstTotalTCP += v
		}
	}
	dstTotalUDP := 0
	if dstHostStats != nil {
		for _, v := range(dstHostStats.SrcConverstation.UDPPortUsage) {
			dstTotalUDP += v
		}
	}

    // Determine the protocol (TCP or UDP) based on the sum of TCP and UDP conversations
    if srcTotalTCP+dstTotalTCP >= srcTotalUDP+dstTotalUDP {
        protocol = "TCP"
    } else {
        protocol = "UDP"
    }

    // Helper function to select a random port based on its usage
	selectRandomPort := func(portUsage map[uint16]int) uint16 {
		var ports []uint16
		var counts []int
		for port, count := range portUsage {
			ports = append(ports, port)
			counts = append(counts, count)
		}
	
		totalSum := 0
		for _, count := range counts {
			totalSum += count
		}
	
		probabilities := make([]float32, len(counts))
		compoundingBase := float32(0)
	
		for idx, count := range counts {
			probVal := float32(count) / float32(totalSum)
			compoundingBase += probVal
			probabilities[idx] = compoundingBase
		}
	
		randVal := rand.Float32()
		selectedIndex := sort.Search(len(probabilities), func(i int) bool { return probabilities[i] > randVal })
	
		if selectedIndex >= len(ports) {
			selectedIndex = len(ports) - 1
		}
		if selectedIndex < 0 {
			selectedIndex = 0
		}

		if len(ports) == 0 {
			return 65336
		}
	
		return ports[selectedIndex]
	}
	
	
	if srcHostStats != nil && dstHostStats != nil {

    // Select source and destination ports based on the chosen protocol
    if protocol == "TCP" {
        srcPort = selectRandomPort(srcHostStats.SrcConverstation.TCPPortUsage)
        dstPort = selectRandomPort(dstHostStats.DstConverstation.TCPPortUsage)
    } else {
        srcPort = selectRandomPort(srcHostStats.SrcConverstation.UDPPortUsage)
        dstPort = selectRandomPort(dstHostStats.DstConverstation.UDPPortUsage)
    }
	}
    return srcPort, dstPort, protocol
}



func prepareSelection(statsMap interface{}) ([]string, []float32) {
    var labels []string
    var totalCounts []int

    switch stats := statsMap.(type) {
    case map[string]*SubnetStats:
        for k, v := range stats {
            labels = append(labels, k)
            totalCounts = append(totalCounts, v.SrcConverstation.TotalConvos+v.DstConverstation.TotalConvos)
        }
    case map[string]*HostStats:
        for k, v := range stats {
            labels = append(labels, k)
            totalCounts = append(totalCounts, v.SrcConverstation.TotalConvos+v.DstConverstation.TotalConvos)
        }
    default:
        panic("unsupported statsMap type")
    }

    totalSum := 0
    for _, count := range totalCounts {
        totalSum += count
    }

    probabilities := make([]float32, len(totalCounts))
    compoundingBase := float32(0)

    for idx, count := range totalCounts {
        probVal := float32(count) / float32(totalSum)
        compoundingBase += probVal
        probabilities[idx] = compoundingBase
    }

	if len(labels) != len(probabilities) {
		return nil, nil
	}

    return labels, probabilities
}

func filterIPsBySubnet(ips []string, subnet string) []string {
    var filteredIPs []string
    for _, ip := range ips {
        ipSubnet := net.ParseIP(ip).Mask(net.ParseIP(ip).DefaultMask()).String()
        if ipSubnet == subnet {
            filteredIPs = append(filteredIPs, ip)
        }
    }
    return filteredIPs
}



func IntExtDecisonPoint(s *SubnetStats) float32{
	return float32(s.SrcConverstation.TotalInternal/s.SrcConverstation.TotalConvos)
}

func displaySimulatedEvents(simulatedEvents []string) {
	fmt.Println("Simulated Network Events:")
	for _, event := range simulatedEvents {
		fmt.Println(event)
	}
}
