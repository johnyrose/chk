package chk

import "net/http"

func CheckHTTP(address string) (*http.Response, error) {
	return http.Get(address)
}
