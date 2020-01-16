package account

import (
    "fmt"

    "github.com/transchain/sdk-go/crypto/ed25519"

    "github.com/katena-chain/sdk-go/entity/common"
)

const (
    TypeKey = "key"

    DefaultRoleID      = "default"
    CompanyAdminRoleID = "company_admin"
)

type KeyCreateV1 struct {
    Id        string            `json:"id" validate:"required,txid"`
    PublicKey ed25519.PublicKey `json:"public_key" validate:"required,len=32"`
    Role      string            `json:"role" validate:"required,min=1"`
}

func NewKeyCreateV1(id string, publicKey ed25519.PublicKey, role string) KeyCreateV1 {
    return KeyCreateV1{
        Id:        id,
        PublicKey: publicKey,
        Role:      role,
    }
}

func (kc KeyCreateV1) GetType() string {
    return GetTypeKeyCreateV1()
}

func (kc KeyCreateV1) GetId() string {
    return kc.Id
}

func (kc KeyCreateV1) GetNamespace() string {
    return Namespace
}

func (kc KeyCreateV1) GetSubNamespace() string {
    return GetKeyCreateSubNamespace()
}

func GetTypeKeyCreateV1() string {
    return fmt.Sprintf("%s.%s", GetKeyCreateSubNamespace(), "v1")
}

func GetKeyCreateSubNamespace() string {
    return fmt.Sprintf("%s.%s.%s", Namespace, TypeKey, common.TypeCreate)
}

type KeyRevokeV1 struct {
    Id        string            `json:"id" validate:"required,txid"`
    PublicKey ed25519.PublicKey `json:"public_key" validate:"required,len=32"`
}

func NewKeyRevokeV1(id string, publicKey ed25519.PublicKey) KeyRevokeV1 {
    return KeyRevokeV1{
        Id:        id,
        PublicKey: publicKey,
    }
}

func (akc KeyRevokeV1) GetType() string {
    return GetTypeKeyRevokeV1()
}

func (akc KeyRevokeV1) GetId() string {
    return akc.Id
}

func (akc KeyRevokeV1) GetNamespace() string {
    return Namespace
}

func (akc KeyRevokeV1) GetSubNamespace() string {
    return GetKeyRevokeSubNamespace()
}

func GetTypeKeyRevokeV1() string {
    return fmt.Sprintf("%s.%s", GetKeyRevokeSubNamespace(), "v1")
}

func GetKeyRevokeSubNamespace() string {
    return fmt.Sprintf("%s.%s.%s", Namespace, TypeKey, common.TypeRevoke)
}
