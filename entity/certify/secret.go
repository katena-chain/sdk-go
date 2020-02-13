/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package certify

import (
    "fmt"

    "github.com/transchain/sdk-go/crypto/nacl"
)

const (
    TypeSecret  = "secret"
    TypeNaclBox = "nacl_box"
)

// SecretNaclBoxV1 is the first version of a nacl box secret.
type SecretNaclBoxV1 struct {
    Id      string         `json:"id" validate:"required,txid"`
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

// GetType returns the type string representation.
func (s SecretNaclBoxV1) GetType() string {
    return GetTypeSecretNaclBoxV1()
}

// GetId returns the id value.
func (s SecretNaclBoxV1) GetId() string {
    return s.Id
}

// GetNamespace returns the certify namespace.
func (s SecretNaclBoxV1) GetNamespace() string {
    return Namespace
}

// GetCategory returns the secret category.
func (s SecretNaclBoxV1) GetCategory() string {
    return GetCategorySecret()
}

// GetCategorySecret returns the secret category.
func GetCategorySecret() string {
    return fmt.Sprintf("%s.%s", Namespace, TypeSecret)
}

// GetTypeSecretNaclBoxV1 returns the secret nacl box v1 type.
func GetTypeSecretNaclBoxV1() string {
    return fmt.Sprintf("%s.%s.%s", GetCategorySecret(), TypeNaclBox, "v1")
}
