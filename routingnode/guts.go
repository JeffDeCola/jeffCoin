// jeffCoin guts.go

package routingnode

import (
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

// getNodeList - Get the NodeList
func getNodeList() NodeSlice {

	s := "START: getNodeList - Get the Blockchain"
	log.Println("ROUTINGNODE GUTS:   " + s)

	s = "END: getNodeList - Get the Blockchain"
	log.Println("ROUTINGNODE GUTS:   " + s)

	return NodeList
}

// getNode - Get a Node in the NodeList
func getNode(id string) NodeStruct {

	s := "START: getNode - Get a Node in the NodeList"
	log.Println("ROUTINGNODE GUTS:   " + s)

	var item NodeStruct

	// SEARCH DATA FOR blockID
	for _, item := range NodeList {
		if strconv.Itoa(item.Index) == id {
			// RETURN ITEM
			s = "END: getNode - Get a Node in the NodeList"
			log.Println("ROUTINGNODE GUTS:   " + s)
			return item
		}
	}

	s = "END: getNode - Get a Node in the NodeList"
	log.Println("ROUTINGNODE GUTS:   " + s)

	return item
}

// appendNode - Append Node to the NodeList
func appendNode(nodeIP string, nodeTCPPort string) NodeStruct {

	s := "START: appendNode - Append Node to the NodeList"
	log.Println("ROUTINGNODE GUTS:   " + s)

	t := time.Now()

	newNode := NodeStruct{
		Index:     0,
		Timestamp: t.String(),
		IP:        nodeIP,
		Port:      nodeTCPPort,
	}

	NodeList = append(NodeList, newNode)

	s = "END: appendNode - Append Node to the NodeList"
	log.Println("ROUTINGNODE GUTS:   " + s)

	return newNode

}
