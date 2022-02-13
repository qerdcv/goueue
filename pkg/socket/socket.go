package socket

import (
	"fmt"
	"net"
)

type Sock struct {
	c net.Conn
}

func New(host, port string) (*Sock, error) {
	c, err := net.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		return nil, err
	}

	return &Sock{c}, err
}

func (s *Sock) Close() error {
	return s.c.Close()
}
