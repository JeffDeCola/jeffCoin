// jeffCoin 5. WEBSERVER webpages.go

package webserver

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"

	routingnode "github.com/JeffDeCola/jeffCoin/routingnode"
	wallet "github.com/JeffDeCola/jeffCoin/wallet"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

type htmlLoginData struct {
	NodeName     string
	ToolVersion  string
	NodeIP       string
	NodeHTTPPort string
	NodeTCPPort  string
}

type htmlValidateData struct {
	NodeName    string
	DoesItMatch string
}

type htmlLogoutData struct {
	NodeName string
}

type htmlIndexData struct {
	NodeName     string
	PublicKeyHex string
	Balance      string
	ToolVersion  string
	NodeIP       string
	NodeHTTPPort string
	NodeTCPPort  string
	WalletOnly   string
	NodeComment  string
}

type htmlAPIData struct {
	NodeName string
}

type htmlSendData struct {
	NodeName     string
	PublicKeyHex string
	Balance      string
}

type htmlConfirmData struct {
	NodeName                  string
	PublicKeyHex              string
	DestinationAddressComma   string
	DestinationAddressNewLine string
	ValueComma                string
	ValueNewLine              string
}

type htmlTransactionRequestData struct {
	Status string
}

// HTML PAGES *************************************************************************************************************

// loginHandler - GET: /login
func loginHandler(res http.ResponseWriter, req *http.Request) {

	s := "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)
	s = "HTTP SERVER - LOGIN"
	log.Info("WEBSERVER:                   " + s)
	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)

	s = "START  loginHandler() - GET: /login"
	log.Debug("WEBSERVER:            " + s)

	t, err := template.ParseFiles("webserver/login.html")
	checkErr(err)

	// GET THIS NODE
	thisNode := routingnode.GetThisNode()

	htmlTemplateData := htmlLoginData{
		NodeName:     thisNode.NodeName,
		NodeIP:       thisNode.NodeIP,
		NodeHTTPPort: thisNode.NodeHTTPPort,
		NodeTCPPort:  thisNode.NodeTCPPort,
		ToolVersion:  thisNode.ToolVersion,
	}

	// Merge data and execute
	err = t.Execute(res, htmlTemplateData)
	checkErr(err)

	s = "END    loginHandler() - GET: /login"
	log.Debug("WEBSERVER:            " + s)

	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)
	s = "HTTP SERVER - COMPLETE LOGIN"
	log.Info("WEBSERVER:                   " + s)
	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)

}

// validateHandler - GET: /validate
func validateHandler(res http.ResponseWriter, req *http.Request) {

	s := "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)
	s = "HTTP SERVER - DISPLAY VALIDATE"
	log.Info("WEBSERVER:                   " + s)
	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)

	s = "START  validateHandler() - GET: /validate"
	log.Debug("WEBSERVER:            " + s)

	t, err := template.ParseFiles("webserver/validate.html")
	checkErr(err)

	// GET THE PARAMATERS SENT VIA POST FORM
	// Parses the request body
	req.ParseForm()
	passwordEntered := req.Form.Get("Password")

	// GET PASSWORD
	thePassword := GetPassword()

	// COMPARE PASSWORDS
	doesItMatch := "Your password is incorrect, try again"
	if passwordEntered == thePassword.Password {
		doesItMatch = "Your password is valid"

		// CREATE A RANDOM TOKEN - Universally Unique Identifier (UUID) and store in sessionTokenString
		s = "CREATE A RANDOM TOKEN - Universally Unique Identifier (UUID)"
		log.Info("WEBSERVER:                   " + s)
		sessionToken, err := uuid.NewV4()
		checkErr(err)
		sessionTokenString = sessionToken.String()

		// SET COOKIE ON CLIENT CALLED jeffCoin_session_token
		// Good for 1 hour
		s = "SET COOKIE ON CLIENT CALLED jeffCoin_session_token"
		log.Info("WEBSERVER:                   " + s)
		http.SetCookie(res, &http.Cookie{
			Name:    "jeffCoin_session_token",
			Value:   sessionTokenString,
			Expires: time.Now().Add(3600 * time.Second),
		})

		s = "VALIDATED - REDIRECTING to /"
		log.Info("WEBSERVER:                   " + s)
		http.Redirect(res, req, "/", http.StatusFound)
		s = "END    validateHandler() - GET: /validate"
		log.Debug("WEBSERVER:            " + s)
		return

	}

	s = "PASSWORD NOT VALID"
	log.Warn("WEBSERVER:                   " + s)

	// GET THIS NODE
	thisNode := routingnode.GetThisNode()

	htmlTemplateData := htmlValidateData{
		NodeName:    thisNode.NodeName,
		DoesItMatch: doesItMatch,
	}

	// Merge data and execute
	err = t.Execute(res, htmlTemplateData)
	checkErr(err)

	s = "END    validateHandler() - GET: /validate"
	log.Debug("WEBSERVER:            " + s)

	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)
	s = "HTTP SERVER - COMPLETE VALIDATE"
	log.Info("WEBSERVER:                   " + s)
	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)

}

