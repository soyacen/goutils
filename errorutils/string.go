package errorutils

func String(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
