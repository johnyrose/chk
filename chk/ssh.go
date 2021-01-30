package chk

import (
	"time"

	"golang.org/x/crypto/ssh"
)

func CheckSSH(address string, timeout int) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User: "user",
		Auth: []ssh.AuthMethod{
			ssh.Password("password"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Duration(timeout * int(time.Second)),
	}

	con, err := ssh.Dial("tcp", address, config)
	if err != nil {
		return nil, err
	}
	return con, nil
}
