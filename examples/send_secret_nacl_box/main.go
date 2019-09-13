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
    "github.com/transchain/sdk-go/crypto/nacl"

    "github.com/katena-chain/sdk-go/client"
)

func main() {
    // Alice wants to send a nacl box secret to Bob to encrypt an off-chain data

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

    // Nacl box information
    aliceCryptPrivateKeyBase64 := "nyCzhimWnTQifh6ucXLuJwOz3RgiBpo33LcX1NjMAsP1ZkQcdlDq64lTwxaDx0lq6LCQAUeYywyMUtfsvTUEeQ=="
    aliceCryptPrivateKey, err := nacl.NewPrivateKeyFromBase64(aliceCryptPrivateKeyBase64)
    if err != nil {
        panic(err)
    }
    bobCryptPublicKeyBase64 := "KiT9KIwaHOMELcqtPMsMVJLE5Hc9P60DZDrBGQcKlk8="
    bobCryptPublicKey, err := nacl.NewPublicKeyFromBase64(bobCryptPublicKeyBase64)
    if err != nil {
        panic(err)
    }

    // Create a Katena API helper
    transactor := client.NewTransactor(apiUrl, chainId, aliceCompanyChainId, &aliceSignPrivateKey)

    // Off-chain information Alice wants to send
    secretUuid := "2075c941-6876-405b-87d5-13791c0dc53a"
    content := []byte("off_chain_secret_to_crypt_from_go")

    // Alice will use its private key and Bob's public key to encrypt a message
    encryptedMessage, boxNonce, err := aliceCryptPrivateKey.Seal(content, bobCryptPublicKey)
    if err != nil {
        panic(err)
    }

    // Send a version 1 of a secret nacl box on Katena
    txStatus, err := transactor.SendSecretNaclBoxV1(secretUuid, aliceCryptPrivateKey.GetPublicKey(), boxNonce, encryptedMessage)
    if err != nil {
        panic(err)
    }

    fmt.Println("Transaction status")
    fmt.Println(fmt.Sprintf("  Code    : %d", txStatus.Code))
    fmt.Println(fmt.Sprintf("  Message : %s", txStatus.Message))

}
