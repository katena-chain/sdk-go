/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package entity

import (
	"github.com/transchain/sdk-go/crypto/ed25519"
)

// TxSigner contains all information about a Tx signer.
type TxSigner struct {
	FqId       string
	PrivateKey *ed25519.PrivateKey
}

// TxSigner constructor.
func NewTxSigner(fqId string, privateKey *ed25519.PrivateKey) *TxSigner {
	return &TxSigner{
		FqId:       fqId,
		PrivateKey: privateKey,
	}
}
