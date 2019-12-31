// jeffCoin 4. WALLET guts.go

package wallet

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"reflect"

	"github.com/btcsuite/btcutil/base58"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ripemd160"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

// WALLET ****************************************************************************************************************

// getWallet - Gets the wallet
func getWallet() walletStruct {

	s := "START  getWallet() - Gets the wallet"
	log.Debug("WALLET:      GUTS     " + s)

	s = "END    getWallet() - Gets the wallet"
	log.Debug("WALLET:      GUTS     " + s)

	return wallet
}

// makeWallet - Creates the wallet and writes to file (Keys and jeffCoin Address)
func makeWallet(nodeName string) walletStruct {

	s := "START  makeWallet() - Creates the wallet and writes to file (Keys and jeffCoin Address)"
	log.Debug("WALLET:      GUTS     " + s)

	// GENERATE ECDSA KEYS
	privateKeyHex, publicKeyHex := generateECDSASKeys()

	// GENERATE JEFFCOIN ADDRESS
	jeffCoinAddressHex := generatejeffCoinAddress(publicKeyHex)

	// LOAD KEYS & JEFFCOIN ADDRESS IN WALLET
	wallet = walletStruct{privateKeyHex, publicKeyHex, jeffCoinAddressHex}

	// ENCRYPT privateKeyHex using key
	keyText := "myverystrongpasswordo32bitlength"
	keyByte := []byte(keyText)
	additionalData := "Jeff's additional data for authorization"
	privateKeyHexEncrypted := EncryptAES(keyByte, privateKeyHex, additionalData)

	// ENCRYPTED STRUCT FOR FILE
	walletStructEncrypted := walletStruct{privateKeyHexEncrypted, publicKeyHex, jeffCoinAddressHex}

	// WRITE WALLET STRUCT TO JSON FILE
	filedata, _ := json.MarshalIndent(walletStructEncrypted, "", " ")
	filename := "wallet/" + nodeName + "-wallet.json"
	_ = ioutil.WriteFile(filename, filedata, 0644)
	s = "Wrote wallet to " + filename
	log.Info("WALLET:      GUTS            " + s)

	s = "END    makeWallet() - Creates the wallet and writes to file (Keys and jeffCoin Address)"
	log.Debug("WALLET:      GUTS     " + s)

	return wallet

}

// readWalletFile - Reads the wallet from a file and puts in struct
func readWalletFile(nodeName string) walletStruct {

	s := "START  readWalletFile() - Reads the wallet from a file and puts in struct"
	log.Debug("WALLET:      GUTS     " + s)

	var walletStructEncrypted walletStruct

	// READ WALLET STRUCT TO JSON FILE
	filename := "wallet/" + nodeName + "-wallet.json"
	filedata, _ := ioutil.ReadFile(filename)
	_ = json.Unmarshal([]byte(filedata), &walletStructEncrypted)
	s = "Read wallet from " + filename
	log.Info("WALLET:      GUTS            " + s)

	// DECRYPT privateKeyHex using key
	keyText := "myverystrongpasswordo32bitlength"
	keyByte := []byte(keyText)
	additionalData := "Jeff's additional data for authorization"
	privateKeyHex := DecryptAES(keyByte, walletStructEncrypted.PrivateKey, additionalData)

	// PLACE privateKeyHex IN STRUCT
	wallet.PrivateKey = privateKeyHex

	s = "END    readWalletFile() - Reads the wallet from a file and puts in struct"
	log.Debug("WALLET:      GUTS     " + s)

	return wallet

}

// KEYS ******************************************************************************************************************

// generateECDSASKeys - Generate privateKeyHex and publicKeyHex
func generateECDSASKeys() (string, string) {

	s := "START  generateECDSASKeys() - Generate privateKeyHex and publicKeyHex"
	log.Debug("WALLET:      GUTS     " + s)

	// GET PRIVATE & PUBLIC KEY PAIR
	curve := elliptic.P256()
	privateKeyRaw, err := ecdsa.GenerateKey(curve, rand.Reader)
	checkErr(err)

	// EXTRACT PUBLIC KEY
	publicKeyRaw := &privateKeyRaw.PublicKey

	// ENCODE
	privateKeyHex, publicKeyHex := encodeKeys(privateKeyRaw, publicKeyRaw)

	// DECODE TO CHECK
	privateKeyRawCheck, publicKeyRawCheck := decodeKeys(privateKeyHex, publicKeyHex)

	// CHECK THEY ARE THE SAME
	if !reflect.DeepEqual(privateKeyRaw, privateKeyRawCheck) {
		s = "ERROR: Private keys do not match"
		log.Error("WALLET:      GUTS     " + s)
	}
	if !reflect.DeepEqual(publicKeyRaw, publicKeyRawCheck) {
		s = "ERROR: Public keys do not match"
		log.Error("WALLET:      GUTS     " + s)
	}

	s = "END    generateECDSASKeys() - Generate privateKeyHex and publicKeyHex"
	log.Debug("WALLET:      GUTS     " + s)

	return privateKeyHex, publicKeyHex

}

