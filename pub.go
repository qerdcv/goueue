package main

import (
	"fmt"
	"log"
	"net"

	"github.com/qerdcv/goueue/pkg/socket"
	"github.com/urfave/cli/v2"
)

func createSocket(host, port string) (net.Conn, error) {
	return net.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
}

func PubAction(c *cli.Context) error {
	host := c.String(hostFlag)
	port := c.String(portFlag)
	message := c.String(messageFlag)
	topic := c.String(topicFlag)
	sock, err := socket.New(host, port)
	if err != nil {
		return fmt.Errorf("create socket: %w", err)
	}

	defer func() {
		err := sock.Close()
		if err != nil {
			log.Fatal(fmt.Errorf("close socket: %w", err))
		}
	}()

	err = sock.Connect()
	if err != nil {
		return fmt.Errorf("connect: %w", err)
	}

	err = sock.Publish(topic, message)
	if err != nil {
		return fmt.Errorf("pub: %w", err)
	}

	err = sock.Disconnect()
	if err != nil {
		return fmt.Errorf("disconnect: %w", err)
	}

	return nil
}
