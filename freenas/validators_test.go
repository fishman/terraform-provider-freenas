package freenas

import "testing"

func TestValidateAllowedStringValue(t *testing.T) {
	exceptValues := Toggles
	validValues := []string{"on"}
	for _, v := range validValues {
		_, errors := validateAllowedStringValue(exceptValues)(v, "allowvalue")
		if len(errors) != 0 {
			t.Fatalf("%q should be a valid value in %#v: %q", v, exceptValues, errors)
		}
	}

	invalidValues := []string{"terraform"}
	for _, v := range invalidValues {
		_, errors := validateAllowedStringValue(exceptValues)(v, "allowvalue")
		if len(errors) == 0 {
			t.Fatalf("%q should be an invalid value", v)
		}
	}
}
