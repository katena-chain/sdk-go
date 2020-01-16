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
    // Alice wants to certify raw off-chain information

    // Common Katena network information
    apiUrl := "https://api.test.katena.transchain.io/api/v1"
    chainId := "katena-chain-test"

    // Alice Katena network information
    aliceSignPrivateKeyBase64 := "7C67DeoLnhI6jvsp3eMksU2Z6uzj8sqZbpgwZqfIyuCZbfoPcitCiCsSp2EzCfkY52Mx58xDOyQLb1OhC7cL5A=="
    aliceCompanyChainId := "abcdef"
    aliceSignPrivateKey, err := ed25519.NewPrivateKeyFromBase64(aliceSignPrivateKeyBase64)
    if err != nil {
        panic(err)
    }

    // Create a Katena API helper
    transactor := client.NewTransactor(apiUrl, chainId, aliceCompanyChainId, &aliceSignPrivateKey)

    keyRevokeUuid := "f4866d77-a083-4ecb-9ec0-a2a152b6ce14"
    publicKeyBase64 := "miJmzMTa68aFFIle91bCo8qeQkanmgu2suA4E7dL2Fs="
    publicKey, err := ed25519.NewPublicKeyFromBase64(publicKeyBase64)
    if err != nil {
        panic(err)
    }

    // Send a version 1 of a certificate raw on Katena
    txStatus, err := transactor.SendKeyRevokeV1(keyRevokeUuid, publicKey)
    if err != nil {
        panic(err)
    }
    fmt.Println("Transaction Data")
    fmt.Println(fmt.Sprintf("  Tx Uuid        : %s", keyRevokeUuid))
    fmt.Println(fmt.Sprintf("  Public Key     : %s", publicKey.String()))

    fmt.Println("Transaction status")
    fmt.Println(fmt.Sprintf("  Code    : %d", txStatus.Code))
    fmt.Println(fmt.Sprintf("  Message : %s", txStatus.Message))

}
