/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package certify

import (
    "fmt"

    "github.com/transchain/sdk-go/crypto/ed25519"
)

const (
    TypeCertificate = "certificate"
    TypeRaw         = "raw"
    TypeEd25519     = "ed25519"
)

// CertificateRawV1 is the first version of a raw certificate.
type CertificateRawV1 struct {
    Id    string `json:"id" validate:"required,txid"`
    Value []byte `json:"value" validate:"required,min=1,max=128"`
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
    return NamespaceCertify
}

// GetSubNamespace returns the certificate sub namespace.
func (c CertificateRawV1) GetSubNamespace() string {
    return GetCertificateSubNamespace()
}

// CertificateEd25519V1 is the first version of an ed25519 certificate.
type CertificateEd25519V1 struct {
    Id        string            `json:"id" validate:"required,txid"`
    Signer    ed25519.PublicKey `json:"signer" validate:"required,len=32"`
    Signature ed25519.Signature `json:"signature" validate:"required,len=64"`
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
    return NamespaceCertify
}

// GetSubNamespace returns the certificate sub namespace.
func (ce CertificateEd25519V1) GetSubNamespace() string {
    return GetCertificateSubNamespace()
}

// GetCertificateSubNamespace returns the certificate sub namespace.
func GetCertificateSubNamespace() string {
    return fmt.Sprintf("%s.%s", NamespaceCertify, TypeCertificate)
}

// GetTypeCertificateRawV1 returns the certificate raw v1 type.
func GetTypeCertificateRawV1() string {
    return fmt.Sprintf("%s.%s.%s", GetCertificateSubNamespace(), TypeRaw, "v1")
}

// GetTypeCertificateRawV1 returns the certificate ed25519 v1 type.
func GetTypeCertificateEd25519V1() string {
    return fmt.Sprintf("%s.%s.%s", GetCertificateSubNamespace(), TypeEd25519, "v1")
}
