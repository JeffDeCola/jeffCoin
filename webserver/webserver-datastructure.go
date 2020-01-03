// jeffCoin 5. WEBSERVER webserver-datastructures.go

package webserver

// PASSWORDS **************************************************************************************************************

// passwordStruct is your password hash
type passwordStruct struct {
	PasswordHash string `json:"passwordHash"`
}

// password - The Password
var password = passwordStruct{}

// SESSION TOKENS *********************************************************************************************************

// The users session token
var sessionTokenString string
