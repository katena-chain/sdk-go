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

// GetCategoryCertificate returns the certificate category.
func GetCategoryCertificate() string {
	return fmt.Sprintf("%s.%s", Namespace, TypeCertificate)
}

// GetTypeCertificateRawV1 returns the certificate raw v1 type.
func GetTypeCertificateRawV1() string {
	return fmt.Sprintf("%s.%s.%s", GetCategoryCertificate(), TypeRaw, "v1")
}

// GetTypeCertificateRawV1 returns the certificate ed25519 v1 type.
func GetTypeCertificateEd25519V1() string {
	return fmt.Sprintf("%s.%s.%s", GetCategoryCertificate(), TypeEd25519, "v1")
}

// GetCategorySecret returns the secret category.
func GetCategorySecret() string {
	return fmt.Sprintf("%s.%s", Namespace, TypeSecret)
}

// GetTypeSecretNaclBoxV1 returns the secret nacl box v1 type.
func GetTypeSecretNaclBoxV1() string {
	return fmt.Sprintf("%s.%s.%s", GetCategorySecret(), TypeNaclBox, "v1")
}
