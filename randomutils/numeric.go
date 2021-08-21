package randomutils

func Intn(n int) int {
	return r.Intn(n)
}

func PickInt32(n ...int32) int32 {
	return n[Intn(len(n))]
}
