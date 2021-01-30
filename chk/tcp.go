package chk

import "net"

func CheckTCP(address string) (*net.TCPConn, error) {
	fullAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTCP("tcp", nil, fullAddr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
