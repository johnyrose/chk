package chk

import (
	"net"
	"time"
)

func CheckTCP(address string, timeout int) (net.Conn, error) {
	d := net.Dialer{Timeout: time.Duration(timeout * int(time.Second))}
	conn, err := d.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	conn.Close()
	return conn, nil
}
