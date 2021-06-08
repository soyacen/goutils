package sliceutils

// RemoveElementString removes the first occurrence of the specified element from the
// specified slice.
// All subsequent elements are shifted to the left (subtracts one from their indices).
// If the array doesn't contains such an element, no elements are removed from the array.
func RemoveElementString(slice []string, element string) []string {
	index := IndexOfString(slice, element, 0)
	if index == IndexNotFound {
		return CloneString(slice)
	}
	result, _ := RemoveString(slice, index)
	return result
}
