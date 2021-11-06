package helpers

// RequiredFieldsIncluded checks whether the given set of fields are contained within the
// request body.
//
// Returns the first missing field found
func RequiredFieldsIncluded(req map[string]string, fields []string) (string, bool) {
	for _, field := range fields {
		if _, ok := req[field]; !ok {
			return field, false
		}
	}
	return "", true
}
