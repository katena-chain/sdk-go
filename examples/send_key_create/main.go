/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package main

import (
	"fmt"

	"github.com/transchain/sdk-go/crypto/ed25519"

	"github.com/katena-chain/sdk-go/client"
	"github.com/katena-chain/sdk-go/entity"
	"github.com/katena-chain/sdk-go/entity/account"
	"github.com/katena-chain/sdk-go/examples/common"
)

func main() {
	// Alice wants to create a key for its company

	// Load default configuration
	settings := common.DefaultSettings()

	// Common Katena network information
	apiUrl := settings.ApiUrl
	chainId := settings.ChainId

	// Alice Katena network information
	aliceCompanyBcId := settings.Company.BcId
	aliceSignKeyInfo := settings.Company.Ed25519Keys["alice"]
	aliceSignPrivateKey := ed25519.NewPrivateKeyFromBase64(aliceSignKeyInfo.PrivateKeyStr)
	aliceSignPrivateKeyId := aliceSignKeyInfo.Id

	// Create a Katena API helper
	txSigner := entity.NewTxSigner(aliceSignPrivateKeyId, &aliceSignPrivateKey, aliceCompanyBcId)
	transactor := client.NewTransactor(apiUrl, chainId, txSigner)

	// Information Alice wants to send
	keyId := settings.KeyId
	newPrivateKey := ed25519.GenerateNewPrivateKey()
	newPublicKey := newPrivateKey.GetPublicKey()

	fmt.Println("New key info")
	fmt.Println(fmt.Sprintf("  Private key : %s", newPrivateKey.String()))
	fmt.Println(fmt.Sprintf("  Public key  : %s", newPublicKey.String()))

	// Choose role between account.DefaultRoleId or account.CompanyAdminRoleId
	role := account.DefaultRoleId

	// Send a version 1 of a key create on Katena
	txResult, err := transactor.SendKeyCreateV1Tx(keyId, newPublicKey, role)
	if err != nil {
		panic(err)
	}

	fmt.Println("Result :")
	err = common.PrintlnJSON(txResult)
	if err != nil {
		panic(err)
	}
}
