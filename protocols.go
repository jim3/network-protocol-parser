package main

type Packets []struct {
	Index  string `json:"_index"`
	Type   string `json:"_type"`
	Score  any    `json:"_score"`
	Source Source `json:"_source"`
}

// Counts TCP packets
func (p Packets) packetCount() int {
	count := 0
	for _, v := range p {
		if v.Source.Layers.TCP.TCPSourcePort != "" {
			count++
		}
	}
	return count
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
