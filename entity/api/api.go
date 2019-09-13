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

// Error allows to wrap API errors.
type Error struct {
    Code    uint32 `json:"code"`
    Message string `json:"message"`
}

// Error returns the error formatted as a string (error interface requirement).
func (e Error) Error() string {
    return fmt.Sprintf(`api error:
  Code    : %d
  Message : %s`, e.Code, e.Message)
}
