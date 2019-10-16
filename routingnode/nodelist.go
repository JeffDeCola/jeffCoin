// jeffCoin nodelist.go

package routingnode

// NodeStruct is your node info
type nodeStruct struct {
	Index     int    `json:"index"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	IP        string `json:"ip"`
	Port      string `json:"port"`
}

// thisNode - The Current Node Info
var thisNode = nodeStruct{}

// NodeSlice is my type
type nodeSlice []nodeStruct

// NodeList is the the Node List
var nodeList = nodeSlice{}
