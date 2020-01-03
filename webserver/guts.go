// jeffCoin 5. WEBSERVER guts.go

package webserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	bcrypt "golang.org/x/crypto/bcrypt"
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

// writePasswordFile - Writes the password hash to a file and puts in struct
func writePasswordFile(nodeName string, passwordString string) passwordStruct {

	s := "START  writePasswordFile() - Writes the password hash to a file and puts in struct"
	log.Debug("WEBSERVER:   GUTS     " + s)

	// HASH password using bcrypt
	passwordHash, _ := hashPassword(passwordString)

	// LOAD passwordHash IN password struct
	password = passwordStruct{passwordHash}

	// WRITE PASSWORD STRUCT TO JSON FILE
	filedata, _ := json.MarshalIndent(password, "", " ")
	filename := "credentials/" + nodeName + "-password-hash.json"
	_ = ioutil.WriteFile(filename, filedata, 0644)
	s = "Wrote password to " + filename
	log.Info("WEBSERVER:   GUTS            " + s)

	s = "END    writePasswordFile() - Writes the password hash to a file and puts in struct"
	log.Debug("WEBSERVER:   GUTS     " + s)

	return password

}

// readPasswordFile - Reads the password hash from a file and puts in struct
func readPasswordFile(nodeName string, passwordString string) passwordStruct {

	s := "START  readPasswordFile() - Reads the password hash from a file and puts in struct"
	log.Debug("WEBSERVER:   GUTS     " + s)

	// READ PASSWORD STRUCT TO JSON FILE and place in password
	filename := "credentials/" + nodeName + "-password-hash.json"
	filedata, _ := ioutil.ReadFile(filename)
	_ = json.Unmarshal([]byte(filedata), &password)
	s = "Read password from " + filename
	log.Info("WEBSERVER:   GUTS            " + s)

	// CHECK HASH
	status := checkPasswordHash(passwordString, password.PasswordHash)
	if !status {
		s = "Your password did not pass check hash in filename" + filename
		log.Fatal("WEBSERVER:   GUTS            " + s)
	} else {
		s = "Your password passed check hash in filename" + filename
		log.Info("WEBSERVER:   GUTS            " + s)
	}

	s = "END    readPasswordFile() - Reads the password hash from a file and puts in struct"
	log.Debug("WEBSERVER:   GUTS     " + s)

	return password

}

// hashPassword - Hashes a password using bcrypt
func hashPassword(password string) (string, error) {

	s := "START  hashPassword() - Hashes a password using bcrypt"
	log.Debug("WEBSERVER:   GUTS     " + s)

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	s = "END    hashPassword() - Hashes a password using bcrypt"
	log.Debug("WEBSERVER:   GUTS     " + s)

	return string(bytes), err

}

// checkPasswordHash - Checks a password with a hash using bcrypt
func checkPasswordHash(password, hash string) bool {

	s := "START  checkPasswordHash() - Checks a password with a hash using bcrypt"
	log.Debug("WEBSERVER:   GUTS     " + s)

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	s = "END    checkPasswordHash() - Checks a password with a hash using bcrypt"
	log.Debug("WEBSERVER:   GUTS     " + s)

	return err == nil

}
