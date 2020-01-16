/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package main

import (
    "fmt"

    "github.com/katena-chain/sdk-go/api"
    "github.com/katena-chain/sdk-go/client"
    "github.com/katena-chain/sdk-go/entity/account"
)

func main() {
    // Alice wants to retrieve a certificate

    // Common Katena network information
    apiUrl := "https://api.test.katena.transchain.io/api/v1"

    // Alice Katena network information
    aliceCompanyChainId := "abcdef"

    // Create a Katena API helper
    transactor := client.NewTransactor(apiUrl, "", "", nil)

    // KeyRevoke uuid Alice wants to retrieve
    keyRevokeUuid := "f4866d77-a083-4ecb-9ec0-a2a152b6ce14"

    // Retrieve a version 1 of a KeyRevoke from Katena
    txWrappers, err := transactor.RetrieveKeysRevoke(aliceCompanyChainId, keyRevokeUuid, 1, api.DefaultPerPageParam)
    if err != nil {
        panic(err)
    }
    for _, txWrapper := range txWrappers.Txs {
        txData := txWrapper.Tx.Data

        fmt.Println("Transaction status")
        fmt.Println(fmt.Sprintf("  Code    : %d", txWrapper.Status.Code))
        fmt.Println(fmt.Sprintf("  Message : %s", txWrapper.Status.Message))
        fmt.Println(fmt.Sprintf("  NonceTime : %s", txWrapper.Tx.NonceTime))

        switch txData := txData.(type) {
        case *account.KeyRevokeV1:
            fmt.Println("KeyRevokeV1")
            fmt.Println(fmt.Sprintf("  Id : %s", txData.Id))
            fmt.Println(fmt.Sprintf("  PublicKey : %s", txData.PublicKey))
            break
        default:
            panic("Unexpected txData type")
        }

        fmt.Println()
    }
}
