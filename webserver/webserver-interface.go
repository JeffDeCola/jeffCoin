// jeffCoin 5. WEBSERVER webserver-interface.go

package webserver

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

// PASWORD ***************************************************************************************************************

// GetPassword - Gets the Password
func GetPassword() passwordStruct {

	s := "START  GetPassword() - Gets the Password"
	log.Debug("WEBSERVER:   I/F      " + s)

	thePassword := getPassword()

	s = "END    GetPassword() - Gets the Password"
	log.Debug("WEBSERVER:   I/F      " + s)

	return thePassword

}

// WritePasswordFile - Writes the password hash to a file and puts in struct
func WritePasswordFile(nodeName string, passwordString string) {

	s := "START  WritePasswordFile() - Writes the password hash to a file and puts in struct"
	log.Debug("WEBSERVER:   I/F      " + s)

	writePasswordFile(nodeName, passwordString)

	s = "Congrats, you created your password hash (-loglevel trace to display)"
	log.Info("WEBSERVER:   I/F             " + s)
	js, _ := json.MarshalIndent(password, "", "    ")
	log.Trace("\n\n" + string(js) + "\n\n")

	s = "END    WritePasswordFile() - Writes the password hash to a file and puts in struct"
	log.Debug("WEBSERVER:   I/F      " + s)

}

// ReadPasswordFile - Reads the password hash from a file and puts in struct
func ReadPasswordFile(nodeName string) {

	s := "START  ReadPasswordFile() - Reads the password hash from a file and puts in struct"
	log.Debug("WEBSERVER:   I/F      " + s)

	readPasswordFile(nodeName)

	s = "Congrats, you loaded your password hash from a file (-loglevel trace to display)"
	log.Info("WEBSERVER:   I/F             " + s)
	js, _ := json.MarshalIndent(password, "", "    ")
	log.Trace("\n\n" + string(js) + "\n\n")

	s = "END    ReadPasswordFile() - Reads the password hash from a file and puts in struct"
	log.Debug("WEBSERVER:   I/F      " + s)

}
