package castutils

func IsString(v interface{}) bool {
	_, ok := v.(string)
	return ok
}

func ToString(v interface{}) string {
	result, _ := v.(string)
	return result
}
