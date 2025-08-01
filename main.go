package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type Packets []struct {
	Index  string `json:"_index"`
	Type   string `json:"_type"`
	Score  any    `json:"_score"`
	Source Source `json:"_source"`
}

type Source struct {
	Layers Layers `json:"layers"`
}

type Layers struct {
	Layer    string   `json:"layers"`
	Frames   Frames   `json:"frame"`
	Ethernet Ethernet `json:"eth"`
	IP       IP       `json:"ip"`
	TCP      TCP      `json:"tcp"`
	UDP      UDP      `json:"udp"`
	DHCP     DHCP     `json:"dhcp"`
}

// PROTOCOLS
type Frames struct {
	Frame                string               `json:"frame"`
	FrameProtocols       string               `json:"frame.protocols"`
	FrameLength          string               `json:"frame.len"`
	FrameUTC             string               `json:"frame.time_utc"`
	FrameInterfaceIDTree FrameInterfaceIDTree `json:"frame.interface_id_tree"`
}
type FrameInterfaceIDTree struct {
	FrameInterfaceName        string `json:"frame.interface_name"`
	FrameInterfaceDescription string `json:"frame.interface_description"`
}

type Ethernet struct {
	EthernetSourceTree EthernetSourceTree `json:"eth.src_tree"`
	EthernetDST        string             `json:"eth.dst"`
	EthernetSRC        string             `json:"eth.src"`
}
type EthernetSourceTree struct {
	OUIResolved string `json:"eth.src.oui_resolved"`
}

type IP struct {
	IPSource      string `json:"ip.src"`
	IPDestination string `json:"ip.dst"`
}

type TCP struct {
	TCPSourcePort      string `json:"tcp.srcport"`
	TCPDestinationPort string `json:"tcp.dstport"`
}

type UDP struct {
	UDPSourcePort      string `json:"udp.srcport"`
	UDPDestinationPort string `json:"udp.dstport"`
	UDPPayload         string `json:"udp.payload"`
}

type DHCP struct {
	DHCPMacAddress string `json:"dhcp.hw.mac_addr"`
	DHCPCookie     string `json:"dhcp.cookie"`
}

func main() {
	// Setup flags
	inputFile := flag.String("file", "ws.json", "Path to the JSON file to parse")
	verbose := flag.Bool("verbose", false, "Enable verbose output") // /flag#Bool
	showTCP := flag.Bool("tcp", false, "Show TCP port information")
	flag.Parse()

	if _, err := os.Stat(*inputFile); os.IsNotExist(err) {
		log.Fatalf("Input file does not exist: %s", *inputFile)
	}

	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Error reading file: %v\n", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Println("Error reading json data", err)
	}

	var p Packets
	err = json.Unmarshal(data, &p)
	if err != nil {
		log.Println("Error reading json data", err)
	}

	// Give user flags to control the output
	for _, v := range p {
		if *verbose {
			fmt.Printf("Packet Details:\n%+v\n", v)
		}
		if *showTCP {
			fmt.Printf("TCP Ports - Source: %s, Destination: %s\n",
				v.Source.Layers.TCP.TCPSourcePort,
				v.Source.Layers.TCP.TCPDestinationPort)
		} else {
			fmt.Printf("Packet from %s to %s\n", v.Source.Layers.IP.IPSource, v.Source.Layers.IP.IPDestination)
		}
	}
}
