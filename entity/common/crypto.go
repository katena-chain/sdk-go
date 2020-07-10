package common

import (
	"crypto/rand"
	"encoding/base64"

	oasisEd25519 "github.com/oasisprotocol/ed25519"
	"golang.org/x/crypto/nacl/box"

	"github.com/katena-chain/sdk-go/crypto/ed25519"
	"github.com/katena-chain/sdk-go/crypto/nacl"
)

// CreatePrivateKeyEd25519FromBase64 accepts a base64 encoded Ed25519 private key (88 chars) and returns an Ed25519 private key.
func CreatePrivateKeyEd25519FromBase64(privateKeyBase64 string) ed25519.PrivateKey {
	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyBase64)
	if err != nil {
		panic(ed25519.ErrBadPrivateKeyBase64Format)
	}
	return ed25519.NewPrivateKey(privateKeyBytes)
}

// CreatePublicKeyEd25519FromBase64 accepts a base64 encoded Ed25519 public key (44 chars) and returns an Ed25519 public key.
func CreatePublicKeyEd25519FromBase64(publicKeyBase64 string) ed25519.PublicKey {
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyBase64)
	if err != nil {
		panic(ed25519.ErrBadPublicKeyBase64Format)
	}
	return ed25519.NewPublicKey(publicKeyBytes)
}

// GenerateNewPrivateKeyEd25519 generates a new ed25519 private key.
func GenerateNewPrivateKeyEd25519() ed25519.PrivateKey {
	_, privKey, err := oasisEd25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	return ed25519.NewPrivateKey(privKey)
}

// CreatePrivateKeyX25519FromBase64 accepts a base64 encoded X25519 private key (88 chars) and returns an X25519 private key.
func CreatePrivateKeyX25519FromBase64(privateKeyBase64 string) nacl.PrivateKey {
	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyBase64)
	if err != nil {
		panic(nacl.ErrBadPrivateKeyBase64Format)
	}
	return nacl.NewPrivateKey(privateKeyBytes)
}

// CreatePublicKeyX25519FromBase64 accepts a base64 encoded X25519 public key (44 chars) and returns an X25519 public key.
func CreatePublicKeyX25519FromBase64(publicKeyBase64 string) nacl.PublicKey {
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyBase64)
	if err != nil {
		panic(nacl.ErrBadPublicKeyBase64Format)
	}
	return nacl.NewPublicKey(publicKeyBytes)
}

// GenerateNewPrivateKeyX25519 generates a new x25519 private key.
func GenerateNewPrivateKeyX25519() nacl.PrivateKey {
	pubKey, privKey, err := box.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	privateKeyBytes := make([]byte, nacl.PrivateKeySize)
	copy(privateKeyBytes, privKey[:])
	copy(privateKeyBytes[32:], pubKey[:])
	return nacl.NewPrivateKey(privateKeyBytes)
}