/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package entity

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// HexBytes allows to json marshal/unmarshal a byte array into an hex encoded string.
type HexBytes []byte

// MarshalJSON encodes a byte array to an hex string.
func (hb HexBytes) MarshalJSON() ([]byte, error) {
	str := hb.String()
	data := make([]byte, len(str)+2)
	data[0] = '"'
	copy(data[1:], []byte(str))
	data[len(data)-1] = '"'
	return data, nil
}

// UnmarshalJSON converts an hex encoded string into a byte array.
func (hb *HexBytes) UnmarshalJSON(data []byte) error {
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return fmt.Errorf("invalid hex string: %s", data)
	}
	hexData, err := hex.DecodeString(string(data[1 : len(data)-1]))
	if err != nil {
		return err
	}
	*hb = hexData
	return nil
}

// String returns the hex string representation of a byte array.
func (hb HexBytes) String() string {
	return strings.ToUpper(hex.EncodeToString(hb))
}
