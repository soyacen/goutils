package sliceutils

// CloneString copies the given slice and adds the given element at the beginning of the new slice.
func CloneString(slice []string) []string {
	if slice == nil {
		return nil
	}
	dst := make([]string, len(slice))
	copy(dst, slice)
	return dst
}
