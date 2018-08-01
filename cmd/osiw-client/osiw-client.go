package main

import (
	"context"
	"os"
	"time"

	"io/ioutil"

	"github.com/jirwin/operations-slack-webhooks/src/osiw"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

const (
	DefaultServerAddress = "127.0.0.1:57438"
)

func main() {
	app := cli.App{
		Name:        "osiw-client",
		Description: "post operational slack incoming webhooks",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "server",
				Usage: "the address to send requests to",
				Value: DefaultServerAddress,
			},

			cli.StringFlag{
				Name:  "hostname",
				Usage: "the hostname sending the requests",
			},

			cli.StringFlag{
				Name:  "title",
				Usage: "the title of the message",
				Value: "Status Update",
			},
		},
		Action: func(c *cli.Context) error {
			server := c.String("server")

			hostname := c.String("hostname")
			var err error
			if hostname == "" {
				hostname, err = os.Hostname()
				if err != nil {
					hostname = "unknown-host"
				}
			}

			conn, err := grpc.Dial(server, grpc.WithInsecure())
			if err != nil {
				return cli.NewExitError(err, -1)
			}
			defer conn.Close()

			client := osiw.NewOswiClient(conn)

			msg, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return cli.NewExitError(err, -1)
			}

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			_, err = client.Post(ctx, &osiw.PostRequest{
				Hostname: hostname,
				Title:    c.String("title"),
				Text:     string(msg),
			})
			if err != nil {
				return cli.NewExitError(err, -1)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
