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
	"github.com/katena-chain/sdk-go/entity"
	entityCommon "github.com/katena-chain/sdk-go/entity/common"
	"github.com/katena-chain/sdk-go/examples/common"
)

func main() {
	// Alice wants to revoke a key for its company

	// Load default configuration
	settings := common.DefaultSettings()

	// Common Katena network information
	apiUrl := settings.ApiUrl
	chainId := settings.ChainId

	// Alice Katena network information
	aliceCompanyBcId := settings.Company.BcId
	aliceSignKeyInfo := settings.Company.Ed25519Keys["alice"]
	aliceSignPrivateKey := entityCommon.CreatePrivateKeyEd25519FromBase64(aliceSignKeyInfo.PrivateKeyStr)
	aliceSignPrivateKeyId := aliceSignKeyInfo.Id

	// Create a Katena API helper
	txSigner := entity.NewTxSigner(entityCommon.ConcatFqId(aliceCompanyBcId, aliceSignPrivateKeyId), &aliceSignPrivateKey)
	transactor := client.NewTransactor(apiUrl, chainId, txSigner)

	// Information Alice wants to send
	keyId := settings.KeyId

	// Send a version 1 of a key revoke on Katena
	txResult, err := transactor.SendKeyRevokeV1Tx(keyId)
	if err != nil {
		panic(err)
	}

	fmt.Println("Result :")
	err = common.PrintlnJSON(txResult)
	if err != nil {
		panic(err)
	}
}
