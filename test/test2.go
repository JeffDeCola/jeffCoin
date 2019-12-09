// my-go-examples ecdsa-asymmetric-cryptography.go

package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"

	"github.com/btcsuite/btcutil/base58"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ripemd160"
)

var filename string

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error is %+v\n", err)
		log.Fatal("ERROR:", err)
	}
}

func readFile(filename string) string {

	plainTextByte, err := ioutil.ReadFile(filename)
	checkErr(err)
	// Convert to string
	plainText := string(plainTextByte)
	return plainText

}

func generateECDSAKeys() (string, string) {

	// GENERATE PRIVATE & PUBLIC KEY PAIR
	curve := elliptic.P256()
	privateKeyRaw, err := ecdsa.GenerateKey(curve, rand.Reader)
	checkErr(err)

	// EXTRACT PUBLIC KEY
	publicKeyRaw := &privateKeyRaw.PublicKey

	// ENCODE
	privateKeyHex, publicKeyHex := encodeKeys(privateKeyRaw, publicKeyRaw)

	return privateKeyHex, publicKeyHex

}

// encodeKeys - Encodes privateKeyRaw & publicKeyRaw to privateKeyHex & publicKeyHex
func encodeKeys(privateKeyRaw *ecdsa.PrivateKey, publicKeyRaw *ecdsa.PublicKey) (string, string) {

	s := "START  encodeKeys() - Encodes privateKeyRaw & publicKeyRaw to privateKeyHex & publicKeyHex"
	log.Trace("WALLET:      GUTS     " + s)

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
	log.Trace("WALLET:      GUTS     " + s)

	return privateKeyHex, publicKeyHex

}

func createSignature(privateKeyHex string, plainText string) string {

	// DECODE
	privateKeyPEM, _ := hex.DecodeString(privateKeyHex)
	block, _ := pem.Decode([]byte(privateKeyPEM))
	privateKeyx509Encoded := block.Bytes
	privateKeyRaw, _ := x509.ParseECPrivateKey(privateKeyx509Encoded)

	// HASH plainText
	hashedPlainText := sha256.Sum256([]byte(plainText))
	hashedPlainTextByte := hashedPlainText[:]

	r := big.NewInt(0)
	s := big.NewInt(0)

	// CREATE SIGNATURE
	r, s, err := ecdsa.Sign(
		rand.Reader,
		privateKeyRaw,
		hashedPlainTextByte,
	)
	checkErr(err)

	signatureByte := r.Bytes()
	signatureByte = append(signatureByte, s.Bytes()...)

	// ENCODE - RETURN HEX
	signature := hex.EncodeToString(signatureByte)

	return signature

}

func verifySignature(publicKeyHex string, signature string, plainText string) bool {

	// DECODE
	publicKeyPEM, _ := hex.DecodeString(publicKeyHex)
	blockPub, _ := pem.Decode([]byte(publicKeyPEM))
	publicKeyx509Encoded := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(publicKeyx509Encoded)
	publicKeyRaw := genericPublicKey.(*ecdsa.PublicKey)

	// HASH plainText
	hashedPlainText := sha256.Sum256([]byte(plainText))
	hashedPlainTextByte := hashedPlainText[:]

	// DECODE signature
	signatureByte, _ := hex.DecodeString(signature)

	// EXTRACT R & S
	r := big.NewInt(0)
	s := big.NewInt(0)
	sigLen := len(signatureByte)
	r.SetBytes(signatureByte[:(sigLen / 2)])
	s.SetBytes(signatureByte[(sigLen / 2):])

	// VERIFY SIGNATURE
	verifyStatus := ecdsa.Verify(
		publicKeyRaw,
		hashedPlainTextByte,
		r,
		s,
	)

	return verifyStatus

}

func init() {

	// GET FILE NAME FROM ARGS
	flag.Parse()
	filenameSlice := flag.Args()
	if len(filenameSlice) != 1 {
		err := errors.New("Only one filename allowed")
		checkErr(err)
	}

	filename = filenameSlice[0] // Make it a string

}

// generatejeffCoinAddress - Creates a jeffCoin Address
func generatejeffCoinAddress(publicKeyHex string) string {

	verPublicKeyHash := hashPublicKey(publicKeyHex)

	checkSum := checksumKeyHash(verPublicKeyHash)

	jeffCoinAddressHex := encodeKeyHash(verPublicKeyHash, checkSum)

	return jeffCoinAddressHex

}

// hashPublicKey - Hashes publicKeyHex
func hashPublicKey(publicKeyHex string) []byte {

	// 1 - SHA-256 HASH
	publicKeyByte, _ := hex.DecodeString(publicKeyHex) // Don't use []byte(publicKeyHex)
	publicSHA256 := sha256.Sum256(publicKeyByte)
	s := "1 - SHA-256 HASH:    " + hex.EncodeToString(publicSHA256[:])
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

	return verPublicKeyHash

}

// checksumKeyHash - Checksums verPublicKeyHash
func checksumKeyHash(verPublicKeyHash []byte) []byte {

	// 4 - SHA-256 HASH
	firstpublicSHA256 := sha256.Sum256(verPublicKeyHash)
	s := "4 - SHA-256 HASH     " + hex.EncodeToString(firstpublicSHA256[:])
	log.Info("WALLET:      GUTS            " + s)

	// 5 - SHA-256 HASH
	secondPublicSHA256 := sha256.Sum256(firstpublicSHA256[:])
	s = "5 - SHA-256 HASH     " + hex.EncodeToString(secondPublicSHA256[:])
	log.Info("WALLET:      GUTS            " + s)

	// 6 - FIRST FOUR BYTE
	checkSum := secondPublicSHA256[:4]
	s = "6 - FIRST FOUR BYTE  " + hex.EncodeToString(checkSum)
	log.Info("WALLET:      GUTS            " + s)

	return checkSum

}

// encodeKeyHash - Encodes verPublicKeyHash & checkSum
func encodeKeyHash(verPublicKeyHash []byte, checkSum []byte) string {

	// 7 - CONCAT
	addressHex := append(verPublicKeyHash, checkSum...)
	s := "7 - CONCAT           " + hex.EncodeToString(addressHex)
	log.Info("WALLET:      GUTS            " + s)

	// 8 - BASE58 ENCODING
	jeffCoinAddressHex := base58.Encode(addressHex)
	s = "8 - BASE58 ENCODING  " + jeffCoinAddressHex
	log.Info("WALLET:      GUTS            " + s)

	return jeffCoinAddressHex

}

func main() {

	fmt.Println(" ")

	// READ FILE INTO STRING
	plainText := readFile(filename)
	fmt.Printf("The original message contains:\n\n%s\n\n", plainText)

	// SENDER GENERATE RSA KEYS
	privateKeyHex, publicKeyHex := generateECDSAKeys()
	fmt.Printf("Public key is:\n\n%s\n\n", publicKeyHex)

	// GET ADDRESS
	address := generatejeffCoinAddress(publicKeyHex)
	fmt.Printf("\nPublic address is:\n\n%s\n\n", address)

	// CREATE SIGNATURE
	signature := createSignature(privateKeyHex, plainText)
	fmt.Printf("The senders signature:\n\n%s\n\n", signature)

	// VERIFY SIGNATURE
	verifyStatus := verifySignature(address, signature, plainText)
	fmt.Printf("The senders signature verification is: %v\n\n", verifyStatus)

}
