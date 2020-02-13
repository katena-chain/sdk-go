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
)

func main() {
    // Alice wants to certify an ed25519 signature of an off-chain data

    // Common Katena network information
    apiUrl := "https://api.test.katena.transchain.io/api/v1"
    chainId := "katena-chain-test"

    // Alice Katena network information
    aliceSignPrivateKeyBase64 := "7C67DeoLnhI6jvsp3eMksU2Z6uzj8sqZbpgwZqfIyuCZbfoPcitCiCsSp2EzCfkY52Mx58xDOyQLb1OhC7cL5A=="
    aliceCompanyChainId := "abcdef"
    aliceSignPrivateKey := ed25519.NewPrivateKeyFromBase64(aliceSignPrivateKeyBase64)

    // Create a Katena API helper
    transactor := client.NewTransactor(apiUrl, chainId, aliceCompanyChainId, &aliceSignPrivateKey)

    // Off-chain information Alice wants to send
    certificateUuid := "2075c941-6876-405b-87d5-13791c0dc53a"
    dataSignature := aliceSignPrivateKey.Sign([]byte("off_chain_data_to_sign_from_go"))

    // Send a version 1 of a certificate ed25519 on Katena
    txStatus, err := transactor.SendCertificateEd25519V1(certificateUuid, aliceSignPrivateKey.GetPublicKey(), dataSignature)
    if err != nil {
        panic(err)
    }

    fmt.Println("Transaction status")
    fmt.Println(fmt.Sprintf("  Code    : %d", txStatus.Code))
    fmt.Println(fmt.Sprintf("  Message : %s", txStatus.Message))
}
