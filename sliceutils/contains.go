package sliceutils

import "strings"

const (
	// IndexNotFound means that no elements were found
	IndexNotFound = -1
)

// IsEmptyString Checks if an slice is nil or length equals 0
func IsEmptyString(slice []string) bool {
	return len(slice) <= 0
}

// IndexOfString Finds the index of the given value in the slice.
func IndexOfString(slice []string, value string, startIndex int) int {
	if IsEmptyString(slice) {
		return IndexNotFound
	}
	if startIndex < 0 {
		startIndex = 0
	}
	length := len(slice)
	for i := startIndex; i < length; i++ {
		if value == slice[i] {
			return i
		}
	}
	return IndexNotFound
}

// ContainsString Checks if the value is in the given slice.
func ContainsString(slice []string, value string) bool {
	return IndexOfString(slice, value, 0) != IndexNotFound
}

// PrefixIndexOfString Finds the index of the prefix of the given value in the prefix slice.
func PrefixIndexOfString(slice []string, value string, startIndex int) int {
	if IsEmptyString(slice) {
		return IndexNotFound
	}
	if startIndex < 0 {
		startIndex = 0
	}
	length := len(slice)
	for i := startIndex; i < length; i++ {
		prefix := slice[i]
		if strings.HasPrefix(value, prefix) {
			return i
		}
	}
	return IndexNotFound
}
