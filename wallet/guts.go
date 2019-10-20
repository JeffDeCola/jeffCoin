// jeffCoin guts.go

package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"

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

// WALLET **************************************************************************************************************

// getWallet - Gets the wallet
func getWallet() walletStruct {

	s := "START: getWallet - Gets the wallet"
	log.Trace("WALLET:      GUTS   " + s)

	s = "END:   getWallet - Gets the wallet"
	log.Trace("WALLET:      GUTS   " + s)

	return wallet
}

func makeWallet() walletStruct {

	// GENERATE KEYS
	privateKeyByte, publicKeyByte := newKeyPair()

	// GET ADDRESS
	address := getAddress(publicKeyByte)

	// PLACE KEYS & ADDRESS IN WALLET
	wallet = walletStruct{privateKeyByte, publicKeyByte, address}

	return wallet

}

func newKeyPair() ([]byte, []byte) {

	// GET PRIVATE & PUBLIC KEY PAIR
	curve := elliptic.P256()
	privateKeyRaw, err := ecdsa.GenerateKey(curve, rand.Reader)
	checkErr(err)

	// CHANGE FORMAT OF PRIVATE KEY
	privateKeys509Binary, _ := x509.MarshalECPrivateKey(privateKeyRaw)
	privateKeyByte := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: privateKeys509Binary,
		})
	//privateKeyHex := hex.EncodeToString(privateKeyByte)
	fmt.Println("PRIVATE HEX", string(privateKeyByte))


    // GET PUBLIC
    // publicKeyByte := append(privateKeyRaw.PublicKey.X.Bytes(), privateKeyRaw.PublicKey.Y.Bytes()...)

	// CHANGE FORMAT OF PUBLIC KEY
	publicKeys509Binary, _ := x509.MarshalPKIXPublicKey(publicKeyRaw)
	publicKeyByte := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: publicKeys509Binary,
		})
	//publicKeyHex := hex.EncodeToString(publicKeyByte)
	fmt.Println("PUBLIC HEX", string(publicKeyByte))

	return privateKeyByte, publicKeyByte
}

func getAddress(PublicKeyByte []byte) string {

	version := make([]byte, 1)
	version[0] = 0x00 // prefix with 00 if it's mainnet

	pubKeyHash := HashPubKey(PublicKeyByte)

	versionedPayload := append(version, pubKeyHash...)
	checksum := checksum(versionedPayload)

	fullPayload := append(versionedPayload, checksum...)
	address := base58.Encode(fullPayload) // encode to base58

	return address
}

func HashPubKey(pubKey []byte) []byte {
	publicSHA256 := sha256.Sum256(pubKey)

	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	checkErr(err)

	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)

	return publicRIPEMD160
}

func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])

	return secondSHA[:4]
}
