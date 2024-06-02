// Package checkvalue - check the value
package checkvalue

import (
	"fmt"
	"strings"
)

// CheckEmptyString - check if string is empty
//
// param: item
// param: val
// return: error
func CheckEmptyString(item string, val string) error {
	if len(strings.TrimSpace(val)) == 0 {
		return fmt.Errorf("the value of %s is empty", item)
	}
	return nil
}
