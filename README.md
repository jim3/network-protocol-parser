# Network Protocol Parser

## Description
Parses JSON-exported Wireshark packets.
This tool is designed to parse network protocol data exported from Wireshark in JSON format. It extracts relevant fields and provides a structured output for further analysis.

## Installation
```bash
git clone https://github.com/jim3/network-protocol-parser.git
cd network-protocol-parser
go mod tidy
```

## Usage
```bash
# Example of how to use the tool
go run main.go
```

## Features
- Parses Wireshark JSON exports
- Extracts frame information

## Requirements
- Wireshark-exported JSON file
