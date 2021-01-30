package chk

import "net"

func CheckTCP(address string) (net.Conn, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	conn.Close()
	return conn, nil
}
