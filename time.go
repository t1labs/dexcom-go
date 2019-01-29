package dexcom

import (
	"time"
)

// DexcomTime is a type alias for time.Time, that allows us to implement a special JSON unmarshalling function
type DexcomTime time.Time

// UnmarshalJSON will unmarshal a timestamp that is returned from the Dexcom API
func (t *DexcomTime) UnmarshalJSON(b []byte) error {
	val, err := time.Parse(`"2006-01-02T15:04:05"`, string(b))
	if err != nil {
		return err
	}

	*t = DexcomTime(val)

	return nil
}
