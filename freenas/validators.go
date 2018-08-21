package freenas

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func validateAllowedStringValue(ss []string) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		value := Trim(v.(string))
		existed := false
		for _, s := range ss {
			if s == value {
				existed = true
				break
			}
		}
		if !existed {
			errors = append(errors, fmt.Errorf(
				"%q must contain a valid string value should be in array %#v, got %q",
				k, ss, value))
		}
		return

	}
}
