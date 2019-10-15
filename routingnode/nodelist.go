// jeffCoin nodelist.go

package routingnode

// NodeStruct is your node info
type NodeStruct struct {
    Index      int      `json:"index"`
	Timestamp string `json:"timestamp"`
	IP        string `json:"ip"`
	Port      string `json:"port"`
}

// NodeSlice is my type
type NodeSlice []NodeStruct

// NodeList is the the Node List
var NodeList = NodeSlice{}
