package socket

func (s *Sock) Publish(topic, message string) error {
	payload := []byte{
		0x30,                                // PUB DUP=0 QoS=0 RETAIN=0
		byte(len(topic) + len(message) + 4), // remain length,
		0x0, byte(len(topic)),               // set length of topic
	}

	payload = append(payload, []byte(topic)...)   // set topic name
	payload = append(payload, 0x0)                // packet identifier
	payload = append(payload, byte(len(message))) // set len of message
	payload = append(payload, []byte(message)...) // set message

	_, err := s.c.Write(payload)
	return err
}
