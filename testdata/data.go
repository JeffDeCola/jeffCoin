// jeffCoin TESTDATA data.go

package testdata

const (
	foundersPubKey       = "2d2d2d2d2d424547494e205055424c4943204b45592d2d2d2d2d0a4d466b77457759484b6f5a497a6a3043415159494b6f5a497a6a3044415163445167414535727549786d5146754548414c747437517778636247446b3863705a0a49684971512b4f4f336a492b31723177734a33686c693133414763686f66523639574b4c354b394a34574278696e2f37736c6f6f7030665268513d3d0a2d2d2d2d2d454e44205055424c4943204b45592d2d2d2d2d0a"
	jeffPubKey           = "2d2d2d2d2d424547494e205055424c4943204b45592d2d2d2d2d0a4d466b77457759484b6f5a497a6a3043415159494b6f5a497a6a304441516344516741457668696a41504c7844367071746a4539356c5272536748782b2b33350a634230674443395746445a703236325276337569464770414b625143422b796d577a5a665a6c66696a6f4e4942616a6f306f414e7873515a73413d3d0a2d2d2d2d2d454e44205055424c4943204b45592d2d2d2d2d0a"
	mattPubKey           = "2d2d2d2d2d424547494e205055424c4943204b45592d2d2d2d2d0a4d466b77457759484b6f5a497a6a3043415159494b6f5a497a6a30444151634451674145315772756f54335a4b77466d676e45426851353163354d6a6f5950430a5a61333845476873793675435739576a3771514d6a7874512b6f44534d3648536631367363706e3367394a2b63566a2b6f694b337179736e70413d3d0a2d2d2d2d2d454e44205055424c4943204b45592d2d2d2d2d0a"
	coinVaultPubKey      = "CoinVaultsPubKey"
	jillPubKey           = "JillsPubKey"
	signatureDataString1 = "1516a84bf05592225aff22d7b8b2b98de5f763bf2d62f291860f864234efe8de5125e7c3e3f987baeae80a790b950892c747eeabd9ea690334ecc681f98d9d48"
	signatureDataString2 = "bb873b906093bef32eedd37e6b37a591f52373eb77843cbf006d65cd9b2cfa494ff8786b120ff1c3d6f8645fbbce1da3dab1dbb6a9e9ab7c89adc22e8456334d"
	signatureDataString3 = "b607df0a311676afc1a271ed1266f5f0212f950e7cf2ffe4f24955345f025d4fe3dcfd34b66f87310924c613bf944b14914e08efeeb46b988f627e9f1f9b5c9e"
	signatureDataString4 = "a30b74e73ad8b159b5c772f5e95c5c40d96a48c76a52714d3b369b5eb91e77c082f95647ac66f240c3baf41bc58729bde5febbccc9f2aed65794b5558c921765"
	signatureDataString5 = "ef7ca38a6696f9b19cbf06c6ea49d37aae8100819a111aefc785601b6e9434dbe2d4a81ebbf506d295c82cb7d7449b290f986f11f5ded1ef0b2b70d9d9b5b721"
	signatureDataString6 = "bc8bab302108f55c6c61f01ade294c579e50e9929af653f3456a4376046adcdfac12403369cbdd2dbda80dc23d0bd49267d2eea0efc6313709661c14fcdef291"
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
