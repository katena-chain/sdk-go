/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package entity

import (
	"fmt"
	"strings"

	kcJson "github.com/transchain/sdk-go/json"
)

// TxData interface defines the methods a concrete TxData must implement.
type TxData interface {
	GetId() string
	GetNamespace() string
	GetCategory() string
	GetType() string
}

// txDataState wraps a TxData and additional values in order to define the unique state to be signed.
type txDataState struct {
	ChainID   string                `json:"chain_id"`
	NonceTime Time                  `json:"nonce_time"`
	Data      kcJson.MarshalWrapper `json:"data"`
}

// GetTxDataStateBytes returns the sorted and marshaled json representation of a TxData ready to be signed.
func GetTxDataStateBytes(chainID string, nonceTime Time, txData TxData) []byte {
	data := txDataState{
		ChainID:   chainID,
		NonceTime: nonceTime,
		Data: kcJson.MarshalWrapper{
			Type:  txData.GetType(),
			Value: txData,
		},
	}
	return kcJson.MustMarshalAndSortJSON(data)
}

// SplitTxid separates a txid field into a company bcid and a uuid.
func SplitTxid(id string) (string, string) {
	split := strings.SplitN(id, "-", 2)
	if len(split) < 2 {
		return "", ""
	}
	return split[0], split[1]
}

// FormatTxid concatenates a company bcid and a uuid into a txid.
func FormatTxid(companyBcid string, uuid string) string {
	return fmt.Sprintf("%s-%s", companyBcid, uuid)
}
