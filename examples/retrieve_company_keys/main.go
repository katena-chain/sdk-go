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
	"github.com/katena-chain/sdk-go/examples/common"
)

func main() {
	// Alice wants to retrieve the keys of its company

	// Load default configuration
	settings := common.DefaultSettings()

	// Common Katena network information
	apiUrl := settings.ApiUrl

	// Alice Katena network information
	aliceCompanyBcId := settings.Company.BcId

	// Create a Katena API helper
	transactor := client.NewTransactor(apiUrl, "", nil)

	// Retrieve a list of keys for a company from the state
	keys, err := transactor.RetrieveCompanyKeys(aliceCompanyBcId, 1, settings.TxPerPage)
	if err != nil {
		panic(err)
	}

	fmt.Println("Keys list :")
	err = common.PrintlnJSON(keys)
	if err != nil {
		panic(err)
	}
}
