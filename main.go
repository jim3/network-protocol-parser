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
	// Command-line flags
	inputFile := flag.String("file", "ws.json", "Path to the JSON file to parse")
	verbose := flag.Bool("verbose", false, "Enable verbose output") // /flag#Bool
	tcpPorts := flag.Bool("tcp", false, "Show TCP port information")
	ipAddr := flag.Bool("ip", false, "Show IP src and dst addresses")
	ipLookUp := flag.Bool("iplookup", false, "Lookup IP address info")

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

	// Unmarshal JSON data into Packets struct
	var p Packets
	err = json.Unmarshal(data, &p)
	if err != nil {
		log.Println("Error reading json data", err)
	}
	fmt.Printf("TOTAL TCP PACKET COUNT IS: %d\n", p.packetCount())

	// Print packet details based on flags
	for _, v := range p {
		if *verbose {
			fmt.Printf("Packet Details:\n%+v\n", v)
		}
		if *tcpPorts {
			src, dst := v.Source.Layers.TCP.GetSourcePort()
			fmt.Println(src, dst)
		}
		if *ipAddr {
			s, d := v.Source.Layers.IP.GetIpAddress()
			fmt.Println(s, d)
		}
		if *ipLookUp {
			// TODO
		}
	}

}
