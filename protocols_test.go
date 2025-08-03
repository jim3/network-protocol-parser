package main

import "testing"

func TestPackets_GetSourcePort(t *testing.T) {
	// Arrange
	tcp := &TCP{
		TCPSourcePort: "443",
	}

	// Act
	sourcePort := tcp.GetSourcePort()

	// Assert
	if sourcePort != "443" {
		t.Errorf("GetSourcePort() returned %s, expected %s", sourcePort, "443")
	}
}
