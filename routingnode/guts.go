// jeffCoin 3. ROUTINGNODE guts.go

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

// NODELIST **************************************************************************************************************

// getNodeList - Gets the nodeList
func getNodeList() nodeSlice {

	s := "START  getNodeList() - Gets the nodeList"
	log.Debug("ROUTINGNODE: GUTS     " + s)

	s = "END    getNodeList() - Gets the nodeList"
	log.Debug("ROUTINGNODE: GUTS     " + s)

	return nodeList

}

// loadNodeList - Loads the entire nodeList
func loadNodeList(message string) {

	s := "START  loadNodeList() -  Loads the entire nodeList"
	log.Debug("ROUTINGNODE: GUTS     " + s)

	// LOAD
	json.Unmarshal([]byte(message), &nodeList)

	s = "END    loadNodeList() -  Loads the entire nodeList"
	log.Debug("ROUTINGNODE: GUTS     " + s)

}

// NODE ******************************************************************************************************************

// getNode - Gets a Node in the nodeList
func getNode(id string) NodeStruct {

	s := "START  getNode() - Gets a Node in the nodeList"
	log.Debug("ROUTINGNODE: GUTS     " + s)

	var item NodeStruct

	// Just a special case to get the Last Node
	if id == "last" {
		id = strconv.Itoa(len(nodeList) - 1)
	}

	// SEARCH DATA FOR id
	for _, item := range nodeList {
		if strconv.Itoa(item.Index) == id {
			// RETURN ITEM
			s = "END    getNode() - Gets a Node in the nodeList"
			log.Debug("ROUTINGNODE: GUTS     " + s)
			return item
		}
	}

	s = "END    getNode() - FAILED: Did Not get a Node in the nodeList"
	log.Debug("ROUTINGNODE: GUTS     " + s)

	return item

}

// appendNewNode - Appends a new Node to the nodeList
func appendNewNode(messageNewNode string) NodeStruct {

	s := "START  appendNewNode() - Appends a new Node to the nodeList"
	log.Debug("ROUTINGNODE: GUTS     " + s)

	newNode := NodeStruct{}
	json.Unmarshal([]byte(messageNewNode), &newNode)

	newNode.Index = len(nodeList)
	nodeList = append(nodeList, newNode)

	s = "END    appendNewNode() - Appends a new Node to the nodeList"
	log.Debug("ROUTINGNODE: GUTS     " + s)

	return newNode

}

// THIS NODE *************************************************************************************************************

// getThisNode - Gets thisNode
func getThisNode() NodeStruct {

	s := "START  getThisNode() - Gets thisNode"
	log.Debug("ROUTINGNODE: GUTS     " + s)

	s = "END    getThisNode() - Gets thisNode"
	log.Debug("ROUTINGNODE: GUTS     " + s)

	return thisNode

}

// loadThisNode - Loads thisNode
func (n NodeStruct) loadThisNode() {

	s := "START  loadThisNode() - Loads thisNode"
	log.Debug("ROUTINGNODE: GUTS     " + s)

	t := time.Now()

	thisNode = n
	thisNode.Index = 0
	thisNode.Status = "active"
	thisNode.Timestamp = t.String()

	s = "END    loadThisNode() - Loads thisNode"
	log.Debug("ROUTINGNODE: GUTS     " + s)

}

// appendThisNode - Appends thisNode to the nodeList
func appendThisNode() NodeStruct {

	s := "START  appendThisNode() - Appends thisNode to the nodeList"
	log.Debug("ROUTINGNODE: GUTS     " + s)

	thisNode.Index = len(nodeList)

	// APPEND
	nodeList = append(nodeList, thisNode)

	s = "END    appendThisNode() - Appends thisNode to the nodeList"
	log.Debug("ROUTINGNODE: GUTS     " + s)

	return thisNode

}

// checkIfThisNodeinNodeList - Check if thisNode is already in the nodeList
func checkIfThisNodeinNodeList() bool {

	s := "START  checkIfThisNodeinNodeList() - Check if thisNode is already in the nodeList"
	log.Debug("ROUTINGNODE: GUTS     " + s)

	// FOR EACH NODE IN NODELIST
	for _, item := range nodeList {

		// DO YOU FIND IT
		if item.NodeIP == thisNode.NodeIP && item.NodeTCPPort == thisNode.NodeTCPPort {

			s = "thisNode is already in the nodeList"
			log.Warn("ROUTINGNODE: GUTS            " + s)

			s = "END    checkIfThisNodeinNodeList() - Check if thisNode is already in the nodeList"
			log.Debug("ROUTINGNODE: GUTS     " + s)

			return true
		}

	}

	s = "thisNode is NOT in the nodeList"
	log.Info("ROUTINGNODE: GUTS            " + s)

	s = "END    checkIfThisNodeinNodeList() - Check if thisNode is already in the nodeList"
	log.Debug("ROUTINGNODE: GUTS     " + s)

	return false

}
