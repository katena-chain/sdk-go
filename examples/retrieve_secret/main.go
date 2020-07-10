/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package main

import (
	"fmt"

	"github.com/katena-chain/sdk-go/client"
	"github.com/katena-chain/sdk-go/entity/certify"
	"github.com/katena-chain/sdk-go/examples/common"
)

func main() {
	// Bob wants to read a nacl box secret from Alice to decrypt an off-chain data

	// Load default configuration
	settings := common.DefaultSettings()

	// Common Katena network information
	apiUrl := settings.ApiUrl

	// Alice Katena network information
	aliceCompanyBcId := settings.Company.BcId

	// Create a Katena API helper
	transactor := client.NewTransactor(apiUrl, "", nil)

	// Nacl box information
	bobCryptKeyInfo := settings.OffChain.X25519Keys["bob"]
	bobCryptPrivateKey := common.CreatePrivateKeyX25519FromBase64(bobCryptKeyInfo.PrivateKeyStr)

	// Secret id Bob wants to retrieve
	secretId := settings.SecretId

	// Retrieve txs related to the secret fqid
	txResults, err := transactor.RetrieveSecretTxs(aliceCompanyBcId, secretId, 1, settings.TxPerPage)
	if err != nil {
		panic(err)
	}

	fmt.Println("Tx list :")
	err = common.PrintlnJSON(txResults)
	if err != nil {
		panic(err)
	}

	// Retrieve the last tx related to the secret fqid
	txResult, err := transactor.RetrieveLastSecretTx(aliceCompanyBcId, secretId)
	if err != nil {
		panic(err)
	}

	fmt.Println("Last Tx :")
	err = common.PrintlnJSON(txResult)
	if err != nil {
		panic(err)
	}

	// Retrieve the last state of a secret with that fqid
	secret, err := transactor.RetrieveSecret(aliceCompanyBcId, secretId)
	if err != nil {
		panic(err)
	}

	fmt.Println("Secret :")
	err = common.PrintlnJSON(secret)
	if err != nil {
		panic(err)
	}

	secretNaclBox := secret.(*certify.SecretNaclBoxV1)
	// Bob will use its private key and the sender's public key (needs to be Alice's) to decrypt a message
	decryptedContent, ok := bobCryptPrivateKey.Open(secretNaclBox.Content, secretNaclBox.Nonce, secretNaclBox.Sender)
	if !ok {
		decryptedContent = []byte("Unable to decrypt")
	}
	fmt.Println(fmt.Sprintf("Decrypted content : %s", decryptedContent))
}