// encodeKeys - Encodes privateKeyRaw & publicKeyRaw to privateKeyHex & publicKeyHex
func encodeKeys(privateKeyRaw *ecdsa.PrivateKey, publicKeyRaw *ecdsa.PublicKey) (string, string) {

	s := "START  encodeKeys() - Encodes privateKeyRaw & publicKeyRaw to privateKeyHex & publicKeyHex"
	log.Debug("WALLET:      GUTS     " + s)

	privateKeyx509Encoded, _ := x509.MarshalECPrivateKey(privateKeyRaw)
	privateKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: privateKeyx509Encoded,
		})
	privateKeyHex := hex.EncodeToString(privateKeyPEM)

	publicKeyx509Encoded, _ := x509.MarshalPKIXPublicKey(publicKeyRaw)
	publicKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: publicKeyx509Encoded,
		})
	publicKeyHex := hex.EncodeToString(publicKeyPEM)

	s = "END    encodeKeys() - Encodes privateKeyRaw & publicKeyRaw to privateKeyHex & publicKeyHex"
	log.Debug("WALLET:      GUTS     " + s)

	return privateKeyHex, publicKeyHex

}

// decodeKeys - Decodes privateKeyHex & publicKeyHex to privateKeyRaw & publicKeyRaw
func decodeKeys(privateKeyHex string, publicKeyHex string) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {

	s := "START  decodeKeys() - Decodes privateKeyHex & publicKeyHex to privateKeyRaw & publicKeyRaw"
	log.Debug("WALLET:      GUTS     " + s)

	privateKeyPEM, _ := hex.DecodeString(privateKeyHex)
	block, _ := pem.Decode([]byte(privateKeyPEM))
	privateKeyx509Encoded := block.Bytes
	privateKeyRaw, _ := x509.ParseECPrivateKey(privateKeyx509Encoded)

	publicKeyPEM, _ := hex.DecodeString(publicKeyHex)
	blockPub, _ := pem.Decode([]byte(publicKeyPEM))
	publicKeyx509Encoded := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(publicKeyx509Encoded)
	publicKeyRaw := genericPublicKey.(*ecdsa.PublicKey)

	s = "END    decodeKeys() - Decodes privateKeyHex & publicKeyHex to privateKeyRaw & publicKeyRaw"
	log.Debug("WALLET:      GUTS     " + s)

	return privateKeyRaw, publicKeyRaw

}

// JEFFCOIN ADDRESS ******************************************************************************************************

// generatejeffCoinAddress - Creates a jeffCoin Address
func generatejeffCoinAddress(publicKeyHex string) string {

	s := "START  generatejeffCoinAddress() - Creates a jeffCoin Address"
	log.Debug("WALLET:      GUTS     " + s)

	verPublicKeyHash := hashPublicKey(publicKeyHex)

	checkSum := checksumKeyHash(verPublicKeyHash)

	jeffCoinAddressHex := encodeKeyHash(verPublicKeyHash, checkSum)

	s = "END    generatejeffCoinAddress() - Creates a jeffCoin Address"
	log.Debug("WALLET:      GUTS     " + s)

	return jeffCoinAddressHex

}

// hashPublicKey - Hashes publicKeyHex
func hashPublicKey(publicKeyHex string) []byte {

	s := "START  hashPublicKey() - Hashes publicKeyHex"
	log.Debug("WALLET:      GUTS     " + s)

	// 1 - SHA-256 HASH
	publicKeyByte, _ := hex.DecodeString(publicKeyHex) // Don't use []byte(publicKeyHex)
	publicSHA256 := sha256.Sum256(publicKeyByte)
	s = "1 - SHA-256 HASH:    " + hex.EncodeToString(publicSHA256[:])
	log.Info("WALLET:      GUTS            " + s)

	// 2 - RIPEMD-160 HASH
	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	checkErr(err)
	publicKeyHash := RIPEMD160Hasher.Sum(nil)
	s = "2 - RIPEMD-160 HASH: " + hex.EncodeToString(publicKeyHash)
	log.Info("WALLET:      GUTS            " + s)

	// VERSION
	version := make([]byte, 1)
	version[0] = 0x00
	s = "VERSION:             " + "00"
	log.Info("WALLET:      GUTS            " + s)

	// 3 - CONCAT
	verPublicKeyHash := append(version, publicKeyHash...)
	s = "3 - CONCAT           " + hex.EncodeToString(verPublicKeyHash)
	log.Info("WALLET:      GUTS            " + s)

	s = "END    hashPublicKey() - Hashes publicKeyHex"
	log.Debug("WALLET:      GUTS     " + s)

	return verPublicKeyHash

}

