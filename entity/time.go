/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package entity

import (
    "encoding/json"
    "time"
)

const RFC3339MicroZeroPadded = "2006-01-02T15:04:05.000000Z07:00"

// Time is a time.Time wrapper.
type Time struct {
    time.Time
}

// MarshalJSON converts a Time into a time.Time and marshals its UTC value with a microseconds precision.
func (t Time) MarshalJSON() ([]byte, error) {
    return json.Marshal(t.Time.UTC().Truncate(time.Microsecond).Format(RFC3339MicroZeroPadded))
}
