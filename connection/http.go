package connection

import (
	"time"

	"github.com/imroc/req"
)

func CheckHTTP(address string, timeout int) (*req.Resp, error) {
	req.SetTimeout(time.Duration(timeout * int(time.Second)))
	return req.Get(address)
}
