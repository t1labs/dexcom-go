package dexcom

import (
	"time"
)

// Time is a type alias for time.Time, that allows us to implement a special JSON unmarshalling function
type Time time.Time

// UnmarshalJSON will unmarshal a timestamp that is returned from the Dexcom API
func (t *Time) UnmarshalJSON(b []byte) error {
	val, err := time.Parse(`"2006-01-02T15:04:05"`, string(b))
	if err != nil {
		return err
	}

	*t = Time(val)

	return nil
}

// IsZero will tell us whether or not the Dexcom time is the zero value
func (t *Time) IsZero() bool {
	return time.Time(*t).IsZero()
}
