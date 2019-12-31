// jeffCoin 5. WEBSERVER guts.go

package webserver

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
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
	passwordEncrypted := encrypt(keyByte, passwordString, additionalData)

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
	passwordString := decrypt(keyByte, passwordStructEncrypted.Password, additionalData)

	// PLACE IN STRUCT
	password.Password = passwordString

	s = "END    readPasswordFile() - Reads the password from a file (AES-256 decrypt) and puts in struct"
	log.Debug("WEBSERVER:   GUTS     " + s)

	return password

}

// ENCRYPTION/DECRYPTION *************************************************************************************************

// encrypt - AES-256 GCM (Galois/Counter Mode) mode encryption
func encrypt(keyByte []byte, plaintext string, additionalData string) string {

	s := "START  encrypt() - AES-256 GCM (Galois/Counter Mode) mode encryption"
	log.Debug("WEBSERVER:   GUTS     " + s)

	plaintextByte := []byte(plaintext)
	additionalDataByte := []byte(additionalData)

	// GET CIPHER BLOCK USING KEY
	block, err := aes.NewCipher(keyByte)
	checkErr(err)

	// GET GCM INSTANCE THAT USES THE AES CIPHER
	gcm, err := cipher.NewGCM(block)
	checkErr(err)

	// CREATE A NONCE
	nonce := make([]byte, gcm.NonceSize())
	// Populates the nonce with a cryptographically secure random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	// ENCRYPT DATA
	// Note how we put the Nonce in the beginging,
	// So we can rip it out when we decrypt
	cipherTextByte := gcm.Seal(nonce, nonce, plaintextByte, additionalDataByte)

	s = "END    encrypt() - AES-256 GCM (Galois/Counter Mode) mode encryption"
	log.Debug("WEBSERVER:   GUTS     " + s)

	// RETURN HEX
	cipherText := hex.EncodeToString(cipherTextByte)
	return cipherText

}

// decrypt - AES-256 GCM (Galois/Counter Mode) mode decryption
func decrypt(keyByte []byte, cipherText string, additionalData string) string {

	s := "START  decrypt() - AES-256 GCM (Galois/Counter Mode) mode decryption"
	log.Debug("WEBSERVER:   GUTS     " + s)

	cipherTextByte, _ := hex.DecodeString(cipherText)
	additionalDataByte := []byte(additionalData)

	// GET CIPHER BLOCK USING KEY
	block, err := aes.NewCipher(keyByte)
	checkErr(err)

	// GET GCM BLOCK
	gcm, err := cipher.NewGCM(block)
	checkErr(err)

	// EXTRACT NONCE FROM cipherTextByte
	// Because I put it there
	nonceSize := gcm.NonceSize()
	nonce, cipherTextByte := cipherTextByte[:nonceSize], cipherTextByte[nonceSize:]

	// DECRYPT DATA
	plainTextByte, err := gcm.Open(nil, nonce, cipherTextByte, additionalDataByte)
	checkErr(err)

	s = "END    decrypt() - AES-256 GCM (Galois/Counter Mode) mode decryption"
	log.Debug("WEBSERVER:   GUTS     " + s)

	// RETURN STRING
	plainText := string(plainTextByte[:])
	return plainText

}
