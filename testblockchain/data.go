// jeffCoin TESTDATA data.go

package testblockchain

// Public Keys (address) and mock signatures
const (
	MockFoundersPubKey   = "2d2d2d2d2d424547494e205055424c4943204b45592d2d2d2d2d0a4d466b77457759484b6f5a497a6a3043415159494b6f5a497a6a304441516344516741456f4979436c746a42504f362b546d644c2b2f31455476614a6a6a33460a4b2f616d3765323276747365445a67442f322f794b767571735959596f7072366b6338736c76582b716a414c75397766314c4d43704a64652f773d3d0a2d2d2d2d2d454e44205055424c4943204b45592d2d2d2d2d0a"
	MockJeffPubKey       = "2d2d2d2d2d424547494e205055424c4943204b45592d2d2d2d2d0a4d466b77457759484b6f5a497a6a3043415159494b6f5a497a6a3044415163445167414577354d774a4d356451377a635536356b4a5a6d746b6c4b47476a31440a335a4a4952787545716d366834316e54717433484a6e6e774258702b586c746f2f313970367032746b76357438455979384136706559336732673d3d0a2d2d2d2d2d454e44205055424c4943204b45592d2d2d2d2d0a"
	MockMattPubKey       = "2d2d2d2d2d424547494e205055424c4943204b45592d2d2d2d2d0a4d466b77457759484b6f5a497a6a3043415159494b6f5a497a6a304441516344516741454c6c4f73766b366c74654d5564796f4b6936743077773377783753750a446a31616c33334a5956696a2b4a534f752f6838757a3030473234653662394d4c344852446c6359595358783968723576724e63784b567730773d3d0a2d2d2d2d2d454e44205055424c4943204b45592d2d2d2d2d0a"
	MockCoinVaultPubKey  = "CoinVaultsPubKeyCoinVaultsPubKeyCoinVaultsPubKeyCoinVaultsPubKeyCoinVaultsPubKeyCoinVaultsPubKeyCoinVaultsPubKeyCoinVaultsPubKeyCoinVaultsPubKey"
	MockJillPubKey       = "JillsPubKeyJillsPubKeyJillsPubKeyJillsPubKeyJillsPubKeyJillsPubKeyJillsPubKeyJillsPubKeyJillsPubKeyJillsPubKeyJillsPubKeyJillsPubKeyJillsPubKey"
	signatureDataString1 = "c2a91bcc9877aa9b408849630320dbc2d539c955748c51bcf96425240c22ff3b1a9c60f285cb2c76b83665b4c9f2caa4dc4f32486ede9177723635f466cd462e"
	signatureDataString2 = "23b68f2781bace70352f868e57dc28eb3fd81d0d2812adc1cb435139d4acb0c6195ad96e7596e6511e1d08312a46473f7cbd6017d7918011057e91231e8ffbbc"
	signatureDataString3 = "dd3d501e3444652272ecbe69b156149601a565c47fe4e988b98eb6923453cd6e70bd47728706ffb9bcd78351249e4f8e703ee9266632baf07f5040bd6a157c8e"
	signatureDataString4 = "a320167fbf827b3ad849b99cb7b12907ac5fa720202c9077a1e0d5a9f0917fb77db67d8ac80bf0e00a0f69c8a01da536a9c2a40bc6cb7aac3081d57bc13b36b1"
	signatureDataString5 = "dd7597b549729f5307c840b4643bd46a24568ee2c4ecb7e8e5aa80de754fbfd28107627d5d746fae0c3f1e6b760591a13f3237dc9b0e086378b104bb6030a067"
	signatureDataString6 = "f8b7955d38475f5b1c4e05ed658e8dfa7b2985d705c05f2a6ddb61327c876626c5aa2e6629f66b6076b4f2704929cef4edfef3bd0cc31f29dc799a7997448bb0"
)

// TRANSACTION REQUEST MESSAGES SIGNED
const (
	txRequestMessageSignedDataString1 = `
        {
            "txRequestMessage": {
                "sourceAddress": "` + MockFoundersPubKey + `",
                "destinations": [
                    {
                        "destinationAddress": "` + MockJeffPubKey + `",
                        "value": 80000
                    }
                ]
            },
            "signature": "` + signatureDataString1 + `"
        }`
	txRequestMessageSignedDataString2 = `
        {
            "txRequestMessage": {
                "sourceAddress": "` + MockJeffPubKey + `",
                "destinations": [
                    {
                        "destinationAddress": "` + MockMattPubKey + `",
                        "value": 50000
                    },
                    {
                        "destinationAddress": "` + MockCoinVaultPubKey + `",
                        "value": 500
                    }
                ]
            },
            "signature": "` + signatureDataString2 + `"
        }`
	txRequestMessageSignedDataString3 = `
        {
            "txRequestMessage": {
                "sourceAddress": "` + MockFoundersPubKey + `",
                "destinations": [
                    {
                        "destinationAddress": "` + MockMattPubKey + `",
                        "value": 250000
                    },
                    {
                        "destinationAddress": "` + MockJeffPubKey + `",
                        "value": 13000
                    }
                ]
            },
            "signature": "` + signatureDataString3 + `"
        }`
	txRequestMessageSignedDataString4 = `
        {
            "txRequestMessage": {
                "sourceAddress": "` + MockMattPubKey + `",
                "destinations": [
                    {
                        "destinationAddress": "` + MockJillPubKey + `",
                        "value": 35000
                    }
                ]
            },
            "signature": "` + signatureDataString4 + `"
        }`
	txRequestMessageSignedDataString5 = `
        {
            "txRequestMessage": {
                "sourceAddress": "` + MockMattPubKey + `",
                "destinations": [
                    {
                        "destinationAddress": "` + MockJeffPubKey + `",
                        "value": 15000
                    }
                ]
            },
            "signature": "` + signatureDataString5 + `"
        }`
	txRequestMessageSignedDataString6 = `
        {
            "txRequestMessage": {
                "sourceAddress": "` + MockJeffPubKey + `",
                "destinations": [
                    {
                        "destinationAddress": "` + MockJillPubKey + `",
                        "value": 33000
                    }
                ]
            },
            "signature": "` + signatureDataString6 + `"
        }`
	txRequestMessageSignedDataStringBad = `
        {
            "txRequestMessage": {
                "sourceAddress": "` + MockFoundersPubKey + `",
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
