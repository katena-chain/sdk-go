/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package common

import (
	"encoding/json"
	"fmt"

	kcJson "github.com/transchain/sdk-go/json"

	"github.com/katena-chain/sdk-go/entity"
)

func PrintlnJSON(data interface{}) error {
	if txData, ok := data.(entity.TxData); ok {
		data = kcJson.MarshalWrapper{
			Value: txData,
			Type:  txData.GetType(),
		}
	}
	encodedData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("%s\n", encodedData))
	return nil
}
