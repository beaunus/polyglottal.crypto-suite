package util

import "net/url"

// SanitizeParameter returns "" if a parameter by the given name doesn't exist.
func SanitizeParameter(name string, values *url.Values) string {
	result := ""
	if (*values)[name] != nil {
		result = (*values)[name][0]
	}
	return result
}
