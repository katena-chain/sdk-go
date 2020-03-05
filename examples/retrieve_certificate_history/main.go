/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package main

import (
	"fmt"

	"github.com/transchain/sdk-go/api"

	"github.com/katena-chain/sdk-go/client"
	"github.com/katena-chain/sdk-go/entity/certify"
)

func main() {
	// Alice wants to retrieve a certificate history

	// Common Katena network information
	apiUrl := "https://nodes.test.katena.transchain.io/api/v1"

	// Alice Katena network information
	aliceCompanyBcid := "abcdef"

	// Create a Katena API helper
	transactor := client.NewTransactor(apiUrl, "", "", nil)

	// Certificate uuid Alice wants to retrieve
	certificateUuid := "2075c941-6876-405b-87d5-13791c0dc53a"

	// Retrieve a version 1 of a certificate history from Katena
	txWrappers, err := transactor.RetrieveCertificatesHistory(aliceCompanyBcid, certificateUuid, 1, api.DefaultPerPageParam)
	if err != nil {
		panic(err)
	}

	for _, txWrapper := range txWrappers.Txs {
		txData := txWrapper.Tx.Data

		fmt.Println("Transaction status")
		fmt.Println(fmt.Sprintf("  Code    : %d", txWrapper.Status.Code))
		fmt.Println(fmt.Sprintf("  Message : %s", txWrapper.Status.Message))

		switch txData := txData.(type) {
		case *certify.CertificateRawV1:
			fmt.Println("CertificateRawV1")
			fmt.Println(fmt.Sprintf("  Id    : %s", txData.Id))
			fmt.Println(fmt.Sprintf("  Value : %s", txData.Value))
			break
		case *certify.CertificateEd25519V1:
			fmt.Println("CertificateEd25519V1")
			fmt.Println(fmt.Sprintf("  Id             : %s", txData.Id))
			fmt.Println(fmt.Sprintf("  Data signer    : %s", txData.Signer))
			fmt.Println(fmt.Sprintf("  Data signature : %s", txData.Signature))
			break
		default:
			panic("Unexpected txData type")
		}

		fmt.Println()
	}
}
