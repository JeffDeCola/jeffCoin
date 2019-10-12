// jeffCoin nodes.go

package routingnode

// NodesStruct is your node info
type NodesStruct struct {
	IPs   []string `json:"ips"`
	Ports []string `json:"ports"`
}

// NodesSlice is my type
type NodesSlice []NodesStruct

// Nodes is the the Nodes
var Nodes = NodesSlice{}
