package freenas

import (
	"fmt"
	"strings"
)

func unconvertibleIdErr(id string, err error) *unconvertibleIdError {
	return &unconvertibleIdError{OriginalId: id, OriginalError: err}
}

type unconvertibleIdError struct {
	OriginalId    string
	OriginalError error
}

func (e *unconvertibleIdError) Error() string {
	return fmt.Sprintf("Unexpected ID format (%q), expected numerical ID. %s",
		e.OriginalId, e.OriginalError.Error())
}

// Remove useless blank in the string.
func Trim(v string) string {
	if len(v) < 1 {
		return v
	}
	return strings.Trim(v, " ")
}
