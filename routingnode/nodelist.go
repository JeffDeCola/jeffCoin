// jeffCoin nodelist.go

package routingnode

// NodesStruct is your node info
type NodeStruct struct {
	Index     int      `json:"index"`
	Timestamp string   `json:"timestamp"`
	IP        []string `json:"ip"`
	Port      []string `json:"port"`
}

// NodesSlice is my type
type NodeSlice []NodesStruct

// Nodes is the the Node List
var NodeList = NodesSlice{}
