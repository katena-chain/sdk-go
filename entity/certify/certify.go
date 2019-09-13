/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package certify

import (
    "fmt"
    "strings"
)

const (
    NamespaceCertify = "certify"
)

// SplitBcid separates a bcid field into a company chain id and a uuid.
func SplitBcid(id string) (string, string) {
    split := strings.SplitN(id, "-", 2)
    if len(split) < 2 {
        return "", ""
    }
    return split[0], split[1]
}

// FormatBcid concatenates a company chain id and a uuid into a bcid.
func FormatBcid(companyChainId string, uuid string) string {
    return fmt.Sprintf("%s-%s", companyChainId, uuid)
}
