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
)

func main() {
	// Alice wants to retrieve the keys of its company

	// Common Katena network information
	apiUrl := "https://nodes.test.katena.transchain.io/api/v1"

	// Alice Katena network information
	aliceCompanyBcid := "abcdef"

	// Create a Katena API helper
	transactor := client.NewTransactor(apiUrl, "", "", nil)

	// Retrieve the keys from Katena
	keys, err := transactor.RetrieveCompanyKeys(aliceCompanyBcid, 1, api.DefaultPerPageParam)
	if err != nil {
		panic(err)
	}

	for _, key := range keys {
		fmt.Println("KeyV1")
		fmt.Println(fmt.Sprintf("  Company bcid : %s", key.CompanyBcid))
		fmt.Println(fmt.Sprintf("  Public key   : %s", key.PublicKey))
		fmt.Println(fmt.Sprintf("  Is active    : %t", key.IsActive))
		fmt.Println(fmt.Sprintf("  Role         : %s", key.Role))

		fmt.Println()
	}
}
