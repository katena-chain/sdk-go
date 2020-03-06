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

	DefaultRoleID      = "default"
	CompanyAdminRoleID = "company_admin"
)

// GetTypeKeyCreateV1 returns the key create v1 type.
func GetTypeKeyCreateV1() string {
	return fmt.Sprintf("%s.%s", GetCategoryKeyCreate(), "v1")
}

// GetCategoryKeyCreate returns the key create category.
func GetCategoryKeyCreate() string {
	return fmt.Sprintf("%s.%s.%s", Namespace, TypeKey, common.TypeCreate)
}

// GetTypeKeyRevokeV1 returns the key revoke v1 type.
func GetTypeKeyRevokeV1() string {
	return fmt.Sprintf("%s.%s", GetCategoryKeyRevoke(), "v1")
}

// GetCategoryKeyRevoke returns the key revoke category.
func GetCategoryKeyRevoke() string {
	return fmt.Sprintf("%s.%s.%s", Namespace, TypeKey, common.TypeRevoke)
}
