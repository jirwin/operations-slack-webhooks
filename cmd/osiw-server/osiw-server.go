package main

import (
	"os"

	"context"

	"os/signal"

	"fmt"

	"github.com/jirwin/operations-slack-webhooks/src/osiw"
	"github.com/urfave/cli"
	"golang.org/x/sys/unix"
)

const (
	DefaultListenAddress = "127.0.0.1:57438"
)

func main() {
	app := cli.App{
		Name:        "osiw-server",
		Description: "operational slack incoming webhooks",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "listen",
				Usage: "the address to listen on",
				Value: DefaultListenAddress,
			},

			cli.StringFlag{
				Name:  "webhook-url",
				Usage: "the slack incoming webhook url",
			},
		},
		Action: func(c *cli.Context) error {
			listenAddr := c.String("listen")

			if !c.IsSet("webhook-url") {
				return cli.NewExitError("webhook url is required (--webhook-url)", -1)
			}
			webhookUrl := c.String("webhook-url")

			ctx, canc := context.WithCancel(context.Background())

			stop := make(chan os.Signal, 1)
			signal.Notify(stop, unix.SIGINT, unix.SIGTERM, unix.SIGQUIT, unix.SIGABRT)

			s := osiw.NewServer(listenAddr, webhookUrl)
			s.Start(ctx)

		Outer:
			for {
				select {
				case <-ctx.Done():
					s.Stop(ctx)
					break Outer

				case <-stop:
					canc()
				}
			}

			fmt.Println("Quitting osiw server.")

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
