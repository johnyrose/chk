package chk

import "golang.org/x/crypto/ssh"

func CheckSSH(address string) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User: "user",
		Auth: []ssh.AuthMethod{
			ssh.Password("password"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	con, err := ssh.Dial("tcp", address, config)
	if err != nil {
		return nil, err
	}
	return con, nil
}
