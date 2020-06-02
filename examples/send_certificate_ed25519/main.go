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
	"github.com/katena-chain/sdk-go/examples/common"
)

func main() {
	// Alice wants to certify an ed25519 signature of an off-chain data

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

	// Off-chain information Alice wants to send
	certificateId := settings.CertificateId
	dataSignature := aliceSignPrivateKey.Sign([]byte("off_chain_data_to_sign_from_go"))

	// Send a version 1 of a certificate ed25519 on Katena
	txResult, err := transactor.SendCertificateEd25519V1Tx(certificateId, aliceSignPrivateKey.GetPublicKey(), dataSignature)
	if err != nil {
		panic(err)
	}

	fmt.Println("Result :")
	err = common.PrintlnJSON(txResult)
	if err != nil {
		panic(err)
	}
}
