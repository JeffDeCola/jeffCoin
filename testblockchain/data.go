// jeffCoin TESTDATA data.go

package testblockchain

// Public Keys (address) and mock signatures
const (
	MockFoundersPubKey   = "2d2d2d2d2d424547494e205055424c4943204b45592d2d2d2d2d0a4d466b77457759484b6f5a497a6a3043415159494b6f5a497a6a3044415163445167414535727549786d5146754548414c747437517778636247446b3863705a0a49684971512b4f4f336a492b31723177734a33686c693133414763686f66523639574b4c354b394a34574278696e2f37736c6f6f7030665268513d3d0a2d2d2d2d2d454e44205055424c4943204b45592d2d2d2d2d0a"
	MockJeffPubKey       = "2d2d2d2d2d424547494e205055424c4943204b45592d2d2d2d2d0a4d466b77457759484b6f5a497a6a3043415159494b6f5a497a6a304441516344516741457668696a41504c7844367071746a4539356c5272536748782b2b33350a634230674443395746445a703236325276337569464770414b625143422b796d577a5a665a6c66696a6f4e4942616a6f306f414e7873515a73413d3d0a2d2d2d2d2d454e44205055424c4943204b45592d2d2d2d2d0a"
	MockMattPubKey       = "2d2d2d2d2d424547494e205055424c4943204b45592d2d2d2d2d0a4d466b77457759484b6f5a497a6a3043415159494b6f5a497a6a30444151634451674145315772756f54335a4b77466d676e45426851353163354d6a6f5950430a5a61333845476873793675435739576a3771514d6a7874512b6f44534d3648536631367363706e3367394a2b63566a2b6f694b337179736e70413d3d0a2d2d2d2d2d454e44205055424c4943204b45592d2d2d2d2d0a"
	MockCoinVaultPubKey  = "CoinVaultsPubKeyCoinVaultsPubKeyCoinVaultsPubKeyCoinVaultsPubKeyCoinVaultsPubKeyCoinVaultsPubKeyCoinVaultsPubKeyCoinVaultsPubKeyCoinVaultsPubKey"
	MockJillPubKey       = "JillsPubKeyJillsPubKeyJillsPubKeyJillsPubKeyJillsPubKeyJillsPubKeyJillsPubKeyJillsPubKeyJillsPubKeyJillsPubKeyJillsPubKeyJillsPubKeyJillsPubKey"
	signatureDataString1 = "1516a84bf05592225aff22d7b8b2b98de5f763bf2d62f291860f864234efe8de5125e7c3e3f987baeae80a790b950892c747eeabd9ea690334ecc681f98d9d48"
	signatureDataString2 = "917754ba1c3a7e111357231cf8837c32670d4369ef713dffcc0775fcf0213daf1324e148fb8ed1467853133ed5b82a044ad38a1b6e84470ad2a652dd9b8be9cb"
	signatureDataString3 = "b607df0a311676afc1a271ed1266f5f0212f950e7cf2ffe4f24955345f025d4fe3dcfd34b66f87310924c613bf944b14914e08efeeb46b988f627e9f1f9b5c9e"
	signatureDataString4 = "e9a972e8d04ae557e1ef89173e054fadc2207993742fded193789720b88e0b2a0b7b01432ce5994714eb27310e057ce39c03e91d9fe406258e91fab0256782b1"
	signatureDataString5 = "ef7ca38a6696f9b19cbf06c6ea49d37aae8100819a111aefc785601b6e9434dbe2d4a81ebbf506d295c82cb7d7449b290f986f11f5ded1ef0b2b70d9d9b5b721"
	signatureDataString6 = "3b2a3d2dd1edd7864a46e876765fb99c328936da62c5089b91f581e7450d093b6e20dff50d4776d8a0a2d43bd3e4a5c61a67cb76b8f276552ac6d445726b31b0"
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
