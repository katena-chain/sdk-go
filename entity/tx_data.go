/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package entity

import (
	"reflect"

	kcJson "github.com/transchain/sdk-go/json"

	"github.com/katena-chain/sdk-go/entity/account"
	"github.com/katena-chain/sdk-go/entity/certify"
)

var AvailableTxDataTypes = map[string]reflect.Type{
	certify.GetCertificateRawV1Type():     reflect.TypeOf(certify.CertificateRawV1{}),
	certify.GetCertificateEd25519V1Type(): reflect.TypeOf(certify.CertificateEd25519V1{}),
	certify.GetSecretNaclBoxV1Type():      reflect.TypeOf(certify.SecretNaclBoxV1{}),
	account.GetKeyCreateV1Type():          reflect.TypeOf(account.KeyCreateV1{}),
	account.GetKeyRevokeV1Type():          reflect.TypeOf(account.KeyRevokeV1{}),
}

// TxData interface defines the methods a concrete TxData must implement.
type TxData interface {
	// To fetch all the entity state ids a TxData can create/update/delete.
	// This is also used to index Txs.
	GetStateIds(signerCompanyBcId string) map[string]string

	// To identify which plugin should handle this TxData.
	GetNamespace() string

	// To identify its subtype.
	GetType() string
}

// txDataState wraps a TxData and additional values in order to define a unique state ready to be signed.
type txDataState struct {
	ChainId   string                `json:"chain_id"`
	NonceTime Time                  `json:"nonce_time"`
	Data      kcJson.MarshalWrapper `json:"data"`
}

// GetTxDataStateBytes returns the sorted and marshaled json representation of a TxData ready to be signed.
func GetTxDataStateBytes(chainId string, nonceTime Time, txData TxData) []byte {
	data := txDataState{
		ChainId:   chainId,
		NonceTime: nonceTime,
		Data: kcJson.MarshalWrapper{
			Type:  txData.GetType(),
			Value: txData,
		},
	}
	return kcJson.MustMarshalAndSortJSON(data)
}
