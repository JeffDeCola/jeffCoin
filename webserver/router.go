// jeffCoin 5. WEBSERVER router.go

package webserver

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

// Message takes incoming JSON payload for writing data
type Message struct {
	Data string `json:"data"`
}

// JeffsRouter is the router
func JeffsRouter() *mux.Router {

	// MAKE ROUTER
	router := mux.NewRouter().StrictSlash(true)

	// LOAD ROUTES ONE BY ONE
	for _, route := range routes {

		var handler http.Handler
		handler = route.RouteHandlerFunc

		// WRAP IN LOGGER
		handler = Logger(handler, route.RouteName)

		router.
			Name(route.RouteName).
			Methods(route.RouteHTTPVerb).
			Path(route.RouteEndPoint).
			Handler(handler)
	}

	// ADD THE CSS DIRECTORY - TOOK ME FOREVER TO FIGURE THIS OUT
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("webserver/css"))))
	return router
}
