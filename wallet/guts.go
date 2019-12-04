// jeffCoin 4. WALLET guts.go

package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
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

	s := "START: getWallet - Gets the wallet"
	log.Trace("WALLET:      GUTS   " + s)

	s = "END:   getWallet - Gets the wallet"
	log.Trace("WALLET:      GUTS   " + s)

	return wallet
}

// makeWallet - Creates wallet with Keys and jeffCoin address
func makeWallet() walletStruct {

	s := "START: makeWallet - Creates wallet with Keys and jeffCoin address"
	log.Trace("WALLET:      GUTS   " + s)

	// GENERATE ECDSA KEYS
	privateKeyHex, publicKeyHex := generateECDSASKeys()

	// GET JEFFCOIN ADDRESS
	jeffCoinAddressHex := generatejeffCoinAddress(publicKeyHex)

	// LOAD KEYS & JEFFCOIN ADDRESS IN WALLET
	wallet = walletStruct{privateKeyHex, publicKeyHex, jeffCoinAddressHex}

	s = "END:   makeWallet - Creates wallet with Keys and jeffCoin address"
	log.Trace("WALLET:      GUTS   " + s)

	return wallet

}

// KEYS ******************************************************************************************************************

// generateECDSASKeys - Generate privateKeyHex and publicKeyHex
func generateECDSASKeys() (string, string) {

	s := "START: generateECDSASKeys - Generate privateKeyHex and publicKeyHex"
	log.Trace("WALLET:      GUTS   " + s)

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
		log.Error("WALLET:      GUTS   " + s)
	}
	if !reflect.DeepEqual(publicKeyRaw, publicKeyRawCheck) {
		s = "ERROR: Public keys do not match"
		log.Error("WALLET:      GUTS   " + s)
	}

	s = "END:   generateECDSASKeys - Generate privateKeyHex and publicKeyHex"
	log.Trace("WALLET:      GUTS   " + s)

	return privateKeyHex, publicKeyHex

}

// encodeKeys - Encodes privateKeyRaw & publicKeyRaw to privateKeyHex & publicKeyHex
func encodeKeys(privateKeyRaw *ecdsa.PrivateKey, publicKeyRaw *ecdsa.PublicKey) (string, string) {

	s := "START: encodeKeys - Encodes privateKeyRaw & publicKeyRaw to privateKeyHex & publicKeyHex"
	log.Trace("WALLET:      GUTS   " + s)

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

	s = "END:   encodeKeys - Encodes privateKeyRaw & publicKeyRaw to privateKeyHex & publicKeyHex"
	log.Trace("WALLET:      GUTS   " + s)

	return privateKeyHex, publicKeyHex

}

// decodeKeys - Decodes privateKeyHex & publicKeyHex to privateKeyRaw & publicKeyRaw
func decodeKeys(privateKeyHex string, publicKeyHex string) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {

	s := "START: decodeKeys - Decodes privateKeyHex & publicKeyHex to privateKeyRaw & publicKeyRaw"
	log.Trace("WALLET:      GUTS   " + s)

	privateKeyPEM, _ := hex.DecodeString(privateKeyHex)
	block, _ := pem.Decode([]byte(privateKeyPEM))
	privateKeyx509Encoded := block.Bytes
	privateKeyRaw, _ := x509.ParseECPrivateKey(privateKeyx509Encoded)

	publicKeyPEM, _ := hex.DecodeString(publicKeyHex)
	blockPub, _ := pem.Decode([]byte(publicKeyPEM))
	publicKeyx509Encoded := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(publicKeyx509Encoded)
	publicKeyRaw := genericPublicKey.(*ecdsa.PublicKey)

	s = "END:   decodeKeys - Decodes privateKeyHex & publicKeyHex to privateKeyRaw & publicKeyRaw"
	log.Trace("WALLET:      GUTS   " + s)

	return privateKeyRaw, publicKeyRaw

}

// JEFFCOIN ADDRESS ******************************************************************************************************

// generatejeffCoinAddress - Creates jeffCoinAddress
func generatejeffCoinAddress(publicKeyHex string) string {

	s := "START: generatejeffCoinAddress - Creates jeffCoinAddress"
	log.Trace("WALLET:      GUTS   " + s)

	verPublicKeyHash := hashPublicKey(publicKeyHex)

	checkSum := checksumKeyHash(verPublicKeyHash)

	jeffCoinAddressHex := encodeKeyHash(verPublicKeyHash, checkSum)

	s = "END:   generatejeffCoinAddress - Creates jeffCoinAddress"
	log.Trace("WALLET:      GUTS   " + s)

	return jeffCoinAddressHex

}

