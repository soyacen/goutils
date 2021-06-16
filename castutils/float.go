package castutils

func IsFloat32(v interface{}) bool {
	_, ok := v.(float32)
	return ok
}

func ToFloat32(v interface{}) float32 {
	result, _ := v.(float32)
	return result
}

func IsFloat64(v interface{}) bool {
	_, ok := v.(float64)
	return ok
}

func ToFloat64(v interface{}) float64 {
	result, _ := v.(float64)
	return result
}
