package castutils

func IsBool(v interface{}) bool {
	_, ok := v.(bool)
	return ok
}

func ToBool(v interface{}) bool {
	result, _ := v.(bool)
	return result
}

func IsBoolSlice(v interface{}) bool {
	_, ok := v.([]bool)
	return ok
}

func ToBoolSlice(v interface{}) []bool {
	result, _ := v.([]bool)
	return result
}
