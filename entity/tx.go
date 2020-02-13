/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package entity

import (
    "encoding/json"
    "errors"
    "fmt"
    "reflect"

    "github.com/transchain/sdk-go/crypto/ed25519"
    kcJson "github.com/transchain/sdk-go/json"

    "github.com/katena-chain/sdk-go/entity/account"
    "github.com/katena-chain/sdk-go/entity/certify"
)

var AvailableTxDataTypes = map[string]reflect.Type{
    certify.GetTypeCertificateRawV1():     reflect.TypeOf(certify.CertificateRawV1{}),
    certify.GetTypeCertificateEd25519V1(): reflect.TypeOf(certify.CertificateEd25519V1{}),
    certify.GetTypeSecretNaclBoxV1():      reflect.TypeOf(certify.SecretNaclBoxV1{}),
    account.GetTypeKeyCreateV1():          reflect.TypeOf(account.KeyCreateV1{}),
    account.GetTypeKeyRevokeV1():          reflect.TypeOf(account.KeyRevokeV1{}),
}

// Tx wraps a tx data with its signature information and a nonce time to avoid replay attacks.
type Tx struct {
    NonceTime Time              `json:"nonce_time" validate:"required"`
    Data      TxData            `json:"data" validate:"required"`
    Signer    ed25519.PublicKey `json:"signer" validate:"required,len=32"`
    Signature ed25519.Signature `json:"signature" validate:"required,len=64"`
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
            NonceTime: a.NonceTime,
            Signer:    a.Signer,
            Signature: a.Signature,
        },
        Data: kcJson.MarshalWrapper{
            Type:  a.Data.GetType(),
            Value: a.Data,
        },
    })
}

// UnmarshalJSON converts an encoded Tx and create the concrete TxData according to its type information.
func (a *Tx) UnmarshalJSON(data []byte) error {
    jsonTx := &unmarshalTxAlias{
        TxAlias: (*TxAlias)(a),
    }
    if err := json.Unmarshal(data, &jsonTx); err != nil {
        return err
    }
    if txDataType, ok := AvailableTxDataTypes[jsonTx.Data.Type]; ok {
        txData := reflect.New(txDataType).Interface()
        if err := json.Unmarshal(jsonTx.Data.Value, txData); err != nil {
            return err
        }
        a.Data = txData.(TxData)
    } else {
        return errors.New(fmt.Sprintf("unknown tx data type: %s", jsonTx.Data.Type))
    }
    return nil
}


