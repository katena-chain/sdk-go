/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package common

type Settings struct {
	// API url to dialogue with a Katena network
	ApiUrl string

	// Katena network id
	ChainId string

	// Number of transactions the API should return
	TxPerPage int

	// Dummy company committed on chain
	Company *Company

	// Sample transaction ids used in examples
	// If one id is already used on the Katena test network, feel free to change these values in DefaultSettings below
	CertificateId string
	SecretId      string
	KeyId         string

	// Off chain samples data to do off chain operations
	OffChain *OffChain
}

type Company struct {
	// Unique company identifier on a Katena network
	BcId string

	// Dummy users with their keys to sign transactions
	Ed25519Keys map[string]*Key
}

type OffChain struct {
	// Dummy users with their keys to sign off-chain data
	Ed25519Keys map[string]*KeyPair

	// Dummy users with their keys to seal/open nacl boxes to share secret information
	X25519Keys map[string]*KeyPair
}

type Key struct {
	Id            string
	PrivateKeyStr string
}

type KeyPair struct {
	PrivateKeyStr string
	PublicKeyStr  string
}

func DefaultSettings() *Settings {
	return &Settings{
		ApiUrl:    "https://nodes.test.katena.transchain.io/api/v1",
		ChainId:   "katena-chain-test",
		TxPerPage: 10,
		Company: &Company{
			BcId: "abcdef",
			Ed25519Keys: map[string]*Key{
				"alice": {
					Id:            "36b72ca9-fd58-44aa-b90d-5a855276ff82",
					PrivateKeyStr: "7C67DeoLnhI6jvsp3eMksU2Z6uzj8sqZbpgwZqfIyuCZbfoPcitCiCsSp2EzCfkY52Mx58xDOyQLb1OhC7cL5A==",
				},
				"bob": {
					Id:            "7cf17643-5567-4dfa-9b0c-9cd19c45177a",
					PrivateKeyStr: "3awdq5HUZ2fgV2fM6sbV1yJKIvuTV2OZ5AMfes4ftHUiOpqsicnv+67vLfKLwWR/Bh/hNbJaq6fziXoh+oqxRQ==",
				},
				"carla": {
					Id:            "236f8028-bb87-4c19-b6e0-cbcaea35e764",
					PrivateKeyStr: "p2T1gRu2HHdhcsTVEk6VwpJRkLahvnLsi9miSS1Yg4PSk6jrTRFvtoPzi2z6yn+Ul9+niTHBUvbskbQ2TkDxmQ==",
				},
			},
		},
		CertificateId: "ce492f92-a529-40c1-91e9-2af71e74ebea",
		SecretId:      "3b1cfd5f-d0fe-478c-ba30-17817e29611e",
		KeyId:         "9941bc28-4033-4d5a-a337-76b640223de2",
		OffChain: &OffChain{
			Ed25519Keys: map[string]*KeyPair{
				"david": {
					PrivateKeyStr: "aGya1W2C2bfu1bMA+wJ8kbpZePjKprv4t93EhX+durqOksFaT9pC0054jFeKYFyGzi+1gCp1NZAeCsG/yQEJWA==",
					PublicKeyStr:  "jpLBWk/aQtNOeIxXimBchs4vtYAqdTWQHgrBv8kBCVg=",
				},
			},
			X25519Keys: map[string]*KeyPair{
				"alice": {
					PrivateKeyStr: "nyCzhimWnTQifh6ucXLuJwOz3RgiBpo33LcX1NjMAsP1ZkQcdlDq64lTwxaDx0lq6LCQAUeYywyMUtfsvTUEeQ==",
					PublicKeyStr:  "9WZEHHZQ6uuJU8MWg8dJauiwkAFHmMsMjFLX7L01BHk=",
				},
				"bob": {
					PrivateKeyStr: "quGBP8awD/J3hjSvwGD/sZRcMDks8DPz9Vw0HD4+zecqJP0ojBoc4wQtyq08ywxUksTkdz0/rQNkOsEZBwqWTw==",
					PublicKeyStr:  "KiT9KIwaHOMELcqtPMsMVJLE5Hc9P60DZDrBGQcKlk8=",
				},
			},
		},
	}
}
