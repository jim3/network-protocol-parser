# Network Protocol Parser

## Description
CLI app that parses JSON-exported Wireshark packets. Use the flags to control the output format and verbosity.

## Installation & Example Usage & Output
```bash
git clone https://github.com/jim3/network-protocol-parser.git
cd network-protocol-parser
go run . -file="ws.json" -tcp
```

## An Example Output
```bash
TCP Ports - Source: 43056, Destination: 443
TCP Ports - Source: 443, Destination: 43056
TCP Ports - Source: 12087, Destination: 443
TCP Ports - Source: 12088, Destination: 443
```

## TODO
In order of importance
- [ ] Create methods for each protocol
- [ ] Create tests for each method