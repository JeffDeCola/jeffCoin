// jeffCoin 3. ROUTINGNODE routingnode-datastructures.go

package routingnode

// NODELIST **************************************************************************************************************

// NodeStruct is your node info
type NodeStruct struct {
	Index          int    `json:"index"`
	Status         string `json:"status"`
	Timestamp      string `json:"timestamp"`
	NodeName       string `json:"nodename"`
	NodeIP         string `json:"nodeip"`
	NodeHTTPPort   string `json:"nodehttpport"`
	NodeTCPPort    string `json:"nodetcpport"`
	NetworkIP      string `json:"networkip"`
	NetworkTCPPort string `json:"networktcpport"`
	TestBlockChain bool   `json:"testblockchain"`
	WalletOnly     bool   `json:"walletonly"`
	ToolVersion    string `json:"toolversion"`
}

// thisNode - The Current Node Info
var thisNode = NodeStruct{}

// nodeSlice is my type
type nodeSlice []NodeStruct

// nodeList is the the Node List
var nodeList = nodeSlice{}
