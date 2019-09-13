/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package main

import (
    "encoding/base64"
    "fmt"

    "github.com/transchain/sdk-go/crypto/nacl"

    "github.com/katena-chain/sdk-go/client"
    "github.com/katena-chain/sdk-go/entity/certify"
)

func main() {
    // Bob wants to read a nacl box secret from Alice to decrypt an off-chain data

    // Common Katena network information
    apiUrl := "https://api.test.katena.transchain.io/api/v1"

    // Alice Katena network information
    aliceCompanyChainId := "abcdef"

    // Create a Katena API helper
    transactor := client.NewTransactor(apiUrl, "", "", nil)

    // Nacl box information
    bobCryptPrivateKeyBase64 := "quGBP8awD/J3hjSvwGD/sZRcMDks8DPz9Vw0HD4+zecqJP0ojBoc4wQtyq08ywxUksTkdz0/rQNkOsEZBwqWTw=="
    bobCryptPrivateKey, err := nacl.NewPrivateKeyFromBase64(bobCryptPrivateKeyBase64)
    if err != nil {
        panic(err)
    }

    // Secret uuid Bob wants to retrieve
    secretUuid := "2075c941-6876-405b-87d5-13791c0dc53a"

    // Retrieve version 1 of secrets from Katena
    txWrappers, err := transactor.RetrieveSecrets(aliceCompanyChainId, secretUuid)
    if err != nil {
        panic(err)
    }

    for _, txWrapper := range txWrappers.Txs {
        txData := txWrapper.Tx.Data.(*certify.SecretNaclBoxV1)

        fmt.Println("Transaction status")
        fmt.Println(fmt.Sprintf("  Code    : %d", txWrapper.Status.Code))
        fmt.Println(fmt.Sprintf("  Message : %s", txWrapper.Status.Message))

        fmt.Println("SecretNaclBoxV1")
        fmt.Println(fmt.Sprintf("  Id                : %s", txData.Id))
        fmt.Println(fmt.Sprintf("  Data sender       : %s", txData.Sender))
        fmt.Println(fmt.Sprintf("  Data nonce        : %s", txData.Nonce))
        fmt.Println(fmt.Sprintf("  Data content      : %s", base64.StdEncoding.EncodeToString(txData.Content)))

        // Bob will use its private key and the sender's public key (needs to be Alice's) to decrypt a message
        decryptedContent, ok := bobCryptPrivateKey.Open(txData.Content, txData.Nonce, txData.Sender)
        if !ok {
            decryptedContent = []byte("Unable to decrypt")
        }

        fmt.Println(fmt.Sprintf("  Decrypted content : %s", decryptedContent))
        fmt.Println()
    }
}
