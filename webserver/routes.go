// jeffCoin routes.go

package webserver

import "net/http"

// Route - The struct for the route endpoints (e.g. /jeff)
type Route struct {
	RouteName        string
	RouteHTTPVerb    string
	RouteEndPoint    string
	RouteHandlerFunc http.HandlerFunc
}

// Routes is slice
type Routes []Route

var routes = Routes{
	Route{
		"GetIndex",
		"GET",
		"/",
		indexHandler,
	},
	Route{
		"ShowBlock",
		"GET",
		"/showblock/{blockID}",
		showBlockHandler,
	},
	Route{
		"ShowChain",
		"GET",
		"/showchain",
		showChainHandler,
	},
}
