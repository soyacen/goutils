package urlutils

import (
	"net/url"
	"strings"
)

func EncodeURLParams(url string, params url.Values) string {
	if len(params) == 0 {
		return url
	}
	return strings.Join([]string{url, "?", params.Encode()}, "")
}
