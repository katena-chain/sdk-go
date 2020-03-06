package account

import (
	"github.com/transchain/sdk-go/crypto/ed25519"
)

// KeyV1 is the first version of a key.
type KeyV1 struct {
	CompanyBcid string            `json:"company_bcid" validate:"required,len=6,alpha"`
	PublicKey   ed25519.PublicKey `json:"public_key" validate:"required,len=32"`
	IsActive    bool              `json:"is_active" validate:"required"`
	Role        string            `json:"role"`
}

// KeyV1 constructor.
func NewKeyV1(companyBcid string, publicKey ed25519.PublicKey, isActive bool, role string) *KeyV1 {
	return &KeyV1{
		CompanyBcid: companyBcid,
		PublicKey:   publicKey,
		IsActive:    isActive,
		Role:        role,
	}
}

// KeyCreateV1 is the first version of a key create message.
type KeyCreateV1 struct {
	Id        string            `json:"id" validate:"required,txid"`
	PublicKey ed25519.PublicKey `json:"public_key" validate:"required,len=32"`
	Role      string            `json:"role" validate:"required,min=1"`
}

// KeyCreateV1 constructor.
func NewKeyCreateV1(id string, publicKey ed25519.PublicKey, role string) *KeyCreateV1 {
	return &KeyCreateV1{
		Id:        id,
		PublicKey: publicKey,
		Role:      role,
	}
}

// GetType returns the type string representation.
func (kc KeyCreateV1) GetType() string {
	return GetTypeKeyCreateV1()
}

// GetId returns the id value.
func (kc KeyCreateV1) GetId() string {
	return kc.Id
}

// GetNamespace returns the certify namespace.
func (kc KeyCreateV1) GetNamespace() string {
	return Namespace
}

// GetCategory returns the key create category.
func (kc KeyCreateV1) GetCategory() string {
	return GetCategoryKeyCreate()
}

// KeyRevokeV1 is the first version of a key revoke message.
type KeyRevokeV1 struct {
	Id        string            `json:"id" validate:"required,txid"`
	PublicKey ed25519.PublicKey `json:"public_key" validate:"required,len=32"`
}

// KeyCreateV1 constructor.
func NewKeyRevokeV1(id string, publicKey ed25519.PublicKey) *KeyRevokeV1 {
	return &KeyRevokeV1{
		Id:        id,
		PublicKey: publicKey,
	}
}

// GetType returns the type string representation.
func (kr KeyRevokeV1) GetType() string {
	return GetTypeKeyRevokeV1()
}

// GetId returns the id value.
func (kr KeyRevokeV1) GetId() string {
	return kr.Id
}

// GetNamespace returns the certify namespace.
func (kr KeyRevokeV1) GetNamespace() string {
	return Namespace
}

// GetCategory returns the key revoke category.
func (kr KeyRevokeV1) GetCategory() string {
	return GetCategoryKeyRevoke()
}
