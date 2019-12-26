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
		"GetAPI",
		"GET",
		"/api",
		apiHandler,
    },
    Route{
		"GetSend",
		"GET",
		"/send",
		sendHandler,
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
		"ShowPendingBlock",
		"GET",
		"/showpendingblock",
		showPendingBlockHandler,
	},
	Route{
		"ShowLockedBlock",
		"GET",
		"/showlockedblock/",
		showLockedBlockHandler,
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
		"TransactionRequest",
		"GET",
		"/transactionrequest/{destinationaddress}/{value}",
		transactionRequestHandler,
	},
	Route{
		"ShowAddressBalance",
		"GET",
		"/showaddressbalance/{jeffCoinAddress}",
		showAddressBalanceHandler,
	},
}
