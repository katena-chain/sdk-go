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
	"github.com/katena-chain/sdk-go/entity/account"
)

func main() {
	// Alice wants to create a key for its company

	// Common Katena network information
	apiUrl := "https://nodes.test.katena.transchain.io/api/v1"
	chainID := "katena-chain-test"

	// Alice Katena network information
	aliceSignPrivateKeyBase64 := "7C67DeoLnhI6jvsp3eMksU2Z6uzj8sqZbpgwZqfIyuCZbfoPcitCiCsSp2EzCfkY52Mx58xDOyQLb1OhC7cL5A=="
	aliceCompanyBcid := "abcdef"
	aliceSignPrivateKey := ed25519.NewPrivateKeyFromBase64(aliceSignPrivateKeyBase64)

	// Create a Katena API helper
	transactor := client.NewTransactor(apiUrl, chainID, aliceCompanyBcid, &aliceSignPrivateKey)

	// Information Alice wants to send
	keyCreateUuid := "2075c941-6876-405b-87d5-13791c0dc53a"
	newPublicKeyBase64 := "gaKih+STp93wMuKmw5tF5NlQvOlrGsahpSmpr/KwOiw="
	newPublicKey := ed25519.NewPublicKeyFromBase64(newPublicKeyBase64)
	// Choose role between account.DefaultRoleID or account.CompanyAdminRoleID
	role := account.DefaultRoleID

	// Send a version 1 of a key create on Katena
	txStatus, err := transactor.SendKeyCreateV1(keyCreateUuid, newPublicKey, role)
	if err != nil {
		panic(err)
	}

	fmt.Println("Transaction status")
	fmt.Println(fmt.Sprintf("  Code    : %d", txStatus.Code))
	fmt.Println(fmt.Sprintf("  Message : %s", txStatus.Message))
}
