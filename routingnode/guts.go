// jeffCoin guts.go

package routingnode

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

// getNodeList - Get the nodeList
func getNodeList() nodeSlice {

	s := "START: getNodeList - Get the blockchain"
	log.Trace("ROUTINGNODE: GUTS   " + s)

	s = "END:   getNodeList - Get the blockchain"
	log.Trace("ROUTINGNODE: GUTS   " + s)

	return nodeList

}

// getNode - Get a Node in the nodeList
func getNode(id string) nodeStruct {

	s := "START: getNode - Get a Node in the nodeList"
	log.Trace("ROUTINGNODE: GUTS   " + s)

	var item nodeStruct

	// Just a special case to get the Last Node
	if id == "last" {
		id = strconv.Itoa(len(nodeList) - 1)
	}

	// SEARCH DATA FOR id
	for _, item := range nodeList {
		if strconv.Itoa(item.Index) == id {
			// RETURN ITEM
			s = "END:   getNode - Get a Node in the nodeList"
			log.Trace("ROUTINGNODE: GUTS   " + s)
			return item
		}
	}

	s = "END:   getNode - FAILED - Did Not get a Node in the nodeList"
	log.Trace("ROUTINGNODE: GUTS   " + s)

	return item

}

// loadThisNode - Load thisNode
func loadThisNode(ip string, tcpPort string) {

	s := "START: loadThisNode - Load thisNode"
	log.Trace("ROUTINGNODE: GUTS   " + s)

	t := time.Now()

	thisNode = nodeStruct{
		Index:     0,
		Status:    "active",
		Timestamp: t.String(),
		IP:        ip,
		Port:      tcpPort,
	}

	s = "END:   loadThisNode - Load thisNode"
	log.Trace("ROUTINGNODE: GUTS   " + s)

}

// appendThisNode - Append thisNode to nodeList
func appendThisNode() nodeStruct {

	s := "START: appendThisNode - Append thisNode to nodeList"
	log.Trace("ROUTINGNODE: GUTS   " + s)

	thisNode.Index = len(nodeList)

	// APPEND
	nodeList = append(nodeList, thisNode)

	s = "END:   appendThisNode - Append thisNode to nodeList"
	log.Trace("ROUTINGNODE: GUTS   " + s)

	return thisNode

}

// getThisNode - Get thisNode
func getThisNode() nodeStruct {

	s := "START: getThisNode - Get thisNode"
	log.Trace("ROUTINGNODE: GUTS   " + s)

	s = "END:   getThisNode - Get thisNode"
	log.Trace("ROUTINGNODE: GUTS   " + s)

	return thisNode

}

// checkIfThisNodeinNodeList - Check if thisNode is already in the nodeList
func checkIfThisNodeinNodeList() bool {

	s := "START: checkIfThisNodeinNodeList - check if thisNode is already in the nodeList"
	log.Trace("ROUTINGNODE: GUTS   " + s)

	// FOR EACH NODE IN NODELIST
	for _, item := range nodeList {

		// DO YOU FIND IT
		if item.IP == thisNode.IP && item.Port == thisNode.Port {

			s = "ThisNode is already in the nodeList"
			log.Warn("ROUTINGNODE: GUTS          " + s)

			s = "END:   checkIfThisNodeinNodeList - check if thisNode is already in the nodeList"
			log.Trace("ROUTINGNODE: GUTS   " + s)

			return true
		}

	}

	s = "END:   checkIfThisNodeinNodeList - check if thisNode is already in the nodeList"
	log.Trace("ROUTINGNODE: GUTS   " + s)

	return false

}

// loadNodeList - Load nodeList
func loadNodeList(message string) {

	s := "START: loadNodeList -  Load nodeList"
	log.Trace("ROUTINGNODE: GUTS   " + s)

	// LOAD
	json.Unmarshal([]byte(message), &nodeList)

	s = "END:   loadNodeList -  Load nodeList"
	log.Trace("ROUTINGNODE: GUTS   " + s)

}

// appendNewNode - Append New Node to the nodeList
func appendNewNode(messageNewNode string) nodeStruct {

	s := "START: appendNewNode - Append New Node to the nodeList"
	log.Trace("ROUTINGNODE: GUTS   " + s)

	newNode := nodeStruct{}
	json.Unmarshal([]byte(messageNewNode), &newNode)

	newNode.Index = len(nodeList)
	nodeList = append(nodeList, newNode)

	s = "END:   appendNewNode - Append New Node to the nodeList"
	log.Trace("ROUTINGNODE: GUTS   " + s)

	return newNode

}
