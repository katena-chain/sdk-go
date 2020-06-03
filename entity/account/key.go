package account

import (
	"github.com/transchain/sdk-go/crypto/ed25519"

	"github.com/katena-chain/sdk-go/entity/common"
)

// KeyV1 is the first version of a key.
type KeyV1 struct {
	FqId      string            `json:"fqid" validate:"required,fqid"`
	PublicKey ed25519.PublicKey `json:"public_key" validate:"required,len=32"`
	IsActive  bool              `json:"is_active" validate:"required"`
	Role      string            `json:"role" validate:"required,min=1"`
}

// KeyV1 constructor.
func NewKeyV1(fqId string, publicKey ed25519.PublicKey, isActive bool, role string) *KeyV1 {
	return &KeyV1{
		FqId:      fqId,
		PublicKey: publicKey,
		IsActive:  isActive,
		Role:      role,
	}
}

// KeyCreateV1 is the first version of a key create message.
type KeyCreateV1 struct {
	Id        string            `json:"id" validate:"required,uuid4"`
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

// GetStateIds returns key-value pairs of id keys and id values.
func (kc KeyCreateV1) GetStateIds(signerCompanyBcId string) map[string]string {
	return map[string]string{
		GetKeyIdKey(): common.ConcatFqId(signerCompanyBcId, kc.Id),
	}
}

// GetNamespace returns the account namespace.
func (kc KeyCreateV1) GetNamespace() string {
	return Namespace
}

// GetType returns the type string representation.
func (kc KeyCreateV1) GetType() string {
	return GetKeyCreateV1Type()
}

// KeyRotateV1 is the first version of a key rotate message.
type KeyRotateV1 struct {
	Id        string            `json:"id" validate:"required,uuid4"`
	PublicKey ed25519.PublicKey `json:"public_key" validate:"required,len=32"`
}

// KeyRotateV1 constructor.
func NewKeyRotateV1(id string, publicKey ed25519.PublicKey) *KeyRotateV1 {
	return &KeyRotateV1{
		Id:        id,
		PublicKey: publicKey,
	}
}

// GetStateIds returns key-value pairs of id keys and id values.
func (kr KeyRotateV1) GetStateIds(signerCompanyBcId string) map[string]string {
	return map[string]string{
		GetKeyIdKey(): common.ConcatFqId(signerCompanyBcId, kr.Id),
	}
}

// GetNamespace returns the account namespace.
func (kr KeyRotateV1) GetNamespace() string {
	return Namespace
}

// GetType returns the type string representation.
func (kr KeyRotateV1) GetType() string {
	return GetKeyRotateV1Type()
}

// KeyRevokeV1 is the first version of a key revoke message.
type KeyRevokeV1 struct {
	Id string `json:"id" validate:"required,uuid4"`
}

// KeyRevokeV1 constructor.
func NewKeyRevokeV1(id string) *KeyRevokeV1 {
	return &KeyRevokeV1{
		Id: id,
	}
}

// GetStateIds returns key-value pairs of id keys and id values.
func (kr KeyRevokeV1) GetStateIds(signerCompanyBcId string) map[string]string {
	return map[string]string{
		GetKeyIdKey(): common.ConcatFqId(signerCompanyBcId, kr.Id),
	}
}

// GetNamespace returns the account namespace.
func (kr KeyRevokeV1) GetNamespace() string {
	return Namespace
}

// GetType returns the type string representation.
func (kr KeyRevokeV1) GetType() string {
	return GetKeyRevokeV1Type()
}
