/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package certify

import (
	"github.com/katena-chain/sdk-go/crypto/nacl"
	"github.com/katena-chain/sdk-go/entity/common"
)

// SecretNaclBoxV1 is the first version of a nacl box secret.
type SecretNaclBoxV1 struct {
	Id      string         `json:"id" validate:"required,uuid4"`
	Sender  nacl.PublicKey `json:"sender" validate:"required,len=32"`
	Nonce   nacl.BoxNonce  `json:"nonce" validate:"required,len=24"`
	Content []byte         `json:"content" validate:"required,min=1,max=128"`
}

// SecretNaclBoxV1 constructor.
func NewSecretNaclBoxV1(id string, sender nacl.PublicKey, nonce nacl.BoxNonce, content []byte) *SecretNaclBoxV1 {
	return &SecretNaclBoxV1{
		Id:      id,
		Sender:  sender,
		Nonce:   nonce,
		Content: content,
	}
}

// GetStateIds returns key-value pairs of id keys and id values.
func (snb SecretNaclBoxV1) GetStateIds(signerCompanyBcId string) map[string]string {
	return map[string]string{
		GetSecretIdKey(): common.ConcatFqId(signerCompanyBcId, snb.Id),
	}
}

// GetNamespace returns the certify namespace.
func (snb SecretNaclBoxV1) GetNamespace() string {
	return Namespace
}

// GetType returns the type string representation.
func (snb SecretNaclBoxV1) GetType() string {
	return GetSecretNaclBoxV1Type()
}
