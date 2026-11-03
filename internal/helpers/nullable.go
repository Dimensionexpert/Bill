package helpers

import (
	"strings"
)

// this func will return nill for the empty strings, and return normal strings with trimspace
func Nullable(s string) any {
	if strings.TrimSpace(s) == "" {
		return nil
	}
	return strings.TrimSpace(s)
}