// hashPublicKey - Hashes publicKeyHex
func hashPublicKey(publicKeyHex string) []byte {

	s := "START: hashPublicKey - Hashes publicKeyHex"
	log.Trace("WALLET:      GUTS   " + s)

	// 1 - SHA-256 HASH
	publicKeyByte, _ := hex.DecodeString(publicKeyHex) // Don't use []byte(publicKeyHex)
	publicSHA256 := sha256.Sum256(publicKeyByte)
	s = "1 - SHA-256 HASH:    " + hex.EncodeToString(publicSHA256[:])
	log.Info("WALLET:      GUTS          " + s)

	// 2 - RIPEMD-160 HASH
	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	checkErr(err)
	publicKeyHash := RIPEMD160Hasher.Sum(nil)
	s = "2 - RIPEMD-160 HASH: " + hex.EncodeToString(publicKeyHash)
	log.Info("WALLET:      GUTS          " + s)

	// VERSION
	version := make([]byte, 1)
	version[0] = 0x00
	s = "VERSION:             " + "00"
	log.Info("WALLET:      GUTS          " + s)

	// 3 - CONCAT
	verPublicKeyHash := append(version, publicKeyHash...)
	s = "3 - CONCAT           " + hex.EncodeToString(verPublicKeyHash)
	log.Info("WALLET:      GUTS          " + s)

	s = "END:   hashPublicKey - Hashes publicKeyHex"
	log.Trace("WALLET:      GUTS   " + s)

	return verPublicKeyHash

}

// checksumKeyHash() - Checksums verPublicKeyHash
func checksumKeyHash(verPublicKeyHash []byte) []byte {

	s := "START: checksumKeyHash() - Checksums verPublicKeyHash"
	log.Trace("WALLET:      GUTS   " + s)

	// 4 - SHA-256 HASH
	firstpublicSHA256 := sha256.Sum256(verPublicKeyHash)
	s = "4 - SHA-256 HASH     " + hex.EncodeToString(firstpublicSHA256[:])
	log.Info("WALLET:      GUTS          " + s)

	// 5 - SHA-256 HASH
	secondPublicSHA256 := sha256.Sum256(firstpublicSHA256[:])
	s = "5 - SHA-256 HASH     " + hex.EncodeToString(secondPublicSHA256[:])
	log.Info("WALLET:      GUTS          " + s)

	// 6 - FIRST FOUR BYTE
	checkSum := secondPublicSHA256[:4]
	s = "6 - FIRST FOUR BYTE  " + hex.EncodeToString(checkSum)
	log.Info("WALLET:      GUTS          " + s)

	s = "END:   checksumKeyHash() - Checksums verPublicKeyHash"
	log.Trace("WALLET:      GUTS   " + s)

	return checkSum

}

// encodeKeyHash - Encodes verPublicKeyHash & checkSum
func encodeKeyHash(verPublicKeyHash []byte, checkSum []byte) string {

	s := "START: encodeKeyHash - Encodes verPublicKeyHash & checkSum"
	log.Trace("WALLET:      GUTS   " + s)

	// 7 - CONCAT
	addressHex := append(verPublicKeyHash, checkSum...)
	s = "7 - CONCAT           " + hex.EncodeToString(addressHex)
	log.Info("WALLET:      GUTS          " + s)

	// 8 - BASE58 ENCODING
	jeffCoinAddressHex := base58.Encode(addressHex)
	s = "8 - BASE58 ENCODING  " + jeffCoinAddressHex
	log.Info("WALLET:      GUTS          " + s)

	s = "END:   encodeKeyHash - Encodes verPublicKeyHash & checkSum"
	log.Trace("WALLET:      GUTS   " + s)

	return jeffCoinAddressHex

}

// SIGNATURE *************************************************************************************************************

// createSignature - Create a ECDSA Digital Signature
func createSignature(senderPrivateKeyRaw *ecdsa.PrivateKey, plainText string) string {

	s := "START: createSignature - Create a ECDSA Digital Signature"
	log.Trace("WALLET:      GUTS   " + s)

	// HASH plainText
	hashedPlainText := sha256.Sum256([]byte(plainText))
	hashedPlainTextByte := hashedPlainText[:]

	rSign := big.NewInt(0)
	sSign := big.NewInt(0)

	// CREATE SIGNATURE
	rSign, sSign, err := ecdsa.Sign(
		rand.Reader,
		senderPrivateKeyRaw,
		hashedPlainTextByte,
	)
	checkErr(err)

	signatureByte := rSign.Bytes()
	signatureByte = append(signatureByte, sSign.Bytes()...)

	// ENCODE - RETURN HEX
	signature := hex.EncodeToString(signatureByte)

	s = "END:   createSignature - Create a ECDSA Digital Signature"
	log.Trace("WALLET:      GUTS   " + s)

	return signature

}