// logoutHandler - GET: /logout
func logoutHandler(res http.ResponseWriter, req *http.Request) {

	s := "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)
	s = "HTTP SERVER - LOGOUT"
	log.Info("WEBSERVER:                   " + s)
	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)

	s = "START  logoutHandler() - GET: /logout"
	log.Debug("WEBSERVER:            " + s)

	t, err := template.ParseFiles("webserver/logout.html")
	checkErr(err)

	// PUT A RANDOM STRING IN HERE
	sessionToken, err := uuid.NewV4()
	sessionTokenString = sessionToken.String()

	// GET THIS NODE
	thisNode := routingnode.GetThisNode()

	htmlTemplateData := htmlLogoutData{
		NodeName: thisNode.NodeName,
	}

	// Merge data and execute
	err = t.Execute(res, htmlTemplateData)
	checkErr(err)

	s = "LOGOUT - REDIRECTING to /login"
	log.Info("WEBSERVER:                   " + s)
	http.Redirect(res, req, "/login", http.StatusFound)
	s = "END    validateHandler() - GET: /login"
	log.Debug("WEBSERVER:            " + s)

	s = "END   logoutHandler() - GET: /logout"
	log.Debug("WEBSERVER:            " + s)

	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)
	s = "HTTP SERVER - COMPLETE LOGOUT"
	log.Info("WEBSERVER:                   " + s)
	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)

	return

}

