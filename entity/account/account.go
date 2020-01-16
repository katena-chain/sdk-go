package account

import (
    "fmt"

    "github.com/katena-chain/sdk-go/entity/common"
)

const (
    Namespace = "account"
)


func GetQueryKeyCreateID() string {
    return fmt.Sprintf("%s.%s", GetKeyCreateSubNamespace(), common.AttributeId)
}

func GetQueryKeyRevokeID() string {
    return fmt.Sprintf("%s.%s", GetKeyRevokeSubNamespace(), common.AttributeId)
}
