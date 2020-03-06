/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package certify

import (
	"github.com/transchain/sdk-go/crypto/ed25519"
)

// CertificateRawV1 is the first version of a raw certificate.
type CertificateRawV1 struct {
	Id    string `json:"id" validate:"required,txid"`
	Value []byte `json:"value" validate:"required,min=1,max=128"`
}

// CertificateRawV1 constructor.
func NewCertificateRawV1(id string, value []byte) *CertificateRawV1 {
	return &CertificateRawV1{
		Id:    id,
		Value: value,
	}
}

// GetType returns the type string representation.
func (c CertificateRawV1) GetType() string {
	return GetTypeCertificateRawV1()
}

// GetId returns the id value.
func (c CertificateRawV1) GetId() string {
	return c.Id
}

// GetNamespace returns the certify namespace.
func (c CertificateRawV1) GetNamespace() string {
	return Namespace
}

// GetCategory returns the certificate category.
func (c CertificateRawV1) GetCategory() string {
	return GetCategoryCertificate()
}

// CertificateEd25519V1 is the first version of an ed25519 certificate.
type CertificateEd25519V1 struct {
	Id        string            `json:"id" validate:"required,txid"`
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

// GetType returns the type string representation.
func (ce CertificateEd25519V1) GetType() string {
	return GetTypeCertificateEd25519V1()
}

// GetId returns the id value.
func (ce CertificateEd25519V1) GetId() string {
	return ce.Id
}

// GetNamespace returns the certify namespace.
func (ce CertificateEd25519V1) GetNamespace() string {
	return Namespace
}

// GetCategory returns the certificate category.
func (ce CertificateEd25519V1) GetCategory() string {
	return GetCategoryCertificate()
}
