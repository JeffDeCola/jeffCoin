// jeffCoin TESTDATA data.go

package testdata

const (
	foundersPubKey  = "Founders PubKey"
	jeffPubKey      = "Jeffs PubKey"
	mattPubKey      = "Matts PubKey"
	coinVaultPubKey = "CoinVaults PubKey"
	jillPubKey      = "Jills PubKey"
	foundersSig     = "Founders Signature"
	jeffSig         = "Jeffs Signature"
	mattSig         = "Matts Signature"
)

// TRANSACTION REQUEST MESSAGES SIGNED
const (
	txRequestMessageSignedDataString1 = `
        {
            "txRequestMessage": {
                "sourceAddress": "` + foundersPubKey + `",
                "destinations": [
                    {
                        "destinationAddress": "` + jeffPubKey + `",
                        "value": 80000
                    }
                ]
            },
            "signature": "` + foundersSig + `"
        }`
	txRequestMessageSignedDataString2 = `
        {
            "txRequestMessage": {
                "sourceAddress": "` + jeffPubKey + `",
                "destinations": [
                    {
                        "destinationAddress": "` + mattPubKey + `",
                        "value": 50000
                    },
                    {
                        "destinationAddress": "` + coinVaultPubKey + `",
                        "value": 500
                    }
                ]
            },
            "signature": "` + jeffSig + `"
        }`
	txRequestMessageSignedDataString3 = `
        {
            "txRequestMessage": {
                "sourceAddress": "` + foundersPubKey + `",
                "destinations": [
                    {
                        "destinationAddress": "` + mattPubKey + `",
                        "value": 250000
                    },
                    {
                        "destinationAddress": "` + jeffPubKey + `",
                        "value": 13000
                    }
                ]
            },
            "signature": "` + foundersSig + `"
        }`
	txRequestMessageSignedDataString4 = `
        {
            "txRequestMessage": {
                "sourceAddress": "` + mattPubKey + `",
                "destinations": [
                    {
                        "destinationAddress": "` + jillPubKey + `",
                        "value": 35000
                    }
                ]
            },
            "signature": "` + mattSig + `"
        }`
	txRequestMessageSignedDataString5 = `
        {
            "txRequestMessage": {
                "sourceAddress": "` + mattPubKey + `",
                "destinations": [
                    {
                        "destinationAddress": "` + jeffPubKey + `",
                        "value": 15000
                    }
                ]
            },
            "signature": "` + mattSig + `"
        }`
	txRequestMessageSignedDataString6 = `
        {
            "txRequestMessage": {
                "sourceAddress": "` + jeffPubKey + `",
                "destinations": [
                    {
                        "destinationAddress": "` + jillPubKey + `",
                        "value": 33000
                    }
                ]
            },
            "signature": "` + jeffSig + `"
        }`
	txRequestMessageSignedDataStringBad = `
        {
            "txRequestMessage": {
                "sourceAddress": "` + foundersPubKey + `",
                "destinations": [
                    {
                        "destinationAddress": "Bad Pub Key",
                        "value": 80000
                    }
                ]
            },
            "signature": "Bad"
        }`
)
