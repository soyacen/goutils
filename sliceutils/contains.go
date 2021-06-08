package sliceutils

// ContainsString Checks if the value is in the given slice.
func ContainsString(slice []string, value string) bool {
	return IndexOfString(slice, value, 0) != IndexNotFound
}
