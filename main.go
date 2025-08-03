package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Setup flags
	inputFile := flag.String("file", "ws.json", "Path to the JSON file to parse")
	verbose := flag.Bool("verbose", false, "Enable verbose output") // /flag#Bool
	tcpPorts := flag.Bool("tcp", false, "Show TCP port information")
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

	// Declares a instance of Packets `type Packets []struct {...}`
	var p Packets
	err = json.Unmarshal(data, &p)
	if err != nil {
		log.Println("Error reading json data", err)
	}

	if *tcpPorts {
		fmt.Printf("TCP PACKET COUNT —→ %d\n", p.packetCount())
		for _, v := range p {
			if *verbose {
				fmt.Printf("Packet Details:\n%+v\n", v)
			}
			if *tcpPorts {
				fmt.Printf("TCP Ports - Source: %s, Destination: %s\n",
					v.Source.Layers.TCP.TCPSourcePort,
					v.Source.Layers.TCP.TCPDestinationPort)
			}
		}
	}
}
