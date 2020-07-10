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
	// Alice wants to send a nacl box secret to Bob to encrypt an off-chain data

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

	// Nacl box information
	bobCryptKeyInfo := settings.OffChain.X25519Keys["bob"]
	bobCryptPublicKey := entityCommon.CreatePublicKeyX25519FromBase64(bobCryptKeyInfo.PublicKeyStr)
	aliceCryptKeyInfo := settings.OffChain.X25519Keys["alice"]
	aliceCryptPrivateKey := entityCommon.CreatePrivateKeyX25519FromBase64(aliceCryptKeyInfo.PrivateKeyStr)

	// Create a Katena API helper
	txSigner := entity.NewTxSigner(entityCommon.ConcatFqId(aliceCompanyBcId, aliceSignPrivateKeyId), &aliceSignPrivateKey)
	transactor := client.NewTransactor(apiUrl, chainId, txSigner)

	// Off-chain information Alice wants to send
	secretId := settings.SecretId
	content := []byte("off_chain_secret_to_crypt_from_go")

	// Alice will use its private key and Bob's public key to encrypt a message
	encryptedMessage, boxNonce, err := aliceCryptPrivateKey.Seal(content, bobCryptPublicKey)
	if err != nil {
		panic(err)
	}

	// Send a version 1 of a secret nacl box on Katena
	txResult, err := transactor.SendSecretNaclBoxV1Tx(secretId, aliceCryptPrivateKey.GetPublicKey(), boxNonce, encryptedMessage)
	if err != nil {
		panic(err)
	}

	fmt.Println("Result :")
	err = common.PrintlnJSON(txResult)
	if err != nil {
		panic(err)
	}
}
