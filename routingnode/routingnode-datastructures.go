// jeffCoin 3. ROUTINGNODE routingnode-datastructures.go

package routingnode

// NODELIST **************************************************************************************************************

// NodeStruct is your node info
type nodeStruct struct {
	Index     int    `json:"index"`
	Status    string `json:"status"`
    Timestamp string `json:"timestamp"`
    NodeName  string `json:"nodename"`
	IP        string `json:"ip"`
	Port      string `json:"port"`
}

// thisNode - The Current Node Info
var thisNode = nodeStruct{}

// nodeSlice is my type
type nodeSlice []nodeStruct

// nodeList is the the Node List
var nodeList = nodeSlice{}
