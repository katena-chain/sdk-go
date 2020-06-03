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
	entityCommon "github.com/katena-chain/sdk-go/entity/common"
	"github.com/katena-chain/sdk-go/examples/common"
)

func main() {
	// Alice wants to certify raw off-chain information

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

	// Create a transactor instance to dialogue with a Katena API
	txSigner := entity.NewTxSigner(entityCommon.ConcatFqId(aliceCompanyBcId, aliceSignPrivateKeyId), &aliceSignPrivateKey)
	transactor := client.NewTransactor(apiUrl, chainId, txSigner)

	// Off-chain information Alice wants to send
	certificateId := settings.CertificateId
	dataRawSignature := []byte("off_chain_data_raw_signature_from_go")

	// Send a version 1 of a certificate raw on Katena
	txResult, err := transactor.SendCertificateRawV1Tx(certificateId, dataRawSignature)
	if err != nil {
		panic(err)
	}

	fmt.Println("Result :")
	err = common.PrintlnJSON(txResult)
	if err != nil {
		panic(err)
	}
}
