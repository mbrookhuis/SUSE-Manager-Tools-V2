// Package base64conv - encode and decode string base64
package base64conv

import (
	b64 "encoding/base64"
)

// Encode - encode string
//
// param: data
// return: string
func Encode(data string) string {
	return b64.StdEncoding.EncodeToString([]byte(data))
}

// Decode - decode string
//
// param: data
// return: output
// return: error
func Decode(data string) (output string, err error) {
	ot, e := b64.StdEncoding.DecodeString(data)
	if e != nil {
		return "", e
	}
	return string(ot), nil
}
