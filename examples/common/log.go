/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package common

import (
	"encoding/json"
	"fmt"
)

func PrintlnJSON(data interface{}) error {
	encodedData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("%s\n", encodedData))
	return nil
}
