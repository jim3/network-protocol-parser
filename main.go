package main

import (
	"encoding/json"
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
	Layer  string `json:"layers"`
	Frames Frames `json:"frame"`
}

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

// ----------------------------------------------

func main() {
	file, err := os.Open("ws.json") // Read the JSON file
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
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

	for _, v := range p {
		fmt.Println(v) // {packets-2025-07-28 doc <nil> {{ { {\Device\NPF_{D24C971D-E38B-41A2-832D-D09F94CEF56C} Wi-Fi}}}}}
	}
}
