package socket

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func (s *Sock) Connect() error {
	payload := []byte{
		0x10,     // connect
		0x37,     // remaining length
		0x0, 0x4, // length of protocol name (version not incluede)
		0x4D, 0x51, 0x54, 0x54, 0x4, // protocol name + version
		0x2,       // connection flags
		0x0, 0x3C, // keep alive
		0x0, 0x2B, // payload length (prefix[goueue-] + uuid = 43 or 2B in hex)
	}

	payload = append(payload, []byte(strings.ToUpper(fmt.Sprintf("goueue-%s", uuid.New().String())))...)

	_, err := s.c.Write(payload)
	return err
}
