package castutils

func IsUInt(v interface{}) bool {
	_, ok := v.(uint)
	return ok
}

func ToUInt(v interface{}) uint {
	result, _ := v.(uint)
	return result
}

func IsUInt8(v interface{}) bool {
	_, ok := v.(uint8)
	return ok
}

func ToUInt8(v interface{}) uint8 {
	result, _ := v.(uint8)
	return result
}

func IsUInt16(v interface{}) bool {
	_, ok := v.(uint16)
	return ok
}

func ToUInt16(v interface{}) uint16 {
	result, _ := v.(uint16)
	return result
}

func IsUInt32(v interface{}) bool {
	_, ok := v.(uint32)
	return ok
}

func ToUInt32(v interface{}) uint32 {
	result, _ := v.(uint32)
	return result
}

func IsUInt64(v interface{}) bool {
	_, ok := v.(uint64)
	return ok
}

func ToUInt64(v interface{}) uint64 {
	result, _ := v.(uint64)
	return result
}

func IsInt(v interface{}) bool {
	_, ok := v.(int)
	return ok
}

func ToInt(v interface{}) int {
	result, _ := v.(int)
	return result
}

func IsInt8(v interface{}) bool {
	_, ok := v.(int8)
	return ok
}

func ToInt8(v interface{}) int8 {
	result, _ := v.(int8)
	return result
}

func IsInt16(v interface{}) bool {
	_, ok := v.(int16)
	return ok
}

func ToInt16(v interface{}) int16 {
	result, _ := v.(int16)
	return result
}

func IsInt32(v interface{}) bool {
	_, ok := v.(int32)
	return ok
}

func ToInt32(v interface{}) int32 {
	result, _ := v.(int32)
	return result
}

func IsInt64(v interface{}) bool {
	_, ok := v.(int64)
	return ok
}

func ToInt64(v interface{}) int64 {
	result, _ := v.(int64)
	return result
}

func IsByte(v interface{}) bool {
	_, ok := v.(byte)
	return ok
}

func ToByte(v interface{}) byte {
	result, _ := v.(byte)
	return result
}

func IsRune(v interface{}) bool {
	_, ok := v.(rune)
	return ok
}

func ToRune(v interface{}) rune {
	result, _ := v.(rune)
	return result
}
