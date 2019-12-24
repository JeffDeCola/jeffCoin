// jeffCoin TESTDATA data.go

package testdata

const (
	foundersPubKey       = "2d2d2d2d2d424547494e205055424c4943204b45592d2d2d2d2d0a4d466b77457759484b6f5a497a6a3043415159494b6f5a497a6a3044415163445167414535727549786d5146754548414c747437517778636247446b3863705a0a49684971512b4f4f336a492b31723177734a33686c693133414763686f66523639574b4c354b394a34574278696e2f37736c6f6f7030665268513d3d0a2d2d2d2d2d454e44205055424c4943204b45592d2d2d2d2d0a"
	jeffPubKey           = "Jeffs PubKey"
	mattPubKey           = "Matts PubKey"
	coinVaultPubKey      = "CoinVaults PubKey"
	jillPubKey           = "Jills PubKey"
	signatureDataString1 = "5baca9c91aca2224a0110838769ccc00a798411e155648e3232110372dc7ee7dffc621bd71f008aaf1e2f4940cd318c51162f5009bea6df1cef3e0b3185dfa1f"
	signatureDataString2 = "aa"
	signatureDataString3 = "aa"
	signatureDataString4 = "aa"
	signatureDataString5 = "aa"
	signatureDataString6 = "aa"
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
            "signature": "` + signatureDataString1 + `"
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
            "signature": "` + signatureDataString2 + `"
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
            "signature": "` + signatureDataString3 + `"
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
            "signature": "` + signatureDataString4 + `"
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
            "signature": "` + signatureDataString5 + `"
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
            "signature": "` + signatureDataString6 + `"
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
