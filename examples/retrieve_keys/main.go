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
	// Alice wants to retrieve txs related to a key

	// Load default configuration
	settings := common.DefaultSettings()

	// Common Katena network information
	apiUrl := settings.ApiUrl

	// Alice Katena network information
	aliceCompanyBcId := settings.Company.BcId

	// Create a Katena API helper
	transactor := client.NewTransactor(apiUrl, "", nil)

	// Key id Alice wants to retrieve
	keyId := settings.KeyId

	// Retrieve txs related to the key fqid
	txResults, err := transactor.RetrieveKeyTxs(aliceCompanyBcId, keyId, 1, settings.TxPerPage)
	if err != nil {
		panic(err)
	}

	fmt.Println("Tx list :")
	err = common.PrintlnJSON(txResults)
	if err != nil {
		panic(err)
	}

	// Retrieve the last tx related to the key fqid
	txResult, err := transactor.RetrieveLastKeyTx(aliceCompanyBcId, keyId)
	if err != nil {
		panic(err)
	}

	fmt.Println("Last Tx :")
	err = common.PrintlnJSON(txResult)
	if err != nil {
		panic(err)
	}

	// Retrieve the last state of a key with that fqid
	key, err := transactor.RetrieveKey(aliceCompanyBcId, keyId)
	if err != nil {
		panic(err)
	}

	fmt.Println("Key :")
	err = common.PrintlnJSON(key)
	if err != nil {
		panic(err)
	}
}
