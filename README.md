# Network Protocol Parser

## Description
Parses JSON-exported Wireshark packets.
This tool is designed to parse network protocol data exported from Wireshark in JSON format. It extracts relevant fields and provides a structured output for further analysis.

## Installation & Usage
```bash
git clone https://github.com/jim3/network-protocol-parser.git
cd network-protocol-parser
go run main.go
```

## Requirements
- Wireshark-exported JSON file

## Features
- Parses Wireshark JSON exports
- Extracts frame information
- This is just the starting code so...

## TODO
- Focus on *less* features and more robustness
- Include more protocols (e.g., IP, TCP, DHCP, etc.)
- Improve error handling
- Add methods for specific protocol analysis
- Implement unit tests