// indexHandler - GET: /
func indexHandler(res http.ResponseWriter, req *http.Request) {

	s := "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)
	s = "HTTP SERVER - DISPLAY MAIN WEBPAGE"
	log.Info("WEBSERVER:                   " + s)
	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)

	s = "START  indexHandler() - GET: /"
	log.Debug("WEBSERVER:            " + s)

	// CHECK AUTHENTICATION
	if !checkAuthentication(req) {
		s = "FAILED AUTHENTICATION - REDIRECTING to /login"
		log.Warn("WEBSERVER:                   " + s)
		http.Redirect(res, req, "/login", http.StatusFound)
		s = "END    indexHandler() - GET: /"
		log.Debug("WEBSERVER:            " + s)
		return
	}

	t, err := template.ParseFiles("webserver/index.html")
	checkErr(err)

	// GET THIS NODE
	thisNode := routingnode.GetThisNode()

	// CHECK IF WALLET ONLY
	// GET IP & TCPPort from thisNode
	var IP, TCPPort, walletOnly, nodeComment string
	if checkIfWalletOnly() {

		IP = thisNode.NetworkIP
		TCPPort = thisNode.NetworkTCPPort
		walletOnly = "WALLET ONLY"
		nodeComment = ""
		s = "WALLET ONLY - Use network IP:Port to get balance"
		log.Info("WEBSERVER:                   " + s)

	} else {

		IP = thisNode.NodeIP
		TCPPort = thisNode.NodeTCPPort
		walletOnly = ""
		nodeComment = "Node &"

	}

	// GET WALLET
	theWallet := wallet.GetWallet()

	// GET address PublicKeyHex from wallet
	addressPublicKeyHex := theWallet.PublicKeyHex

	// GET ADDRESS BALANCE
	gotAddressBalance, err := wallet.RequestAddressBalance(IP, TCPPort, addressPublicKeyHex)
	gotAddressBalance = strings.Trim(gotAddressBalance, "\"")
	checkErr(err)
	gotAddressBalanceInt, err := strconv.ParseFloat(gotAddressBalance, 64)
	checkErr(err)
	balance := gotAddressBalanceInt / float64(1000)
	gotAddressBalance = strconv.FormatFloat(balance, 'f', 3, 64)

	htmlTemplateData := htmlIndexData{
		NodeName:     thisNode.NodeName,
		PublicKeyHex: addressPublicKeyHex,
		Balance:      gotAddressBalance,
		NodeIP:       thisNode.NodeIP,
		NodeHTTPPort: thisNode.NodeHTTPPort,
		NodeTCPPort:  thisNode.NodeTCPPort,
		ToolVersion:  thisNode.ToolVersion,
		WalletOnly:   walletOnly,
		NodeComment:  nodeComment,
	}

	// Merge data and execute
	err = t.Execute(res, htmlTemplateData)
	checkErr(err)

	s = "END    indexHandler() - GET: /"
	log.Debug("WEBSERVER:            " + s)

	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)
	s = "HTTP SERVER - COMPLETE DISPLAY MAIN WEBPAGE"
	log.Info("WEBSERVER:                   " + s)
	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)

}

// apiHandler - GET: /api
func apiHandler(res http.ResponseWriter, req *http.Request) {

	s := "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)
	s = "HTTP SERVER - DISPLAY API COMMANDS"
	log.Info("WEBSERVER:                   " + s)
	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)

	s = "START  apiHandler() - GET: /api"
	log.Debug("WEBSERVER:            " + s)

	// CHECK AUTHENTICATION
	if !checkAuthentication(req) {
		s = "FAILED AUTHENTICATION - REDIRECTING to /login"
		log.Warn("WEBSERVER:                   " + s)
		http.Redirect(res, req, "/login", http.StatusFound)
		s = "END    indexHandler() - GET: /"
		log.Debug("WEBSERVER:            " + s)
		return
	}

	t, err := template.ParseFiles("webserver/api.html")
	checkErr(err)

	// GET THIS NODE
	thisNode := routingnode.GetThisNode()

	htmlTemplateData := htmlAPIData{
		NodeName: thisNode.NodeName,
	}

	// Merge data and execute
	err = t.Execute(res, htmlTemplateData)
	checkErr(err)

	s = "END    apiHandler() - GET: /api"
	log.Debug("WEBSERVER:            " + s)

	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)
	s = "HTTP SERVER - COMPLETE DISPLAY API COMMANDS"
	log.Info("WEBSERVER:                   " + s)
	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)

}

