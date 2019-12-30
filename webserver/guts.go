// jeffCoin 5. WEBSERVER guts.go

package webserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

// PASWORD ***************************************************************************************************************

// getPassword - Gets the password
func getPassword() passwordStruct {

	s := "START  getPassword() - Gets the password"
	log.Debug("PASSWORD:      GUTS     " + s)

	s = "END    getPassword() - Gets the password"
	log.Debug("PASSWORD:      GUTS     " + s)

	return password

}

// writePassword - Writes the password to file (AES-256 encryption)
func writePassword(nodeName string, password string) passwordStruct {

	s := "START  writePassword() -  Writes the password to file (AES-256 encryption)"
	log.Debug("PASSWORD:      GUTS     " + s)

	// LOAD password IN PASSWORD
	password = passwordStruct{password}

	// WRITE PASSWORD STRUCT TO JSON FILE
	filedata, _ := json.MarshalIndent(password, "", " ")
	filename := "password/" + nodeName + "-password.json"
	_ = ioutil.WriteFile(filename, filedata, 0644)
	s = "Wrote password to " + filename
	log.Info("PASSWORD:      GUTS            " + s)

	s = "END    writePassword() -  Writes the password to file (AES-256 encryption)"
	log.Debug("PASSWORD:      GUTS     " + s)

	return password

}

// readPasswordFile - Reads the password from a file
func readPasswordFile(nodeName string) passwordStruct {

	s := "START  readPasswordFile() - Reads the password from a file (AES-256 decrypt)"
	log.Debug("PASSWORD:      GUTS     " + s)

	// READ PASSWORD STRUCT TO JSON FILE
	filename := "password/" + nodeName + "-password.json"
	filedata, _ := ioutil.ReadFile(filename)
	_ = json.Unmarshal([]byte(filedata), &password)
	s = "Read password from " + filename
	log.Info("PASSWORD:      GUTS            " + s)

	s = "END    readPasswordFile() - Reads the password from a file (AES-256 decrypt)"
	log.Debug("PASSWORD:      GUTS     " + s)

	return password

}
