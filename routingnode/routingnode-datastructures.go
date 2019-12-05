// jeffCoin 3. ROUTINGNODE routingnode-datastructures.go

package routingnode

// NODELIST **************************************************************************************************************

// NodeStruct is your node info
type nodeStruct struct {
	Index       int    `json:"index"`
	Status      string `json:"status"`
	Timestamp   string `json:"timestamp"`
	NodeName    string `json:"nodename"`
	ToolVersion string `json:"toolversion"`
	IP          string `json:"ip"`
	HTTPPort    string `json:"httpport"`
	TCPPort     string `json:"tcpport"`
}

// thisNode - The Current Node Info
var thisNode = nodeStruct{}

// nodeSlice is my type
type nodeSlice []nodeStruct

// nodeList is the the Node List
var nodeList = nodeSlice{}
