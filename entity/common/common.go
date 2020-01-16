package common

import (
    "fmt"
    "strings"
)

const (
    TypeCreate = "create"
    TypeRevoke = "revoke"

    AttributeId = "id"
)

// SplitTxid separates a txid field into a company chain id and a uuid.
func SplitTxid(id string) (string, string) {
    split := strings.SplitN(id, "-", 2)
    if len(split) < 2 {
        return "", ""
    }
    return split[0], split[1]
}

// FormatTxid concatenates a company chain id and a uuid into a txid.
func FormatTxid(companyChainId string, uuid string) string {
    return fmt.Sprintf("%s-%s", companyChainId, uuid)
}
