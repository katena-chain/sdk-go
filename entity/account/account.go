/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package account

import (
	"fmt"

	"github.com/katena-chain/sdk-go/entity/common"
)

const (
	Namespace = "account"
	TypeKey   = "key"

	DefaultRoleId      = "default"
	CompanyAdminRoleId = "company_admin"
)

// GetKeyIdKey returns the id key to index a key.
func GetKeyIdKey() string {
	return fmt.Sprintf("%s.%s", Namespace, TypeKey)
}

// GetKeyCreateV1Type returns the type string representation of a KeyCreateV1.
func GetKeyCreateV1Type() string {
	return fmt.Sprintf("%s.%s.%s", GetKeyIdKey(), common.TypeCreate, "v1")
}

// GetKeyRotateV1Type returns the type string representation of a KeRotateV1.
func GetKeyRotateV1Type() string {
	return fmt.Sprintf("%s.%s.%s", GetKeyIdKey(), common.TypeRotate, "v1")
}

// GetKeyRevokeV1Type returns the type string representation of a KeyRevokeV1.
func GetKeyRevokeV1Type() string {
	return fmt.Sprintf("%s.%s.%s", GetKeyIdKey(), common.TypeRevoke, "v1")
}
