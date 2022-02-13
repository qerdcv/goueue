package socket

func (s *Sock) Disconnect() error {
	payload := []byte{
		0xE0, // disconnect
		0x00, // normal disconnect
	}

	_, err := s.c.Write(payload)
	return err
}
