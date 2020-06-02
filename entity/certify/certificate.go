/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package certify

import (
	"github.com/transchain/sdk-go/crypto/ed25519"

	"github.com/katena-chain/sdk-go/entity/common"
)

// CertificateRawV1 is the first version of a raw certificate.
type CertificateRawV1 struct {
	Id    string `json:"id" validate:"required,uuid4"`
	Value []byte `json:"value" validate:"required,min=1,max=128"`
}

// CertificateRawV1 constructor.
func NewCertificateRawV1(id string, value []byte) *CertificateRawV1 {
	return &CertificateRawV1{
		Id:    id,
		Value: value,
	}
}

// GetStateIds returns key-value pairs of id keys and id values.
func (cr CertificateRawV1) GetStateIds(signerCompanyBcId string) map[string]string {
	return map[string]string{
		GetCertificateIdKey(): common.ConcatFqId(signerCompanyBcId, cr.Id),
	}
}

// GetNamespace returns the certify namespace.
func (cr CertificateRawV1) GetNamespace() string {
	return Namespace
}

// GetType returns the type string representation.
func (cr CertificateRawV1) GetType() string {
	return GetCertificateRawV1Type()
}

// CertificateEd25519V1 is the first version of an ed25519 certificate.
type CertificateEd25519V1 struct {
	Id        string            `json:"id" validate:"required,uuid4"`
	Signer    ed25519.PublicKey `json:"signer" validate:"required,len=32"`
	Signature ed25519.Signature `json:"signature" validate:"required,len=64"`
}

// CertificateEd25519V1 constructor.
func NewCertificateEd25519V1(id string, signer ed25519.PublicKey, signature ed25519.Signature) *CertificateEd25519V1 {
	return &CertificateEd25519V1{
		Id:        id,
		Signer:    signer,
		Signature: signature,
	}
}

// GetStateIds returns key-value pairs of id keys and id values.
func (ce CertificateEd25519V1) GetStateIds(signerCompanyBcId string) map[string]string {
	return map[string]string{
		GetCertificateIdKey(): common.ConcatFqId(signerCompanyBcId, ce.Id),
	}
}

// GetNamespace returns the certify namespace.
func (ce CertificateEd25519V1) GetNamespace() string {
	return Namespace
}

// GetType returns the type string representation.
func (ce CertificateEd25519V1) GetType() string {
	return GetCertificateEd25519V1Type()
}
