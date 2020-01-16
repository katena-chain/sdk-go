/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package api

import (
    "fmt"

    "github.com/katena-chain/sdk-go/entity"
)

// TxWrappers wraps a list of TxWrapper with the total txs available.
type TxWrappers struct {
    Txs   []*TxWrapper `json:"txs"`
    Total uint32       `json:"total"`
}

// TxWrapper wraps a tx and its status.
type TxWrapper struct {
    Tx     *entity.Tx `json:"tx"`
    Status *TxStatus  `json:"status"`
}

// TxStatus is a tx blockchain status.
// 0: OK
// 1: PENDING
// >1: ERROR WITH CORRESPONDING CODE
type TxStatus struct {
    Code    uint32 `json:"code"`
    Message string `json:"message"`
}
// PublicError allows to wrap API errors.
type PublicError struct {
    Codespace string `json:"codespace,omitempty"`
    Code      uint32 `json:"code"`
    Message   string `json:"message"`
}

// PublicError constructor.
func NewPublicError(codespace string, code uint32, message string) *PublicError {
    return &PublicError{
        Codespace: codespace,
        Code:      code,
        Message:   message,
    }
}

// PublicError returns the error formatted as a string (error interface requirement).
func (pe PublicError) Error() string {
    return fmt.Sprintf(`ERROR:
Codespace: %s
Code: %d
Message: %s
`, pe.Codespace, pe.Code, pe.Message)
}
