// jeffCoin 5. WEBSERVER webserver-datastructures.go

package webserver

// ***********************************************************************************************************************

// passwordStruct is your password
type passwordStruct struct {
	Password   string `json:"password"`
}

// password - The Password
var password = passwordStruct{}
