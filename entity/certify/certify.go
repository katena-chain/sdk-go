/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package certify

import (
	"fmt"
)

const (
	Namespace       = "certify"
	TypeCertificate = "certificate"
	TypeRaw         = "raw"
	TypeEd25519     = "ed25519"
	TypeSecret      = "secret"
	TypeNaclBox     = "nacl_box"
)

// GetCertificateIdKey returns the id key to index a certificate.
func GetCertificateIdKey() string {
	return fmt.Sprintf("%s.%s", Namespace, TypeCertificate)
}

// GetCertificateRawV1Type returns the type string representation of a CertificateRawV1.
func GetCertificateRawV1Type() string {
	return fmt.Sprintf("%s.%s.%s", GetCertificateIdKey(), TypeRaw, "v1")
}

// GetCertificateRawV1Type returns the type string representation of a CertificateEd25519V1.
func GetCertificateEd25519V1Type() string {
	return fmt.Sprintf("%s.%s.%s", GetCertificateIdKey(), TypeEd25519, "v1")
}

// GetSecretIdKey returns returns the id key to index a secret.
func GetSecretIdKey() string {
	return fmt.Sprintf("%s.%s", Namespace, TypeSecret)
}

// GetSecretNaclBoxV1Type returns the type string representation of a SecretNaclBoxV1.
func GetSecretNaclBoxV1Type() string {
	return fmt.Sprintf("%s.%s.%s", GetSecretIdKey(), TypeNaclBox, "v1")
}
