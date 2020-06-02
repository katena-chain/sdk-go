/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package common

import (
	"fmt"
	"strings"
)

const (
	TypeCreate = "create"
	TypeRotate = "rotate"
	TypeRevoke = "revoke"
)

// SplitFqId splits a fully qualified id into a company bcid and a uuid.
func SplitFqId(fqid string) (string, string) {
	split := strings.SplitN(fqid, "-", 2)
	if len(split) < 2 {
		return "", ""
	}
	return split[0], split[1]
}

// ConcatFqId concatenates a company bcid and a uuid into a fully qualified id.
func ConcatFqId(companyBcId string, uuid string) string {
	return fmt.Sprintf("%s-%s", companyBcId, uuid)
}