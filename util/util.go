package util

// SafeString takes a string pointer and always returns a value string. If the pointer is nil
// an empty string is returned.
func SafeString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
