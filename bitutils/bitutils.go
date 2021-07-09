package bitutils

// GetBitValue get the value of the specified position of the operand
func GetBitValue(source, position int) int {
	return (source >> position) & 1
}

// SetBitValue set the value of the specified position of the opera to the specified value
func SetBitValue(source, position, value int) int {
	mask := 1 << position
	if value > 0 {
		source |= mask
	} else {
		source &= ^mask
	}
	return source
}

// ReverseBitValue reverse the specified position of the operand
func ReverseBitValue(source, position int) int {
	mask := 1 << position
	return source ^ mask
}