// checksumKeyHash - Checksums verPublicKeyHash
func checksumKeyHash(verPublicKeyHash []byte) []byte {

	s := "START  checksumKeyHash() - Checksums verPublicKeyHash"
	log.Debug("WALLET:      GUTS     " + s)

	// 4 - SHA-256 HASH
	firstpublicSHA256 := sha256.Sum256(verPublicKeyHash)
	s = "4 - SHA-256 HASH     " + hex.EncodeToString(firstpublicSHA256[:])
	log.Info("WALLET:      GUTS            " + s)

	// 5 - SHA-256 HASH
	secondPublicSHA256 := sha256.Sum256(firstpublicSHA256[:])
	s = "5 - SHA-256 HASH     " + hex.EncodeToString(secondPublicSHA256[:])
	log.Info("WALLET:      GUTS            " + s)

	// 6 - FIRST FOUR BYTE
	checkSum := secondPublicSHA256[:4]
	s = "6 - FIRST FOUR BYTE  " + hex.EncodeToString(checkSum)
	log.Info("WALLET:      GUTS            " + s)

	s = "END    checksumKeyHash() - Checksums verPublicKeyHash"
	log.Debug("WALLET:      GUTS     " + s)

	return checkSum

}

// encodeKeyHash - Encodes verPublicKeyHash & checkSum
func encodeKeyHash(verPublicKeyHash []byte, checkSum []byte) string {

	s := "START  encodeKeyHash() - Encodes verPublicKeyHash & checkSum"
	log.Debug("WALLET:      GUTS     " + s)

	// 7 - CONCAT
	addressHex := append(verPublicKeyHash, checkSum...)
	s = "7 - CONCAT           " + hex.EncodeToString(addressHex)
	log.Info("WALLET:      GUTS            " + s)

	// 8 - BASE58 ENCODING
	jeffCoinAddressHex := base58.Encode(addressHex)
	s = "8 - BASE58 ENCODING  " + jeffCoinAddressHex
	log.Info("WALLET:      GUTS            " + s)

	s = "END    encodeKeyHash() - Encodes verPublicKeyHash & checkSum"
	log.Debug("WALLET:      GUTS     " + s)

	return jeffCoinAddressHex

}

// SIGNATURE *************************************************************************************************************

// createSignature - Creates a ECDSA Digital Signature
func createSignature(privateKeyHex string, plainText string) string {

	s := "START  createSignature() - Creates a ECDSA Digital Signature"
	log.Debug("WALLET:      GUTS     " + s)

	// DECODE PRIVATE KEY
	privateKeyPEM, _ := hex.DecodeString(privateKeyHex)
	block, _ := pem.Decode([]byte(privateKeyPEM))
	privateKeyx509Encoded := block.Bytes
	privateKeyRaw, _ := x509.ParseECPrivateKey(privateKeyx509Encoded)

	// HASH plainText
	hashedPlainText := sha256.Sum256([]byte(plainText))
	hashedPlainTextByte := hashedPlainText[:]

	r := big.NewInt(0)
	ss := big.NewInt(0)

	// CREATE SIGNATURE
	r, ss, err := ecdsa.Sign(
		rand.Reader,
		privateKeyRaw,
		hashedPlainTextByte,
	)
	checkErr(err)

	signatureByte := r.Bytes()
	signatureByte = append(signatureByte, ss.Bytes()...)

	// ENCODE - RETURN HEX
	signature := hex.EncodeToString(signatureByte)

	s = "END    createSignature() - Creates a ECDSA Digital Signature"
	log.Debug("WALLET:      GUTS     " + s)

	return signature

}

// ENCRYPT/DECRYPT TEXT **************************************************************************************************

// EncryptAES - AES-256 GCM (Galois/Counter Mode) mode encryption
func EncryptAES(keyByte []byte, plaintext string, additionalData string) string {

	s := "START  EncryptAES() - AES-256 GCM (Galois/Counter Mode) mode encryption"
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

	s = "END    EncryptAES() - AES-256 GCM (Galois/Counter Mode) mode encryption"
	log.Debug("WEBSERVER:   GUTS     " + s)

	// RETURN HEX
	cipherText := hex.EncodeToString(cipherTextByte)
	return cipherText

}

// DecryptAES - AES-256 GCM (Galois/Counter Mode) mode decryption
func DecryptAES(keyByte []byte, cipherText string, additionalData string) string {

	s := "START  DecryptAES() - AES-256 GCM (Galois/Counter Mode) mode decryption"
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

	s = "END    DecryptAES() - AES-256 GCM (Galois/Counter Mode) mode decryption"
	log.Debug("WEBSERVER:   GUTS     " + s)

	// RETURN STRING
	plainText := string(plainTextByte[:])
	return plainText

}
