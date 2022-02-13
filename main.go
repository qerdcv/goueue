package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	hostFlag    = "host"
	portFlag    = "port"
	topicFlag   = "topic"
	messageFlag = "message"
)

func main() {
	app := &cli.App{
		Name:        "GOUEUE",
		Description: "message queue writen with go",
		Authors: []*cli.Author{
			{
				Name:  "Vadym Tishchenko",
				Email: "vadym.tishchenko2@gmail.com",
			},
		},
		Commands: []*cli.Command{
			{
				Name:        "pub",
				Description: "Pub to the topic",
				Action:      PubAction,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  hostFlag,
						Value: "localhost",
						Usage: "Host",
					},
					&cli.StringFlag{
						Name:    portFlag,
						Aliases: []string{"p"},
						Value:   "1883",
						Usage:   "Port",
					},
					&cli.StringFlag{
						Name:     topicFlag,
						Aliases:  []string{"t"},
						Usage:    "Queue topic",
						Required: true,
					},
					&cli.StringFlag{
						Name:    messageFlag,
						Aliases: []string{"m"},
						Usage:   "Message topic",
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
