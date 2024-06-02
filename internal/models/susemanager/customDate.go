// Package sumamodels - structs needed for SUSE Manager API Calls
package sumamodels

import (
	"encoding/json"
	"strings"
	"time"
)

// CustomDate - date format used in api calls
type CustomDate time.Time

// UnmarshalJSON - unmarshal JSON
//
// param: b
// return: error
func (j *CustomDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("Jan 2, 2006, 15:04:05 PM", s)
	if err != nil {
		return err
	}
	*j = CustomDate(t)
	return nil
}

// MarshalJSON - marshal JSON
//
// return: time
// return: error
func (j CustomDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}
