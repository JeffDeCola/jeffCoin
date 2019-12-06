// jeffCoin 5. WEBSERVER routes.go

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
		"ShowBlockchain",
		"GET",
		"/showblockchain",
		showBlockchainHandler,
	},
	Route{
		"ShowBlock",
		"GET",
		"/showblock/{blockID}",
		showBlockHandler,
	},
	Route{
		"ShowLockedBlock",
		"GET",
		"/showlockedblock/",
		showLockedBlockHandler,
	},
	Route{
		"ShowCurrentBlock",
		"GET",
		"/showcurrentblock",
		showCurrentBlockHandler,
	},
	Route{
		"ShowNodeList",
		"GET",
		"/shownodelist",
		showNodeListHandler,
	},
	Route{
		"ShowNode",
		"GET",
		"/shownode/{nodeID}",
		showNodeHandler,
	},
	Route{
		"ShowThisNode",
		"GET",
		"/showthisnode",
		showThisNodeHandler,
	},
	Route{
		"ShowWallet",
		"GET",
		"/showwallet",
		showWalletHandler,
	},
	Route{
		"ShowJeffCoinAddress",
		"GET",
		"/showjeffcoinaddress",
		showJeffCoinAddressHandler,
	},
	Route{
		"ShowBalance",
		"GET",
		"/showbalance",
		showBalanceHandler,
	},
	Route{
		"ShowAddressBalance",
		"GET",
		"/showaddressbalance/{jeffCoinAddress}",
		showAddressBalanceHandler,
	},
	Route{
		"TransactionRequest",
		"GET",
		"/transactionrequest/{destinationaddress}/{value}",
		transactionRequestHandler,
	},
}