// sendHandler - GET: /send
func sendHandler(res http.ResponseWriter, req *http.Request) {

	s := "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)
	s = "HTTP SERVER - DISPLAY API COMMANDS"
	log.Info("WEBSERVER:                   " + s)
	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)

	s = "START  sendHandler() - GET: /send"
	log.Debug("WEBSERVER:            " + s)

	// CHECK AUTHENTICATION
	if !checkAuthentication(req) {
		s = "FAILED AUTHENTICATION - REDIRECTING to /login"
		log.Warn("WEBSERVER:                   " + s)
		http.Redirect(res, req, "/login", http.StatusFound)
		s = "END    indexHandler() - GET: /"
		log.Debug("WEBSERVER:            " + s)
		return
	}

	t, err := template.ParseFiles("webserver/send.html")
	checkErr(err)

	// GET THIS NODE
	thisNode := routingnode.GetThisNode()

	// CHECK IF WALLET ONLY
	// GET IP & TCPPort from thisNode
	var IP, TCPPort string
	if checkIfWalletOnly() {

		IP = thisNode.NetworkIP
		TCPPort = thisNode.NetworkTCPPort
		s = "WALLET ONLY - Use network IP:Port to get balance"
		log.Info("WEBSERVER:                   " + s)

	} else {

		IP = thisNode.NodeIP
		TCPPort = thisNode.NodeTCPPort

	}

	// GET WALLET
	theWallet := wallet.GetWallet()

	// GET address PublicKeyHex from wallet
	addressPublicKeyHex := theWallet.PublicKeyHex

	// GET ADDRESS BALANCE
	gotAddressBalance, err := wallet.RequestAddressBalance(IP, TCPPort, addressPublicKeyHex)
	gotAddressBalance = strings.Trim(gotAddressBalance, "\"")
	checkErr(err)
	gotAddressBalanceInt, err := strconv.ParseFloat(gotAddressBalance, 64)
	checkErr(err)
	balance := gotAddressBalanceInt / float64(1000)
	gotAddressBalance = strconv.FormatFloat(balance, 'f', 3, 64)

	htmlTemplateData := htmlSendData{
		NodeName:     thisNode.NodeName,
		PublicKeyHex: addressPublicKeyHex,
		Balance:      gotAddressBalance,
	}

	// Merge data and execute
	err = t.Execute(res, htmlTemplateData)
	checkErr(err)

	s = "END    sendHandler() - GET: /send"
	log.Debug("WEBSERVER:            " + s)

	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)
	s = "HTTP SERVER - COMPLETE DISPLAY API COMMANDS"
	log.Info("WEBSERVER:                   " + s)
	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)

}

// confirmHandler - GET: /confirm
func confirmHandler(res http.ResponseWriter, req *http.Request) {

	s := "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)
	s = "HTTP SERVER - DISPLAY CONFIRM SEND"
	log.Info("WEBSERVER:                   " + s)
	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)

	s = "START  confirmHandler() - GET: /confirm"
	log.Debug("WEBSERVER:            " + s)

	// CHECK AUTHENTICATION
	if !checkAuthentication(req) {
		s = "FAILED AUTHENTICATION - REDIRECTING to /login"
		log.Warn("WEBSERVER:                   " + s)
		http.Redirect(res, req, "/login", http.StatusFound)
		s = "END    indexHandler() - GET: /"
		log.Debug("WEBSERVER:            " + s)
		return
	}

	t, err := template.ParseFiles("webserver/confirm.html")
	checkErr(err)

	// GET THE PARAMATERS SENT VIA POST FORM
	// Parses the request body
	// It may or may not have a comma
	req.ParseForm()
	destinationAddressComma := req.Form.Get("DestinationAddress")
	valueComma := req.Form.Get("Value")

	// ADD NEWLINE TO COMMAS (Just added whitespace but would like to figure this out)
	destinationAddressNewline := strings.Replace(destinationAddressComma, ",", ", ", -1)
	valueNewLine := strings.Replace(valueComma, ",", ", ", -1)
	// GET THIS NODE
	thisNode := routingnode.GetThisNode()

	// GET WALLET
	theWallet := wallet.GetWallet()

	htmlTemplateData := htmlConfirmData{
		NodeName:                  thisNode.NodeName,
		PublicKeyHex:              theWallet.PublicKeyHex,
		DestinationAddressComma:   destinationAddressComma,
		DestinationAddressNewLine: destinationAddressNewline,
		ValueComma:                valueComma,
		ValueNewLine:              valueNewLine,
	}

	// Merge data and execute
	err = t.Execute(res, htmlTemplateData)
	checkErr(err)

	s = "END    confirmHandler() - GET: /confirm"
	log.Debug("WEBSERVER:            " + s)

	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)
	s = "HTTP SERVER - COMPLETE CONFIRM SEND"
	log.Info("WEBSERVER:                   " + s)
	s = "----------------------------------------------------------------"
	log.Info("WEBSERVER:                   " + s)

}
