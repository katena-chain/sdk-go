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
    // Alice wants to revoke a key for its company

    // Common Katena network information
    apiUrl := "https://api.test.katena.transchain.io/api/v1"
    chainId := "katena-chain-test"

    // Alice Katena network information
    aliceSignPrivateKeyBase64 := "7C67DeoLnhI6jvsp3eMksU2Z6uzj8sqZbpgwZqfIyuCZbfoPcitCiCsSp2EzCfkY52Mx58xDOyQLb1OhC7cL5A=="
    aliceCompanyChainId := "abcdef"
    aliceSignPrivateKey := ed25519.NewPrivateKeyFromBase64(aliceSignPrivateKeyBase64)

    // Create a Katena API helper
    transactor := client.NewTransactor(apiUrl, chainId, aliceCompanyChainId, &aliceSignPrivateKey)

    // Information Alice wants to send
    keyRevokeUuid := "2075c941-6876-405b-87d5-13791c0dc53a"
    publicKeyBase64 := "gaKih+STp93wMuKmw5tF5NlQvOlrGsahpSmpr/KwOiw="
    publicKey := ed25519.NewPublicKeyFromBase64(publicKeyBase64)

    // Send a version 1 of a key create on Katena
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
