// jeffCoin routingnode-interface.go

package routingnode

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"time"

	log "github.com/sirupsen/logrus"
)

// GenesisNodeList - Creates the NodeList (Only run once)
func GenesisNodeList(nodeIP string, nodeTCPPort string) {

	s := "START: GenesisNodeList - Creates the NodeList (Only run once)"
	log.Println("ROUTINGNODE I/F:    " + s)

	newNode := appendNode(nodeIP, nodeTCPPort)

	fmt.Printf("\nCongrats, your first Node in your Network is:\n\n")
	js, _ := json.MarshalIndent(newNode, "", "    ")
	fmt.Printf("%v\n\n", string(js))

	s = "END: GenesisNodeList - Creates the NodeList (Only run once)"
	log.Println("ROUTINGNODE I/F:    " + s)

}

// GetNodeList - Gets the NodeList
func GetNodeList() NodeSlice {

	s := "START: GetNodeList - Gets the NodeList"
	log.Println("ROUTINGNODE I/F:    " + s)

	theNodeList := getNodeList()

	s = "END: GetNodeList - Gets the NodeList"
	log.Println("ROUTINGNODE I/F:    " + s)

	return theNodeList

}

// GetNode - Get a Node (via Index number) from the NodeList
func GetNode(id string) NodeStruct {

	s := "START: GetNode - Get a Node (via Index number) from the NodeList"
	log.Println("ROUTINGNODE I/F:    " + s)

	theNode := getNode(id)

	// RETURN NOT FOUND
	s = "END: GetNode - Get a Node (via Index number) from the NodeList"
	log.Println("ROUTINGNODE I/F:    " + s)

	return theNode

}

// LoadNodeList - Loads the NodeList from the Network Node
func LoadNodeList(networkIP string, networkTCPPort string) error {

	s := "START: LoadNewNode - Loads the NodeList from the Network Node"
	log.Println("ROUTINGNODE I/F:    " + s)

	// SETUP THE CONNECTION
	conn, err := net.Dial("tcp", networkIP+":"+networkTCPPort)
	checkErr(err)

	// GET THE RESPONSE MESSAGE
	message, _ := bufio.NewReader(conn).ReadString('\n')
	s = "Message from Network Node: " + message
	log.Println("ROUTINGNODE I/F:    " + s)
	if message == "ERROR" {
		s = "ERROR: Could not get blockchain from node"
		log.Println("ROUTINGNODE I/F:    " + s)
		return errors.New(s)
	}

	// SEND THE REQUEST
	fmt.Fprintf(conn, "SENDNODELIST\n")

	// GET THE NODELIST
	message, _ = bufio.NewReader(conn).ReadString('\n')
	s = "Message from Network Node: " + message
	log.Println("ROUTINGNODE I/F:    " + s)
	if message == "ERROR" {
		s = "ERROR: Could not get blockchain from node"
		log.Println("ROUTINGNODE I/F:    " + s)
		return errors.New(s)
	}

	fmt.Println(message)
	// LOAD UP THE NODELIST FROM THE STRING STARTING AT index 1
	// CHANGE SEND TO GUTS????????????????????????????????????
	NewNodeSlice := NodeSlice{}
	json.Unmarshal([]byte(message), &NewNodeSlice)
	// This was tricky but I figured it out
	NodeList = append(NodeList, NewNodeSlice...)

	// CLOSE CONNECTION
	fmt.Fprintf(conn, "EOF\n")
	time.Sleep(2 * time.Second)
	conn.Close()

	s = "END: LoadNodeList - Loads the NodeList from the Network Node"
	log.Println("ROUTINGNODE I/F:    " + s)

	return nil

}

// BroadcastNewNode	- Broadcasts this Node to the Network to add to their NodeLists
func BroadcastNewNode(nodeIP string, nodeTCPPort string, networkIP string, networkTCPPort string) error {

	s := "START: BroadcastNewNode	- Broadcase this Nod to the Network to add to their NodeLists"
	log.Println("ROUTINGNODE I/F:    " + s)

	// SETUP THE CONNECTION
	conn, err := net.Dial("tcp", networkIP+":"+networkTCPPort)
	checkErr(err)

	// GET THE RESPONSE MESSAGE
	message, _ := bufio.NewReader(conn).ReadString('\n')
	s = "Message from Network Node: " + message
	log.Println("ROUTINGNODE I/F:    " + s)
	if message == "ERROR" {
		s = "ERROR: Could not get blockchain from node"
		log.Println("ROUTINGNODE I/F:    " + s)
		return errors.New(s)
	}

	// SEND THE REQUEST
	fmt.Fprintf(conn, "ADDNEWNODE\n")

	// GET THE ????????????
	message, _ = bufio.NewReader(conn).ReadString('\n')
	s = "Message from Network Node: " + message
	log.Println("ROUTINGNODE I/F:    " + s)
	if message == "ERROR" {
		s = "ERROR: Could not get blockchain from node"
		log.Println("ROUTINGNODE I/F:    " + s)
		return errors.New(s)
	}

	fmt.Println(message)
	// LOAD UP THE BLOCKCHAIN FROM THE STRING
	// CHANGE SEND TO GUTS????????????????????????????????????

	// CLOSE CONNECTION
	fmt.Fprintf(conn, "EOF\n")
	time.Sleep(2 * time.Second)
	conn.Close()

	s = "END: BroadcastNewNode	- Broadcase this Nod to the Network to add to their NodeLists"
	log.Println("ROUTINGNODE I/F:    " + s)

	return nil

}
