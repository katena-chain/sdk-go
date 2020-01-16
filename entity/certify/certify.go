/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package certify

import (
    "fmt"

    "github.com/katena-chain/sdk-go/entity/common"
)

const (
    NamespaceCertify = "certify"
)

func GetQueryCertifyCertificateID() string {
    return fmt.Sprintf("%s.%s", GetCertificateSubNamespace(), common.AttributeId)
}

func GetQueryCertifySecretID() string {
    return fmt.Sprintf("%s.%s", GetSecretSubNamespace(), common.AttributeId)
}
