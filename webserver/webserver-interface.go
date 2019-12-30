// jeffCoin 5. WEBSERVER webserver-interface.go

package webserver

import (
	"bufio"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"time"

	errors "github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

// PASWORD ***************************************************************************************************************

// GetPassword - Gets the Password
func GetPassword() passwordStruct {

	s := "START  GetPassword() - Gets the Password"
	log.Debug("WEBSERVER:   I/F     " + s)

	thePassword := getPassword()

	s = "END    GetPassword() - Gets the Password"
	log.Debug("WEBSERVER:   I/F     " + s)

	return thePassword

}

// WritePassword - Writes the password to file (AES-256 encryption)
func WritePassword(nodeName string, password string) string {

	s := "START  WritePassword() - Writes the password to file (AES-256 encryption)"
	log.Debug("WEBSERVER:   I/F     " + s)

	thePassword := writePassword(nodeName, password)
    showPassword := thePassword
    showPassword.Password = showPassword.Password[0:2} + "..."

	s = "Congrats, you created your Password (-loglevel trace to display)"
    log.Info("WEBSERVER:   I/F             " + s)
	js, _ := json.MarshalIndent(showPassword, "", "    ")
	log.Trace("\n\n" + string(js) + "\n\n")

	s = "END    WritePassword() - Writes the password to file (AES-256 encryption)"
	log.Debug("WEBSERVER:   I/F     " + s)

    return thePassword.Password
    
}

// ReadPasswordFile - Reads the password from a file (AES-256 decrypt)
func ReadPasswordFile(nodeName string) string {

	s := "START  ReadPasswordFile() - Reads the password from a file (AES-256 decrypt)"
	log.Debug("WEBSERVER:   I/F     " + s)

    thePassword := readPasswordFile(nodeName)
    showPassword := thePassword
    showPassword.Password = showPassword.Password[0:2} + "..."

	s = "Congrats, you loaded your old Password from a file (-loglevel trace to display)"
	log.Info("WEBSERVER:   I/F     " + s)
	js, _ := json.MarshalIndent(showPassword, "", "    ")
	log.Trace("\n\n" + string(js) + "\n\n")

	s = "END    ReadPasswordFile() - Reads the password from a file (AES-256 decrypt)"
	log.Debug("WEBSERVER:   I/F     " + s)

    return thePassword.Password
    
}
