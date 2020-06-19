/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package entity

import (
	"encoding/json"

	"github.com/transchain/sdk-go/crypto/ed25519"
	kcJson "github.com/transchain/sdk-go/json"
)

// Tx wraps a TxData with its signature, its signer id and a nonce time.
type Tx struct {
	// To avoid replay attack, this value is checked in the replay protector.
	NonceTime Time `json:"nonce_time" validate:"required"`

	// Transaction data, this value is handled by the application.
	Data TxData `json:"data" validate:"required"`

	// Transaction signer and signature info to identify the sender.
	SignerFqId string            `json:"signer_fqid" validate:"required,fqid"`
	Signature  ed25519.Signature `json:"signature" validate:"required,len=64"`
}

// TxAlias is useful to avoid triggering the custom Tx MarshalJSON method.
type TxAlias Tx

// marshalTxAlias wraps a Tx for marshalling operations.
type marshalTxAlias struct {
	Data kcJson.MarshalWrapper `json:"data" validate:"required"`
	*TxAlias
}

// unmarshalTxAlias wraps a Tx for unmarshalling operations.
type unmarshalTxAlias struct {
	Data kcJson.UnmarshalWrapper `json:"data" validate:"required"`
	*TxAlias
}

// MarshalJSON encodes a Tx to add the TxData type information.
func (a Tx) MarshalJSON() ([]byte, error) {
	return json.Marshal(marshalTxAlias{
		TxAlias: &TxAlias{
			NonceTime:  a.NonceTime,
			SignerFqId: a.SignerFqId,
			Signature:  a.Signature,
		},
		Data: kcJson.MarshalWrapper{
			Type:  a.Data.GetType(),
			Value: a.Data,
		},
	})
}

// UnmarshalJSON converts an encoded Tx and create the concrete TxData according to its type information.
func (a *Tx) UnmarshalJSON(data []byte) error {
	jsonTx := unmarshalTxAlias{
		TxAlias: (*TxAlias)(a),
	}
	if err := json.Unmarshal(data, &jsonTx); err != nil {
		return err
	}
	txData, err := UnmarshalTxData(&jsonTx.Data)
	if err != nil {
		return err
	}
	a.Data = txData
	return nil
}
