package model

// WSS is the WebSocket Secure protocol.
type WSS string

// String returns the string representation of the WSS.
func (w WSS) String() string {
	return string(w)
}

// Relay is the data model of the relay.
type Relay struct {
	WSS    WSS  `json:"wss"`    // WSS is the WebSocket Secure protocol for the relay server.
	Read   bool `json:"read"`   // Read is true if the relay server can read.
	Write  bool `json:"write"`  // Write is true if the relay server can write.
	Search bool `json:"search"` // Search is true if the relay server can search.
}
