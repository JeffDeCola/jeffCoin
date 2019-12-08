// jeffCoin 5. WEBSERVER http-server.go

package webserver

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// HTTP SERVER ***********************************************************************************************************

// StartHTTPServer - Start the WebServer
func StartHTTPServer(nodeIP string, nodeWebPort string) {

	s := "START  StartHTTPServer() - Start the WebServer"
	log.Trace("WEBSERVER:   SERVER   " + s)

	s = "HTTP Server listening on " + nodeIP + ":" + nodeWebPort
	log.Info("WEBSERVER:   SERVER          " + s)

	// CREATE ROUTER
	myRouter := JeffsRouter()

	// LISTEN ON IP AND PORT
	log.Fatal(http.ListenAndServe(nodeIP+":"+nodeWebPort, myRouter))

}
