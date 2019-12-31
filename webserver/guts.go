// jeffCoin 5. WEBSERVER guts.go

package webserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	wallet "github.com/JeffDeCola/jeffCoin/wallet"
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
	log.Debug("WEBSERVER:   GUTS     " + s)

	s = "END    getPassword() - Gets the password"
	log.Debug("WEBSERVER:   GUTS     " + s)

	return password

}

// writePasswordFile - Writes the password to file (AES-256 encryption) and puts in struct
func writePasswordFile(nodeName string, passwordString string) passwordStruct {

	s := "START  writePasswordFile() -  Writes the password to file (AES-256 encryption) and puts in struct"
	log.Debug("WEBSERVER:   GUTS     " + s)

	// LOAD password IN PASSWORD
	password = passwordStruct{passwordString}

	// ENCRYPT PASSWORD using key
	keyText := "myverystrongpasswordo32bitlength"
	keyByte := []byte(keyText)
	additionalData := "Jeff's additional data for authorization"
	passwordEncrypted := wallet.EncryptAES(keyByte, passwordString, additionalData)

	// ENCRYPTED STRUCT FOR FILE
	passwordStructEncrypted := passwordStruct{passwordEncrypted}

	// WRITE PASSWORD STRUCT TO JSON FILE
	filedata, _ := json.MarshalIndent(passwordStructEncrypted, "", " ")
	filename := "credentials/" + nodeName + "-password.json"
	_ = ioutil.WriteFile(filename, filedata, 0644)
	s = "Wrote password to " + filename
	log.Info("WEBSERVER:   GUTS            " + s)

	s = "END    writePasswordFile() -  Writes the password to file (AES-256 encryption) and puts in struct"
	log.Debug("WEBSERVER:   GUTS     " + s)

	return password

}

// readPasswordFile - Reads the password from a file (AES-256 decrypt) and puts in struct
func readPasswordFile(nodeName string) passwordStruct {

	s := "START  readPasswordFile() - Reads the password from a file (AES-256 decrypt) and puts in struct"
	log.Debug("WEBSERVER:   GUTS     " + s)

	var passwordStructEncrypted passwordStruct

	// READ PASSWORD STRUCT TO JSON FILE
	filename := "credentials/" + nodeName + "-password.json"
	filedata, _ := ioutil.ReadFile(filename)
	_ = json.Unmarshal([]byte(filedata), &passwordStructEncrypted)
	s = "Read password from " + filename
	log.Info("WEBSERVER:   GUTS            " + s)

	// DECRYPT PASSWORD using key
	keyText := "myverystrongpasswordo32bitlength"
	keyByte := []byte(keyText)
	additionalData := "Jeff's additional data for authorization"
	passwordString := wallet.DecryptAES(keyByte, passwordStructEncrypted.Password, additionalData)

	// PLACE passwordString IN STRUCT
	password.Password = passwordString

	s = "END    readPasswordFile() - Reads the password from a file (AES-256 decrypt) and puts in struct"
	log.Debug("WEBSERVER:   GUTS     " + s)

	return password

}
